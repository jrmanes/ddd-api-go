FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/jrmanes/ddd-api-go
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/jrmanes cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/jrmanes /go/bin/jrmanes
ENTRYPOINT ["/go/bin/jrmanes"]