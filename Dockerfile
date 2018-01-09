FROM golang:1.9
WORKDIR /go/src/github.com/heptiolabs/namespace-deleter

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/heptiolabs/namespace-deleter

FROM alpine:latest
COPY --from=0 /go/bin/namespace-deleter /bin/namespace-deleter
