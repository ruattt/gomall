.PHONY: gen-frontend
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module app/frontend -I ../../idl/