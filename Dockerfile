FROM ubuntu
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && \
    apt-get -y install gcc  && \
    rm -rf /var/lib/apt/lists/*
COPY dropper.c /
RUN mkdir output; gcc dropper.c -o output/dropper
