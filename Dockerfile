FROM golang:1.9 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build

FROM golang:1.9
COPY --from=build /go/src/app/s3-query .
CMD ["./s3-query"]