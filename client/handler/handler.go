package handler

import(
	"context"
	"log"
	pb "tlsgrpc/protos/greetingproto"
)


func GreetUser(client pb.GreetingServiceClient){

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "user"})
	if err != nil {
	  log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
	
}