FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/go-app ./cmd/main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /bin/go-app ./
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["./go-app"]
