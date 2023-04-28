FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY . .
COPY ./ ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
RUN sed -i 's/\r$//' run.sh
RUN chmod +x 'run.sh'

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]