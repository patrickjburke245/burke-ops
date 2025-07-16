FROM golang:1.24

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /app .

#4
CMD ["/app"]