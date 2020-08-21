package main

import (
	"context"
	//"microSocket/util"
	"google.golang.org/grpc"
	"log"
	"microSocket/pb"
)
func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:8565")
	//if err != nil {
	//	return
	//}
	//defer conn.Close()

	//person1 := &pb.Person{
	//	Id: 32,
	//	Name: "haha",
	//	Phones:[]*pb.Phone{
	//		{
	//			Type:   pb.PhoneType_WORK,
	//			Number: "",
	//		},
	//		{
	//			Type:   pb.PhoneType_HOME,
	//			Number: "3333",
	//		},
	//	},
	//}
	//person2 := &pb.Person{
	//	Id: 32,
	//	Name: "haha",
	//	Phones:[]*pb.Phone{
	//		{
	//			Type:   pb.PhoneType_WORK,
	//			Number: "",
	//		},
	//		{
	//			Type:   pb.PhoneType_HOME,
	//			Number: "3333",
	//		},
	//	},
	//}
	//
	//book := &pb.ContactBook{};
	//book.Persons = append(book.Persons,person1);
	//book.Persons = append(book.Persons,person2);
	//
	////编码数据
	//data , _  := proto.Marshal(book);
	//log.Print(data);
	//
	//ubook := &pb.ContactBook{};
	//proto.Unmarshal(data,ubook);
	//log.Print(ubook);

	//tt := make(map[string]string)
	//tt["module"] = "test"
	//tt["method"] = "Hello"
	//tt["name"] = "jd"
	//把map转化为string
	//a := []byte(util.Map2String(tt))
	//log.Print(a)

	//把string打包
	//sock := &ms.CommSocket{}
	//b := sock.Pack(data)
	//
	////发送数据
	//conn.Write(b)
	//
	//res := make([]byte, 20)
	//conn.Read(res)
	//fmt.Println(string(res))\
	conn , err := grpc.Dial(":9001",grpc.WithInsecure());
	if err != nil {
		log.Fatal("grpc.Dial err :%v",err);
	}
	defer conn.Close();

	client := pb.NewSearchServiceClient(conn);
	resp ,err := client.Search(context.Background(),&pb.SearchRequest{Request:"haha"})
	if err != nil {
		log.Fatal("client.search err :%v",err);
	}
	log.Printf("resp:%s",resp.GetResponse())


}
