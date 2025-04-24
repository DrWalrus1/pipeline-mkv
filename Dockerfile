FROM ubuntu:latest
ENV ACCEPT_EULA=Y

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential pkg-config libc6-dev libssl-dev libexpat1-dev libavcodec-dev libgl1-mesa-dev qtbase5-dev zlib1g-dev wget less software-properties-common && \
    add-apt-repository ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install -y golang-go

WORKDIR /makemkv
RUN wget https://www.makemkv.com/download/makemkv-bin-1.18.1.tar.gz && \
    wget https://www.makemkv.com/download/makemkv-oss-1.18.1.tar.gz && \
    mkdir makemkv-bin && \
    mkdir makemkv-oss && \
    tar -xvf makemkv-bin-1.18.1.tar.gz -C ./makemkv-bin --strip-components 1 && \
    tar -xvf makemkv-oss-1.18.1.tar.gz -C ./makemkv-oss --strip-components 1 && \
    rm makemkv-bin-1.18.1.tar.gz makemkv-oss-1.18.1.tar.gz

WORKDIR /makemkv/makemkv-oss
RUN ./configure && \
    make install

WORKDIR /makemkv/makemkv-bin
RUN yes yes | make && \
    make install

WORKDIR /code
COPY . .
RUN go build .

EXPOSE 8080
CMD [ "/code/servermakemkv"]
