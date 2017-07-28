package main

import (
	"archive/tar"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
	//"compress/gzip"
)

func fetch(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	var urls []string

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			//fmt.Println(link)
			urls = append(urls, link)
			//return urls
		} else {
			fmt.Println("not found")
		}
	})
	return urls, nil
}

func cleanUrl(taskq chan string, ur string, urls []string) {
	//var urll []string
	u, err := url.Parse(ur)
	if err != nil {
		fmt.Println("base url parse err")
	}
	base_scheme := u.Scheme
	base_host := u.Host
	base_path := strings.Split(u.Path, "/")[1]

	for _, i := range urls {
		switch {
		case strings.HasPrefix(i, "http"):
		case strings.HasPrefix(i, "//"):
			i = "https:" + i
		case strings.HasPrefix(i, "/"):
			i = base_scheme + "://" + base_host + "/" + i
		default:
			i = base_scheme + "://" + base_host + "/" + base_path + "/" + i
		}

		taskq <- i
	}
}

func downimag(urls chan string, wg *sync.WaitGroup, tmpdir string) error {

	for link := range urls {
		reps, err := http.Get(link)

		if err != nil {
			fmt.Println("create img file err: ", err)
			continue
		}
		defer reps.Body.Close()
		s := string(path.Base(link))
		//tmpdir = tmpdir + "\"
		f, err := os.Create(tmpdir + "\\" + s)
		defer f.Close()
		if err != nil {
			fmt.Println("get url error")
			continue
		}
		io.Copy(f, reps.Body)
	}
	wg.Done()
	return nil
}

func maketar(dir string, dstTar string) error {
	// 创建空的目标文件
	fw, er := os.Create(dstTar + ".tar")
	if er != nil {
		return er
	}
	defer fw.Close()

	//uncompress := gzip.NewWriter(fw)		// 压缩
	//defer uncompress.Close()

	tw := tar.NewWriter(fw) // 打包
	defer tw.Close()

	b_dir := filepath.Base(dir)

	filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			fmt.Println("tar FileInfoHeader err")
			return err
		}

		p, _ := filepath.Rel(dir, name)
		header.Name = filepath.Join(b_dir, p)

		if err = tw.WriteHeader(header); err != nil {
			fmt.Println("tw WriteHeader err", err)
			return err
		}
		// 打开要打包的文件，准备读取
		fr, er := os.Open(name)
		if er != nil {
			fmt.Println("open file error", er)
			return er
		}
		defer fr.Close()

		// 将文件数据写入 tw 中
		//if _, er = io.Copy(tw, fr); er != nil {
		//	fmt.Println("io copy err", er)
		//	return er
		//}
		io.Copy(tw, fr) // 这里如果收集报错 将会终止walk 注意！！！！

		return nil
	})
	return nil
}

func main() {
	t1 := time.Now().Unix()

	//url := "http://59.110.12.72:7070/golang-spider/img.html"
	url := "http://www.quanjing.com/"
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("E:\\gopro\\TEST\\lesson8", "spider")
	if err != nil {
		log.Fatal(err)
	}

	taskq := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go downimag(taskq, wg, tmpdir)
	}
	cleanUrl(taskq, url, urls)

	fmt.Println(tmpdir)
	//defer os.RemoveAll(tmpdir)

	close(taskq)
	wg.Wait()

	fmt.Println("sub time seconds>>>>>", int(time.Now().Unix())-int(t1))
	TarFile := filepath.Base(tmpdir)

	if err := maketar(tmpdir, TarFile); err != nil {
		fmt.Println(err)
	}
}
