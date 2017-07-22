package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Untar(path string, r io.Reader) error {
	uncompress, err := gzip.NewReader(r) //读取.gz(gzip)文件
	if err != nil {
		return err
	}
	tr := tar.NewReader(uncompress) ////读取.tar文件
	for {
		hdr, err := tr.Next() //获取头部信息
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fullpath := filepath.Join(path, hdr.Name) //拼接路径 (解压的目录+ header的路径)
		info := hdr.FileInfo()                    //查看文件详细信息
		//判断是否目录
		if info.IsDir() {
			os.MkdirAll(fullpath, 0755)
			continue
		}
		dir := filepath.Dir(fullpath) //获取文件父目录
		os.MkdirAll(dir, 755)

		f, err := os.Create(fullpath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, tr) //将tar的文件 copy到f
		if err != nil {
			f.Close()
			return err
		}
		f.Chmod(info.Mode()) //附权
		f.Close()

	}

}

func main() {
	Untar(os.Args[1], os.Stdin)
}
