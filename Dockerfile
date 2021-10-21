FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/yokoe/cloudrunexp
COPY . .
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/yokoe/cloudrunexp /app

CMD /app/main $PORT
