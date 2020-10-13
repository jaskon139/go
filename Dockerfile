FROM golang:1.10.0

WORKDIR /app
ADD . .

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

ENTRYPOINT ["app"]

