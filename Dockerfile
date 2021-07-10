FROM golang:1.16-alpine
WORKDIR /app
COPY . .

# for go binary to run on alpine image 
ENV CGO_ENABLED=0 \
    GIN_MODE=release \
    PORT=8080

RUN go build -o /dictionaryapi

CMD ["/dictionaryapi"]
