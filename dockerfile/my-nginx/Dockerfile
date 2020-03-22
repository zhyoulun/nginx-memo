FROM nginx:1.14

RUN apt update
RUN apt-get install procps -y
RUN apt-get install vim -y
RUN apt-get install strace -y
RUN apt-get install curl -y
RUN apt-get install net-tools -y
RUN apt-get install golang -y
RUN apt-get install telnet -y
RUN apt-get install lsof -y
RUN apt-get install tcpdump -y

COPY my.conf /etc/nginx/conf.d/
COPY static/ /www/static/
