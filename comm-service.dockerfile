FROM alpine:latest

RUN mkdir /app

COPY communicationApp /app
COPY templates /templates

CMD [ "/app/communicationApp"]