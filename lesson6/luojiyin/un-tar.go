package main

import (
	"fmt"
	"archive/tar"
	"compress/gzip"
	"io"
	"path/filepath"
)

func  tar(src string, writers...io.Writer) error {
	if _, err := os.Stat(src); err != nil {
		return fmt.Errorf("nable to tar files -%v", err.Error()
		}
		mw := io.MultiWriter(writers...)

		gzw := gzip.NewWriter(mw)
		defer gzw.Close()

		tw := tar.NewWriter(gzw)
		defer  tw.Close()

		return filepath.Walk(src, func(file string, info os.FileInfo, err error)error{
			if err != nil {
				return err
			}





} 
