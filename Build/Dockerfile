FROM centos:centos8

ARG USER_ID
ARG GROUP_ID

LABEL maintainer="Andrii Anpilogov (andrii.anpilogov@autodesk.com)"

RUN [ -e /etc/yum.conf ] && sed -i '/tsflags=nodocs/d' /etc/yum.conf || true

# RUN yum install -y icu

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN yum update -y

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

# # Initialize Toolkit
RUN echo "source /opt/rh/gcc-toolset-9/enable" >> /etc/bashrc
ENV PATH="${PATH}:/opt/go/bin"

# Initialize Toolkit
# RUN : "${USER_ID:?Build argument needs to be set and non-empty.}" \
#     : "${GROUP_ID:?Build argument needs to be set and non-empty.}"

# Create user
RUN if [ $USER_ID != "" ] ; then groupadd docker \
    && useradd -ms /bin/bash --uid $USER_ID -g docker -G wheel user \
    && printf "user:user" | chpasswd \
    && printf "user ALL= NOPASSWD: ALL\\n" >> /etc/sudoers ; fi

# Initialize Toolkit
RUN if [ $USER_ID != "" ] ; then echo "source /opt/rh/gcc-toolset-9/enable" >> /home/user/.bash_profile ; fi
RUN if [ $USER_ID != "" ] ; then echo "export PATH=\$PATH:/opt/go/bin" >> /home/user/.bash_profile ; fi

ENTRYPOINT ["./entrypoint.sh"]
