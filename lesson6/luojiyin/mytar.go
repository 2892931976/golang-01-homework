package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func checkerror(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	destFile := flag.Arg(0)

	if destFile == "" {
		fmt.Println("Usage : destFile.tar.gz source")
		os.Exit(1)
	}

	sourcedir := flag.Arg(1)

	if sourcedir == "" {
		fmt.Println("Usage : gotar destFile.tar.gz source dir")
		os.Exit(1)
	}

	dir, err := os.Open(sourcedir)
	checkerror(err)

	defer dir.Close()

	files, err := dir.Readdir(0)
	checkerror(err)
	tarfile, err := os.Create(destFile)

	checkerror(err)

	defer tarfile.Close()
	var fileWriter io.WriteCloser = tarfile

	if strings.HasSuffix(destFile, ".gz") {
		fileWriter = gzip.NewWriter(tarfile)
		defer fileWriter.Close()
	}

	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
		checkerror(err)

		defer file.Close()

		header := new(tar.Header)
		header.Name = file.Name()
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()

		err = tarfileWriter.WriteHeader(header)
		checkerror(err)

		_, err = io.Copy(tarfileWriter, file)

		checkerror(err)

	}
}
