FROM alpine:latest

RUN mkdir /app

COPY baseService /app

CMD [ "/app/baseService"]