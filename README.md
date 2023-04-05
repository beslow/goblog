[![go-test](https://github.com/beslow/goblog/actions/workflows/go-test.yml/badge.svg)](https://github.com/beslow/goblog/actions/workflows/go-test.yml) [![Deployment](https://github.com/beslow/goblog/actions/workflows/CD.yml/badge.svg?branch=main)](https://github.com/beslow/goblog/actions/workflows/CD.yml)

# Example
http://1.117.145.187/blog
# Goblog deploy steps:
### 1. upload the docker-compose-no-build.yml file to server
`scp docker-compose-no-build.yml root@x.x.x.x:/root`
### 2. create mysql data dir
`mkdir -p /var/lib/mysql`
### 3. create redis data dir
`mkdir -p /var/lib/data`
### 4. create rocketmq related dir
```
mkdir -p /var/lib/rocketmq/data/namesrv/logs
mkdir -p /var/lib/rocketmq/data/namesrv/store
mkdir -p /var/lib/rocketmq/data/broker/logs
mkdir -p /var/lib/rocketmq/rocketmq/data/broker/store
mkdir -p /var/lib/rocketmq/conf
```

### 5. create broker.conf file
`mkdir -p /var/lib/rocketmq/conf/broker.conf`

with content
```
brokerClusterName = DefaultCluster
brokerName = broker-a
brokerId = 0
deleteWhen = 04
fileReservedTime = 48
brokerRole = ASYNC_MASTER
flushDiskType = ASYNC_FLUSH
namesrvAddr = namesrv:9876
#brokerIP1 = 172.17.0.4
brokerIP1=broker-a
```
### 6. install docker and docker-compose
```
yum install docker -y
yum install docker-compose -y
```
### 7. start docker
`systemctl start docker`
### 8. run goblog
`docker-compose -f docker-compose-no-build.yml up -d`

# Update image
### 1. build image
`docker build -t beslow/goblog .`
### 2. push to docker.io, need docker login
`docker push beslow/goblog`
### 3. update image locally
`docker pull beslow/goblog:latest`

# Test
### 1. create test database
`CONFIG_DIR=test go run main.go -db:create`
### 2. migrate test database
`CONFIG_DIR=test go run main.go -db:migrate`
### 3. seed test database
`CONFIG_DIR=test go run main.go -db:seed`
### 4. generate assets
`go generate test/*.go`
### 5. run test
`go test ./...`