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
	"path/filepath"
)

var dir_list, file_list []string //创建两个个动态字符串数组，即切片。用来存取文件和目录。

func walkFunc(path string, info os.FileInfo, err error) error { /*“walkFunc”可以获取3个参数信息，即：文件的绝对路径，
	通过“os.FileInfo”获取文件信息，用“err”返回错误信息，最后需要返回一个“error”类型的数据。*/
	if info.IsDir() { //判断文件类型如果是目录就把他放在目录的动态数组中，
		dir_list = append(dir_list, path)
	} else { //如果不是目录那就按照文件处理，将它放在文件的目录中去。
		file_list = append(file_list, path)
	}
	return nil //返回空值。
}

func main() {
	filepath.Walk(os.Args[2], walkFunc) /*将命令行参数的第三个参数传递给“walkFunc”函数。即用“filepath.Walk”遍历“os.Args[2]”目录下的所有的文件名*/

	f, err := os.Create(os.Args[1]) //创建一个“*.tar”的文件。
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() //有的小伙伴总是忘记关文件，我们可以用defer关键字帮我们忘记关闭文件的坏习惯。

	tw := tar.NewWriter(f) //向“*.tar”文件中写入数据是通过“tar.Writer”完成的，所以首先要创建“tar.Writer”。我们通过“tar.NewWriter”创建他需要提供一个可写的对象，我们上面创建的文件就得到用处。
	defer tw.Close()

	for _, d_list := range dir_list {
		fileinfo, err := os.Stat(d_list) //获取目录的信息
		if err != nil {
			fmt.Println(err)
		}
		hdr, err := tar.FileInfoHeader(fileinfo, "") /*“tar.FileInfoHeader”其实是调用“os.FileInfo ”方法获取文件的信息的，你要知道文件有两个属性，
		一个是文件信息，比如大小啊，编码格式，修改时间等等，还有一个就是文件内容，就是我们所看到的具体内容。 */
		if err != nil {
			fmt.Println(err)
		}
		err = tw.WriteHeader(hdr) //由于是目录，里面的内容我们就不用管理，只记录目录的文件信息。
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, f_list := range file_list {
		fileinfo, err := os.Stat(f_list) //同理，我们将文件也做相应的梳理，获取文件的头部信息，将其传给“tar.Writer”处理。
		if err != nil {
			fmt.Println(err)
		}
		hdr, err := tar.FileInfoHeader(fileinfo, "")
		if err != nil {
			fmt.Println(err)
		}
		err = tw.WriteHeader(hdr)
		if err != nil {
			fmt.Println(err)
		}
		f1, err := os.Open(f_list) //由于是文件，我们就可以看其内容，将头部信息写入后还是不够的，还需要将具体的内容写进去，这样我们得到的才是一个完整的文件。
		if err != nil {
			fmt.Println(err)
		}
		io.Copy(tw, f1) //用io.Copy方法将读到的内容传给“tar.Writer”，让其进行写入到他的对象f中去（也就是“tw := tar.NewWriter(f)”中的“f”）
	}
}
