package main //引入包，放在程序的第一行，两种package，一种是库package, 一种是二进制package;二进制package使用main表示，库package的名字跟go文件所在的目录名一样

import ( //引入第三方库
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
