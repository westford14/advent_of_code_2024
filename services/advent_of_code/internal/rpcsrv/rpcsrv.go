package rpcsrv

import (
	"context"
	"log/slog"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	day_v1 "github.com/westford14/advent_of_code_2024/pkg/gen/westford14/advent_of_code/v1"
)

type AdventServer struct {
	day_v1.UnimplementedAdventServiceServer
}

type AdventServerConfig struct{}

func New(ctx context.Context, opts AdventServerConfig) (*AdventServer, error) {
	return &AdventServer{}, nil
}

func (s *AdventServer) Day(ctx context.Context, req *day_v1.DayRequest) (*day_v1.DayResponse, error) {

	if err := protovalidate.Validate(req); err != nil {
		validateErr, ok := err.(*protovalidate.ValidationError)
		if !ok {
			slog.ErrorContext(ctx, "error validating DayRequest", err.Error())
			return nil, status.Errorf(codes.Internal, "Unknown error occurred")
		}
		return nil, status.Errorf(codes.InvalidArgument, validateErr.Error())
	}
	return nil, status.Errorf(codes.Unimplemented, "method Day not implemented")
}
