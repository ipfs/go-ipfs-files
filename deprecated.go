// Package files: This package and has been deprecated and its contents moved to github.com/ipfs/go-libipfs/files
//
// All content in this package is a thin wrapper around the functionality in the new package location.
package files

import (
	"github.com/ipfs/go-libipfs/files"
)

// Errors
var (
	ErrNotDirectory          = files.ErrNotDirectory
	ErrNotReader             = files.ErrNotReader
	ErrNotSupported          = files.ErrNotSupported
	ErrInvalidDirectoryEntry = files.ErrInvalidDirectoryEntry
	ErrPathExistsOverwrite   = files.ErrPathExistsOverwrite
	ErrUnixFSPathOutsideRoot = files.ErrUnixFSPathOutsideRoot
)

// Interfaces
type (
	Node        = files.Node
	File        = files.File
	DirEntry    = files.DirEntry
	DirIterator = files.DirIterator
	Directory   = files.Directory
	FileInfo    = files.FileInfo
)

// Structs
type (
	Filter          = files.Filter
	Symlink         = files.Symlink
	MultiFileReader = files.MultiFileReader
	ReaderFile      = files.ReaderFile
	SliceFile       = files.SliceFile
	TarWriter       = files.TarWriter
	WebFile         = files.WebFile
)

// Helpers
var (
	WriteTo                 = files.WriteTo
	NewFilter               = files.NewFilter
	NewLinkFile             = files.NewLinkFile
	ToSymlink               = files.ToSymlink
	NewMultiFileReader      = files.NewMultiFileReader
	NewFileFromPartReader   = files.NewFileFromPartReader
	NewBytesFile            = files.NewBytesFile
	NewReaderFile           = files.NewReaderFile
	NewReaderStatFile       = files.NewReaderStatFile
	NewReaderPathFile       = files.NewReaderPathFile
	NewSerialFile           = files.NewSerialFile
	NewSerialFileWithFilter = files.NewSerialFileWithFilter
	FileEntry               = files.FileEntry
	NewMapDirectory         = files.NewMapDirectory
	NewSliceDirectory       = files.NewSliceDirectory
	NewTarWriter            = files.NewTarWriter
	ToFile                  = files.ToFile
	ToDir                   = files.ToDir
	FileFromEntry           = files.FileFromEntry
	DirFromEntry            = files.DirFromEntry
	Walk                    = files.Walk
	NewWebFile              = files.NewWebFile
)
