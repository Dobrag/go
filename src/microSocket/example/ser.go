package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	msf "microSocket"
	"microSocket/pb"
	"net"
)

var ser = msf.NewMsf(&msf.CommSocket{})

//框架事件
//----------------------------------------------------------------------------------------------------------------------
type event struct {
}

//客户端握手成功事件
func (this event) OnHandel(fd uint32, conn net.Conn) bool {
	log.Println(fd, "链接成功类")
	return true
}

//断开连接事件
func (this event) OnClose(fd uint32) {
	log.Println(fd, "链接断开类")
}

//接收到消息事件
func (this event) OnMessage(fd uint32, msg map[string]string) bool {
	log.Println("这个是接受消息事件",msg)
	return true
}
//----------------------------------------------------------------------------------------------------------------------
//框架业务逻辑
type Test struct {
}

func (this Test) Default(fd uint32,data map[string]string) bool {
	log.Println("default")
	return true
}

func (this Test) BeforeRequest(fd uint32,data map[string]string) bool {
	log.Println("before")
	return true
}

func (this Test) AfterRequest(fd uint32,data map[string]string) bool{
	log.Println("after")
	return true
}

func (this Test) Hello(fd uint32,data map[string]string) bool {
	log.Println("收到消息了")
	log.Println(data)
	ser.SessionMaster.WriteByid(fd,[]byte("我是一个可小i奥的石头"))
	return true
}

//protoc --go_out="." --proto_path =./proto test.proto
//--go_out=plugins=grpc:. --proto_path=./proto test.proto

//----------------------------------------------------------------------------------------------------------------------
type SearchService struct{}
func (s *SearchService)Search(ctx context.Context,r *pb.SearchRequest)(*pb.SearchResponse,error){
	return &pb.SearchResponse{Response: r.GetRequest() + "Server"},nil
}
const PORT = "9001"
func main() {
	//log.SetFlags(log.Lshortfile | log.LstdFlags | log.Llongfile)
	//log.Print("hah")
	server := grpc.NewServer();
	pb.RegisterSearchServiceServer(server,&SearchService{});
	lis,err := net.Listen("tcp",":"+PORT);
	if err != nil {
		log.Fatalf("failed to listen:%v",err);
	}
	server.Serve(lis);


	//ser.EventPool.RegisterEvent(&event{})
	//ser.EventPool.RegisterStructFun("test", &Test{})
	//ser.Listening(":8565")
}
