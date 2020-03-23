FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak
RUN mv sources.list /etc/apt/

RUN apt update -y
RUN apt install procps -y
RUN apt install vim -y
RUN apt install strace -y
RUN apt install curl -y
RUN apt install net-tools -y
RUN apt install telnet -y
RUN apt install lsof -y
RUN apt install tcpdump -y
RUN apt install ethtool -y

EXPOSE 8080
CMD go run /go/src/app/go_server.go