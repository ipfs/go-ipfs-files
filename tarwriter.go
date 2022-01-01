package files

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

type TarWriter struct {
	TarW *tar.Writer
}

// NewTarWriter wraps given io.Writer into a new tar writer
func NewTarWriter(w io.Writer) (*TarWriter, error) {
	return &TarWriter{
		TarW: tar.NewWriter(w),
	}, nil
}

func (w *TarWriter) writeDir(f Directory, fpath string) error {
	if err := writeDirHeader(w.TarW, fpath); err != nil {
		return err
	}

	it := f.Entries()
	for it.Next() {
		if err := w.WriteFile(it.Node(), path.Join(fpath, it.Name())); err != nil {
			return err
		}
	}
	return it.Err()
}

func (w *TarWriter) writeFile(f File, fpath string) error {
	size, err := f.Size()
	if err != nil {
		return err
	}

	mode, mtime, err := getFileMeta(f)
	if err != nil {
		return err
	}

	err = w.TarW.WriteHeader(&tar.Header{
		Name:     fpath,
		Size:     size,
		Typeflag: tar.TypeReg,
		Mode:     int64(mode),
		ModTime:  mtime,
	})

	if err != nil {
		return err
	}

	if _, err := io.Copy(w.TarW, f); err != nil {
		return err
	}
	w.TarW.Flush()
	return nil
}

// WriteNode adds a node to the archive.
func (w *TarWriter) WriteFile(nd Node, fpath string) error {
	switch nd := nd.(type) {
	case *Symlink:
		return writeSymlinkHeader(w.TarW, nd.Target, fpath)
	case File:
		return w.writeFile(nd, fpath)
	case Directory:
		return w.writeDir(nd, fpath)
	default:
		return fmt.Errorf("file type %T is not supported", nd)
	}
}

// Close closes the tar writer.
func (w *TarWriter) Close() error {
	return w.TarW.Close()
}

func writeDirHeader(w *tar.Writer, fpath string) error {
	return w.WriteHeader(&tar.Header{
		Name:     fpath,
		Typeflag: tar.TypeDir,
		Mode:     0777,
		ModTime:  time.Now().Truncate(time.Second),
		// TODO: set mode, dates, etc. when added to unixFS
	})
}

func writeSymlinkHeader(w *tar.Writer, target, fpath string) error {
	return w.WriteHeader(&tar.Header{
		Name:     fpath,
		Linkname: target,
		Mode:     0777,
		Typeflag: tar.TypeSymlink,
	})
}

func getFileMeta(f File) (os.FileMode, time.Time, error) {
	var err error = nil
	mode := f.Mode()
	mtime := f.ModTime()

	if mode == 0 {
		switch nd := f.(type) {
		case File:
			mode = 0644
		case Directory:
			mode = 0777
		default:
			err = fmt.Errorf("mode for file type %T is not supported", nd)
		}
	}

	if !mtime.IsZero() {
		mtime = time.Now()
	}

	return mode, mtime, err
}
