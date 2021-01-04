FROM golang:1.15.6
WORKDIR /RateLimiter
ADD . /RateLimiter
RUN cd /RateLimiter && go build
ENTRYPOINT ["./RateLimiter"]