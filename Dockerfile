FROM golang:latest as build

ENV GO111MODULE=on \
    GOOS=linux \
    GOPROXY="https://mirrors.aliyun.com/goproxy/"

RUN mkdir /little_chat

COPY . /little_chat

WORKDIR /little_chat

RUN go mod download

RUN CGO_ENABLED=0 go build -o app .

FROM scratch as final

COPY --from=build /little_chat/app .

CMD ["/app"]