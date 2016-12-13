
## Installation on host

	- Install mongodb on server

	- Update configuration: 
		Modify IP address for server in $GOPATH/src/github.com/pascallimeux/urmmongo/server/config/config.json

	- Update script: 
		Modify account username in $GOPATH/src/github.com/pascallimeux/urmmongo/scripts/start_urmmongo.sh

	- Install dependencies with or without proxy :
		https_proxy=proxy:8080 go get gopkg.in/mgo.v2
		https_proxy=proxy:8080 go get github.com/gorilla/mux

	- Compile binary
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/server
		./build_urmmongo.sh

	- Start component
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/dist
		./start_urmmongo.sh

	- Stop component
		./stop_urmmongo.sh


## Installation in docker container
	
	- Update Dockerfile
		setup or remove proxy in this file:  $GOPATH/src/github.com/pascallimeux/urmmongo/Dockerfile

	- Compile binary
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/server
		./build_urmmongo.sh

	- Build container
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/dist
		./build_docker.sh

	- Start container
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/dist
		./start_docker.sh

	- Stop container
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/dist
		./stop_docker.sh

	- Deploy autostart service
		check path in urmmongo.service
		cd $GOPATH/src/github.com/pascallimeux/urmmongo/dist
		./deploy_service.sh
