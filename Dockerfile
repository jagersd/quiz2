FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/ui

COPY quiz2 /app
COPY prod.config.ini /app/conf.ini
ADD ui /app/ui/

WORKDIR /app

CMD ["./quiz2","-m","./quiz2","-d","./quiz2"]

