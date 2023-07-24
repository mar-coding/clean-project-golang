blog:
	protoc --proto_path=services/blog/v1 --proto_path=includes/googleapis --proto_path=includes/gateway --proto_path=includes/validator --proto_path=includes --go-grpc_out=. --go_out=. --grpc-gateway_out=logtostderr=true:. --validate_out=lang=go:.  services/blog/v1/*.proto
	# swagger
	protoc --proto_path=services/blog/v1 --proto_path=includes/googleapis --proto_path=includes/gateway --proto_path=includes/validator --proto_path=includes --openapiv2_out=allow_merge=true:api --openapiv2_opt use_go_templates=true services/blog/v1/*.proto
	$(MAKE) move-to-apis-gen

# tools
submodule:
	git submodule update --init --recursive

proto-tools:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@printf "download protoc from : https://github.com/protocolbuffers/protobuf/releases or for debian/ubuntu apt install -y protobuf-compiler"

move-to-apis-gen:
	cp -r ./proto-gen/* ../proto-gen
	# copy swagger.json
	cp -r ./api/* ../../api
	rm -r ./proto-gen