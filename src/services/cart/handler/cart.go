/**
*FileName: handler
*Create on 2018/12/1 上午4:57
*Create by mok
*/

package handler

import (
	proto "MokShop/src/share/proto/cart"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"MokShop/src/services/cart/db"
	"github.com/bwmarrin/snowflake"
	"time"
)

type CartHandler struct {

}
/*
Add(context.Context, *AddReq) (*AddResp, error)
	Del(context.Context, *DelReq) (*DelResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	DelAll(context.Context, *DelAllReq) (*DelAllResp, error)
	Display(context.Context, *DisplayReq) (*DisplayResp, error)
*/

func(h * CartHandler)Add(ctx context.Context, req *proto.AddReq) (*proto.AddResp, error){
	if req.Price < 0 || req.Quantity <= 0{
		return nil,status.Error(codes.InvalidArgument,"请求参数错误")
	}
	node,err  := snowflake.NewNode(1)
	if err != nil{
		return nil,status.Error(codes.Internal,"服务发生故障了")
	}
	id := node.Generate()
	createTime := time.Now().Unix()
	modifyTime := time.Now().Unix()
	if err = db.AddCargo(req.UID,uint64(id),req.GoodID,req.GoodName,req.Price,req.Quantity,createTime,modifyTime);err != nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	return new(proto.AddResp),status.Error(codes.OK,"OK")
}


func(h *CartHandler)Del(ctx context.Context, req *proto.DelReq) (*proto.DelResp, error){
	if err := db.DelCargo(req.ItemID);err != nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	return new(proto.DelResp),status.Error(codes.OK,"OK")
}


func(h *CartHandler)DelAll(ctx context.Context, req *proto.DelAllReq) (*proto.DelAllResp, error){
	if err := db.DelAllCargo(req.UID);err != nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	return new(proto.DelAllResp),status.Error(codes.OK,"OK")
}


func(h *CartHandler)Display(ctx context.Context, req *proto.DisplayReq) (*proto.DisplayResp, error){
	items,err := db.Display(req.UID)
	if err != nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	var resp = &proto.DisplayResp{
		UID:req.UID,
		Items:items,
	}
	return resp,status.Error(codes.OK,"OK")
}

func(h *CartHandler)Update(ctx context.Context, req *proto.UpdateReq) (*proto.UpdateResp, error){
	if err := db.Update(req.ItemID,req.Quantity);err != nil{
		return nil,status.Error(codes.Internal,err.Error())
	}
	return new(proto.UpdateResp),status.Error(codes.OK,"OK")
}