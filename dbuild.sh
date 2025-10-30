# /bin/sh

# docker build configuration
echo \
'FROM alpine:3.15
MAINTAINER SHU <free1139@163.com>

RUN apk add --update ca-certificates
RUN mkdir -p /app/var/log/
COPY ./publish/'$PRJ_NAME' /app

CMD ["/app/bin/supd", "-c","/app/etc/supd.conf"]
'>Dockerfile


# build bin data
export CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go install -v github.com/gwaysys/supd/cmd/supd@v1.0.9||exit 1
# publish build
sup publish all

echo "# Building Dockerfile"
# remove old images
sudo docker rmi -f $PRJ_NAME||exit 1
# build images
sudo docker build -t $PRJ_NAME .||exit 1
# rm tmp data
# rm app

# show images build result
sudo docker images $PRJ_NAME||exit 1

# remove dockerfile
rm Dockerfile||exit 1
