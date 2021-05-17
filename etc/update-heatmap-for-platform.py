from pandas import DataFrame
import matplotlib.pyplot as plt
import seaborn as sns
import csv,sys

ExperimentName=sys.argv[1]

with open(ExperimentName+'.csv', newline='') as f:
    reader = csv.reader(f)
    data = list(reader)

if ExperimentName == "pod-delete":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)','N/A']
    Cols = ['Experiment Supported']
    plt.title("Pod Delete Experiment", fontsize =10)
elif ExperimentName == "container-kill":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Container Kill Experiment", fontsize =10)
elif ExperimentName == "disk-fill":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Disk Fill Experiment", fontsize =10)
elif ExperimentName == "pod-cpu-hog":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod CPU Hog Experiment", fontsize =10)   
elif ExperimentName == "pod-memory-hog":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Memory Hog Experiment", fontsize =10)
elif ExperimentName == "pod-network-corruption":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Corruption Experiment", fontsize =10)
elif ExperimentName == "pod-network-duplication":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Duplication Experiment", fontsize =10)
elif ExperimentName == "pod-network-latency":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Latency Experiment", fontsize =10)    
elif ExperimentName == "pod-network-loss":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Loss Experiment", fontsize =10)
elif ExperimentName == "pod-autoscaler":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod Autoscaler Experiment", fontsize =10)
elif ExperimentName == "kubelet-service-kill":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Kubelet Service Kill", fontsize =10)
elif ExperimentName == "node-cpu-hog":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Node CPU Hog", fontsize =10)
elif ExperimentName == "node-memory-hog":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Node Memory Hog", fontsize =10)
elif ExperimentName == "node-drain":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Node Drain Experiment", fontsize =10)   
elif ExperimentName == "node-taint":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Node Taint Experiment", fontsize =10) 
elif ExperimentName == "node-io-stress":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Node IO Stress", fontsize =10)         
elif ExperimentName == "pod-io-stress":
    Index = ['GKE','Konvoy','Packet(Kubeadm)','Minikube','EKS','AKS','Kind','Rancher','OpenShift(OKD)','k3s','microk8s','AWS(KOPS)']
    Cols = ['Experiment Supported']
    plt.title("Pod IO Stress", fontsize =10)        
else:
    print("Experiment %s not supported",ExperimentName)

df = DataFrame(data, index=Index, columns=Cols)
df = df[df.columns].astype(float)

print(df)
svm = sns.heatmap(df, cmap="Reds")
figure = svm.get_figure()
svm.set_xticklabels(svm.get_xmajorticklabels(), fontsize = 18)
svm.set_yticklabels(svm.get_ymajorticklabels(), fontsize = 15)
plt.subplots_adjust(left=0.278,bottom=0.095,right=0.9,top=0.88,wspace=0.2,hspace=0.2)
figure.set_figheight(10)
figure.set_figwidth(15)
plt.savefig(ExperimentName+'-platform-heatmap.png', dpi=450)
