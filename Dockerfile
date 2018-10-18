FROM golang:1.10 as builder
ARG DOCK_PKG_DIR=/go/src/github.com/weather-project/
ADD . $DOCK_PKG_DIR
WORKDIR $DOCK_PKG_DIR
RUN go get -t -d -v -insecure ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go test ./...


FROM scratch
# ADD main / # This might be needed
LABEL source=git@github.com:kyma-project/examples.git
WORKDIR /app/
COPY --from=builder /go/src/github.com/weather-project/main /app/
COPY --from=builder /go/src/github.com/weather-project/docs/api/api.yaml /app/
CMD ["/main"]

EXPOSE 8017:8017