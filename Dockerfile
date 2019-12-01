FROM ubuntu:bionic AS resource
ADD check /opt/resource/check
ADD in /opt/resource/in
ADD out /opt/resource/out
RUN chmod +x /opt/resource/*
RUN apt-get update \
      && DEBIAN_FRONTEND=noninteractive \
      apt-get install -y --no-install-recommends \
        tzdata \
        ca-certificates \
        git \
        jq \
        openssh-client \
      && rm -rf /var/lib/apt/lists/*
RUN git config --global user.email "git@localhost"
RUN git config --global user.name "git"

FROM resource