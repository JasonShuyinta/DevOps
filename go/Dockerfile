FROM golang:1.22

# Set the home directory for the new user 
ENV HOME /home/gouser 
WORKDIR /app 


# Set destination for COPY 
COPY go.mod ./ 

RUN go mod download 

COPY . .

CMD ["go", "run", "main.go"]
