FROM golang:latest as build

WORKDIR /app
COPY semesta-app2/go.mod ./
RUN go mod download
COPY semesta-app2/* .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest 
WORKDIR /app

COPY --from=build /app/app .

EXPOSE 3001

CMD ["./app"]
