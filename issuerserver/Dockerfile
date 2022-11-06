FROM        golang
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN git config --global https.proxy http://127.0.0.1:7890
RUN git config --global https.proxy https://127.0.0.1:7800
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN         go mod download
RUN         go build -o app

ENTRYPOINT  ["./app"]