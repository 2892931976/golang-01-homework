package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	sourcefile := flag.Arg(0)

	if sourcefile == "" {
		fmt.Println("Usage : need a sourcefile.tar ")
		os.Exit(1)
	}
	file, err := os.Open(sourcefile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var fileReader io.ReadCloser = file

	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}
	tarBallReader := tar.NewReader(fileReader)

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Println("Creating directory:", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		case tar.TypeReg:
			fmt.Println("Untarring:", filename)
			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}
			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}
			writer.Close()
		default:
			fmt.Println("Unable to untar type: %c in file %s ", header.Typeflag, filename)
		}
	}
}
