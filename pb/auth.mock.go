// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

package grpc

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type V1AuthMock struct{}

func (m *V1AuthMock) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	res :=
		&LoginReply{
			Token: "qui",
		}
	return res, nil
}