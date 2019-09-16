FROM golang:1.12.0-alpine3.9

# RUN GOROOT=/usr/local/go
# Run GOPATH=$HOME/Projects/Proj1
# RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN mkdir /usr/local/go/src/api

ADD . /api

WORKDIR /api

RUN go run main.go 


CMD ["/api/main"]