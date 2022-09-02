FROM golang:1.18-alpine

# allow go test to work properly
ENV CGO_ENABLED=0

ENV GOBIN=/usr/app/bin
WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go test -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["bin/1.0"]