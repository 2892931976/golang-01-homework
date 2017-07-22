package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func tarFun(desc, src string) error {
	fd, err := os.Create(desc)
	if err != nil {
		return err
	}
	defer fd.Close()

	gw := gzip.NewWriter(fd)
	defer gw.Close()

	tr := tar.NewWriter(gw)
	defer tr.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = path
		err = tr.WriteHeader(hdr)
		if err != nil {
			return err
		}
		fs, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fs.Close()
		if info.Mode().IsRegular() {
			io.Copy(tr, fs)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mytar [desc] [src]")
		return
	}
	tarFun(os.Args[1], os.Args[2])
}
