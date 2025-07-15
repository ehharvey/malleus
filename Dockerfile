FROM golang:1.24 AS builder
WORKDIR /build
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -C cmd/server -o /build/app


FROM scratch
WORKDIR /app
COPY --from=builder /build/app ./app
ENTRYPOINT [ "./app" ]