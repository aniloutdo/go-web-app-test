
FROM golang:1.22 as builder

WORKDIR /src

#COPY go.mod .
COPY . ./
RUN go mod download 
#COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-web-app



#FROM alpine:latest



#COPY --from=builder /go-web-app /go-web-app

EXPOSE 8080

#CMD ["./go-web-app"]
ENTRYPOINT ["/go-web-app"]
