#!/bin/bash

export GO111MODULE=on
go_root="/usr/local/go"
if [ -d "$go_root" ]; then
    export GOROOT="$go_root"
fi
# 将GOBIN加入PATH
bin_path=$GOROOT/bin:

rep=${PATH/bin_path/""}
if [ ! "$PATH" = "$rep" ]; then
    PATH=$rep # 重新设定原值的位置
fi
export PATH=$bin_path$PATH

echo -n '#' "Your Project Name : "
read project_name

cat env.sh|awk '{gsub("\"goapp\"","\"'$project_name'\"");print $0}'>tmp.sh
cp tmp.sh env.sh
rm tmp.sh

echo '#' "Init Done"
echo '#'
echo '#' "Using '. env.sh' to loading environment of project."
echo '#' "Then using 'sup help' for building help."
echo '#'

# This example shows how to prompt for user's input.
echo -n '#' "Clean template data?[Y/n]"
read ANS
case $ANS in
    Y)
        rm -rf .git
        rm -rf cmd/*
        rm -rf module/*
        rm -rf version
        rm go.mod
        rm go.sum
        rm init.sh
        go mod init $project_name
        exit 0
        ;;
esac

