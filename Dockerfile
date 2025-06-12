FROM alpine:latest
RUN mkdir /app
COPY helloBetsApp /app

CMD ["/app/helloBetsApp"]