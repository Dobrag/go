package microSocket

import (
	"log"
	"net"
)

type SocketTypes interface{
	ConnHandle(msf *Msf,sess *Session)
	Pack(data []byte)[]byte
}

type Msf struct {
	EventPool     *RoutersMap
	SessionMaster *SessionM
	SocketType    SocketTypes
}

func NewMsf(socketType SocketTypes) *Msf {
	msf := &Msf{
		EventPool :     NewRoutersMap(),
		SocketType:     socketType,
	}
	msf.SessionMaster = NewSessonM(msf)
	return msf
}

func (this *Msf) Listening(address string) {
	tcpListen, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}
	go this.SessionMaster.HeartBeat(2)  //开启心跳
	fd := uint32(0) //初始化fd
	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//调用握手事件
		log.Println("11111")
		if this.EventPool.OnHandel(fd, conn) == false {
			continue
		}
		log.Println("2222")
		this.SessionMaster.SetSession(fd, conn)
		go this.SocketType.ConnHandle(this,this.SessionMaster.GetSessionById(fd))
		fd++
	}
}
//解包得到数据之后，拿到数据在这里执行
func (this *Msf) Hook(fd uint32,requestData map[string]string)bool {
	//调用接收消息事件
	log.Print(requestData)
	if this.EventPool.OnMessage(fd, requestData) == false {
		return false
	}
	//requestData["fd"] = fmt.Sprintf("%d", fd)
	//路由
	if actionName, exit := requestData["action"]; exit {
		if this.EventPool.HookAction(actionName, fd, requestData) == false {
			return false
		}
	} else {
		if this.EventPool.HookModule(requestData["module"],requestData["method"], fd, requestData) == false {
			return false
		}
	}
	return true
}


