FROM golang:alpine3.12

WORKDIR /dist

COPY chat .

COPY /config/app.yaml .

RUN mkdir /dist/images/ && mkdir /dist/log/

CMD ["./chat"]