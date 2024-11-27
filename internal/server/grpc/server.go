package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/interceptors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/metrics"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/db/postgres"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	CFG     *config.Config
	PG      postgres.Client
	S3      *minio.Client
	Logger  logger.Logger
	Metrics *metrics.Metrics

	RegisterServices func(server *grpc.Server)
}

func New(
	cfg *config.Config,
	pg postgres.Client,
	s3 *minio.Client,
	logger logger.Logger,
	metrics *metrics.Metrics,
	registerServices func(server *grpc.Server),
) *Server {
	return &Server{
		CFG:              cfg,
		PG:               pg,
		S3:               s3,
		Logger:           logger,
		Metrics:          metrics,
		RegisterServices: registerServices,
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", s.CFG.Service.GRPC.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	im := interceptors.NewInterceptorManager(s.Logger, s.CFG)
	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.CFG.Service.GRPC.MaxConnectionIdle * time.Minute,
		Timeout:           s.CFG.Service.GRPC.Timeout * time.Second,
		MaxConnectionAge:  s.CFG.Service.GRPC.MaxConnectionAge * time.Minute,
		Time:              s.CFG.Service.GRPC.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			im.Logger,
			im.PanicRecover,
		),
	)

	if os.Getenv("ENV") != "prod" {
		reflection.Register(server)
	}

	if s.RegisterServices != nil {
		s.RegisterServices(server)
	}

	go func() {
		if err := server.Serve(l); err != nil {
			log.Fatalf("failed to start grpc server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	server.GracefulStop()

	return nil
}
