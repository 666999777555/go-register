package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("port", 50051, "the port to serve on")

type ecServer struct {
	pb.UnimplementedEchoServer
}

func Register(port int, res func(s *grpc.Server)) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
		return err
	}
	cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	s := grpc.NewServer()
	//反射接口支持查询
	reflection.Register(s)
	res(s)
	log.Printf("sever listening at %v", err)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen:%v\",err")
		return err
	}
	return err
}
