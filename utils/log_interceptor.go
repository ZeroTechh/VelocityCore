package utils

import (
	"context"

	"github.com/ZeroTechh/blaze"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LogInterceptor logs all communications
type LogInterceptor struct {
	log *zap.Logger
}

// LogGRPCInterceptor logs all grpc communications
func (l LogInterceptor) LogGRPCInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	funcLog := blaze.NewFuncLog(
		info.FullMethod,
		l.log,
		zap.Any("Request", req),
	)
	funcLog.Completed()

	output, err := handler(ctx, req)
	if err != nil {
		funcLog.Error(err)
		return output, err
	}

	funcLog.Completed(zap.Any("Output", output))
	return output, err
}
