# Example
http://1.117.145.187/blog
# Goblog deploy steps:
### 1. upload the docker-compose-no-build.yml file to server
`scp docker-compose-no-build.yml root@x.x.x.x:/root`
### 2. create mysql data dir
`mkdir -p /var/lib/mysql`
### 3. create redis data dir
`mkdir -p /var/lib/data`
### 4. install docker and docker-compose
```
yum install docker -y
yum install docker-compose -y
```
### 5. start docker
`systemctl start docker`
### 6. run goblog
`docker-compose -f docker-compose-no-build.yml up -d`

### 7. set acl user admin
```
$ docker exec -it root_redis_1 redis-cli
> acl setuser admin on >admin ~* +@all
```

# Update image
### 1. build image
`docker build -t beslow/goblog .`
### 2. push to docker.io, need docker login
`docker push beslow/goblog`
### 3. update image locally
`docker pull beslow/goblog:latest`
