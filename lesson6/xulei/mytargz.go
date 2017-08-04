package main

import (
	"archive/tar"

	"fmt"
	"io"

	"os"

	"compress/gzip"
)

func main() {

	if len(os.Args) != 3 {

		fmt.Println("mytargz  *.tar.gz dirname")
	}
	TarFile := os.Args[1]
	Src := os.Args[2]

	TarGzFunc(TarFile, Src)

}

func TarGzFunc(TarFile string, Src string) error {

	zipfile, err := os.Create(TarFile) //创建tar.gz文件
	if err != nil {

		return err
	}

	defer zipfile.Close()
	file := gzip.NewWriter(zipfile)

	defer file.Close()

	tarfile := tar.NewWriter(file)

	dir, err := os.Open(Src)

	if err != nil {
		return err

	}
	defer dir.Close()
	files, err := dir.Readdir(0)

	if err != nil {

		return err
	}

	for _, fil := range files {

		if fil.IsDir() {

			continue
		}
		fmt.Println(fil.Name())

		fi, err := os.Open(dir.Name() + "/" + fil.Name())

		if err != nil {

			return err
		}
		defer fi.Close()

		fileheader := new(tar.Header)
		fileheader.Name = fil.Name()
		fileheader.Size = fil.Size()
		fileheader.Mode = int64(fil.Mode())
		fileheader.ModTime = fil.ModTime()

		err = tarfile.WriteHeader(fileheader)

		if err != nil {

			return err
		}
		_, err = io.Copy(tarfile, fi)

		if err != nil {

			return err
		}
	}
	fmt.Println("tar.gz ok")
	return nil

}
