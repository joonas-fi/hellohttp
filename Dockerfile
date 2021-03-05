# would work with "FROM scratch", but using alpine to enable debug tooling
FROM alpine:latest

CMD ["/hellohttp"]

ADD rel/hellohttp_linux-amd64 /hellohttp
