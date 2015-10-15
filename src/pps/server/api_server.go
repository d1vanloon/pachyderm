package server

import (
	"fmt"

	"go.pachyderm.com/pachyderm/src/pkg/container"
	"go.pachyderm.com/pachyderm/src/pps"
	"go.pachyderm.com/pachyderm/src/pps/persist"
	"go.pedge.io/google-protobuf"
	"go.pedge.io/protolog"
	"golang.org/x/net/context"
)

var (
	emptyInstance = &google_protobuf.Empty{}
)

type apiServer struct {
	persistAPIClient persist.APIClient
	containerClient  container.Client
}

func newAPIServer(persistAPIClient persist.APIClient, containerClient container.Client) *apiServer {
	return &apiServer{persistAPIClient, containerClient}
}

func (a *apiServer) CreateJob(ctx context.Context, request *pps.CreateJobRequest) (response *pps.Job, err error) {
	persistJob, err := a.persistAPIClient.CreateJob(ctx, jobToPersist(request.Job))
	if err != nil {
		return nil, err
	}
	return persistToJob(persistJob), nil
}

func (a *apiServer) GetJob(ctx context.Context, request *pps.GetJobRequest) (response *pps.Job, err error) {
	persistJob, err := a.persistAPIClient.GetJobByID(ctx, &google_protobuf.StringValue{Value: request.JobId})
	if err != nil {
		return nil, err
	}
	return persistToJob(persistJob), nil
}

func (a *apiServer) GetJobsByPipelineName(ctx context.Context, request *pps.GetJobsByPipelineNameRequest) (response *pps.Jobs, err error) {
	persistPipelines, err := a.persistAPIClient.GetPipelinesByName(ctx, &google_protobuf.StringValue{Value: request.PipelineName})
	if err != nil {
		return nil, err
	}
	var jobs []*pps.Job
	for _, persistPipeline := range persistPipelines.Pipeline {
		persistJobs, err := a.persistAPIClient.GetJobsByPipelineID(ctx, &google_protobuf.StringValue{Value: persistPipeline.Id})
		if err != nil {
			return nil, err
		}
		iJobs := persistToJobs(persistJobs)
		jobs = append(jobs, iJobs.Job...)
	}
	return &pps.Jobs{
		Job: jobs,
	}, nil
}

func (a *apiServer) StartJob(ctx context.Context, request *pps.StartJobRequest) (response *google_protobuf.Empty, err error) {
	persistJob, err := a.persistAPIClient.GetJobByID(ctx, &google_protobuf.StringValue{Value: request.JobId})
	if err != nil {
		return nil, err
	}
	if err := a.startPersistJob(persistJob); err != nil {
		return nil, err
	}
	return emptyInstance, nil
}

func (a *apiServer) GetJobStatus(ctx context.Context, request *pps.GetJobStatusRequest) (response *pps.JobStatus, err error) {
	persistJobStatuses, err := a.persistAPIClient.GetJobStatusesByJobID(ctx, &google_protobuf.StringValue{Value: request.JobId})
	if err != nil {
		return nil, err
	}
	if len(persistJobStatuses.JobStatus) == 0 {
		return nil, fmt.Errorf("pachyderm.pps.server: no job statuses for %s", request.JobId)
	}
	return persistToJobStatus(persistJobStatuses.JobStatus[0]), nil
}

