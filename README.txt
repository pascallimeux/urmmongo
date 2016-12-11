
dependencies:

https_proxy=proxy:8080 go get gopkg.in/mgo.v2
https_proxy=proxy:8080 go get github.com/gorilla/mux

go command line options
to build and compress binary:  	go build -ldflags "-s" main.go urmmongo
to test one function: 			go test -run "TestValueCreateNominal"
to lunch all tests:				go test