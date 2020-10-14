FROM debian:buster

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

WORKDIR /app
ADD . .

RUN chmod +x /app/configure 
RUN apt update
RUN apt -y install curl git bash
RUN git clone https://github.com/jaskon139/oktetofortest.git
RUN chmod +x ./oktetofortest/sswork/*

ENTRYPOINT ["app"]

