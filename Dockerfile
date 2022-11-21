FROM golang:1.19-alpine3.16 AS builder
RUN apk update && \
    apk --no-cache add build-base
LABEL stage=builder
WORKDIR /building-ground
COPY . /building-ground
RUN cd /building-ground && \
#    go test ./... -cover -v  && \
    go build -o main .

FROM alpine:3.16 as production
RUN apk update && \
    apk --no-cache add ca-certificates
COPY --from=builder /building-ground/main .
COPY /configs /configs/
ENTRYPOINT ./main