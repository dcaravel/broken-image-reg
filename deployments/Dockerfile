FROM golang:1.21 as builder

WORKDIR /go/src/app

COPY broken-image-reg broken-image-reg
COPY go-containerregistry go-containerregistry

WORKDIR /go/src/app/broken-image-reg

RUN CGO_ENABLED=0 go build -o ../brokenreg . && chmod 755 ../brokenreg

FROM registry.access.redhat.com/ubi8-minimal

COPY --from=builder /go/src/app/brokenreg .

ENTRYPOINT [ "/brokenreg" ]