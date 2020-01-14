package api

import (
	"fmt"
	"net"
	"strconv"

	"github.com/go-xorm/xorm"
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	interceptors "github.com/happilymarrieddad/learning-go-gRPC/api/interceptors"
	pb "github.com/happilymarrieddad/learning-go-gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Routes
	v1users "github.com/happilymarrieddad/learning-go-gRPC/api/v1/users"
)

func Run(port int, db *xorm.Engine) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcmw.ChainUnaryServer(
				interceptors.GlobalRepoInjector(db),
			),
		),
	)

	pb.RegisterV1UsersServer(s, v1users.GetRoutes())

	reflection.Register(s)

	fmt.Printf("Server is running on port %d\n", port)
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
