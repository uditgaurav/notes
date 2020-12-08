### Fannel

For flannel to work correctly, you must pass --pod-network-cidr=10.244.0.0/16 to kubeadm init.

Set /proc/sys/net/bridge/bridge-nf-call-iptables to 1 by running sysctl net.bridge.bridge-nf-call-iptables=1 to pass bridged IPv4 traffic to iptables’ chains. This is a requirement for some CNI plugins to work, for more information please see here.

Make sure that your firewall rules allow UDP ports 8285 and 8472 traffic for all hosts participating in the overlay network. see here .

Note that flannel works on amd64, arm, arm64, ppc64le and s390x under Linux. Windows (amd64) is claimed as supported in v0.11.0 but the usage is undocumented.
```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/62e44c867a2846fefb68bd5f178daf4da3095ccb/Documentation/kube-flannel.yml
```

### Cilium

For more information about using Cilium with Kubernetes, see Kubernetes Install guide for Cilium.

For Cilium to work correctly, you must pass --pod-network-cidr=10.217.0.0/16 to kubeadm init.

These commands will deploy Cilium with its own etcd managed by etcd operator.

Note: If you are running kubeadm in a single node please untaint it so that etcd-operator pods can be scheduled in the control-plane node.

kubectl taint nodes <node-name> node-role.kubernetes.io/master:NoSchedule-
To deploy Cilium you just need to run:
```
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.5/examples/kubernetes/1.14/cilium.yaml
```

### Calio

For more information about using Calico, see Quickstart for Calico on Kubernetes, Installing Calico for policy and networking, and other related resources.

For Calico to work correctly, you need to pass --pod-network-cidr=192.168.0.0/16 to kubeadm init or update the calico.yml file to match your Pod network. Note that Calico works on amd64, arm64, and ppc64le only.
```
kubectl apply -f https://docs.projectcalico.org/v3.8/manifests/calico.yaml
```

### Weave Net

Set /proc/sys/net/bridge/bridge-nf-call-iptables to 1 by running sysctl net.bridge.bridge-nf-call-iptables=1 to pass bridged IPv4 traffic to iptables’ chains. This is a requirement for some CNI plugins to work, for more information please see here.

The official Weave Net set-up guide is here.

Weave Net works on amd64, arm, arm64 and ppc64le without any extra action required. Weave Net sets hairpin mode by default. This allows Pods to access themselves via their Service IP address if they don’t know their PodIP.
```
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
```

