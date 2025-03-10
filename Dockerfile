FROM golang:1.23.6

WORKDIR /usr/src/gomall

ENV GOPROXY=https://Goproxy.cn,direct

COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
COPY common common


RUN cd app/frontend/ && go mod download && go mod verify

COPY app/frontend app/frontend
RUN go build -v -o /opt/gomall/frontend/server

COPY app/frontend/conf /opt/gomall/frontend
COPY app/frontend/static /opt/gomall/frontend
COPY app/frontend/template /opt/gomall/frontend

EXPOSE 8080
CMD ["/opt/gomall/frontend/server"]