FROM golang:1.10.1-stretch as main
RUN go get "github.com/golang/lint/golint"
RUN go get "github.com/golang/dep/cmd/dep"
WORKDIR /go/src/stelligent/hello-go
COPY *.go Gopkg.* ./
RUN dep ensure
RUN go vet ./...
RUN golint -set_exit_status $(go list ./... | grep -v /vendor/)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main *.go

FROM golang:1.10.1-stretch as healthcheck
RUN go get "github.com/golang/lint/golint"
WORKDIR /go/src/stelligent/hello-go
COPY healthcheck/*.go ./
RUN ls -R
RUN go vet ./...
RUN golint -set_exit_status $(go list ./... | grep -v /vendor/)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o healthcheck *.go

FROM scratch
COPY --from=main /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=main /go/src/stelligent/hello-go/main /
COPY --from=healthcheck /go/src/stelligent/hello-go/healthcheck /
ENTRYPOINT [ "/main"] 
HEALTHCHECK --interval=10s --timeout=3s CMD ["/healthcheck"]
EXPOSE 8080
ARG SOURCE_COMMIT
ENV SOURCE_COMMIT=$SOURCE_COMMIT