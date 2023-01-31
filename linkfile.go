package files

import (
	"github.com/ipfs/go-libipfs/files"
)

type Symlink = files.Symlink

var NewLinkFile = files.NewLinkFile
var ToSymlink = files.ToSymlink
