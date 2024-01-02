export tag=v1.3-metrics
export registry_server="registry.cn-shanghai.aliyuncs.com/luyanfei"

build:
	echo "build httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t ${registry_server}/httpserver:${tag} .

push: release
	echo "pushing registry.cn-shanghai.aliyuncs.com/luyanfei/httpserver"
	docker push ${registry_server}/httpserver:${tag}
