package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type InterceptorManager struct {
	logger logger.Logger
	cfg    *config.Config
}

func NewInterceptorManager(logger logger.Logger, cfg *config.Config) *InterceptorManager {
	return &InterceptorManager{logger: logger, cfg: cfg}
}

func (im *InterceptorManager) Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.logger.Infof("Method: %s, Metadata: %v, Err: %v", info.FullMethod, md, err)

	return reply, err
}

func (im *InterceptorManager) PanicRecover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if err := recover(); err != nil {
			md, _ := metadata.FromIncomingContext(ctx)
			im.logger.Errorf("Method: %s, Metadata: %v, Err: %v", info.FullMethod, md, err)
		}
	}()

	return handler(ctx, req)
}
