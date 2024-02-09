FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod download

# RUN go build -o bin .

CMD go run .

# ENTRYPOINT [ "/app/bin" ]

EXPOSE 8080