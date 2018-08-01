FROM golang:1.10.1-stretch as builder

RUN go get "github.com/golang/lint/golint"
RUN go get "github.com/golang/dep/cmd/dep"

WORKDIR /go/src/stelligent/hello-go
COPY *.go Gopkg.* ./
RUN dep ensure
RUN go vet ./...
RUN golint -set_exit_status $(go list ./... | grep -v /vendor/)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main *.go


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/stelligent/hello-go/main /
ENTRYPOINT [ "/main"] 
EXPOSE 8080