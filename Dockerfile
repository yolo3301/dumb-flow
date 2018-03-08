# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/yolo3301/dumb-flow
RUN apk add --update alpine-sdk
RUN cd /go/src/github.com/yolo3301/dumb-flow && make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/yolo3301/dumb-flow/bin /app/
EXPOSE 13301
CMD ["./df-server"]