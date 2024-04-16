package walletadapter

import (
	"app/contract/goproto/wallet"
	"app/param/walletparam"
	"app/pkg/protobufmapper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	address string
}

func New(address string) Client {
	return Client{
		address: address,
	}
}

func (c Client) CreateTransaction(ctx context.Context, request walletparam.CreateTransactionRequest) (walletparam.CreateTransactionResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return walletparam.CreateTransactionResponse{}, err
	}
	defer conn.Close()

	client := wallet.NewCreateTransactionServiceClient(conn)

	resp, err := client.CreateTransaction(ctx,
		&wallet.CreateTransactionRequest{
			User:   request.User,
			Type:   protobufmapper.MapTypeWalletParamToProtobuf(request.Type),
			Amount: float32(request.Amount),
		})
	if err != nil {
		return walletparam.CreateTransactionResponse{}, err
	}

	return protobufmapper.MapProtobufToResponseParam(resp), nil
}
