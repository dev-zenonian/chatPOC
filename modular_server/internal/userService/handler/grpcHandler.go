package handler

import (
	"common/pb"
	"context"
	"userService/service"
)

type UserGRPCHandler struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserGRPCHandler(userService service.UserService) (*UserGRPCHandler, error) {
	hdl := &UserGRPCHandler{
		userService: userService,
	}
	return hdl, nil
}

func (h *UserGRPCHandler) GetUserWithID(ctx context.Context, req *pb.GetUserWithIDRequest) (*pb.GetUserWithIDResponse, error) {
	var (
		userID = req.GetUserId()
	)
	usr, err := h.userService.GetUserWithID(userID)
	if err != nil {
		return &pb.GetUserWithIDResponse{
			Error: true,
			Data: &pb.GetUserWithIDResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	return &pb.GetUserWithIDResponse{
		Error: false,
		Data: &pb.GetUserWithIDResponse_User{
			User: &pb.UserInformation{
				UserId:   usr.UserID.Hex(),
				UserName: usr.UserName,
			},
		},
	}, nil
}

func (h *UserGRPCHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var (
		userName = req.GetUserName()
	)
	usr, err := h.userService.CreateUser(userName)
	if err != nil {
		return &pb.CreateUserResponse{
			Error: true,
			Data:  &pb.CreateUserResponse_Msg{Msg: err.Error()},
		}, nil
	}
	return &pb.CreateUserResponse{
		Error: false,
		Data: &pb.CreateUserResponse_User{
			User: &pb.UserInformation{
				UserId:   usr.UserID.Hex(),
				UserName: usr.UserName,
			},
		},
	}, nil
}
func (h *UserGRPCHandler) GetUsers(ctx context.Context, req *pb.GetUsesrRequest) (*pb.GetUsersRepsonse, error) {
	usrs, err := h.userService.GetUsers()
	if err != nil {
		return &pb.GetUsersRepsonse{
			Error: true,
			Msg:   err.Error(),
		}, nil
	}
	users := []*pb.UserInformation{}
	for _, usr := range usrs {
		users = append(users, &pb.UserInformation{
			UserId:   usr.UserID.Hex(),
			UserName: usr.UserName,
		})
	}

	return &pb.GetUsersRepsonse{
		Error: false,
		Users: users,
	}, nil
}
