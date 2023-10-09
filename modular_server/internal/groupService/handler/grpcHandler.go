package handler

import (
	"common/pb"
	"context"
	"groupService/service"
	"log"
)

type GroupGRPCHandler struct {
	groupService service.GroupService
	pb.UnimplementedGroupServiceServer
}

func NewGroupGRPCHandler(groupService service.GroupService) (*GroupGRPCHandler, error) {
	grpcHandler := &GroupGRPCHandler{
		groupService: groupService,
	}
	return grpcHandler, nil
}

func (h *GroupGRPCHandler) GetGroupWithID(ctx context.Context, in *pb.GetGroupWithIDRequest) (*pb.GetGroupWithIDResponse, error) {
	groupID := in.GetGroupId()
	log.Println(in)
	if groupID == "" {
		return &pb.GetGroupWithIDResponse{
			Error: true,
			Data: &pb.GetGroupWithIDResponse_Msg{
				Msg: "groupID must not be null",
			},
		}, nil
	}
	group, err := h.groupService.GetGroupWithID(groupID)
	if err != nil {
		return &pb.GetGroupWithIDResponse{
			Error: true,
			Data: &pb.GetGroupWithIDResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	clients := []string{}
	for _, client := range group.ClientsID {
		clients = append(clients, client)
	}
	return &pb.GetGroupWithIDResponse{
		Error: false,
		Data: &pb.GetGroupWithIDResponse_Group{
			Group: &pb.GroupInformation{
				GroupId:   groupID,
				GroupName: group.Name,
				ClientIDs: clients,
			},
		},
	}, nil
}

func (h *GroupGRPCHandler) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	var (
		clientIDs = req.GetClientIDs()
		adminID   = req.GetAdminID()
		groupName = req.GetGroupName()
		isPrivate = req.GetIsPrivate()
	)
	if groupName == "" {
		return &pb.CreateGroupResponse{
			Error: true,
			Data: &pb.CreateGroupResponse_Msg{
				Msg: "Group name must not be null",
			},
		}, nil
	}
	group, err := h.groupService.CreateGroup(groupName, adminID, clientIDs, isPrivate)
	if err != nil {
		return &pb.CreateGroupResponse{
			Error: true,
			Data: &pb.CreateGroupResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	clients := []string{}
	for _, client := range group.ClientsID {
		clients = append(clients, client)
	}
	rsp := &pb.CreateGroupResponse{
		Error: false,
		Data: &pb.CreateGroupResponse_Group{
			Group: &pb.GroupInformation{
				GroupId:   group.GroupID.Hex(),
				GroupName: group.Name,
				ClientIDs: clients,
			}}}
	return rsp, nil
}

func (h *GroupGRPCHandler) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	var (
		groupID = req.GetGroupID()
		adminID = req.GetAdminID()
	)
	if err := h.groupService.DeleteGroup(groupID, adminID); err != nil {
		rsp := &pb.DeleteGroupResponse{
			Error: true,
			Data:  err.Error(),
		}
		return rsp, nil
	}
	return &pb.DeleteGroupResponse{
		Error: false,
		Data:  "Deleted",
	}, nil
}
