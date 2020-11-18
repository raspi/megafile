// MegaFile (.MEG) reader
package lib

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const MagicFlag = 0xFFFFFFFF

// File header
type header struct {
	Flags               uint32 // 0xFFFF
	Size1               uint32
	Size2               uint32
	FileCount1          uint32 // How many files there are
	FileCount2          uint32 // How many files there are
	FilenamesDataLength uint32 // Length for file count information table
}

// File record
type filerec struct {
	Flags          uint16
	CRC32          uint32
	Index          uint32
	Length         uint32 // Data size
	StartOffset    uint32 // Starting offset inside archive
	NameTableIndex uint16 // File name in the file index table
}

// File representation for humans
type File struct {
	Path        string
	Name        string
	StartOffset int
	Length      int
	Index       int
}

type IOReaderSeekerCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

type MegaFileReader struct {
	r      IOReaderSeekerCloser
	header header
}

func New(path string) MegaFileReader {
	mf := MegaFileReader{}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	mf.r = f

	mf.readHeader()

	return mf
}

// Read file header
func (mf *MegaFileReader) readHeader() {
	err := binary.Read(mf.r, binary.LittleEndian, &mf.header)
	if err != nil {
		panic(err)
	}

	if mf.header.Flags != MagicFlag {
		panic(fmt.Errorf(`Invalid flags: %[1]x %[1]b`, mf.header.Flags))
	}
}

// Get list of files
func (mf *MegaFileReader) ListFiles() (files []File) {
	// Read list of filenames from filename table
	var filenames []string

	for i := 0; i < int(mf.header.FileCount1); i++ {
		var filenamelength uint16 // How many characters in filename?
		err := binary.Read(mf.r, binary.LittleEndian, &filenamelength)
		if err != nil {
			panic(err)
		}

		// Read the filename
		fnamebuffer := make([]byte, filenamelength)
		readBytes, err := mf.r.Read(fnamebuffer)
		if readBytes != int(filenamelength) {
			panic(`invalid size`)
		}

		// Add to filename list
		filenames = append(filenames, strings.ToLower(string(fnamebuffer)))
	}

	// Read file metainformation (size, etc, ..)
	for i := 0; i < len(filenames); i++ {
		var frec filerec
		err := binary.Read(mf.r, binary.LittleEndian, &frec)
		if err != nil {
			panic(err)
		}

		// Assign the proper filename from previously read filename table
		fidx := int(frec.NameTableIndex)
		fpath := filenames[fidx]
		dname, fname := path.Split(strings.ReplaceAll(fpath, `\`, `/`))

		// Add to file list
		files = append(files, File{
			Path:        dname,
			Name:        fname,
			StartOffset: int(frec.StartOffset),
			Length:      int(frec.Length),
			Index:       fidx,
		})
	}

	return files
}

func (mf *MegaFileReader) Close() {
	mf.r.Close()
}

// Extract a file
func (mf *MegaFileReader) Extract(src File) error {
	// Seek to position where data begins
	offset, err := mf.r.Seek(int64(src.StartOffset), io.SeekStart)
	if err != nil {
		return err
	}

	if offset != int64(src.StartOffset) {
		return fmt.Errorf(`At wrong offset %d, expected %d`, offset, src.StartOffset)
	}

	// Create a buffer for data
	buffer := make([]byte, src.Length)

	_, err = mf.r.Read(buffer)
	if err != nil {
		return err
	}

	// Write to temporary file
	f, err := ioutil.TempFile(`.`, `data-`)
	if err != nil {
		return err
	}

	// Write contents to file
	_, err = f.Write(buffer)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	// Rename temporary file to proper filename
	err = os.Rename(f.Name(), src.Name)
	if err != nil {
		return err
	}

	return nil
}
