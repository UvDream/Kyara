all: build deploy  clean
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .
deploy:
	scp -r server root@124.223.8.237:/server
#deploy_config:
#	scp  config.yml root@124.223.8.237:/server
clean:
	rm -rf server
