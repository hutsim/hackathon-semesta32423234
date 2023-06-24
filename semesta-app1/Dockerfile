FROM golang:latest as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest 
WORKDIR /app

COPY --from=build /app/app .
COPY --from=build /app/.env .

EXPOSE 3000

CMD ["./app"]
