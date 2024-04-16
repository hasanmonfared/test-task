package walletgrpcserver

import (
	"app/config"
	"app/contract/goproto/wallet"
	"app/model/walletmodel"
	"app/param/walletparam"
	"app/service/walletservice"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	conf config.Config
	wallet.UnimplementedPresenceServiceServer
	svc walletservice.Service
}

func New(cfg config.Config) Server {
	return Server{
		conf:                               cfg,
		UnimplementedPresenceServiceServer: wallet.UnimplementedPresenceServiceServer{},
	}
}

func (s Server) CreateTransaction(ctx context.Context, req *wallet.CreateTransactionRequest) (*wallet.CreateTransactionResponse, error) {
	var transactionType walletmodel.Type
	switch req.Type {
	case wallet.Type_DEPOSIT:

		transactionType = walletmodel.Deposit
	case wallet.Type_WITHDRAWAL:

		transactionType = walletmodel.Withdrawal
	default:
		return nil, errors.New("unsupported transaction type")
	}
	resp, err := s.svc.CreateTransaction(ctx, walletparam.CreateTransactionRequest{
		User:   req.User,
		Type:   transactionType,
		Amount: float64(req.Amount),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s Server) Start() {
	address := fmt.Sprintf(":%d", s.conf.WalletGrpcServer.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	log.Print("presence grpc server started on ", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("couldn't server presence grpc server", err)
	}
	log.Printf("presence grpc server started on %d", address)
}
