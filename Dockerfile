FROM golang AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o revprox ./...

FROM alpine:latest as production

COPY --from=builder /app .
CMD ["./revprox"]

