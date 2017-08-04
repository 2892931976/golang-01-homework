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

	"errors"

	"github.com/PuerkitoBio/goquery"
	"sync"
)

//          需要打包的文件 ,
func tarFun(dir string, w io.Writer) error {
	basedir := filepath.Base(dir) //spider020144075

	gw := gzip.NewWriter(w) //写入.gz文件
	defer gw.Close()
	tr := tar.NewWriter(gw) //写入.tar文件
	defer tr.Close()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error { //遍历src文件
		//fmt.Println(path)
		fi, err := os.Stat(path) //获取包含时间戳和权限标志的os.FileInfo值,传递给FileInfoHeader
		if err != nil {
			return err
		}
		hdr, err := tar.FileInfoHeader(fi, "") //获取文件的头部信息
		if err != nil {
			return err
		}
		//
		p, _ := filepath.Rel(dir, path) //获取相对路径
		hdr.Name = filepath.Join(basedir, p)

		//hdr.Name = path           //替换文件的Name信息 (使其包含之前的目录结构)
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
			////FileModel的方法主要用来进行判断和输出权限
			////判断m是否是目录，也就是检查文件是否有设置的ModeDir位
			////判断m是否是普通文件，也就是说检查m中是否有设置mode type

			//if fi.Mode().IsRegular() {
			io.Copy(tr, fs)
			//}
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

func Wget_Img(c chan string, wg *sync.WaitGroup, desc string) error {
	for link := range c {
		u, err := url.Parse(link)
		if err != nil {
			log.Fatal(err)
		}
		img_resp, err := http.Get(link)
		if err != nil {
			fmt.Println("图片下载失败")
			return err
		}
		defer img_resp.Body.Close() //关闭连接
		if img_resp.StatusCode != http.StatusOK {
			return errors.New("图片下载失败")
		}
		imgname := filepath.Join(desc, u.Path)
		dir := filepath.Dir(imgname) //获取父目录
		os.MkdirAll(dir, 0755)       //创建父目录
		f, err := os.Create(imgname) //创建文件
		if err != nil {
			return err
		}
		defer f.Close()
		//body, _ := ioutil.ReadAll(img_resp.Body)
		io.Copy(f, img_resp.Body)
		//f.Write(body) //将文件写入

	}
	wg.Done()
	return nil
}
//实现多线程下载
func xc_download(urls  []string,tmp_dir string){
	wg := new(sync.WaitGroup)
	wg.Add(len(urls))
	taskch := make(chan string)
	for i := 0; i < len(urls); i++ {
		go Wget_Img(taskch, wg,tmp_dir)
	}
	for _, url := range urls {
		taskch <- url
	}
	close(taskch)
	wg.Wait()
}


func main() {
	access_url := "https://daily.zhihu.com/"
	urls, err := GetImgUrl(access_url)
	if err != nil {
		log.Fatal(err)
	}
	access_url_info, err := url.Parse(access_url)
	if err != nil {
		log.Fatal(err)
	}
	//desc := "./img/"
	//创建一个临时文件		     自动查找一个目录,前缀
	tmp_dir, err := ioutil.TempDir("", "spider")
	if err != nil {
		log.Fatal(err)
	}

	//defer os.RemoveAll(desc) 删除临时目录
	var urls_ok []string
	for _, link := range urls {
		img_link := CleanUrl(access_url_info, link)
		urls_ok = append(urls_ok, img_link)
	}

	//多进程下载
	xc_download(urls_ok,tmp_dir)

	//压缩
	f_tar, err := os.Create(os.Args[1]) //创建目标文件
	if err != nil {
		log.Fatal(err)
	}
	defer f_tar.Close()
	tarFun(tmp_dir, f_tar)

}
