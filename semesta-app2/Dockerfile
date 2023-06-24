FROM golang:latest as build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest 
WORKDIR /app

COPY --from=build /app/app .

EXPOSE 3001

CMD ["./app"]
