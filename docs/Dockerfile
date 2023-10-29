FROM --platform=linux/arm64 busybox:latest

RUN adduser -D static
USER static
WORKDIR /home/static

COPY book .

CMD ["busybox", "httpd", "-f", "-v", "-p", "3069"]
