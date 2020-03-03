package files

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReaderFile is a implementation of File created from an `io.Reader`.
// ReaderFiles are never directories, and can be read from and closed.
type ReaderFile struct {
	abspath string
	reader  io.ReadCloser
	stat    os.FileInfo

	fsize int64
}

// NewBytesFile ...
func NewBytesFile(b []byte) File {
	return &ReaderFile{"", NewReaderFile(bytes.NewReader(b)), nil, int64(len(b))}
}

// NewReaderFile ...
func NewReaderFile(reader io.Reader) File {
	return NewReaderStatFile(reader, nil)
}

// NewReaderStatFile ...
func NewReaderStatFile(reader io.Reader, stat os.FileInfo) File {
	rc, ok := reader.(io.ReadCloser)
	if !ok {
		rc = ioutil.NopCloser(reader)
	}

	return &ReaderFile{"", rc, stat, -1}
}

// NewReaderPathFile ...
func NewReaderPathFile(path string, reader io.ReadCloser, stat os.FileInfo) (*ReaderFile, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return &ReaderFile{abspath, reader, stat, -1}, nil
}

// AbsPath ...
func (f *ReaderFile) AbsPath() string {
	return f.abspath
}

// Read ...
func (f *ReaderFile) Read(p []byte) (int, error) {
	return f.reader.Read(p)
}

// Close ...
func (f *ReaderFile) Close() error {
	return f.reader.Close()
}

// Stat ...
func (f *ReaderFile) Stat() os.FileInfo {
	return f.stat
}

// Size ...
func (f *ReaderFile) Size() (int64, error) {
	if f.stat == nil {
		if f.fsize >= 0 {
			return f.fsize, nil
		}
		return 0, ErrNotSupported
	}
	return f.stat.Size(), nil
}

// Seek ...
func (f *ReaderFile) Seek(offset int64, whence int) (int64, error) {
	if s, ok := f.reader.(io.Seeker); ok {
		return s.Seek(offset, whence)
	}

	return 0, ErrNotSupported
}

var _ File = &ReaderFile{}
var _ FileInfo = &ReaderFile{}
