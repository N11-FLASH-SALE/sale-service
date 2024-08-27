CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
	
run-service:
	go run cmd/main.go
