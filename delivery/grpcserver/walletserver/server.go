package presenceserver

import (
	"app/contract/protobuf/wallet"
	"app/param/walletparam"
	"app/service/walletservice"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	wallet.UnimplementedPresenceServiceServer
	svc walletservice.Service
}

func New(svc walletservice.Service) Server {
	return Server{
		UnimplementedPresenceServiceServer: wallet.UnimplementedPresenceServiceServer{},
		svc:                                svc,
	}
}
func (s Server) CreateTransaction(ctx context.Context, req *wallet.CreateTransactionRequest) (*wallet.CreateTransactionResponse, error) {
	resp, err := s.svc.CreateTransaction(ctx, walletparam.CreateTransactionRequest{
		User:   req.User,
		Type:   req.Type,
		Amount: float64(req.Amount),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s Server) Start() {
	address := fmt.Sprintf(":%d", 8086)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("presence grpc server started on %d", address)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("couldn't server presence grpc server", err)
	}
	log.Printf("presence grpc server started on %d", address)
}
