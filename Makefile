setup-init:
	brew install protobuf

setup:
	go mod tidy
	# dev tools
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	# db
	go get -u gopkg.in/gorp.v1
	go get -u github.com/go-sql-driver/mysql
	# grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/envoyproxy/protoc-gen-validate
	go get -u google.golang.org/grpc
	# others
	go get -u gopkg.in/yaml.v2
	go get -u github.com/sirupsen/logrus
	go get -u github.com/PuerkitoBio/goquery

lint:
	golangci-lint run --golint.min-confidence 1.1
fmt:
	goimports -e -d -local github.com server
# gen:
# 	sh openapi/generator.sh
# 	goimports -w web
# 	cp -r openapi-web/openapi/out/go web/openapi
# 	rm -rf openapi-web/openapi/out
# run-server:
# 	go run openapi-web/main.go
proto:
	protoc -I ./server/proto ./server/proto/*.proto \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:./server/gen/go \
		--go_out=plugins=grpc,paths=source_relative:./server/gen/go \
		--validate_out="lang=go,paths=source_relative:./server/gen/go"
