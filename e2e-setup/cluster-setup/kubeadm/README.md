## Setup Container runtime as Docker

#### INSTALL DOCKER ENGINE

1. On each of your nodes, install the Docker for your Linux distribution as per [Install Docker Engine](https://docs.docker.com/engine/install/#server).
- Install Docker Engine on Ubuntu

```bash
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
```
- Add Docker’s official GPG key:

```bash
 curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```
_For x86_64 / amd64_
```bash
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

**Install Docker Engine**
```bash
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

2. Configure the Docker daemon, in particular to use systemd for the management of the container’s cgroups

```bash
sudo mkdir /etc/docker
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF
```

3. Restart Docker and enable on boot:

```bash
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker
```

## Setup Kubeadm Cluster

#### Installing kubeadm, kubelet and kubectl
```bash
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl

sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

#### To initialize the control-plane node run:

```bash
kubeadm init 
```

Output:

```bash
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.31.18.100:6443 --token 2ekqtg.o88pcx8ikggp7h88 \
    --discovery-token-ca-cert-hash sha256:23ad32acc342526aebcd37bbd83d473f8b22fed3bb71677895178c77585d6167 
```

Checking the node status:

```
root@ip-172-31-18-100:~# kubectl get nodes
NAME               STATUS     ROLES                  AGE   VERSION
ip-172-31-18-100   NotReady   control-plane,master   71s   v1.20.4
```

Add Netowrk CNI:

```bash
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
```
Now again check the node status:
```bash
root@ip-172-31-18-100:~# kubectl get nodes
NAME               STATUS   ROLES                  AGE     VERSION
ip-172-31-18-100   Ready    control-plane,master   3m18s   v1.20.4
```

#### Join a worker Node to the Kubeadm Cluster

