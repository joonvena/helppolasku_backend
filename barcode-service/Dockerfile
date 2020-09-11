FROM golang:1.14-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
 
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /build/app .

FROM scratch
COPY --from=build /dist/app /

ENTRYPOINT ["/app"]

