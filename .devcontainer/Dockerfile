# Base: lightweight Ubuntu Dev Container
FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

# Ensure noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Install common tools
RUN apt-get update && apt-get install -y \
    curl wget unzip zip git make build-essential gnupg lsb-release \
    openjdk-17-jdk \
    python3 python3-pip python3-venv \
    golang-go \
    nodejs npm \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Bazelisk (manages Bazel versions)
RUN curl -Lo /usr/local/bin/bazelisk https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-amd64 \
    && chmod +x /usr/local/bin/bazelisk \
    && ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel

# Install buildifier + buildozer (Bazel tooling)
RUN go install github.com/bazelbuild/buildtools/buildifier@latest \
    && go install github.com/bazelbuild/buildtools/buildozer@latest \
    && mv /root/go/bin/* /usr/local/bin/

# Set Go path
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# Default user (matches devcontainer expectation)
USER vscode
