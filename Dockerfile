FROM golang:1.11
WORKDIR /go/src/app
# COPY .bashrc /root
COPY ./app /go/src/app
# RUN go get -u github.com/onsi/ginkgo/ginkgo
# RUN go get -u github.com/onsi/gomega/...
RUN go get gopkg.in/jarcoal/httpmock.v1
# RUN go get github.com/onsi/ginkgo
# RUN go get github.com/onsi/gomega
RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/onsi/gomega/...

# RUN go get gopkg.in/cheggaaa/pb.v1
RUN go get ./...
# RUN go-wrapper download
# RUN go-wrapper download app_test.go
# RUN go-wrapper install -v
CMD ["app"]
# RUN mkdir /app
# WORKDIR /app
