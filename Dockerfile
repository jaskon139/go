FROM golang:1.10.0

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

WORKDIR /app
ADD . .

RUN chmod +x /app/configure 
RUN apt update
RUN apt -y install curl git

ENTRYPOINT ["app"]

