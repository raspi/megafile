package main

import (
	"flag"
	"fmt"
	"github.com/raspi/megafile/lib"
	"os"
	"path"
)

func main() {

	// Command line arguments
	extractFilesArg := flag.Bool(`x`, false, `Extract file(s) from archive`)

	flag.Usage = func() {
		f := os.Args[0]
		_, _ = fmt.Fprintf(os.Stdout, `Extract MegaFile (.MEG) file(s)`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, "\n")
		_, _ = fmt.Fprintf(os.Stdout, `List file(s):`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s <file.meg>`+"\n", f)
		_, _ = fmt.Fprintf(os.Stdout, `Extract file(s):`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s -x <file.meg>`+"\n", f)
		_, _ = fmt.Fprintf(os.Stdout, `Example:`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, `  %s -x MUSIC.MEG`+"\n", f)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		_, _ = fmt.Fprintf(os.Stderr, `see --help`)
		os.Exit(1)
	}

	// Read file
	mf := lib.New(flag.Args()[flag.NArg()-1])

	files := mf.ListFiles()
	fcount := len(files)
	for i, f := range files {
		fmt.Printf(`file #%03d(idx:%d)/%03d: %s %d bytes at offset %d`+"\n", i+1, f.Index+1, fcount, path.Join(f.Path, f.Name), f.Length, f.StartOffset)

		if *extractFilesArg {
			err := mf.Extract(f)
			if err != nil {
				panic(err)
			}
		}
	}

	mf.Close()
}
