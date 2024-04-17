package walletadapter

import (
	"app/contract/goproto/wallet"
	"app/pkg/protobufmapper"
	"context"
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

func (c Client) CreateTransaction(ctx context.Context, user string, typeT string, amount float64, meta string) error {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := wallet.NewCreateTransactionServiceClient(conn)

	_, eErr := client.CreateTransaction(ctx,
		&wallet.CreateTransactionRequest{
			User:   user,
			Type:   protobufmapper.MapTypeStringToProtobuf(typeT),
			Amount: float32(amount),
			Meta:   meta,
		})
	if eErr != nil {
		return eErr
	}

	return nil
}

//func (c Client) CheckUserUsageDiscount(ctx context.Context, discount string, user string) bool {
//
//}
//func CheckUsageCountDiscountCode(ctx context.Context, code string) uint64 {
//
//}
