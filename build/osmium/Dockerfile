FROM debian:bookworm-slim as builder

RUN apt-get update
RUN apt-get update && apt-get install -y \
    libosmium2-dev libprotozero-dev libboost-program-options-dev libbz2-dev zlib1g-dev liblz4-dev libexpat1-dev cmake pandoc build-essential

RUN mkdir /var/install
WORKDIR /var/install

COPY libosmium libosmium

RUN cd libosmium && \
    mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_EXAMPLES=OFF -DBUILD_TESTING=OFF -DINSTALL_PROTOZERO=ON .. && \
    make

COPY osmium-tool osmium-tool

RUN cd osmium-tool && \
    mkdir build && cd build && \
    cmake -DOSMIUM_INCLUDE_DIR=/var/install/libosmium/include/ .. && \
    make

RUN mv /var/install/osmium-tool/build/src/osmium /usr/bin/osmium