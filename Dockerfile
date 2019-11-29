FROM alpine:latest

WORKDIR /app

COPY max-go .
COPY resource/* ./resource/
COPY resource-public/* ./resource-public/
COPY tmp/* ./tmp/

EXPOSE 9001
ENTRYPOINT ["/app/max-go"]
