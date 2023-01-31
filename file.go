package files

import (
	"github.com/ipfs/go-libipfs/files"
)

var ErrNotDirectory = files.ErrNotDirectory
var ErrNotReader = files.ErrNotReader
var ErrNotSupported = files.ErrNotSupported

type Node = files.Node
type File = files.File
type DirEntry = files.DirEntry
type DirIterator = files.DirIterator
type Directory = files.Directory
type FileInfo = files.FileInfo
