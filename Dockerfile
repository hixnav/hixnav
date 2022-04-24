From centos
COPY hixnav-linux-amd64 /usr/local/bin/hixnav
ENTRYPOINT ["/usr/local/bin/hixnav"]