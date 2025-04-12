FROM ubuntu:latest
ENV ACCEPT_EULA=Y

RUN apt-get update

RUN apt-get install build-essential pkg-config libc6-dev libssl-dev libexpat1-dev libavcodec-dev libgl1-mesa-dev qtbase5-dev zlib1g-dev wget less -y
WORKDIR /makemkv
RUN wget https://www.makemkv.com/download/makemkv-bin-1.18.1.tar.gz
RUN wget https://www.makemkv.com/download/makemkv-oss-1.18.1.tar.gz
RUN mkdir makemkv-bin
RUN mkdir makemkv-oss
RUN tar -xvf makemkv-bin-1.18.1.tar.gz -C ./makemkv-bin --strip-components 1
RUN tar -xvf makemkv-oss-1.18.1.tar.gz -C ./makemkv-oss --strip-components 1
RUN rm makemkv-bin-1.18.1.tar.gz makemkv-oss-1.18.1.tar.gz

WORKDIR /makemkv/makemkv-oss
RUN ./configure
RUN make install

WORKDIR /makemkv/makemkv-bin
RUN yes yes | make
RUN make install

RUN apt install software-properties-common -y
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt update
RUN apt install golang-go -y

WORKDIR /code
COPY . .
RUN go build .

CMD ["/code/servermakemkv"]
