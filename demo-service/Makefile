build:
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-service .

run:
	docker run -d -p 9091:9999 -e MICRO_SERVER_ADDRESS=:9091 laracom-demo-service