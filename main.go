package main

import (
	"flag"
	"fmt"
	"github.com/raspi/megafile/lib"
	"os"
	"path"
)

var (
	// These are set with Makefile -X=main.VERSION, etc
	VERSION   = `v0.0.0`
	BUILD     = `dev`
	BUILDDATE = `0000-00-00T00:00:00+00:00`
)

func main() {

	// Command line arguments
	extractionDirectory := flag.String(`d`, `extracted`, `Extraction directory`)
	extractFilesArg := flag.Bool(`x`, false, `Extract file(s) from archive`)
	showVersionArg := flag.Bool(`V`, false, `Show version`)

	flag.Usage = func() {
		f := os.Args[0]
		_, _ = fmt.Fprintf(os.Stdout, `Extract MegaFile (.MEG) file(s)`+"\n")
		_, _ = fmt.Fprintf(os.Stdout, "\n")

		_, _ = fmt.Fprintf(os.Stdout, "Parameters:\n")
		flag.VisitAll(func(f *flag.Flag) {
			_, _ = fmt.Fprintf(os.Stdout, "  -%s   %s   default: %q\n", f.Name, f.Usage, f.DefValue)
		})

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
		_, _ = fmt.Fprintf(os.Stdout, `See --help`)
		os.Exit(1)
	}

	if *showVersionArg {
		// Show version information
		_, _ = fmt.Fprintf(os.Stdout, `Version %s build %s built on %s`+"\n", VERSION, BUILD, BUILDDATE)
		os.Exit(0)
	}

	// Read file
	mf := lib.New(flag.Args()[flag.NArg()-1])

	files := mf.ListFiles()
	fcount := len(files)
	for i, f := range files {
		fmt.Printf(`File #%03d/%03d (idx:%03d): %s %d bytes at offset %d`+"\n", i+1, fcount, f.Index+1, path.Join(f.Path, f.Name), f.Length, f.StartOffset)

		if *extractFilesArg {
			savedAs, err := mf.Extract(f, *extractionDirectory)
			if err != nil {
				panic(err)
			}

			fmt.Printf(`Saved as %q`+"\n", savedAs)
		}
	}

	mf.Close()
}
