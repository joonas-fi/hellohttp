FROM scratch

CMD ["/hellohttp"]

ADD rel/hellohttp_linux-amd64 /hellohttp
