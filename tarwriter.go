package files

import (
	"archive/tar"
	"fmt"
	"io"
	"path"
	"time"
)

type TarWriter struct {
	TarW   *tar.Writer
	format tar.Format
}

// NewTarWriter wraps given io.Writer into a new tar writer
func NewTarWriter(w io.Writer) (*TarWriter, error) {
	return &TarWriter{
		TarW:   tar.NewWriter(w),
		format: tar.FormatUnknown,
	}, nil
}

func (w *TarWriter) writeDir(f Directory, fpath string) error {
	if err := w.writeHeader(f, fpath, 0); err != nil {
		return err
	}

	it := f.Entries()
	for it.Next() {
		if err := w.WriteNode(it.Node(), path.Join(fpath, it.Name())); err != nil {
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

	if err = w.writeHeader(f, fpath, size); err != nil {
		return err
	}

	if _, err := io.Copy(w.TarW, f); err != nil {
		return err
	}
	w.TarW.Flush()
	return nil
}

// WriteNode adds a node to the archive.
func (w *TarWriter) WriteNode(nd Node, fpath string) error {
	switch nd := nd.(type) {
	case *Symlink:
		return w.writeHeader(nd, fpath, 0)
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

func (w *TarWriter) writeHeader(n Node, fpath string, size int64) error {
	hdr := &tar.Header{
		Format: w.format,
		Name:   fpath,
		Size:   size,
		Mode:   int64(UnixPermsOrDefault(n)),
	}

	switch nd := n.(type) {
	case *Symlink:
		hdr.Typeflag = tar.TypeSymlink
		hdr.Linkname = nd.Target
	case Directory:
		hdr.Typeflag = tar.TypeDir
	default:
		hdr.Typeflag = tar.TypeReg
	}

	if m := n.ModTime(); m.IsZero() {
		hdr.ModTime = time.Now()
	} else {
		hdr.ModTime = m
	}

	return w.TarW.WriteHeader(hdr)
}

func (w *TarWriter) SetFormat(format tar.Format) {
	w.format = format
}
