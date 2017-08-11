package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func downloadImgs(ch chan string, wg *sync.WaitGroup, dir string) error {
	for u := range ch {
		resp, _ := http.Get(u)
		defer resp.Body.Close()
		name := filepath.Base(u)
		fullname := filepath.Join(dir, name)
		f, _ := os.Create(fullname)
		io.Copy(f, resp.Body)
		fmt.Println(fullname)
	}
	wg.Done()
	return nil
}

func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("src")
		fmt.Println(link)
		urls = append(urls, link)
	})

	return urls, nil
}

func cleanUrls(u string, urls []string) []string {
	var ret []string
	uri, _ := url.Parse(u)
	for i := range urls {
		ret = append(ret, cleanUrl(uri, urls[i]))
	}
	return ret
}

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

func cleanUrl(uri *url.URL, link string) string {
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

func maketar(dir string, w io.Writer) error {
	base := filepath.Base(dir)
	compress := gzip.NewWriter(w) // 压缩
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()

		h, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, name)
		h.Name = filepath.Join(base, p)
		if err = tr.WriteHeader(h); err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			io.Copy(tr, f)
		}
		return nil
	})
	return nil
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
	url := os.Args[1]
	urls, err := fetch(url)
	fmt.Println("urls", urls)
	if err != nil {
		log.Fatal(err)
	}

	urls = cleanUrls(url, urls)
	for _, u := range urls {
		fmt.Println(u)
	}

	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}

	// 5个协程
	wg := new(sync.WaitGroup)
	wg.Add(5)
	taskch := make(chan string)
	for i := 0; i < 5; i++ {
		go downloadImgs(taskch, wg, tmpdir)
	}
	for _, url := range urls {
		taskch <- url
	}
	close(taskch)
	wg.Wait()
	//

	f, err := os.Create("a.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	err = maketar(tmpdir, f)
	if err != nil {
		log.Fatal(err)
	}

}
