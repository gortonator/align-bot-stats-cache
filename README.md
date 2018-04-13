# align-bot-stats-cache
Redis Cache Layer for the NEU ALIGN Chatbot

## Instructions to setup on a Fresh VM.
1. Download and Install [Redis](https://redis.io/topics/quickstart)
1. Install Go.
    1. Download [Go](https://golang.org/dl/) Archive.
    1. Extract it into `/usr/local` creating a Go tree in `/usr/local/go` (Typically these commands must be run as root or through sudo.) 
        - example `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
        - **NOTE:** To install to a custom location, refer documentaion on official Golang [website](https://golang.org/doc/install#install).
    1. Add `/usr/local/go/bin` to the PATH environment variable
        - example `export PATH=$PATH:/usr/local/go/bin`
1. Set-up the Go working directory structure.
    1. create a folder in your home directory to house all go source code.
        - It is preferable to create the folder in your home directory: `~/go` (where ~ is shortcut to your home directory)
    1. under that directory, create further three directories `~/go/src`, `~/go/pkg` and `~/go/bin`
    1. now create the following structure of directories `~/go/src/github.com/rapidclock` and go to that directory.
    1. clone the align-bot-stats-cache repo here. (Final resting place : `~/go/src/github.com/rapidclock/align-bot-stats-cache/`)
1. Install dependencies:
    - `go get -u github.com/buger/jsonparser`
    - `go get -u github.com/json-iterator/go`
    - `go get -u github.com/gomodule/redigo/redis`
    - `go get -u github.com/mailru/easyjson/...`
1. Install the align-bot-stats-cache Application.
    - run the command `go install`
1. Start the Redis Server using default .conf file.
    - run this command `redis-server`
1. Run the align-bot-stats-cache Application.
    - go to `~/go/bin` and run the output file called **align-bot-stats-cache** as `./align-bot-stats-cache`

**NOTE :** 
  - The application is running on localhost right now. The set port is 15000.
  - Remember to Start the Redis server before starting the applicaton. It requires the redis server to be running.
