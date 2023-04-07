FROM --platform=linux/arm64 peaceiris/mdbook as bookbuilder

COPY . /book

RUN mdbook build

FROM --platform=linux/arm64 busybox:latest

RUN adduser -D static
USER static
WORKDIR /home/static

COPY --from=bookbuilder /book/book .

CMD ["busybox", "httpd", "-f", "-v", "-p", "3069"]
