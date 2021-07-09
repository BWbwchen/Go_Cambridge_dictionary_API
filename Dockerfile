FROM golang:1.16-alpine AS build
WORKDIR /app
COPY . .

# for go binary to run on alpine image 
ENV CGO_ENABLED=0
RUN go mod download 
RUN go build -o /dictionaryapi

FROM alpine:latest AS deploy
WORKDIR /
COPY --from=build /dictionaryapi /
ENV GIN_MODE=release
ENV PORT=8080

CMD ["/dictionaryapi"]