// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pkg/pb/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	// Without author
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...client.CallOption) (*ForgetPasswordResponse, error)
	// With author
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*DeleteUserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...client.CallOption) (*ForgetPasswordResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.ForgetPassword", in)
	out := new(ForgetPasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*UpdateUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateUser", in)
	out := new(UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*DeleteUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteUser", in)
	out := new(DeleteUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	// Without author
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	ForgetPassword(context.Context, *ForgetPasswordRequest, *ForgetPasswordResponse) error
	// With author
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	UpdateUser(context.Context, *UpdateUserRequest, *UpdateUserResponse) error
	DeleteUser(context.Context, *DeleteUserRequest, *DeleteUserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, out *ForgetPasswordResponse) error
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error
		DeleteUser(ctx context.Context, in *DeleteUserRequest, out *DeleteUserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *userServiceHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *userServiceHandler) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, out *ForgetPasswordResponse) error {
	return h.UserServiceHandler.ForgetPassword(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

func (h *userServiceHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *UpdateUserResponse) error {
	return h.UserServiceHandler.UpdateUser(ctx, in, out)
}

func (h *userServiceHandler) DeleteUser(ctx context.Context, in *DeleteUserRequest, out *DeleteUserResponse) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}