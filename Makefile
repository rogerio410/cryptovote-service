.PHONY: compile

compile:
		protoc --proto_path=proto proto/**/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:pkg/pb