#!/bin/bash

kubeadm reset -f
apt-mark unhold kubeadm kubectl kubelet
apt-get purge -y kubeadm kubectl kubelet kubernetes-cni 
apt-get autoremove -y
rm -rf /etc/cni /etc/kubernetes /var/lib/dockershim /var/lib/etcd /var/lib/kubelet /var/run/kubernetes ~/.kube/*
iptables -F && iptables -X
iptables -t nat -F && iptables -t nat -X
iptables -t raw -F && iptables -t raw -X
iptables -t mangle -F && iptables -t mangle -X
systemctl restart docker

