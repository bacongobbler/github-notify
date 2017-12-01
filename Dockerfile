FROM golang
WORKDIR /go/src/github.com/bacongobbler/github-notify/
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
RUN go build -o bin/github-notify ./main.go

FROM alpine
COPY --from=0 /go/src/github.com/bacongobbler/github-notify/bin/github-notify /bin/github-notify
RUN chmod 755 /bin/github-notify
ENTRYPOINT /bin/github-notify
