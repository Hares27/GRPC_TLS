	
package main

import(
	"log"
	"google.golang.org/grpc"
	pb "tlsgrpc/protos/greetingproto"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"google.golang.org/grpc/credentials"
    "tlsgrpc/client/handler"
	
	

)


func main(){

	

log.Printf("Client")
	caCert, err := ioutil.ReadFile("../cert/ca-cert.pem")
	
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}
	

	//read client cert
	
	clientCert, err := tls.LoadX509KeyPair("../newCertificates/client-cert.pem", "../newCertificates/client-key.pem")
	if err != nil {
		log.Fatal("Error clientCert",err)
	}
	
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	


	conn, err := grpc.NewClient("0.0.0.0:7777", grpc.WithTransportCredentials(tlsCredential))
	
	if err != nil {
	  log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := pb.NewGreetingServiceClient(conn)
	handler.GreetUser(client)
	
	
}