FROM golang:1.17

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct


RUN go mod download && go mod verify
RUN go build -o /app/agent-container .

#EXPOSE 8080
#EXPOSE 10001

# 执行manager-container
CMD ["./agent-container"]