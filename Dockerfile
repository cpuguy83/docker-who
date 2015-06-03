FROM golang:1.4

COPY . /opt/docker-who
RUN cd /opt/docker-who && go build
RUN cp /opt/docker-who/docker-who /usr/bin/docker-who
ENV FOO bar

ENTRYPOINT ["/usr/bin/docker-who"]
