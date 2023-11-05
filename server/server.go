package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	userData "totality-project-gRPC/data"
	pb "totality-project-gRPC/proto"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 5501, "listening port")

type streamServer struct {
	pb.UnimplementedUserServiceServer
}

type User struct {
	UserId    int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

func (s *streamServer) GetUserByUserId(stream pb.UserService_GetUserByUserIdServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("[Recieved Request From Client] : %v\n", req)

		u, ok := userData.UsersData[req.UserId]

		if !ok {
			u = pb.UserResponse{UserId: req.GetUserId(), FirstName: "User Not Found", City: "USer Not Found", Phone: "User Not Found", Height: 0, Married: false}
		}
		err = stream.Send(&u)
		if err != nil {
			return err
		}

	}
}

func (s *streamServer) GetUserListByIds(stream pb.UserService_GetUserListByIdsServer) error {
	var u pb.UserResponseList
	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	log.Printf("[Recieved Request From Client] : %v\n", req)

	for i := 0; i < len(req.GetUserRequestList()); i++ {
		u1, ok := userData.UsersData[req.GetUserRequestList()[i].UserId]
		if !ok {
			u1 = pb.UserResponse{UserId: req.GetUserRequestList()[i].UserId, FirstName: "User Not Found", City: "User Not Found", Phone: "User Not Found", Height: 0, Married: false}
		}
		u.UserResponseList = append(u.UserResponseList, &u1)

	}

	err = stream.Send(&u)
	if err != nil {
		return err
	}

	return nil

}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("error in listening: %v", err)
	} else {
		log.Printf("server listen at port: %v", *port)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &streamServer{})
	grpcServer.Serve(listener)

}
