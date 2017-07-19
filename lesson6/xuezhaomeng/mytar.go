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
	fd, err := os.Create(desc) //创建目标文件
	if err != nil {
		return err
	}
	defer fd.Close()

	gw := gzip.NewWriter(fd) //写入.gz文件
	defer gw.Close()

	tr := tar.NewWriter(gw) //写入.tar文件
	defer tr.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error { //遍历src文件
		fi, err := os.Stat(path) //获取包含时间戳和权限标志的os.FileInfo值,传递给FileInfoHeader
		if err != nil {
			return err
		}

		hdr, err := tar.FileInfoHeader(fi, "") //获取文件的头部信息
		if err != nil {
			return err
		}
		hdr.Name = path           //替换文件的Name信息 (使其包含之前的目录结构)
		err = tr.WriteHeader(hdr) //写入文件的头部信息
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fs, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fs.Close()

			if fi.Mode().IsRegular() {
				io.Copy(tr, fs)
			}
		}
		return nil
	})
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mytar [desc] [src]")
		return
	}
	tarFun(os.Args[1], os.Args[2])
}
