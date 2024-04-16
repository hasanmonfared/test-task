package walletgrpcserver

import (
	"app/config"
	"app/contract/goproto/wallet"
	"app/param/walletparam"
	"app/pkg/protobufmapper"
	"app/service/walletservice"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	conf config.Config
	wallet.UnimplementedCreateTransactionServiceServer
	svc walletservice.Service
}

func New(cfg config.Config, svc walletservice.Service) Server {
	return Server{
		svc:  svc,
		conf: cfg,
		UnimplementedCreateTransactionServiceServer: wallet.UnimplementedCreateTransactionServiceServer{},
	}
}

func (s Server) CreateTransaction(ctx context.Context, req *wallet.CreateTransactionRequest) (*wallet.CreateTransactionResponse, error) {

	resp, err := s.svc.CreateTransaction(ctx, walletparam.CreateTransactionRequest{
		User:   req.User,
		Type:   protobufmapper.MapProtobufTypeWalletToParam(req.Type),
		Amount: float64(req.Amount),
	})
	if err != nil {
		return nil, err
	}
	return protobufmapper.MapResponseParamToProtobuf(resp), nil
}

func (s Server) Start() {
	address := fmt.Sprintf(":%d", s.conf.WalletGrpcServer.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	wallet.RegisterCreateTransactionServiceServer(grpcServer, &s)
	log.Print("presence grpc server started on ", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("couldn't server presence grpc server", err)
	}
	log.Printf("presence grpc server started on %d", address)
}
