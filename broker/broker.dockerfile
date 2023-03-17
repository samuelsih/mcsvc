# #base go image builder
# FROM golang:1.20-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o broker_service_app ./cmd/api

# RUN chmod +x /app/broker_service_app


# # tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app/broker_service_app /app

# CMD [ "/app/broker_service_app" ]


FROM alpine:latest

RUN mkdir /app

COPY broker_service_app /app

CMD [ "/app/broker_service_app" ]