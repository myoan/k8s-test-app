FROM golang:1.12.7 AS build
ADD . /go
WORKDIR /go
ENV GOPATH=
RUN CGO_ENABLED=0 go build -o app main.go

FROM busybox
COPY --from=build /go/app /go/server/app
ENTRYPOINT ["/go/server/app"]