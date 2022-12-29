FROM golang:1.19

WORKDIR /logging-batch-go
COPY . .

RUN mkdir ./logs
RUN mv ./.aws ~/
RUN go env -w GO111MODULE=auto
RUN go mod download
RUN go get
RUN go build logFile.com/log-file-go

EXPOSE 8000

# CMD [ "go" , "run" , "main.go" ]
CMD [ "./log-file-go" ]