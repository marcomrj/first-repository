FROM golang:1.16 as build
WORKDIR /go/src
COPY . .
ENV PATH="/go/bin:${PATH}"
RUN GOOS=linux GOARCH=386 go build main.go

FROM ubuntu
WORKDIR /app
COPY ./assets ./assets
COPY --from=build /go/src/main ./
EXPOSE 8000
CMD ["./main"]