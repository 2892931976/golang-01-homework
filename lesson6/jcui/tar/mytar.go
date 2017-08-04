package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func tarFile(src, srcs string, tw *tar.Writer, fi os.FileInfo) error {
	full_sfile := src + srcs //获取完整路径
	fmt.Println(full_sfile)
	hdr, err := tar.FileInfoHeader(fi, "") //写入文件头信息
	if err != nil {
		err := fmt.Errorf("%s, %s", fi, "tar file header is error ")
		return err
	}
	hdr.Name = srcs
	err = tw.WriteHeader(hdr)
	if err != nil {
		err := fmt.Errorf("%s, %s", "tar file header write faild", tw)
		return err
	}

	sf, err := os.Open(full_sfile)
	if err != nil {
		err := fmt.Errorf("%s, %s", "open tar src file faild", fi)
		return err
	}
	defer sf.Close()
	_, err = io.Copy(tw, sf)
	if err != nil {
		err := fmt.Errorf("%s, %s", "tar file faild end")
		return err
	}

	return nil
}

func tarDir(src, srcs string, tw *tar.Writer, fi os.FileInfo) error {
	full_sfile := src + srcs
	fmt.Println(full_sfile)
	/*
		得遍历目录得到所有信息,如果是目录继续遍历,如果是文件则执行tarFile
		下面写的貌似不对的了,有错误的逻辑,不知道咋写了
	*/
	last := len(srcs) - 1
	if srcs[last] != os.PathSeparator {
		srcs += string(os.PathSeparator)
	}

	fis, er := ioutil.ReadDir(full_sfile)
	if er != nil {
		return er
	}
	for _, fi := range fis {
		if fi.IsDir() {
			err := tarDir(src, srcs+fi.Name(), tw, fi)
			if err != nil {
				err := fmt.Errorf("%s, %s", "tar file faild end")
				return err
			}
		} else {
			err := tarFile(src, srcs+fi.Name(), tw, fi)
			if err != nil {
				err := fmt.Errorf("%s, %s", "tar file faild end")
				return err
			}
		}
	}
	if len(srcs) > 0 {
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		hdr.Name = srcs
	}
	return nil
}

func tarfile(s []string) error {
	dfile := s[2] //正常逻辑应当判断打包文件是否存在,提示覆盖或者重命名
	sfile := s[3]
	fi, err := os.Stat(sfile) //判断目标文件是否存在,这一步必须要做到
	//对源文件或目录做处理
	if err != nil {
		err := fmt.Errorf("%s, %s", sfile, "is not Exist ")
		return err
	}
	//hdr, err := tar.FileInfoHeader(fi, "")
	//if err != nil {
	//	err := fmt.Errorf("%s, %s", sfile,"tar file header is error ")
	//	return err
	//}

	//对目标文件或目录做处理
	df, err := os.Create(dfile)
	if err != nil {
		err := fmt.Errorf("%s, %s", "tar file create faild", df)
		return err
	}
	gf := gzip.NewWriter(df)
	tw := tar.NewWriter(gf)

	defer gf.Close()
	defer tw.Close()
	src, srcs := path.Split(path.Clean(sfile))
	if fi.IsDir() {
		tarDir(src, srcs, tw, fi)
	} else {
		tarFile(src, srcs, tw, fi)
	}
	/*
		如下这段内容只能实现对单个文件打包,碰到目录,或多个文件就会报错
	*/
	//将tar.header写入到.tar文件中
	//err = tr.WriteHeader(hdr)
	//if err != nil {
	//	err := fmt.Errorf("%s, %s", "tar file header write faild", df)
	//	return err
	//}

	//sf, err := os.Open(sfile)
	//if err != nil {
	//	err := fmt.Errorf("%s, %s", "open tar src file faild", sfile)
	//	return err
	//}
	//defer sf.Close()
	//_, err = io.Copy(tr, sf)
	//if err != nil {
	//	err := fmt.Errorf("%s, %s", "tar file faild end")
	//	return err
	//}
	return nil

}

func untarfile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		err := fmt.Errorf("%s, %s", file, "open tar file faild")
		return err
	}
	g, err := gzip.NewReader(f)
	if err != nil {
		err := fmt.Errorf("%s, %s", file, "open tar file faild")
		return err
	}
	t := tar.NewReader(g)

	for {
		hdr, err := t.Next() //获取文件头信息
		if err != nil {
			return nil
		}
		fmt.Println(hdr.Name)
		info := hdr.FileInfo()
		//io.Copy(ioutil.Discard, tr) //将内容丢弃,等同于>/dev/null
		if info.IsDir() {
			os.Mkdir(hdr.Name, 0755)
			continue
		}
		f, err := os.Create(hdr.Name)
		if err != nil {
			err := fmt.Errorf("%s, %s", err, "files is error")
			return err
		}
		io.Copy(f, t)
		os.Chmod(hdr.Name, info.Mode())
		f.Close()
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("args is error.")
		return
	}
	switch os.Args[1] {
	case "-cz":
		if len(os.Args) != 4 {
			fmt.Println("args is error. Example : ./mytar -cz aa.tar.gz aa")
			return
		}
		err := tarfile(os.Args)
		if err != nil {
			log.Fatal(err)
		}
	case "-xz":
		if len(os.Args) != 3 {
			fmt.Println("args is error. Example : ./mytar -xz aa.tar.gz")
			return
		}
		err := untarfile(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println(`Args is Error.
For Example :
	tar   --  ./mytar -cz aa.tar.gz aa
	untar --  ./mytar -xz aa.tar.gz   `)
		return
	}
}
