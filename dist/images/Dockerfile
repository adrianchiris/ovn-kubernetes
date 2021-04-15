FROM centos:8 as ovn

USER root

ENV PYTHONDONTWRITEBYTECODE yes

# Get openvswitch from Nvidia controlled source
# Based on nv OVS commit: 21dfd42d
RUN curl -s https://3c720378ea793bfa1cdbc4180c64fca33176674914260f6d:@packages.nvidia.com/install/repositories/sdn/nv-ovs-L1/script.rpm.sh | bash
# Based on nv OVN commit: 7d806544
RUN curl -s https://4a6c3c2ab9c3468ee4624cf102ba7dc81436598399db8477:@packages.nvidia.com/install/repositories/sdn/nv-ovn-L1/script.rpm.sh | bash 

ARG K8S_VER

COPY kubernetes.repo /etc/yum.repos.d/kubernetes.repo
# install epel.repo
RUN yum install -y epel-release
RUN INSTALL_PKGS=" \
	bind-utils procps-ng openssl numactl-libs firewalld firewalld-filesystem \
	libpcap iproute iputils strace socat unbound-libs python2 \
        " && \
	yum install -y --setopt=tsflags=nodocs --setopt=skip_missing_names_on_install=False $INSTALL_PKGS && \
    yum -y --setopt=tsflags=nodocs --setopt=skip_missing_names_on_install=False update && \
    yum install -y --setopt=tsflags=nodocs --setopt=skip_missing_names_on_install=False \
        kubectl$K8S_VER openvswitch openvswitch-devel ovn ovn-host ovn-central && \
    yum autoremove -y && \
    yum clean all && rm -rf /var/cache/dnf/*

# ovnkube.sh is the entry point. This script examines environment
# variables to direct operation and configure ovn
COPY ovnkube.sh ovndb-raft-functions.sh git_info /root/
COPY ovn-kube-util /usr/bin/

LABEL io.k8s.display-name="ovn" \
      io.k8s.description="ovn centos image" \
      maintainer="Girish Moodalbail <gmoodalbail@nvidia.com>"

WORKDIR /root
ENTRYPOINT /root/ovnkube.sh

FROM ovn as ovnkube

RUN  mkdir -p /usr/libexec/cni/

# Built in ../../go_controller, then the binaries are copied here.
# put things where they are in the rpm
COPY ovnkube ovndbchecker /usr/bin/
COPY ovn-k8s-cni-overlay /usr/libexec/cni/ovn-k8s-cni-overlay

COPY chassis-del.sh git_info /root/
# override the rpm's ovn_k8s.conf with this local copy
COPY ovn_k8s.conf /etc/openvswitch/ovn_k8s.conf

# iptables wrappers
COPY ./iptables-scripts/ip* /usr/sbin/

LABEL io.k8s.display-name="ovn-kubernetes" \
      io.k8s.description="ovnkube centos image" \
      maintainer="Girish Moodalbail <gmoodalbail@nvidia.com>"

WORKDIR /root
ENTRYPOINT /root/ovnkube.sh
