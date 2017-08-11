/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func fetch(url string) ([]string, error) { //改函数会拿到我们想要的图片的路径。
	var urls []string //定义一个空切片数组
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) //表示当出现错误是，返回空列表，并将错误状态返回。
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link) //将过滤出来的图片路径都追加到urls的数组中去,最终返回给用户。
		} else {
			fmt.Println("抱歉，没有发现该路径。")
		}

	})
	return urls, nil
}

func Clean_urls(root_path string, picture_path []string) []string {
	var Absolute_path []string //定义一个绝对路径数组。
	url_info, err := url.Parse(root_path)
	if err != nil {
		log.Fatal(err)
	}
	Scheme := url_info.Scheme //获取到链接的协议
	Host := url_info.Host     //获取链接的主机名
	for _, souce_path := range picture_path {
		if strings.HasPrefix(souce_path, "https") { //如果当前当前路径是以“https”开头说明是绝对路径，因此我们给一行空代码，表示不执行任何操作，千万别写：“continue”，空着就好。

		} else if strings.HasPrefix(souce_path, "//") { //判断当前路径是否以“//”开头(说明包含主机名)
			souce_path = Scheme + ":" + souce_path //如果是就对其进行拼接操作。以下逻辑相同。
		} else if strings.HasPrefix(souce_path, "/") { //说明不包含主机名和协议，我们进行拼接即可。
			souce_path = Scheme + "://" + Host + souce_path
		} else {
			souce_path = filepath.Dir(root_path) + souce_path //文件名称和用户输入的目录相拼接。
		}
		Absolute_path = append(Absolute_path, souce_path) //不管是否满足上面的条件，最终都会被追加到该数组中来。
	}
	return Absolute_path //最终返回处理后的每个链接的绝对路基。
}

func downloadImgs(urls []string, dir string) error {
	for _, link := range urls {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()               //千万要关闭链接，不然会造成资源泄露（就是为了防止死循环）。
		if resp.StatusCode != http.StatusOK { //如果返回状态出现错误，就抛出错误。也就是你要过滤出来你下载的图片是ok的！
			log.Fatal(resp.Status)
		}
		file_name := filepath.Base(link)           //创建一个文件名也就是，这里我们捕捉网站到原文件名称。
		full_name := filepath.Join(dir, file_name) //这是将下载到文件存放在我们指定到目录中去。
		f, err := os.Create(full_name)             //创建我们定义到文件。
		if err != nil {
			log.Panic("创建文件失败啦！") //创建失败的话，我们就给出自定义的报错。
		}
		io.Copy(f, resp.Body) //将文件到内容拷贝到我们创建的文件中。
		//fmt.Printf("已下载文件至：\033[31;1m%s\033[0m\n",full_name)
		//defer os.RemoveAll(file_name) //删除文件。
	}
	return nil
}

func make_tar(dir string, w io.Writer) error {
	basedir := filepath.Base(dir) //取出文件的目录
	compress := gzip.NewWriter(w) //实现压缩功能
	defer compress.Close()
	tw := tar.NewWriter(w) //表示我们会把数据都写入w中去。而这个w就是我们在主函数中创建的文件。
	defer tw.Close()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error { //用"filepath.Walk"函数去递归"dir"目录下的文件。
		header, err := tar.FileInfoHeader(info, "") //将文件信息读取出来传给"header"变量，注意"info"后面参数是不传值的，除非你的目录是该软连接。
		if err != nil {
			return err
		}
		p, _ := filepath.Rel(dir, path)         //取出文件的上级目录。
		header.Name = filepath.Join(basedir, p) //header.Name = path  //这是将path的相对路径传给"header.Name ",然后在写入到tw中去。不然的话只能拿到"info.Name ()"的名字，也就是如果不来这个操作的话它只会保存文件名，而不会记得路径。
		//fmt.Printf("path=%s,header.name=%s,info.name= %s\n",path,header.Name,info.Name())
		tw.WriteHeader(header) //将文件的信息写入到文件w中去。
		if info.IsDir() {
			return nil
		}
		f1, err := os.Open(path)
		if err != nil {
			log.Panic("创建文件出错！")
		}
		defer f1.Close()
		io.Copy(tw, f1) //再将文件的内容写到tw中去。
		return nil
	})
	return nil
}

func main() {
	root_path := "http://daily.zhihu.com/" //定义一个URl，也就是我们要爬的网站。
	picture_path, err := fetch(root_path)  //“fetch”函数会帮我们拿到picture_path的路径，但是路径可能是相对路径或是绝对路径。不同意。
	if err != nil {
		log.Fatal(err)
	}
	Absolute_path := Clean_urls(root_path, picture_path) //“Clean_urls”函数会帮我们把picture_path的路径做一个统一，最终都拿到了绝对路径Absolute_path数组。
	//for _, Picture_absolute_path := range Absolute_path {
	//	fmt.Println(Picture_absolute_path) //最终我们会得到一个图片的完整路径，我们可以对这个路径进行下载，压缩，加密等等操作。
	//}
	tmpdir, err := ioutil.TempDir("", "yinzhengjie") //创建一个临时目录，注意第一个参数最好设置为空(如果设置的话指定目录必须为空。)，因为系统会随机给你在"yinzhengjie"随机加一串数字。
	defer os.RemoveAll(tmpdir)                       //将临时目录删除掉,但是未了能看到效果我们不要删除，方便我们取验证。
	//fmt.Println(tmpdir)
	err = downloadImgs(Absolute_path, tmpdir)
	if err != nil {
		log.Panic(err)
	}
	//make_tar("..",os.Stdout) //将结果输出到屏幕上。
	f, err := os.Create("img.tar.gz") //创建一个"io.Writer"，即可写对象。
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	make_tar(tmpdir, f) //将下载到文件放到一个临时目录中，然后把这个临时目录生成一个我们自定义的文件f。
}
