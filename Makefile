generate:
	git submodule update --init --recursive
	go install github.com/bufbuild/buf/cmd/buf@latest
	cd schemas/protobuf && buf generate
