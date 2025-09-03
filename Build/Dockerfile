FROM centos:centos8

ARG USER_ID
ARG GROUP_ID

LABEL maintainer="Andrii Anpilogov (andrii.anpilogov@autodesk.com)"

RUN [ -e /etc/yum.conf ] && sed -i '/tsflags=nodocs/d' /etc/yum.conf || true

# CentOS 8 vault repos
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* && \
    sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* && \
    yum update -y

RUN dnf -y install dnf-plugins-core dnf-plugin-config-manager && dnf -y update
RUN dnf -y install epel-release && dnf -y config-manager --set-enabled powertools

# Tooling
RUN dnf -y install gcc-toolset-9-gcc gcc-toolset-9-gcc-c++ cmake ninja-build git python3 python3-pip

# Free Pascal 3.0.4
RUN (cd /opt && curl -O -L 'http://downloads.sourceforge.net/project/freepascal/Linux/3.0.4/fpc-3.0.4-1.x86_64.rpm' \
    && rpm -i fpc-3.0.4-1.x86_64.rpm)

# Go
RUN (cd /opt && curl -O -L https://golang.org/dl/go1.17.2.linux-amd64.tar.gz && tar zxvf go1.17.2.linux-amd64.tar.gz)

# Java & Mono
RUN dnf -y install java-11-openjdk-devel
RUN rpm --import https://download.mono-project.com/repo/xamarin.gpg \
    && dnf config-manager --add-repo https://download.mono-project.com/repo/centos8-stable.repo \
    && dnf -y install mono-complete

# General tools
RUN dnf -y install glibc-common glibc-utils less passwd tar vim-minimal vim-enhanced which sudo bash-completion mc yum-utils && yum clean all

# Shell env for GCC toolset & Go
RUN echo "source /opt/rh/gcc-toolset-9/enable" >> /etc/bashrc
ENV PATH="${PATH}:/opt/go/bin"

# ---- Emscripten SDK (emsdk) ----
ENV EMSDK=/opt/emsdk
RUN git clone https://github.com/emscripten-core/emsdk.git ${EMSDK} && \
    bash -lc "cd ${EMSDK} && ./emsdk install latest && ./emsdk activate latest"

# Source emsdk env in interactive shells + user shell
RUN echo "source ${EMSDK}/emsdk_env.sh" >> /etc/bashrc

# Create user (only if USER_ID provided)
RUN if [ -n "$USER_ID" ]; then \
      groupadd -g ${GROUP_ID:-1000} docker && \
      useradd -ms /bin/bash --uid $USER_ID -g docker -G wheel user && \
      printf 'user:user' | chpasswd && \
      printf 'user ALL= NOPASSWD: ALL\n' >> /etc/sudoers && \
      echo 'source /opt/rh/gcc-toolset-9/enable' >> /home/user/.bash_profile && \
      echo 'export PATH=$PATH:/opt/go/bin' >> /home/user/.bash_profile && \
      echo "source ${EMSDK}/emsdk_env.sh" >> /home/user/.bash_profile ; \
    fi

# (Optional) sanity check in build-time shell that sourcing works
RUN bash -lc "source ${EMSDK}/emsdk_env.sh && emcc --version && em++ --version && node --version"

# Your repo should provide this
ENTRYPOINT ["./entrypoint.sh"]
