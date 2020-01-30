FROM golang:latest AS builder
ADD . /app/backend
WORKDIR /app/backend
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...
# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8000
