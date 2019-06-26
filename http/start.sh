CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build   -o demo .

docker build -t demo:latest .

docker tag demo:latest jessewu/demo:v1.0.0 # 增加一个标准格式的镜像, 请更换为自己的仓库名字

docker run  -p 9090:9090 jessewu/demo:v1.0.0 # 测试你的镜像没有问题

docker login

docker push jessewu/demo:v1.0.0
