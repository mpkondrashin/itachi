FROM ubuntu
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && \
    apt-get -y install mingw-w64 && \
    rm -rf /var/lib/apt/lists/*
COPY dropper.c /
RUN mkdir output; x86_64-w64-mingw32-g++ dropper.c -o output/dropper.exe

#    apt-get -y install gcc  && \
