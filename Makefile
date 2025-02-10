export ROOT_MOD=gomall_study
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-frontend
gen-frontend:
	# @cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/home.proto

	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module  ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/auth_page.proto

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
	@cd rpc_gen && cwgo client --type RPC  --service user --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto

