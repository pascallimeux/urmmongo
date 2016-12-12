FROM golang:latest

# configure proxy 
ENV http_proxy=http://10.193.21.110:8080
ENV https_proxy=http://10.193.21.110:8080
ENV use_proxy=on

# Install our dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
#RUN go get github.com/pascallimeux/urmmongo

# Install this src code
RUN cp -R ../urmmongo /go/src/github.com/pascallimeux/

# Install binary and configurate application within container 
RUN go install github.com/pascallimeux/urmmongo/server
RUN cp /go/src/github.com/pascallimeux/urmmongo/server/config/config4docker.json /go/bin/config.json
RUN mkdir /var/log/mhealth-urm-mongo/

# Set binary as entrypoint
ENTRYPOINT /go/bin/server /go/bin/server/config.json

# Expose default port (3000)
EXPOSE 8088 