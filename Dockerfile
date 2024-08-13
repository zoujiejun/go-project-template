FROM ccr.ccs.tencentyun.com/uranus.reg/dockerhub:golang-1.21 AS builder

COPY . /src
WORKDIR /src

RUN make build


FROM ccr.ccs.tencentyun.com/uranus.reg/dockerhub:debian-stable-slim

#RUN apt-get update && apt-get install -y --no-install-recommends \
#		ca-certificates  \
#        netbase \
#        && rm -rf /var/lib/apt/lists/ \
#        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/config /app/conf
WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /app/conf

CMD ["./go-project-template", "-f", "/app/conf/config.yaml"]
