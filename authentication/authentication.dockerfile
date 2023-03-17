FROM alpine:latest

RUN mkdir /app

COPY auth_service_app /app

CMD [ "/app/auth_service_app" ]