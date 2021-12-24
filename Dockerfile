FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/go-app main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /bin/go-app ./
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["./go-app"]
