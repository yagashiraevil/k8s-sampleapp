FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/sampleapp

FROM golang:alpine3.15
WORKDIR /app
COPY --from=builder /app/sampleapp .
EXPOSE 8080
USER nonroot
ENTRYPOINT ["/app/sampleapp"]
