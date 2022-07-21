FROM golang:1.17-alpine

WORKDIR /go/src/github.com/DavidBagaryan/drone-attack

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go build -o /drone-attack ./cmd/main.go

CMD [ "/drone-attack" ]
