package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("输入的参数不正确，(for example: ./mytar test.tar.gz test.go ...)")
		return
	}
	tf, err := os.Create(os.Args[1]) //获取tar.gz包名称
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()
	gw := gzip.NewWriter(tf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	files := os.Args[2:] //获取将进行打包的文件
	for _, file := range files {
		fdr, err := os.Stat(file) //获取文件头信息
		if err != nil {
			fmt.Println(err)
			continue
		}
		hdr, err := tar.FileInfoHeader(fdr, "") //将文件头信息格式化成tar包文件头信息
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = tw.WriteHeader(hdr) //写入文件头信息
		if err != nil {
			fmt.Println(err)
			continue
		}
		f, err := os.Open(file) //打开文件
		if err != nil {
			fmt.Println(err)
			continue
		}
		m, err := io.Copy(tw, f) //将文件信息写入tar包
		if err != nil {
			fmt.Println(err)
			continue
		}
		f.Close()
		fmt.Println(m)
	}
}
