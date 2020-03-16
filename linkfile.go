package files

import (
	"os"
	"strings"
)

// Symlink ...
type Symlink struct {
	Target string

	stat   os.FileInfo
	reader strings.Reader
}

// NewLinkFile ...
func NewLinkFile(target string, stat os.FileInfo) File {
	lf := &Symlink{Target: target, stat: stat}
	lf.reader.Reset(lf.Target)
	return lf
}

// Close ...
func (lf *Symlink) Close() error {
	return nil
}

// Read ...
func (lf *Symlink) Read(b []byte) (int, error) {
	return lf.reader.Read(b)
}

// Seek ...
func (lf *Symlink) Seek(offset int64, whence int) (int64, error) {
	return lf.reader.Seek(offset, whence)
}

// Size ...
func (lf *Symlink) Size() (int64, error) {
	return lf.reader.Size(), nil
}

// ToSymlink ...
func ToSymlink(n Node) *Symlink {
	l, _ := n.(*Symlink)
	return l
}

var _ File = &Symlink{}
