FROM ubuntu:20.04
ARG DEBIAN_FRONTEND=noninteractive

ENV OSMIUM_VERSION 2.18.0
ENV OSMIUM_TOOL_VERSION 1.14.0

RUN apt-get update
RUN apt-get update && apt-get install -y \
    libosmium2-dev libprotozero-dev libboost-program-options-dev libbz2-dev zlib1g-dev liblz4-dev libexpat1-dev cmake pandoc build-essential

RUN mkdir /var/install
WORKDIR /var/install

COPY third_party/libosmium libosmium

RUN cd libosmium && \
    mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_EXAMPLES=OFF -DBUILD_TESTING=OFF -DINSTALL_PROTOZERO=ON .. && \
    make

COPY third_party/osmium-tool osmium-tool

RUN cd osmium-tool && \
    mkdir build && cd build && \
    cmake -DOSMIUM_INCLUDE_DIR=/var/install/libosmium/include/ .. && \
    make

RUN mv /var/install/osmium-tool/build/src/osmium /usr/bin/osmium