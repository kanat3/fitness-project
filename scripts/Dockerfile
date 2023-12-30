FROM docker:24.0.7-dind

RUN apk update && apk add docker-compose && apk add --no-cache make && apk add --no-cache bash
RUN rm -rf /usr/local/go
RUN wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz

ENV GOPATH=/etc
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

RUN mkdir /etc/entrypoint
COPY entrypoint.sh /etc/entrypoint
RUN chmod +x /etc/entrypoint/entrypoint.sh

ENTRYPOINT ["/etc/entrypoint/entrypoint.sh"]