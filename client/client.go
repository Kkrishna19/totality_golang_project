package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "totality-project-gRPC/proto"
)

var (
	addr = flag.String("addr", "0.0.0.0:5501", "connect server address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connection server error: %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewUserServiceClient(conn)
	userReq, err := grpcClient.GetUserByUserId(context.Background())

	if err != nil {
		log.Fatal("userReq error: %v", err)
	}
	var i int32
	for i = 1; i < 10; i++ {
		userReq.Send(&pb.UserRequest{UserId: i})
		resp, err := userReq.Recv()
		if err != nil {
			log.Fatalf("response error: %v", err)
		}

		log.Printf("[Response Recieved for GetUserByUserId() from server]: %v\n", resp)
	}

	userListReq, err := grpcClient.GetUserListByIds(context.Background())
	if err != nil {
		log.Fatal("userListReq error: %v", err)
	}
	userRequestList := []*pb.UserRequest{{UserId: 1}, {UserId: 5}, {UserId: 10}, {UserId: 100}}
	userListReq.Send(&pb.UserRequestList{
		UserRequestList: userRequestList,
	})
	resp, err := userListReq.Recv()
	if err != nil {
		log.Fatalf("response error: %v", err)
	}

	log.Printf("[Response Received for GetUserListByIds() from sever]: %v\n", resp)
}
