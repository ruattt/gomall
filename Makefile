export ROOT_MOD=gomall_study
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module  ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/order_page.proto

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC  --service user --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto


.PHONY: gen-product
gen-product:
	@cd app/product && cwgo server --type RPC --service product --module  ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/product.proto
	@cd rpc_gen && cwgo client --type RPC --service product --module  ${ROOT_MOD}/rpc_gen  --I ../idl --idl ../idl/product.proto


.PHONY: gen-cart
gen-cart:
	@cd app/cart && cwgo server --type RPC --service cart --module  ${ROOT_MOD}/app/cart --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/cart.proto
	@cd rpc_gen && cwgo client --type RPC --service cart --module  ${ROOT_MOD}/rpc_gen  --I ../idl --idl ../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd app/payment && cwgo server --type RPC --service payment --module  ${ROOT_MOD}/app/payment --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/payment.proto
	@cd rpc_gen && cwgo client --type RPC --service payment --module  ${ROOT_MOD}/rpc_gen  --I ../idl --idl ../idl/payment.proto

.PHONY: gen-checkout
gen-checkout: 
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/checkout.proto

.PHONY: gen-order
gen-order: 
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/order.proto

.PHONY: gen-email
gen-email: 
	@cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/email.proto
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/email.proto
