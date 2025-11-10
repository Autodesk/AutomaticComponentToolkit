FROM centos:centos8

ARG USER_ID
ARG GROUP_ID

LABEL maintainer="Andrii Anpilogov (andrii.anpilogov@autodesk.com)"

RUN [ -e /etc/yum.conf ] && sed -i '/tsflags=nodocs/d' /etc/yum.conf || true

# Vault repos for CentOS 8
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* && \
    sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* && \
    yum update -y

RUN dnf -y install dnf-plugins-core dnf-plugin-config-manager && dnf -y update
RUN dnf -y install epel-release && dnf -y config-manager --set-enabled powertools
RUN dnf -y install libatomic

# Tooling
RUN dnf -y install gcc-toolset-9-gcc gcc-toolset-9-gcc-c++ cmake ninja-build git curl

# ---- Python 3.9 (required by recent emscripten) ----
RUN dnf -y module enable python39 && \
    dnf -y install python39 python39-pip && \
    alternatives --install /usr/bin/python3 python3 /usr/bin/python3.9 2 && \
    alternatives --set python3 /usr/bin/python3.9

# (optional) ICU runtime/headers if you need them
# RUN dnf -y install libicu libicu-devel

# Free Pascal 3.0.4
RUN cd /opt && \
    curl -L -H "Accept: */*" -H "Accept-Language: en-US,en;q=0.9" \
         -H "User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36" \
         -o fpc-3.0.4-1.x86_64.rpm \
         'https://sourceforge.net/projects/freepascal/files/Linux/3.0.4/fpc-3.0.4-1.x86_64.rpm/download' && \
    FILE_SIZE=$(stat -c%s fpc-3.0.4-1.x86_64.rpm 2>/dev/null || echo 0) && \
    if [ "$FILE_SIZE" -lt 1000000 ]; then \
        echo "Downloaded file too small ($FILE_SIZE bytes), trying alternative URL..."; \
        rm -f fpc-3.0.4-1.x86_64.rpm; \
        curl -L -f -o fpc-3.0.4-1.x86_64.rpm 'http://downloads.sourceforge.net/project/freepascal/Linux/3.0.4/fpc-3.0.4-1.x86_64.rpm'; \
    fi && \
    [ $(stat -c%s fpc-3.0.4-1.x86_64.rpm) -gt 1000000 ] && \
    rpm -i fpc-3.0.4-1.x86_64.rpm && \
    rm -f fpc-3.0.4-1.x86_64.rpm

# Go
RUN (cd /opt && curl -O -L https://golang.org/dl/go1.17.2.linux-amd64.tar.gz && tar zxvf go1.17.2.linux-amd64.tar.gz)
ENV PATH="${PATH}:/opt/go/bin"

# Java & Mono
RUN dnf -y install java-11-openjdk-devel
RUN rpm --import https://download.mono-project.com/repo/xamarin.gpg \
    && dnf config-manager --add-repo https://download.mono-project.com/repo/centos8-stable.repo \
    && dnf -y install mono-complete

# General tools
RUN dnf -y install glibc-common glibc-utils less passwd tar vim-minimal vim-enhanced which sudo bash-completion mc yum-utils && yum clean all

# Enable GCC toolset in shells
RUN echo "source /opt/rh/gcc-toolset-9/enable" >> /etc/bashrc

# ---- Emscripten SDK (emsdk) ----
ENV EMSDK=/opt/emsdk
RUN git clone https://github.com/emscripten-core/emsdk.git ${EMSDK} && \
    bash -lc "cd ${EMSDK} && ./emsdk install latest && ./emsdk activate latest"

# Source emsdk env in interactive shells
RUN echo "source ${EMSDK}/emsdk_env.sh" >> /etc/bashrc

# Create user (if requested)
RUN if [ -n "$USER_ID" ]; then \
      groupadd -g ${GROUP_ID:-1000} docker && \
      useradd -ms /bin/bash --uid $USER_ID -g docker -G wheel user && \
      printf 'user:user' | chpasswd && \
      printf 'user ALL= NOPASSWD: ALL\n' >> /etc/sudoers && \
      echo 'source /opt/rh/gcc-toolset-9/enable' >> /home/user/.bash_profile && \
      echo "export EMSDK_PYTHON=${EMSDK_PYTHON}" >> /home/user/.bash_profile && \
      echo "source ${EMSDK}/emsdk_env.sh" >> /home/user/.bash_profile ; \
    fi

# Sanity check (now python3 == 3.9; don't rely on EMSDK_PYTHON which emsdk_env clears)
RUN bash -lc "source ${EMSDK}/emsdk_env.sh && \
  python3 --version && emcc --version && em++ --version && node --version && \
  /opt/emsdk/upstream/bin/wasm-opt --version"

# Your repo should provide this script; it will source the env before running any command.
ENTRYPOINT ["./entrypoint.sh"]