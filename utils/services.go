package utils

import (
	"net"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/ZeroTechh/VelocityCentral/services"
	"github.com/ZeroTechh/blaze"
)

// HandlePanic is used to handle panic when it occurs in a microservice
func HandlePanic(log *blaze.FuncLog) {
	if r := recover(); r != nil {
		log.Panic("Service Paniced", zap.Any("Panic", r))
	}
}

// CreateGRPCServer is used to create a GRPC server for microservice
func CreateGRPCServer(name string, log *zap.Logger) (*grpc.Server, *net.Listener) {
	url := services.ServiceUrls[name]
	lis, err := net.Listen("tcp", url)
	if err != nil {
		err := errors.Wrap(err, "Error While Listening")
		log.Fatal(
			"Error While Creating Service",
			zap.String("Service Name", name),
			zap.String("URL", url),
			zap.Error(err),
		)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	grpc.NewServer()
	return grpcServer, &lis
}

// CreateGRPCClient is used to create a GRPC Service's client
func CreateGRPCClient(name string, log *zap.Logger) *grpc.ClientConn {
	url := services.ServiceUrls[name]
	conn, err := grpc.Dial(
		url,
		grpc.WithInsecure(),
	)
	if err != nil {
		err := errors.Wrap(err, "Error While Dialing")
		log.Fatal(
			"Error While Creating GRPC Client",
			zap.String("URL", url),
			zap.Error(err),
		)
		panic(err)
	}
	return conn
}
