/*
#!/usr/bin/env gorun
@author :yinzhengjie
Blog:http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/
EMAIL:y1053419035@qq.com
*/

package main

import (
	"yinzhengjie/monitor/common"
	"net"
	"fmt"
	"encoding/json"
)

/*
	这个包是专门处理数据发送的。
*/

type Sender struct {  //定义一个Sender结构体，接收Channel发来的数据，同事通过网络IP将数据发送出去。
	addr string  //定义IP地址
	ch chan *common.Metric  //接收到的数据。
}

func NewSender(addr string) *Sender {  //Sender的构造方法。
	return &Sender{
		addr:addr,
		ch:make(chan *common.Metric),
	}
}

func (s *Sender) Start() {  //从Channel中读取数据。
	conn,err := net.Dial("tcp",s.addr)  //建立连接
	if err != nil {
		panic(err)
	}

	for  {  //循环从s.sh里面读取metric
		metric := <- s.ch  //将Channel中的数据取出来
		buf,_ := json.Marshal(metric)  //序列化metric
		fmt.Fprintf(conn,"%s\n",buf) //发送数据
	}

	//for metric := range s.ch{   //上面的for循环也可以这样写，这是2种不同的用channel的传值方式。
	//	buf,_ := json.Marshal(metric)
	//	fmt.Fprintf(conn,"%s\n",buf)
	//}


}

func (s *Sender	)Channel()  chan *common.Metric {  //把自己的Channle暴露出去，让别人可以给它发送数据。
	return s.ch
}
