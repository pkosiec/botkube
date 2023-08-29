package source

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SingleDispatchUnimplemented is used for plugins which doesn't implement HandleSingleDispatch method.
type SingleDispatchUnimplemented struct{}

// HandleSingleDispatch returns unimplemented error.
func (SingleDispatchUnimplemented) HandleSingleDispatch(context.Context, SingleDispatchInput) (SingleDispatchOutput, error) {
	return SingleDispatchOutput{}, status.Errorf(codes.Unimplemented, "method HandleSingleDispatch not implemented")
}

// StreamUnimplemented is used for plugins which doesn't implement Stream method.
type StreamUnimplemented struct{}

// Stream returns unimplemented error.
func (StreamUnimplemented) Stream(context.Context, StreamInput) (StreamOutput, error) {
	return StreamOutput{}, status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
