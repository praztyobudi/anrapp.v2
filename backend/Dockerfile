FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
# Install air untuk hot-reload
#RUN go install github.com/air-verse/air@latest
#RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go build -o anrapp
EXPOSE 8080
CMD ["./anrapp"]
#CMD ["air"]
