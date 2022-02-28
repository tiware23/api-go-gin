FROM golang:1.17-alpine
WORKDIR /app
RUN apk add git
COPY go.mod ./
COPY go.sum ./
#RUN go mod download
COPY controllers ./
COPY models ./
COPY routes ./
#RUN go get github.com/tiware23/api-go-gin/controllers && go get github.com/tiware23/api-go-gin/models
RUN go build -o /gin-api
CMD ["/gin-api"]
