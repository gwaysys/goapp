# /bin/sh

PRJ_ROOT=$PRJ_ROOT
PRJ_NAME=$PRJ_NAME

# debug run
sudo docker run -it --rm \
    -v $PRJ_ROOT/etc:/app/etc \
    -v $PRJ_ROOT/var/log:/app/var/log \
    --net=host \
    $PRJ_NAME

# # for release
# PRJ_ROOT=`pwd`
# PRJ_NAME="goapp"
# ver=0.1
#
# # Publish docker: https://help.aliyun.com/document_detail/51810.html?spm=5176.11065259.1996646101.searchclickresult.260e6e5cpqIABR
# # Docker build：https://blog.csdn.net/boonya/article/details/74906927
# # sudo docker login 
# # sudo docker tag [imageid] gwaypg/$PRJ_NAME:$ver
# # sudo docker push gwaypg/$PRJ_NAME:$ver
# 
# sudo docker pull gwaypg/$PRJ_NAME:$ver
# 
# sudo docker run -d --restart=always \
#     -v /etc/localtime:/etc/localtime:ro \
#     -v $pwd_dir/etc:/app/etc \
#     -v $pwd_dir/var/log:/app/var/log \
#     -e PRJ_ROOT=/app \
#     -e GIN_MODE=release \
#     -w /app/cmd/app \
#     --name $PRJ_NAME.cmd.app \
#     --net=host \
#     gwaypg/$PRJ_NAME:$ver \
#     ./app
# 
# sudo docker run -d --restart=always \
#     -v /etc/localtime:/etc/localtime:ro \
#     -v $pwd_dir/etc:/app/etc \
#     -v $pwd_dir/var/log:/app/var/log \
#     -e PRJ_ROOT=/app \
#     -e GIN_MODE=release \
#     -w /app/src/applet/web \
#     --name $PRJ_NAME.cmd.web \
#     --net=host \
#     gwaypg/$PRJ_NAME:$ver \
#     ./web

