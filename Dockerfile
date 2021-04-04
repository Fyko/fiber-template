FROM golang:1.16 AS build

WORKDIR /go/src/github.com/fyko/fiber-template/ 
COPY go.mod go.sum ./
RUN go mod download
RUN go run github.com/prisma/prisma-client-go prefetch
COPY . .
RUN go run github.com/prisma/prisma-client-go generate
RUN go build -o build/server cmd/server/main.go   

FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /go/src/github.com/fyko/fiber-template/app
COPY --from=build /go/src/github.com/fyko/fiber-template/build/server .

CMD ["./server"]