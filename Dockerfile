FROM golang:alpine as build_container
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logging .
# RUN apt-get update && apt-get install -y logrotate

FROM alpine
WORKDIR /root/
COPY --from=build_container /app/logging .

EXPOSE 8006
ENTRYPOINT ["./logging"]
