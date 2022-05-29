package files

import (
	"errors"
	"io"
	"os"

	"github.com/ipfs/go-cid"
)

var (
	ErrNotDirectory = errors.New("file isn't a directory")
	ErrNotReader    = errors.New("file isn't a regular file")

	ErrNotSupported = errors.New("operation not supported")
)

// Node is a common interface for files, directories and other special files
type Node interface {
	io.Closer

	// Size returns size of this file (if this file is a directory, total size of
	// all files stored in the tree should be returned). Some implementations may
	// choose not to implement this
	Size() (int64, error)
}

// Node represents a regular Unix file
type File interface {
	Node

	io.Reader
	io.Seeker
}

// DirEntry exposes information about a directory entry
type DirEntry interface {
	// Name returns base name of this entry, which is the base name of referenced
	// file
	Name() string

	// Node returns the file referenced by this DirEntry
	Node() Node
}

// FileType is an enum of possible UnixFS file types.
type FileType int32

const (
	// TUnknown means the file type isn't known (e.g., it hasn't been
	// resolved).
	TUnknown FileType = iota
	// TFile is a regular file.
	TFile
	// TDirectory is a directory.
	TDirectory
	// TSymlink is a symlink.
	TSymlink
	// TDirEnd indicates the implict end of a directory listing.
	TDirEnd
)

func (t FileType) String() string {
	switch t {
	case TUnknown:
		return "unknown"
	case TFile:
		return "file"
	case TDirectory:
		return "directory"
	case TSymlink:
		return "symlink"
	default:
		return "<unknown file type>"
	}
}

// An ExtendedDirectoryEntry exposes unixfs semenatics for a directory entry,
// abstracting the concrete semantics in interface-go-ipfs-core.
type ExtendedDirEntry interface {
	DirEntry
	Cid() cid.Cid

	Size() (uint64, error)   // The size of the file in bytes (or the size of the symlink).
	Type() (FileType, error) // The type of the file.
	Target() (string, error) // The symlink target (if a symlink).
}

// DirIterator is a iterator over directory entries.
// See Directory.Entries for more
type DirIterator interface {
	// DirEntry holds information about current directory entry.
	// Note that after creating new iterator you MUST call Next() at least once
	// before accessing these methods. Calling these methods without prior calls
	// to Next() and after Next() returned false may result in undefined behavior
	DirEntry

	// Next advances iterator to the next file.
	Next() bool

	// Err may return an error after previous call to Next() returned `false`.
	// If previous call to Next() returned `true`, Err() is guaranteed to
	// return nil
	Err() error
}

// Directory is a special file which can link to any number of files.
type Directory interface {
	Node

	// Entries returns a stateful iterator over directory entries.
	//
	// Example usage:
	//
	// it := dir.Entries()
	// for it.Next() {
	//   name := it.Name()
	//   file := it.Node()
	//   [...]
	// }
	// if it.Err() != nil {
	//   return err
	// }
	//
	// Note that you can't store the result of it.Node() and use it after
	// advancing the iterator
	Entries() DirIterator
}

// FileInfo exposes information on files in local filesystem
type FileInfo interface {
	Node

	// AbsPath returns full real file path.
	AbsPath() string

	// Stat returns os.Stat of this file, may be nil for some files
	Stat() os.FileInfo
}
