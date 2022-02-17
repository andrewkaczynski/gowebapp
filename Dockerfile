FROM golang:1.17-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY cmd/*.go ./
RUN go build -o ./gowebapp

FROM alpine:3.15.0 as bin

COPY --from=build /src/gowebapp /app/gowebapp
EXPOSE 8080
CMD [ "/app/gowebapp" ]