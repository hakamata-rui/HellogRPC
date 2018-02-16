GLIDE_HOME=./.glide
glide=glide -home ${GLIDE_HOME}

.PHONY: install
install:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	curl https://glide.sh/get | sh

.PHONY: gen-code
gen-code:
	protoc --go_out=plugins=grpc,import_path=pb:./ ./pb/hello.proto

.PHONY: gen-doc
gen-doc:
	protoc --doc_out=./pb/ --doc_opt=html,index.html ./pb/hello.proto

.PHONY: run
run:
	go run server/main.go & 
	sleep 3
	go run client/main.go 
