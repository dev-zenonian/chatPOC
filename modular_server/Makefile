dockerup:
	cd _container && docker-compose up -d --remove-orphans

dockerdown:
	cd _container && docker-compose down 

dockerclean:
	cd _container && docker-compose down && sudo rm -rf _volumes/mongodb/* _volumes/kafka/data/* _volumes/zookeeper/data/* _volumes/zookeeper/log/* 

server:
	go run cmd/main.go

gen:
	protoc --proto_path=./internal/common/proto \
		--go_out=:./internal/common/pb \
		--go-grpc_out=:./internal/common/pb  \
		--grpc-gateway_out ./internal/common/pb \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		./internal/common/proto/*.proto 

copyProto:
	cp -r ./../proto/* ./internal/common/proto/

sync: copyProto gen

clean:
	rm ./internal/common/pb/*

groupMessageService:
	cd internal/groupMessageService && HTTP_PORT="8091" GRPC_PORT="8092" GRPC_ENDPOINT="127.0.0.1:8081" go run main.go

groupService:
	cd internal/groupService && HTTP_PORT="8080" GRPC_PORT="8081" GRPC_ENDPOINT="127.0.0.1:8085" go run main.go

wsHandler:
	cd internal/wsService && HTTP_PORT="8082" GRPC_PORT="8083" GRPC_ENDPOINT="127.0.0.1:8084" go run main.go

sessionService:
	cd internal/sessionService && HTTP_PORT="8089" GRPC_PORT="8084" go run main.go

userService:
	cd internal/userService && HTTP_PORT="8087" GRPC_PORT="8085" go run main.go 

messageService:
	cd internal/messageService && HTTP_PORT="8086" GRPC_PORT="8090" GRPC_ENDPOINT="127.0.0.1:8085" go run main.go

distributor:
	cd internal/distributorService && GRPC_PORT="8088"  GRPC_ENDPOINT="127.0.0.1:8084" go run main.go 

.PHONY: dockerup dockerdown dockerclean server gen copyProto sync clean groupService sessionService wsHandler  messageService userService distributor
