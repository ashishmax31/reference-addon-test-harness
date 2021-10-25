FROM golang:1.16 AS builder
ENV PKG=/go/src/github.com/mt-sre/reference-addon-test-harness/
WORKDIR ${PKG}


COPY . .
RUN make

FROM scratch

COPY --from=builder /go/src/github.com/mt-sre/reference-addon-test-harness/reference-addon-test-harness reference-addon-test-harness

ENTRYPOINT [ "/reference-addon-test-harness" ]

