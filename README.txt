
Classic installation
	install mongodb on server,
	https_proxy=proxy:8080 go get gopkg.in/mgo.v2
	https_proxy=proxy:8080 go get github.com/gorilla/mux
	cd ..src/github.com/pascallimeux/urmmongo/server
	./build_urmmongo.sh
	cd ../dist
	./start_urmmongo.sh
	./stop_urmmongo.sh

Installation with docker
	cd ..src/github.com/pascallimeux/urmmongo/server
	./build_urmmongo.sh
	cd ..
	docker-compose build
	cd ../dist
	./start_docker.sh
	./stop_docker.sh


