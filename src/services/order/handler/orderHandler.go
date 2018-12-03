/**
*FileName: handler
*Create on 2018/12/3 上午12:03
*Create by mok
*/

package handler

import (
	proto "MokShop/src/share/proto/order"
	"context"
	"MokShop/src/services/order/model"
	userproto"MokShop/src/share/proto/user"
	"google.golang.org/grpc"
	"time"
	"sync"

	"MokShop/src/services/order/db"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)



type OrderHandler struct {

}


//新增订单
func(h *OrderHandler)Commit(ctx context.Context, req *proto.CommitReq) (reply *proto.CommitResp, err error){
	var wg sync.WaitGroup
	wg.Add(2)
	//获取商品库存
	go func() {
		defer wg.Done()
		conn,err := grpc.Dial("192.168.50.250:9001",grpc.WithInsecure(),grpc.WithBackoffConfig(grpc.BackoffConfig{MaxDelay:3*time.Second}))
		if err != nil{
			return
		}
		client := userproto.NewCargoServiceClient(conn)
		//todo:这里应该去请求Get方法，获取到用户信息
		userinfo,err = client.GetCargoInfo(context.TODO(),&userproto.LoginReq{},grpc.FailFast(true))
		if err != nil{
			return
		}
	}()
	//获取优惠券信息
	go func() {
		defer wg.Done()
		//todo:获取优惠券信息
	}()
	wg.Wait()
	//处理数据
	userinfo := &model.UserInfo{req.UID,req.Username,req.Email,req.Phone,req.Addr}
	order := &model.OrderDetail{}
	order.UserInfo = userinfo
	order.GoodID = req.GoodID
	order.GoodName = req.GoodName
	order.Price = req.Price
	order.Cost = req.Cost
	order.CreateTime = time.Now().Unix()
	order.Coupons = coupons
	order.OrderID = orderID  //利用雪花算法生成ID

	//将order数据插入数据库中
	err = db.InsertOrder(order)
	if err != nil{
		return nil,status.Error(codes.Internal,"添加订单失败")
	}

	//生成快递信息
	go func() {
		//todo:生成快递信息，利用nsq将信息传送到快递微服务
	}()

	//信息发送
	go func() {
		//todo:将信息发送利用nsq到信息微服务
	}()
	reply.Status = 0
	reply.OrderID = order.OrderID
	err = status.Error(codes.OK,"添加订单成功，我们已经将订单信息发送到你的邮箱中")
	return
}

//订单支付
func(h *OrderHandler)Pay(ctx context.Context, req *proto.PayReq) (resp *proto.PayResp, err error){
	//查询该订单是否存在以及该订单状态是否已经支付


	//调用支付宝接口实现订单支付


	//将支付信息存入缓存或者redis中，轮询然后添加到数据库中


	//订单支付信息发送到用户邮箱

	return
}

//订单信息详情
func(h *OrderHandler)Display(ctx context.Context, *proto.DisplayReq) (resp *proto.DisplayResp, err error){
	//获取订单信息
	return
}


//取消订单
func(h *OrderHandler)Del(ctx context.Context, req *proto.DelReq) (resp *proto.DelResp, err error){
	//删除订单信息

	//发送删除信息到快递微服务

	//发送删除信息到信息微服务
	return
}