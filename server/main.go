package main

import(
	"fmt"
	"net"
	"log"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	pb "tlsgrpc/protos/greetingproto"
	"tlsgrpc/server/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)



	

	


func main(){
	
	log.Printf("Server");

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
	  log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := handler.Newserver()
	// create a gRPC server object


	
	caPem,err:=ioutil.ReadFile("../cert/ca-cert.pem");
	if err!=nil{
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		log.Fatal(err)
	}
	

	serverCert, err := tls.LoadX509KeyPair("../cert/server-cert.pem", "../cert/server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(conf)
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	
    pb.RegisterGreetingServiceServer(grpcServer,s)
	 err = grpcServer.Serve(lis); 
	 log.Println("Error",err);
	 if err !=nil{
		log.Fatal(err)
	 }
	
	
	

	
}