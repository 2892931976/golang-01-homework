package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	_ "reflect"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// 1.获取图片链接
func fetch(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//resp.Body必须关闭，不然会造成资源泄漏
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	//输出到屏幕，不占用内存
	//io.Copy(os.Stdin, resp.Body)

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("src")
		//link_type := reflect.TypeOf(link)
		//fmt.Println(link_type)
		urls = append(urls, link)
	})
	urls = cleanUrls(url, urls)
	return urls, nil
}

// 2.清洗url,格式化链接
func cleanUrls(weburl string, urls []string) []string {
	var resurls []string
	u, err := url.Parse(weburl)
	if err != nil {
		log.Fatal(err)
	}
	s := make([]string, 10)
	s = strings.Split(u.Path, "/")
	s = s[:len(s)-1]
	var path string
	for _, v := range s {
		path += v + "/"
	}
	for _, v := range urls {
		switch {
		case strings.HasPrefix(v, "http") || strings.HasPrefix(v, "https"):
			resurls = append(resurls, v)
		case strings.HasPrefix(v, "//"):
			resurls = append(resurls, u.Scheme+":"+v)
		case strings.HasPrefix(v, "/"):
			resurls = append(resurls, u.Scheme+"://"+u.Host+v)
		default:
			resurls = append(resurls, u.Scheme+"://"+u.Host+path+v)
		}

	}
	return resurls
}

// 3.下载文件到本地(并行下载和串行下载)
//并行下载
func work(ch chan string, wg *sync.WaitGroup, dir string) {
	for u := range ch {
		resp, err := http.Get(u)
		if err != nil {
			log.Print(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatal(errors.New(resp.Status))
		}
		name := path.Base(u)
		fullpath := filepath.Join(dir, name)
		f, err := os.Create(fullpath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(f, resp.Body)
	}
	wg.Done()
}

func downloadImgs(urls []string, dir string) error {
	wg := new(sync.WaitGroup)
	wg.Add(10)
	taskch := make(chan string)
	for i := 0; i < 10; i++ {
		go work(taskch, wg, dir)
	}
	for _, weburl := range urls {
		taskch <- weburl
	}
	close(taskch)
	wg.Wait()
	return nil
}

// 4.打包
func maketar(dir string, w io.Writer) error {
	base := filepath.Base(dir)
	compress := gzip.NewWriter(w)
	defer compress.Close()
	tr := tar.NewWriter(compress)
	defer tr.Close()
	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		//写入tar的FileHeader
		//以读取的方式打开文件
		//判断目录和文件，如果是文件，把文件内容写入body
		if err != nil {
			return err
		}

		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		//header.Name = name
		p, _ := filepath.Rel(dir, name)
		header.Name = filepath.Join(base, p)
		if err = tr.WriteHeader(header); err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			io.Copy(tr, f)
		}
		return nil

	})
	return nil
}

// 5.并发

func main() {
	//url := "http://59.110.12.72:7070/golang-spider/img.html"
	//url := "http://daily.zhihu.com/"
	url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	/*
		for _, u := range urls {
			fmt.Println(u)
		}

	*/
	urls = cleanUrls(url, urls)

	tmpdir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(tmpdir)
	//defer os.RemoveAll(tmpdir)
	err = downloadImgs(urls, tmpdir)
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Create("Img.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	maketar(tmpdir, f)
}
