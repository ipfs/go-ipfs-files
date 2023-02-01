// Package files: This package and has been deprecated and its contents moved to github.com/ipfs/go-libipfs/files
//
// All content in this package is a thin wrapper around the functionality in the new package location.
package files

import (
	"github.com/ipfs/go-libipfs/files"
)

// Errors
var (
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrNotDirectory = files.ErrNotDirectory
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrNotReader = files.ErrNotReader
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrNotSupported = files.ErrNotSupported
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrInvalidDirectoryEntry = files.ErrInvalidDirectoryEntry
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrPathExistsOverwrite = files.ErrPathExistsOverwrite
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ErrUnixFSPathOutsideRoot = files.ErrUnixFSPathOutsideRoot
)

// Interfaces
type (
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	Node = files.Node
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	File = files.File
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	DirEntry = files.DirEntry
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	DirIterator = files.DirIterator
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	Directory = files.Directory
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	FileInfo = files.FileInfo
)

// Structs
type (
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	Filter = files.Filter
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	Symlink = files.Symlink
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	MultiFileReader = files.MultiFileReader
	ReaderFile      = files.ReaderFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	SliceFile = files.SliceFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	TarWriter = files.TarWriter
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	WebFile = files.WebFile
)

// Helpers
var (
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	WriteTo = files.WriteTo
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewFilter = files.NewFilter
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewLinkFile = files.NewLinkFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ToSymlink = files.ToSymlink
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewMultiFileReader = files.NewMultiFileReader
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewFileFromPartReader = files.NewFileFromPartReader
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewBytesFile = files.NewBytesFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewReaderFile = files.NewReaderFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewReaderStatFile = files.NewReaderStatFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewReaderPathFile = files.NewReaderPathFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewSerialFile = files.NewSerialFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewSerialFileWithFilter = files.NewSerialFileWithFilter
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	FileEntry = files.FileEntry
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewMapDirectory = files.NewMapDirectory
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewSliceDirectory = files.NewSliceDirectory
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewTarWriter = files.NewTarWriter
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ToFile = files.ToFile
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	ToDir = files.ToDir
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	FileFromEntry = files.FileFromEntry
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	DirFromEntry = files.DirFromEntry
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	Walk = files.Walk
	// Deprecated: moved to github.com/ipfs/go-libipfs/files
	NewWebFile = files.NewWebFile
)
