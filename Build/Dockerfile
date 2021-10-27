FROM centos:centos8

LABEL maintainer="Andrii Anpilogov (andrii.anpilogov@autodesk.com)"

ARG USER_ID
ARG GROUP_ID

RUN : "${USER_ID:?Build argument needs to be set and non-empty.}" \
    : "${GROUP_ID:?Build argument needs to be set and non-empty.}"

RUN [ -e /etc/yum.conf ] && sed -i '/tsflags=nodocs/d' /etc/yum.conf || true

# RUN yum install -y icu

RUN dnf -y install dnf-plugins-core
RUN dnf -y install dnf-plugin-config-manager
RUN dnf -y update

RUN dnf -y install epel-release
RUN dnf -y config-manager --set-enabled powertools

# GCC
RUN dnf -y install gcc-toolset-9-gcc 
RUN dnf -y install gcc-toolset-9-gcc-c++

# CMake & Ninja
RUN dnf -y install cmake
RUN dnf -y install ninja-build

# Free Pascal
# 3.0.4 should be used for validating compiler base compatibility.
RUN (cd /opt && curl -O -L 'http://downloads.sourceforge.net/project/freepascal/Linux/3.0.4/fpc-3.0.4-1.x86_64.rpm' \
    && rpm -i fpc-3.0.4-1.x86_64.rpm)
# Using 3.2.2 until PolymorphicFactory is reworked to not using generics that are not full functional in 3.0.4.
# RUN (cd /opt && curl -O -L 'http://downloads.sourceforge.net/project/freepascal/Linux/3.2.2/fpc-3.2.2-1.x86_64.rpm' \
#     && rpm -i fpc-3.2.2-1.x86_64.rpm)

# Golang
RUN (cd /opt && curl -O -L https://golang.org/dl/go1.17.2.linux-amd64.tar.gz \
    && tar zxvf go1.17.2.linux-amd64.tar.gz)

# Java
RUN dnf -y install \
           java-11-openjdk-devel

# Mono
RUN rpm --import https://download.mono-project.com/repo/xamarin.gpg \
    && dnf config-manager --add-repo https://download.mono-project.com/repo/centos8-stable.repo \
    && dnf -y install mono-complete


# General purpose tools
RUN dnf -y install \
           glibc-common \
           glibc-utils \
           less \
           passwd \
           tar \
           vim-minimal \
           vim-enhanced \
           which \
           sudo \
           bash-completion \
           mc \
           yum-utils \
           && yum clean all

# Create user
RUN groupadd docker \
    && useradd -ms /bin/bash --uid $USER_ID -g docker -G wheel user \
    && printf "user:user" | chpasswd \
    && printf "user ALL= NOPASSWD: ALL\\n" >> /etc/sudoers

# Initialize Toolkit
RUN echo "source /opt/rh/gcc-toolset-9/enable" >> /home/user/.bash_profile
RUN echo "export PATH=\$PATH:/opt/go/bin" >> /home/user/.bash_profile

VOLUME /data

# Run Container as nonroot
USER user
WORKDIR /data

ENTRYPOINT [ "/bin/bash", "-l" ]
