package service

import (
	"context"
	"gomall_study/app/checkout/infra/rpc"
	"gomall_study/rpc_gen/kitex_gen/cart"
	checkout "gomall_study/rpc_gen/kitex_gen/checkout"
	"gomall_study/rpc_gen/kitex_gen/payment"
	"gomall_study/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
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
	var total float32

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
	}
	var orderId string

	u, _ := uuid.NewRandom()
	orderId = u.String()

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

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})

	if err != nil {
        klog.Error(err.Error())
    }

	paymentResult, er := rpc.PaymentClient.Charge(s.ctx, payReq)
	if er != nil {
        return nil, er
    }
	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
