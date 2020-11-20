from pandas import DataFrame
import matplotlib.pyplot as plt
import seaborn as sns
import csv,sys

ExperimentName=sys.argv[1]

with open(ExperimentName+'.csv', newline='') as f:
    reader = csv.reader(f)
    data = list(reader)

if ExperimentName == "pod-delete":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Delete Experiment", fontsize =20)
elif ExperimentName == "container-kill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Container Kill Experiment", fontsize =20)
elif ExperimentName == "disk-fill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Disk Fill Experiment", fontsize =20)
elif ExperimentName == "pod-cpu-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod CPU Hog Experiment", fontsize =20)   
elif ExperimentName == "pod-memory-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Memory Hog Experiment", fontsize =20)
elif ExperimentName == "pod-network-corruption":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Corruption Experiment", fontsize =20)
elif ExperimentName == "pod-network-duplication":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Duplication Experiment", fontsize =20)
elif ExperimentName == "pod-network-latency":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Latency Experiment", fontsize =20)    
elif ExperimentName == "pod-network-loss":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Loss Experiment", fontsize =20)
elif ExperimentName == "pod-autoscaler":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Autoscaler Experiment", fontsize =20)
elif ExperimentName == "kubelet-service-kill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Kubelet Service Kill", fontsize =20)
elif ExperimentName == "node-cpu-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node CPU Hog", fontsize =20)
elif ExperimentName == "node-memory-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Memory Hog", fontsize =20)
elif ExperimentName == "node-drain":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Drain Experiment", fontsize =20)   
elif ExperimentName == "node-taint":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Taint Experiment", fontsize =20) 
elif ExperimentName == "node-io-stress":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node IO Stress", fontsize =20)         
elif ExperimentName == "pod-io-stress":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod IO Stress", fontsize =20)        
else:
    print("Experiment %s not supported",ExperimentName)

df = DataFrame(data, index=Index, columns=Cols)
df = df[df.columns].astype(float)

print(df)
svm = sns.heatmap(df, cmap="Reds")
figure = svm.get_figure()

plt.subplots_adjust(left=0.218,bottom=0.095,right=0.9,top=0.88,wspace=0.2,hspace=0.2)
figure.set_figheight(10)
figure.set_figwidth(15)
plt.savefig(ExperimentName+'-runtime-heatmap.png', dpi=350)
from pandas import DataFrame
import matplotlib.pyplot as plt
import seaborn as sns
import csv,sys

ExperimentName=sys.argv[1]

with open(ExperimentName+'.csv', newline='') as f:
    reader = csv.reader(f)
    data = list(reader)

if ExperimentName == "pod-delete":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Delete Experiment", fontsize =20)
elif ExperimentName == "container-kill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Container Kill Experiment", fontsize =20)
elif ExperimentName == "disk-fill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Disk Fill Experiment", fontsize =20)
elif ExperimentName == "pod-cpu-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod CPU Hog Experiment", fontsize =20)   
elif ExperimentName == "pod-memory-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Memory Hog Experiment", fontsize =20)
elif ExperimentName == "pod-network-corruption":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Corruption Experiment", fontsize =20)
elif ExperimentName == "pod-network-duplication":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Duplication Experiment", fontsize =20)
elif ExperimentName == "pod-network-latency":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Latency Experiment", fontsize =20)    
elif ExperimentName == "pod-network-loss":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Network Loss Experiment", fontsize =20)
elif ExperimentName == "pod-autoscaler":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod Autoscaler Experiment", fontsize =20)
elif ExperimentName == "kubelet-service-kill":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Kubelet Service Kill", fontsize =20)
elif ExperimentName == "node-cpu-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node CPU Hog", fontsize =20)
elif ExperimentName == "node-memory-hog":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Memory Hog", fontsize =20)
elif ExperimentName == "node-drain":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Drain Experiment", fontsize =20)   
elif ExperimentName == "node-taint":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node Taint Experiment", fontsize =20) 
elif ExperimentName == "node-io-stress":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Node IO Stress", fontsize =20)         
elif ExperimentName == "pod-io-stress":
    Index = ['Cluster Runtime Docker','Cluster Runtime Containerd','Cluster Runtime CRIO']
    Cols = ['Experiment Supported']
    plt.title("Pod IO Stress", fontsize =20)        
else:
    print("Experiment %s not supported",ExperimentName)

df = DataFrame(data, index=Index, columns=Cols)
df = df[df.columns].astype(float)

print(df)
svm = sns.heatmap(df, cmap="Reds")
figure = svm.get_figure()

plt.subplots_adjust(left=0.218,bottom=0.095,right=0.9,top=0.88,wspace=0.2,hspace=0.2)
figure.set_figheight(10)
figure.set_figwidth(15)
plt.savefig(ExperimentName+'-runtime-heatmap.png', dpi=350)
