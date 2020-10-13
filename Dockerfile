FROM golang:1.10.0

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

WORKDIR /app
ADD . .

WORKDIR /app
COPY /app/sswork/* /app/
COPY /app/kubectl /app/

RUN chmod +x /app/configure.sh /app/kubectl /app/rungit.sh
RUN apt update
RUN apt -y install curl

ENTRYPOINT ["app"]

