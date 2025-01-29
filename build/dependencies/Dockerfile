FROM golang:1.19.13 AS build

WORKDIR /build

RUN mkdir -p src

COPY go.mod .
COPY src/ src/

RUN GOCACHE=OFF
RUN go mod tidy
RUN go mod download
RUN go install github.com/gopherjs/gopherjs@v1.19.0-beta1

CMD ["sh","-c", \
    "GOOS=js GOARCH=ecmascript gopherjs build -o ./bin/client/go.js ./src/client/   \
    && CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o ./bin/server/go.wasm ./src/server/   \
    && sed -i '1ivar require = null\n' ./bin/client/go.js \
    && rm ./bin/client/go.js.map " ]