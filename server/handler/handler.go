package handler

import(
	pb "tlsgrpc/protos/greetingproto"
	"golang.org/x/net/context"
	"log"
)

type GreetingServiceServer struct{
	
}





func Newserver()*GreetingServiceServer{
	return &GreetingServiceServer{};

}


func (server *GreetingServiceServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name:=req.GetName()
	log.Println("Request from client is ",name)
	return &pb.HelloReply{Message:name},nil
}