package files

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSliceFiles(t *testing.T) {
	sf := NewMapDirectory(map[string]Node{
		"1": NewBytesFile([]byte("Some text!\n")),
		"2": NewBytesFile([]byte("beep")),
		"3": NewBytesFile([]byte("boop")),
	})

	CheckDir(t, sf, []Event{
		{
			kind:  TFile,
			name:  "1",
			value: "Some text!\n",
		},
		{
			kind:  TFile,
			name:  "2",
			value: "beep",
		},
		{
			kind:  TFile,
			name:  "3",
			value: "boop",
		},
	})
}

func TestReaderFiles(t *testing.T) {
	message := "beep boop"
	rf := NewBytesFile([]byte(message))
	buf := make([]byte, len(message))

	if n, err := rf.Read(buf); n == 0 || err != nil {
		t.Fatal("Expected to be able to read")
	}
	if err := rf.Close(); err != nil {
		t.Fatal("Should be able to close")
	}
	if n, err := rf.Read(buf); n != 0 || err != io.EOF {
		t.Fatal("Expected EOF when reading after close")
	}
}

func TestReaderFileStat(t *testing.T) {
	reader := strings.NewReader("beep boop")
	mode := os.FileMode(0754)
	mtime := time.Date(2020, 11, 2, 12, 27, 35, 55555, time.UTC)
	stat := &mockFileInfo{name: "test", mode: mode, mtime: mtime}

	rf := NewReaderStatFile(reader, stat)
	if rf.Mode() != mode {
		t.Fatalf("Expected file mode to be [%v] but got [%v]", mode, rf.Mode())
	}
	if rf.ModTime() != mtime {
		t.Fatalf("Expected file modified time to be [%v] but got [%v]", mtime, rf.ModTime())
	}
}

func TestMultipartFiles(t *testing.T) {
	data := `
--Boundary!
Content-Type: text/plain
Content-Disposition: form-data; name="file-0?mode=0754&mtime=1604320500&mtime-nsecs=55555;ans=42"; filename="name"
Some-Header: beep

beep
--Boundary!
Content-Type: application/x-directory
Content-Disposition: file; filename="dir"

--Boundary!
Content-Type: text/plain
Content-Disposition: file; filename="dir/nested"

some content
--Boundary!
Content-Type: application/symlink
Content-Disposition: file; filename="dir/simlynk"

anotherfile
--Boundary!
Content-Type: text/plain
Content-Disposition: file; filename="implicit1/implicit2/deep_implicit"

implicit file1
--Boundary!
Content-Type: text/plain
Content-Disposition: file; filename="implicit1/shallow_implicit"

implicit file2
--Boundary!--

`

	reader := strings.NewReader(data)
	mpReader := multipart.NewReader(reader, "Boundary!")
	dir, err := NewFileFromPartReader(mpReader, multipartFormdataType)
	if err != nil {
		t.Fatal(err)
	}

	CheckDir(t, dir, []Event{
		{
			kind:  TFile,
			name:  "name",
			value: "beep",
		},
		{
			kind: TDirStart,
			name: "dir",
		},
		{
			kind:  TFile,
			name:  "nested",
			value: "some content",
		},
		{
			kind:  TSymlink,
			name:  "simlynk",
			value: "anotherfile",
		},
		{
			kind: TDirEnd,
		},
		{
			kind: TDirStart,
			name: "implicit1",
		},
		{
			kind: TDirStart,
			name: "implicit2",
		},
		{
			kind:  TFile,
			name:  "deep_implicit",
			value: "implicit file1",
		},
		{
			kind: TDirEnd,
		},
		{
			kind:  TFile,
			name:  "shallow_implicit",
			value: "implicit file2",
		},
		{
			kind: TDirEnd,
		},
	})
}
