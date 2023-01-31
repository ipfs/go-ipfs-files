package files

import (
	"github.com/ipfs/go-libipfs/files"
)

var ErrUnixFSPathOutsideRoot = files.ErrUnixFSPathOutsideRoot

type TarWriter = files.TarWriter

var NewTarWriter = files.NewTarWriter
