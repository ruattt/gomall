package service

import (
	"context"
	"gomall_study/app/checkout/infra/mq"
	"gomall_study/app/checkout/infra/rpc"
	"gomall_study/rpc_gen/kitex_gen/cart"
	checkout "gomall_study/rpc_gen/kitex_gen/checkout"
	"gomall_study/rpc_gen/kitex_gen/email"
	"gomall_study/rpc_gen/kitex_gen/order"
	"gomall_study/rpc_gen/kitex_gen/payment"
	"gomall_study/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
	// "github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartRedsult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
        return nil, kerrors.NewBizStatusError(5005001, err.Error())
	}
	if cartRedsult == nil || cartRedsult.Items == nil {
		return nil, kerrors.NewBizStatusError(5004001, "Cart is empty")
	}
	var (
		total float32
		oi    []*order.OrderItem
	)

	for _, cartItem := range cartRedsult.Items {
	    productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
            return nil, resultErr
        }

		if productResp.Product == nil  {
			continue
		}

		p := productResp.Product.Price

		cost := p * float32(cartItem.Quantity)
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	// create order

	// &order.PlaceOrderReq{
	// 	UserId: req.UserId,
	// 	Email:  req.Email,
	// 	Address: &order.Address{
	// 		StreetAddress: req.Address.StreetAddress,
	// 		City:          req.Address.City,
	// 		State:         req.Address.State,
	// 		Country:       req.Address.Country,
	// 		ZipCode:       req.Address.ZipCode,
	// 	},
	// 	Items: oi,
	// }
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		Email:        req.Email,
		Items: oi,
	}
	if req.Address != nil {
		addr := req.Address
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       addr.ZipCode,
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}
	klog.Info("orderResult", orderResult)


	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
        klog.Error(err.Error())
    }
	klog.Info(emptyResult)

	// charge
	var orderId string
	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:   req.UserId,
		Amount:  total,
        OrderId: orderId,
        CreditCard: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
            CreditCardCvv:    req.CreditCard.CreditCardCvv,
            CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
        },
	}
	paymentResult, er := rpc.PaymentClient.Charge(s.ctx, payReq)
	if er != nil {
        return nil, er
    }
	klog.Info(paymentResult)

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)

	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
