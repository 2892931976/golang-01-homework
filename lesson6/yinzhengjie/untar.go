/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	//"compress/gzip"
	//"log"
)

func main() {
	/*
		uncompress,err := gzip.NewReader(os.Stdin)  //讲传入的文件解压传给“uncompress”
		if err != nil {
			log.Fatal(err)  //意思是当程序解压失败时，就立即终止程序，“log.Fatal”一般用于程序初始化。
		}
	*/
	tr := tar.NewReader(os.Stdin) /*从 “*.tar”文件中读出数据是通过“tar.Reader”完成的，所以首先要创建“tar.Reader”
	，
	         可以通过“tar.NewReader”方法来创建它，该方法要求提供一个“os.Reader”对象，以便从该对象中读出数据。*/
	for {
		hdr, err := tr.Next() //此时，我们就拥有了一个“tar.Reader”对象 tr，可以用“tr.Next()”来遍历包中的文件.

		if err != nil {
			return
		}
		fmt.Printf("已解压：\033[31;1m%s\033[0m\n", hdr.Name)
		//io.Copy(ioutil.Discard,tr) //表示将读取到到内容丢弃，"ioutil.Discard"可以看作是Linux中的：／dev/null.
		info := hdr.FileInfo() // 获取文件信息
		if info.IsDir() {      //判断文件是否为目录
			os.Mkdir(hdr.Name, 0755) //创建目录并赋予权限。
			continue                 //创建目录后就要跳过当前循环，继续下一次循环了。
		}
		f, _ := os.Create(hdr.Name) //如果不是目录就直接创建该文件
		io.Copy(f, tr)              //最终将读到的内容写入已经创建的文件中去。
		f.Close()                   /*不建议写成“defer f.Close()”因为“f.Close()”会将缓存中的数据写入到文件中，同时“f.Close()”
		                还会向“*.tar”文件的最后写入结束信息，如果不关闭“f”而直接退出程序，那么将导致“.tar”文件不完整。而
		                “defer f.Close()”是在函数结束后再执行关闭文件，那么在这个过程中，内存始终会被占用着，浪费这不必要的资
		源。*/
	}
}