func (a *apiServer) GetJobLogs(request *pps.GetJobLogsRequest, responseServer pps.API_GetJobLogsServer) (err error) {
	persistJobLogs, err := a.persistAPIClient.GetJobLogsByJobID(context.Background(), &google_protobuf.StringValue{Value: request.JobId})
	if err != nil {
		return err
	}
	for _, persistJobLog := range persistJobLogs.JobLog {
		if persistJobLog.OutputStream == request.OutputStream {
			if err := responseServer.Send(&google_protobuf.BytesValue{Value: persistJobLog.Value}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *apiServer) CreatePipeline(ctx context.Context, request *pps.CreatePipelineRequest) (response *pps.Pipeline, err error) {
	persistPipeline, err := a.persistAPIClient.CreatePipeline(ctx, pipelineToPersist(request.Pipeline))
	if err != nil {
		return nil, err
	}
	return persistToPipeline(persistPipeline), nil
}

func (a *apiServer) GetPipeline(ctx context.Context, request *pps.GetPipelineRequest) (response *pps.Pipeline, err error) {
	persistPipelines, err := a.persistAPIClient.GetPipelinesByName(ctx, &google_protobuf.StringValue{Value: request.PipelineName})
	if err != nil {
		return nil, err
	}
	if len(persistPipelines.Pipeline) == 0 {
		return nil, fmt.Errorf("pachyderm.pps.server: no piplines for name %s", request.PipelineName)
	}
	return persistToPipeline(persistPipelines.Pipeline[0]), nil
}

func (a *apiServer) GetAllPipelines(ctx context.Context, request *google_protobuf.Empty) (response *pps.Pipelines, err error) {
	persistPipelines, err := a.persistAPIClient.GetAllPipelines(ctx, request)
	if err != nil {
		return nil, err
	}
	pipelineMap := make(map[string]*pps.Pipeline)
	for _, persistPipeline := range persistPipelines.Pipeline {
		// pipelines are ordered newest to oldest, so if we have already
		// seen a pipeline with the same name, it is newer
		if _, ok := pipelineMap[persistPipeline.Name]; !ok {
			pipelineMap[persistPipeline.Name] = persistToPipeline(persistPipeline)
		}
	}
	pipelines := make([]*pps.Pipeline, len(pipelineMap))
	i := 0
	for _, pipeline := range pipelineMap {
		pipelines[i] = pipeline
		i++
	}
	return &pps.Pipelines{
		Pipeline: pipelines,
	}, nil
}

func (a *apiServer) startPersistJob(persistJob *persist.Job) error {
	if _, err := a.persistAPIClient.CreateJobStatus(
		context.Background(),
		&persist.JobStatus{
			JobId: persistJob.Id,
			Type:  pps.JobStatusType_JOB_STATUS_TYPE_STARTED,
		},
	); err != nil {
		return err
	}
	// TODO(pedge): throttling? worker pool?
	go func() {
		if err := a.runJob(persistJob); err != nil {
			protolog.Errorln(err.Error())
			// TODO(pedge): how to handle the error?
			if _, err = a.persistAPIClient.CreateJobStatus(
				context.Background(),
				&persist.JobStatus{
					JobId: persistJob.Id,
					Type:  pps.JobStatusType_JOB_STATUS_TYPE_ERROR,
				},
			); err != nil {
				protolog.Errorln(err.Error())
			}
		} else {
			// TODO(pedge): how to handle the error?
			if _, err = a.persistAPIClient.CreateJobStatus(
				context.Background(),
				&persist.JobStatus{
					JobId: persistJob.Id,
					Type:  pps.JobStatusType_JOB_STATUS_TYPE_SUCCESS,
				},
			); err != nil {
				protolog.Errorln(err.Error())
			}
		}
	}()
	return nil
}

func (a *apiServer) runJob(persistJob *persist.Job) error {
	switch {
	case persistJob.GetTransform() != nil:
		return a.reallyRunJob(persistJob.GetTransform(), persistJob.JobInput, persistJob.JobOutput)
	case persistJob.GetPipelineId() != "":
		persistPipeline, err := a.persistAPIClient.GetPipelineByID(
			context.Background(),
			&google_protobuf.StringValue{Value: persistJob.GetPipelineId()},
		)
		if err != nil {
			return err
		}
		if persistPipeline.Transform == nil {
			return fmt.Errorf("pachyderm.pps.server: transform not set on pipeline %v", persistPipeline)
		}
		return a.reallyRunJob(persistPipeline.Transform, persistJob.JobInput, persistJob.JobOutput)
	default:
		return fmt.Errorf("pachyderm.pps.server: neither transform or pipeline id set on job %v", persistJob)
	}
}

func (a *apiServer) reallyRunJob(
	transform *pps.Transform,
	jobInputs []*pps.JobInput,
	jobOutputs []*pps.JobOutput,
) error {
	return nil
}
