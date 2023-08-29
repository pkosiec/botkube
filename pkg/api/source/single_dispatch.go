package source

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SingleDispatchUnimplemented struct{}

func (SingleDispatchUnimplemented) HandleSingleDispatch(context.Context, SingleDispatchInput) (SingleDispatchOutput, error) {
	return SingleDispatchOutput{}, status.Errorf(codes.Unimplemented, "method HandleSingleDispatch not implemented")
}
