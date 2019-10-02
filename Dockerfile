# builder image
FROM golang:1.13-alpine3.10 as builder

ENV CGO_ENABLED 0
ENV GO111MODULE on
RUN apk --no-cache add git
WORKDIR /go/src/github.com/linki/wanchain-cli
COPY . .
RUN go test -v ./...
RUN go build -o /bin/wanchain-exporter -v \
  -ldflags "-X github.com/linki/wanchain-exporter/cmd.version=$(git describe --tags --always --dirty) -w -s" \
  ./cmd/ethereum_exporter

# final image
FROM alpine:3.10

RUN apk --no-cache add ca-certificates dumb-init
COPY --from=builder /bin/wanchain-exporter /bin/wanchain-exporter

USER 65534
ENTRYPOINT ["dumb-init", "--", "/bin/wanchain-exporter"]
