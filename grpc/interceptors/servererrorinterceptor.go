package interceptors

import (
	"context"
	"github.com/metaitself/xmeta/metaerror"
	"google.golang.org/grpc"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			err = metaerror.FromError(err).GRPCStatus().Err()
		}
		return resp, err
	}
}
