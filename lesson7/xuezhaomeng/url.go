package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

func CleanUrl(uri *url.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "") //一般情况是这样 ,/static/img/logo.png
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)
	}
}

func GetImgUrl(u string) ([]string, error) {
	var urls []string
	//获取URL的response
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	//io.Copy(os.Stdout,resp.Body)

	//使用goquery获取网页中的图片链接
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}

	//				标签<img>
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src") //attr() 选择的元素
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not  found")
		}
	})
	return urls, nil
}

func Wget_Img(link, desc string) error {
	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}

	img_resp, err := http.Get(link)
	if err != nil {
		fmt.Println("图片下载失败")
		return err
	}

	imgname := filepath.Join(desc, u.Path)
	dir := filepath.Dir(imgname) //获取父目录
	os.MkdirAll(dir, 0755)       //创建父目录
	f, err := os.Create(imgname) //创建文件
	if err != nil {
		return err
	}
	defer f.Close()

	body, _ := ioutil.ReadAll(img_resp.Body)
	f.Write(body) //将文件写入
	return nil

}

func main() {
	access_url := "http://59.110.12.72:7070/golang-spider/img.html"
	urls, err := GetImgUrl(access_url)
	if err != nil {
		log.Fatal(err)
	}
	access_url_info, err := url.Parse(access_url)
	if err != nil {
		log.Fatal(err)
	}
	desc := "./img/"
	for _, link := range urls {
		img_link := CleanUrl(access_url_info, link)
		Wget_Img(img_link, desc)
	}

	tarFun(os.Args[1], desc)

}
