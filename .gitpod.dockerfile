FROM gitpod/workspace-full

RUN sudo apt-get update && \
    sudo apt-get update && \
    sudo apt install build-essential -y && \
    sudo apt install unzip -y && \
    sudo apt install zip -y && \
    sudo apt install openssl -y && \
    sudo apt install libssl-dev -y && \
    sudo apt install git -y && \
    sudo apt install pkg-config -y

USER gitpod

ARG BAT_VERSION="0.22.1"
ARG BAT_ARCH="amd64"

RUN wget https://github.com/sharkdp/bat/releases/download/v${BAT_VERSION}/bat_${BAT_VERSION}_${BAT_ARCH}.deb && \
    sudo dpkg -i bat_${BAT_VERSION}_${BAT_ARCH}.deb && \
    rm bat_${BAT_VERSION}_${BAT_ARCH}.deb && \      
    sudo apt-get install exa -y && \
    sudo apt-get -y install hey

ARG GOLANG_VERSION="1.20"
ARG GOLANG_OS="linux"
ARG GOLANG_ARCH="amd64"
ARG TINYGO_VERSION="0.27.0"
ARG TINYGO_ARCH="amd64"

RUN wget https://go.dev/dl/go${GOLANG_VERSION}.${GOLANG_OS}-${GOLANG_ARCH}.tar.gz
RUN sudo rm -rf /usr/local/go 
RUN sudo tar -C /usr/local -xzf go${GOLANG_VERSION}.${GOLANG_OS}-${GOLANG_ARCH}.tar.gz

RUN echo "" >> ${HOME}/.bashrc
RUN echo "export GOLANG_HOME=\"/usr/local/go\"" >> ${HOME}/.bashrc
RUN echo "export PATH=\"\$GOLANG_HOME/bin:\$PATH\"" >> ${HOME}/.bashrc

RUN source ${HOME}/.bashrc

RUN rm go${GOLANG_VERSION}.${GOLANG_OS}-${GOLANG_ARCH}.tar.gz               

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb
RUN rm tinygo_${TINYGO_VERSION}_${TINYGO_ARCH}.deb


