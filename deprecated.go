// Package files: This package and has been deprecated and its contents moved to github.com/ipfs/boxo/files
//
// All content in this package is a thin wrapper around the functionality in the new package location.
package files

import (
	"github.com/ipfs/boxo/files"
)

// Errors
var (
	// Deprecated: moved to [files.ErrNotDirectory]
	ErrNotDirectory = files.ErrNotDirectory
	// Deprecated: moved to [files.ErrNotReader]
	ErrNotReader = files.ErrNotReader
	// Deprecated: moved to [files.ErrNotSupported]
	ErrNotSupported = files.ErrNotSupported
	// Deprecated: moved to [files.ErrInvalidDirectoryEntry]
	ErrInvalidDirectoryEntry = files.ErrInvalidDirectoryEntry
	// Deprecated: moved to [files.ErrPathExistsOverwrite]
	ErrPathExistsOverwrite = files.ErrPathExistsOverwrite
	// Deprecated: moved to [files.ErrUnixFSPathOutsideRoot]
	ErrUnixFSPathOutsideRoot = files.ErrUnixFSPathOutsideRoot
)

// Interfaces
type (
	// Deprecated: moved to [files.Node]
	Node = files.Node
	// Deprecated: moved to [files.File]
	File = files.File
	// Deprecated: moved to [files.DirEntry]
	DirEntry = files.DirEntry
	// Deprecated: moved to [files.DirIterator]
	DirIterator = files.DirIterator
	// Deprecated: moved to [files.Directory]
	Directory = files.Directory
	// Deprecated: moved to [files.FileInfo]
	FileInfo = files.FileInfo
)

// Structs
type (
	// Deprecated: moved to [files.Filter]
	Filter = files.Filter
	// Deprecated: moved to [files.Symlink]
	Symlink = files.Symlink
	// Deprecated: moved to [files.MultiFileReader]
	MultiFileReader = files.MultiFileReader
	// Deprecated: moved to [files.ReaderFile]
	ReaderFile = files.ReaderFile
	// Deprecated: moved to [files.SliceFile]
	SliceFile = files.SliceFile
	// Deprecated: moved to [files.TarWriter]
	TarWriter = files.TarWriter
	// Deprecated: moved to [files.WebFile]
	WebFile = files.WebFile
)

// Helpers
var (
	// Deprecated: moved to [files.WriteTo]
	WriteTo = files.WriteTo
	// Deprecated: moved to [files.NewFilter]
	NewFilter = files.NewFilter
	// Deprecated: moved to [files.NewLinkFile]
	NewLinkFile = files.NewLinkFile
	// Deprecated: moved to [files.ToSymlink]
	ToSymlink = files.ToSymlink
	// Deprecated: moved to [files.NewMultiFileReader]
	NewMultiFileReader = files.NewMultiFileReader
	// Deprecated: moved to [files.NewFileFromPartReader]
	NewFileFromPartReader = files.NewFileFromPartReader
	// Deprecated: moved to [files.NewBytesFile]
	NewBytesFile = files.NewBytesFile
	// Deprecated: moved to [files.NewReaderFile]
	NewReaderFile = files.NewReaderFile
	// Deprecated: moved to [files.NewReaderStatFile]
	NewReaderStatFile = files.NewReaderStatFile
	// Deprecated: moved to [files.NewReaderPathFile]
	NewReaderPathFile = files.NewReaderPathFile
	// Deprecated: moved to [files.NewSerialFile]
	NewSerialFile = files.NewSerialFile
	// Deprecated: moved to [files.NewSerialFileWithFilter]
	NewSerialFileWithFilter = files.NewSerialFileWithFilter
	// Deprecated: moved to [files.NewSerialFileWithFilter]
	FileEntry = files.FileEntry
	// Deprecated: moved to [files.NewMapDirectory]
	NewMapDirectory = files.NewMapDirectory
	// Deprecated: moved to [files.NewSliceDirectory]
	NewSliceDirectory = files.NewSliceDirectory
	// Deprecated: moved to [files.NewTarWriter]
	NewTarWriter = files.NewTarWriter
	// Deprecated: moved to [files.ToFile]
	ToFile = files.ToFile
	// Deprecated: moved to [files.ToDir]
	ToDir = files.ToDir
	// Deprecated: moved to [files.FileFromEntry]
	FileFromEntry = files.FileFromEntry
	// Deprecated: moved to [files.DirFromEntry]
	DirFromEntry = files.DirFromEntry
	// Deprecated: moved to [files.Walk]
	Walk = files.Walk
	// Deprecated: moved to [files.NewWebFile]
	NewWebFile = files.NewWebFile
)
