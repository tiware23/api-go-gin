FROM golang:1.17-alpine AS build-env
WORKDIR /app
RUN apk add git && go mod init github.com/tiware23/api-go-gin
RUN go get -u github.com/gin-gonic/gin
COPY . ./
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#COPY controllers ./
#COPY models ./
#COPY routes ./
##RUN go get github.com/tiware23/api-go-gin/controllers && go get github.com/tiware23/api-go-gin/models
RUN go build -o /app/gin-api

FROM  golang:1.17-alpine AS runtime
WORKDIR /app
COPY --from=build-env /app/gin-api ./
CMD ["/app/gin-api"]
