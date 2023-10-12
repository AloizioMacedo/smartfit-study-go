FROM golang:1.21.3-alpine3.17 as builder

WORKDIR /app

COPY . .
RUN go build

FROM gcr.io/distroless/cc

WORKDIR /app

COPY --from=builder /app/templates templates
COPY --from=builder /app/smartfit smartfit
COPY --from=builder /app/locations.json locations.json

ENV GIN_MODE=release

CMD ["./smartfit"]
