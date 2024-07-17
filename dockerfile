FROM golang:1.18-alpine
WORKDIR ./backend/proyecto/go-jwt
COPY ./backend/proyecto/go-jwt/go.mod .
COPY ./backend/proyecto/go-jwt/go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist
CMD ./out/dist

