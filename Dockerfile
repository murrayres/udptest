FROM quay.octanner.io/base/oct-golang:1.7
RUN apk update
RUN apk add openssl ca-certificates git
RUN mkdir -p /go/src/udptest
WORKDIR /go/src/udptest
COPY . .
WORKDIR /
ADD build.sh /build.sh
RUN chmod +x /build.sh
RUN /build.sh
WORKDIR /go/src/udptest
CMD ["/go/src/udptest/udpserver"]
EXPOSE 9000

