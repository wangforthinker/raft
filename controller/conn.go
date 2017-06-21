package controller

import (
	"github.com/wangforthinker/raft/utils"
	"time"
	"net"
	"fmt"
)

var(
	connLog = utils.NewLogger("conn")
)

type Message struct {
	MsgType	int
	Data	[]byte
}

type sendMsgWrap struct {
	msg	*Message
	//unique id
	msgID	int64
	//receive msg channel
	rev	chan <- *Message
}

type SendMsgQueue struct {
	cache	map[int64]*sendMsgWrap
}

type Conn struct {
	addr	string
	port	int

	stopCh	<- chan struct{}
	errorCh chan error

	conn	net.Conn
	connected bool
}

func NewConn(addr string, port int, stopChan <- chan struct{}) (*Conn,error) {
	c := &Conn{
		addr: addr,
		port: port,
		stopCh: stopChan,
	}

	go c.loop()

	return c,nil
}

func (c* Conn) loop() {
	timer := time.NewTicker(time.Second * 10)
	defer timer.Stop()

	connOk := make(chan <- bool)

	for{
		err := c.dial(connOk)
		if(err != nil){
			connLog.Errorf("connect error:%s",err.Error())
		}

		select {
		case <-c.stopCh:
			connLog.Info("receive stop signal, stop connection...")
			return
		case <-timer.C:
			connLog.Info("receive timer ticker, try to connect again...")
			break
		case <- connOk:
			connLog.Info("connect success")

		}

	}

	return nil
}

func (c* Conn) SendMessage() {

}

func (c* Conn) sendLoop() {

}

func (c* Conn) recvLoop() {

}

func (c* Conn) dial(ok chan <- bool) error {
	conn,err := net.Dial("tcp",fmt.Sprintf("%s:%d",c.addr, c.port))
	if(err != nil){
		return err
	}

	c.conn = conn
	c.connected = true

	ok <- true

	return nil
}