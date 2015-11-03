// Code generated by protoc-gen-go.
// source: pfs/fuse/fuse.proto
// DO NOT EDIT!

/*
Package fuse is a generated protocol buffer package.

It is generated from these files:
	pfs/fuse/fuse.proto

It has these top-level messages:
	Filesystem
	Node
	Attr
	Dirent
	Root
	DirectoryAttr
	DirectoryLookup
	DirectoryReadDirAll
	DirectoryCreate
	DirectoryMkdir
	FileAttr
	FileRead
	FileOpen
	FileWrite
*/
package fuse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import protolog "go.pedge.io/protolog"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Filesystem struct {
	Shard   uint64 `protobuf:"varint,1,opt,name=shard" json:"shard,omitempty"`
	Modulus uint64 `protobuf:"varint,2,opt,name=modulus" json:"modulus,omitempty"`
}

func (m *Filesystem) Reset()         { *m = Filesystem{} }
func (m *Filesystem) String() string { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()    {}

type Node struct {
	RepoName string `protobuf:"bytes,1,opt,name=repoName" json:"repoName,omitempty"`
	CommitID string `protobuf:"bytes,2,opt,name=commitID" json:"commitID,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Write    bool   `protobuf:"varint,4,opt,name=write" json:"write,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

type Attr struct {
	Mode uint32 `protobuf:"varint,1,opt,name=Mode" json:"Mode,omitempty"`
}

func (m *Attr) Reset()         { *m = Attr{} }
func (m *Attr) String() string { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()    {}

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode" json:"inode,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Dirent) Reset()         { *m = Dirent{} }
func (m *Dirent) String() string { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()    {}

type Root struct {
	Filesystem *Filesystem `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Result     *Node       `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error      string      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Root) Reset()         { *m = Root{} }
func (m *Root) String() string { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()    {}

func (m *Root) GetFilesystem() *Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

func (m *Root) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryAttr struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryAttr) Reset()         { *m = DirectoryAttr{} }
func (m *DirectoryAttr) String() string { return proto.CompactTextString(m) }
func (*DirectoryAttr) ProtoMessage()    {}

func (m *DirectoryAttr) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryLookup struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Result    *Node  `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Err       string `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *DirectoryLookup) Reset()         { *m = DirectoryLookup{} }
func (m *DirectoryLookup) String() string { return proto.CompactTextString(m) }
func (*DirectoryLookup) ProtoMessage()    {}

func (m *DirectoryLookup) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryLookup) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryReadDirAll struct {
	Directory *Node     `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    []*Dirent `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	Error     string    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryReadDirAll) Reset()         { *m = DirectoryReadDirAll{} }
func (m *DirectoryReadDirAll) String() string { return proto.CompactTextString(m) }
func (*DirectoryReadDirAll) ProtoMessage()    {}

func (m *DirectoryReadDirAll) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryReadDirAll) GetResult() []*Dirent {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryCreate struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryCreate) Reset()         { *m = DirectoryCreate{} }
func (m *DirectoryCreate) String() string { return proto.CompactTextString(m) }
func (*DirectoryCreate) ProtoMessage()    {}

func (m *DirectoryCreate) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryCreate) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryMkdir struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryMkdir) Reset()         { *m = DirectoryMkdir{} }
func (m *DirectoryMkdir) String() string { return proto.CompactTextString(m) }
func (*DirectoryMkdir) ProtoMessage()    {}

func (m *DirectoryMkdir) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryMkdir) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileAttr struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Result *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FileAttr) Reset()         { *m = FileAttr{} }
func (m *FileAttr) String() string { return proto.CompactTextString(m) }
func (*FileAttr) ProtoMessage()    {}

func (m *FileAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileRead struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileRead) Reset()         { *m = FileRead{} }
func (m *FileRead) String() string { return proto.CompactTextString(m) }
func (*FileRead) ProtoMessage()    {}

func (m *FileRead) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileOpen struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileOpen) Reset()         { *m = FileOpen{} }
func (m *FileOpen) String() string { return proto.CompactTextString(m) }
func (*FileOpen) ProtoMessage()    {}

func (m *FileOpen) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileWrite struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileWrite) Reset()         { *m = FileWrite{} }
func (m *FileWrite) String() string { return proto.CompactTextString(m) }
func (*FileWrite) ProtoMessage()    {}

func (m *FileWrite) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}
