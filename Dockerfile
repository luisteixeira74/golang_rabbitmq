FROM golang:alpine

RUN apk add --no-cache git

# Move to working directory (/app).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]