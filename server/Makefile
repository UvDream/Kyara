all: build deploy
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
deploy:
	#scp -r /Users/wzj/Desktop/wzj/blog/Kyara/server/config.yaml root@124.223.8.237:/server
	scp -r /Users/wzj/Desktop/wzj/blog/Kyara/server/main root@124.223.8.237:/server
