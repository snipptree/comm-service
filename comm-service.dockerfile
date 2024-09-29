FROM alpine:latest
WORKDIR /app
COPY communicationApp /app/communicationApp
EXPOSE 80
CMD ["./communicationApp"]
