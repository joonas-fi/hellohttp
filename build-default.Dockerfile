FROM fn61/buildkit-golang:20181005_1740_183e9622c00c5c6b

WORKDIR /go/src/github.com/joonas/fi/hellohttp

CMD bin/build.sh
