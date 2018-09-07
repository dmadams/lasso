FROM golang:1.8

RUN mkdir -p ${GOPATH}/src/github.com/dmadams/lasso
WORKDIR ${GOPATH}/src/github.com/dmadams/lasso
    
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN rm -rf ./config ./data \
    && ln -s /config ./config \
    && ln -s /data ./data 

EXPOSE 9090
CMD ["/go/bin/lasso"] 
