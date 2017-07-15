package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Use it like tar command")
		return
	}
	TarFunc(os.Args[1], os.Args[2])
}

func TarFunc(TarFile string, path string) error {
	file, err := os.Create(TarFile)
	if err != nil {
		return err
	}

	defer file.Close()
	gz := gzip.NewWriter(file)
	defer gz.Close()

	tw := tar.NewWriter(gz)
	defer tw.Close()

	if err := tarit(path, tw); err != nil {
		return err
	}

	return nil
}

func tarit(source string, tw *tar.Writer) error {
	info, err := os.Stat(source)
	if err != nil {
		return nil
	}
	var _dir string
	if info.IsDir() {
		_dir = filepath.Base(source)
	}

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var link string
		if info.Mode()&os.ModeSymlink != 0 {
			if link, err = os.Readlink(path); err != nil {
				return err
			}
		}
		header, err := tar.FileInfoHeader(info, link)
		if err != nil {
			return err
		}
		if _dir != "" {
			header.Name = filepath.Join(_dir, strings.TrimPrefix(path, source))
		}

		if !info.Mode().IsRegular() {
			return nil
		}
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		buf := make([]byte, 16)
		if _, err = io.CopyBuffer(tw, file, buf); err != nil {
			return err
		}

		return nil
	})
}
