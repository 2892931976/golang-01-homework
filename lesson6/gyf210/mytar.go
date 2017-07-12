package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
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

	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	hdr, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return err
	}

	err = tr.WriteHeader(hdr)
	if err != nil {
		return err
	}

	fs, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fs.Close()

	_, err = io.Copy(tr, fs)
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
