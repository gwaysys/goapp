# Instruction

This project is a template for building a new go project with go module management, using `. env.sh` or `source env.sh` for setting the project environment variable.

Have been test with debian, centos, mac os.

# Dependency
go1.13 or later

# Building a new project
``` text
# clone the template to a new project
git clone https://github.com/gwaysys/goapp.git --single-branch --branch master --depth=1 test
cd test
./init.sh  
. env.sh # Or source env.sh
mkdir -p cmd/app
cd cmd/app
echo '
    package main
    
    import (
    	"fmt"
    	"os"
    	"os/signal"
    )
    
    func main() {
    	fmt.Println("OK")
    
    	fmt.Println("[ctrl+c to exit]")
    	end := make(chan os.Signal, 2)
    	signal.Notify(end, os.Interrupt, os.Kill)
    	<-end
    }
'>main.go

gofmt -w .
go build # Or sup build
./app

# ctrl+c
```
    
# File structure
``` text
$GOROOT -- Root path of golang。
$PRJ_ROOT -- root path of prject, it will auto set in env.sh
    .gitignore -- gitignore file
    dbuild.sh -- deployment shell for building docker image
    drun.sh   -- shell for running docker image 
    env.sh -- environment variable of project
    go.mod -- go module
    var -- some temp files of this project ran.
    etc -- etc files resource. 
    publish -- publish files when called ./publish.sh
```

# Deployment 

Two ways for built-in deployment
```
1, supd
2, docker image 
```

## Deployment with supd
### Install supd
See [supd](https://github.com/gwaysys/supd)

### Building binary for supd

Install the binary with source code
```shell
. env.sh
sup build all # $BUILD_ALL_PATH paths

sup install all # install all binary of $BUILD_ALL_PATH in env.sh
sup status # running status
sup clean all # remove the binary installed
```

Publish project binary
```
. env.sh
./publish.sh # the binary copy is output in publish directory.

cd publish/$PRJ_NAME # Or copy the release file to the deployment server.
. env.sh
sup install all # install all binary of $BUILD_ALL_PATH in env.sh 
sup status # running status
sup clean all # remove the binary installed
```

## Building docker image
```shell
sudo aptitude install docker.io
. env.sh
./dbuild.sh
```

## Online debug mode
```shell
cd $PRJ_ROOT/cmd/app
sup stop # stop the program in supd.
./app # run the the program in the console.
```

## Some improved commands of supc
```text
sup tail [$cfg_name stdout] # sudo supd ctl tail $cfg_name stdout
sup tail $cfg_name stderr # sudo supd ctl tail $cfg_name stderr
sup tailf [$cfg_name stdout] # sudo supd ctl tail -f $cfg_name stdout
sup tailf $cfg_name stderr # sudo supd ctl tail -f $cfg_name stderr

# More logs see $PRJ_ROOT/var/log/ 
# Or see sup help
```
