#Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .

EXPOSE 9090
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
