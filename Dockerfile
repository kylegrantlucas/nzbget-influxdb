FROM golang:latest as build
WORKDIR /go/src/github.com/kylegrantlucas/nzbget-influxdb
COPY . .
RUN go build -o app .

FROM gcr.io/distroless/base
COPY --from=build /go/src/github.com/kylegrantlucas/nzbget-influxdb /
ENTRYPOINT ["/app"]
