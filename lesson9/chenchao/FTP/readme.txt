
FTPserver
	启动 go run FTPserver.go
	参数-home  默认为base目录， 自定义参数可以定义root目录
	启动server端，会初始化创建home目录，和创建日志文件

1、实现了 list   upload  get  
	list  获取目录下所有文件
	upload  从客户端接受一个文件
	get	客户端申请下载一个文件

2、添加了日志功能
	test.log记录了客户端的一些操作，有时间记录



FTPclient

	go run FTPclient.go  -a  [list|upload|get] -n filename
	-a 默认为list  -n 默认为all
	执行脚本不加参数默认获取目录下所有文件
1、list
	获取根目录下所有的文件   
2、upload
	向服务器上传一个文件
3、get
	下载一个文件
	