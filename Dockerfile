FROM alpine:3.9

ENV CLOUD_SDK_URL https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-200.0.0-linux-x86_64.tar.gz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache \
      go \
      curl \
      musl-dev \
      bash \
      python-dev && \
    curl -sSLo google-cloud-sdk.tar.gz ${CLOUD_SDK_URL} > /dev/null && \
    tar xzf google-cloud-sdk.tar.gz && \
    rm -f google-cloud-sdk.tar.gz && \
    ./google-cloud-sdk/install.sh \
        --usage-reporting false \
        --path-update false \
        -q && \
    mkdir /go

WORKDIR /go/src/invoke

COPY invoke.go .
RUN go install -v

COPY . .

CMD ["invoke"]
