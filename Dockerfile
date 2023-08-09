FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/sampleapp

FROM alpine:3.18.3
EXPOSE 8080
WORKDIR /app
COPY --from=builder /app/sampleapp .
# Create a group and user
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
# Tell docker that all future commands should run as the appuser user
USER nonroot
ENTRYPOINT ["/app/sampleapp"]
