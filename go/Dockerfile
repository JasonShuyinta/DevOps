FROM golang:1.22

# Set destination for COPY 
WORKDIR /app

COPY go.mod ./ 
RUN go mod download 

COPY . .

CMD ["go", "run", "main.go"]
