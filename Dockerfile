FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -ldflags '-s -w' -o goproxy .

FROM alpine
WORKDIR /app
COPY --from=build /app/goproxy /app/
CMD ["/app/goproxy"]