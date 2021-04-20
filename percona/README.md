## Setup Percona XtraDB Cluster

#### Install the Percona XtraDB Cluster operator

```bash
git clone -b v1.6.0 https://github.com/percona/percona-xtradb-cluster-operator
cd percona-xtradb-cluster-operator
kubectl apply -f deploy/bundle.yaml
```

- Verify if the operator is running correctly

```
$ kubectl get pod
NAME                                               READY   STATUS    RESTARTS   AGE
percona-xtradb-cluster-operator-749b86b678-8f4q5   1/1     Running   0          23s
```
 
#### Update Storage and Monitoring specification

In this document, we have made changes in the storage section for PXC and the monitoring section PMM.

Changes done in the Storage section for PXC: 

Update Storage Class name and required storage parameters in deploy/cr.yaml. In this example, we have updated 
`spec.pxc.volumeSpec.persistentVolumeClaim.storageClassName` as “openebs-device” and 
`spec.pxc.volumeSpec.persistentVolumeClaim.resources.requests.storage` as “100Gi”

```yaml
Sample snippet:
    volumeSpec:
      persistentVolumeClaim:
        storageClassName: openebs-device
        accessModes: [ "ReadWriteOnce" ]
        resources:
          requests:
            storage: 100Gi
```
#### Changes done in the Monitoring section for PMM:

- Enable monitoring service and server user name. In this example, we have updated 
`spec.pmm.enabled` as “true”
`spec.pmm.serverUser` as “admin”  

The following is the sample snippet of PVC spec of Percona XtraDB where the Storage Class name and storage capacity has been changed.

```yaml
  pmm:
    enabled: true
    image: percona/percona-xtradb-cluster-operator:1.6.0-pmm
    serverHost: monitoring-service
    serverUser: admin
```

### Install the Percona XtraDB Cluster

There is a dependency if you are enabling a monitoring service(PMM) for your PXC. In this case, you must install the PMM server  using the following command before installing PXC. We have used Percona blog to enable the monitoring service.

Use helm to install PMM Server

Using helm, add the Percona chart repository and update the information for the available charts as follows:

```
helm repo add percona https://percona-charts.storage.googleapis.com
helm repo update
helm install monitoring percona/pmm-server --set platform=kubernetes --version 2.7.0 --set "credentials.password=test123" 
```
Note: In this document, we have used “test123” as the PMM server credential password and base64 encoded form is “dGVzdDEyMw==”. This encoded value will be added in one of the secrets while installing the PXC cluster and when accessing the performance benchmark task.

Now, verify PMM server pod is installed and running.

```
$ kubectl get pod
NAME                                               READY   STATUS    RESTARTS   AGE
monitoring-0                                       1/1     Running   0          70m
```
In the previous section, we have made the required changes on the CR YAML spec. Let’s install the PXC cluster using the following command.

```
$ kubectl apply -f deploy/cr.yaml
```
After applying the above command, you may see that cluster1-pxc-0 pod started in CreateContainerConfigError state. 
```
$  kubectl get pod
NAME                                               READY   STATUS                       RESTARTS   AGE
cluster1-haproxy-0                                 1/2     Running                      0          21s
cluster1-pxc-0                                     0/2     CreateContainerConfigError   0          21s
monitoring-0                                       1/1     Running                      0          107m
percona-xtradb-cluster-operator-749b86b678-8f4q5   1/1     Running                      0          3h50m
```
This is due to the unavailability of the PMM server key in the secret. To resolve this, edit the corresponding secret and add the PMM server key. 

```
$ kubectl get secrets
NAME                                          TYPE                                  DATA   AGE
default-token-s78cq                           kubernetes.io/service-account-token   3      5h15m
internal-cluster1                             Opaque                                6      65s
my-cluster-secrets                            Opaque                                6      65s
my-cluster-ssl                                kubernetes.io/tls                     3      61s
my-cluster-ssl-internal                       kubernetes.io/tls                     3      60s
percona-xtradb-cluster-operator-token-v82b5   kubernetes.io/service-account-token   3      3h50m
sh.helm.release.v1.monitoring.v1              helm.sh/release.v1                    1      107m
```
Let’s edit the secret internal-cluster1 using the following command and add pmmserver value as per the given credential password during PMM server installation time. In this example, we have added  pmmserver: dGVzdDEyMw==  in the secret internal-cluster1. 

```
$ kubectl edit secret internal-cluster1
```
Sample spec of the modified secret content.

```yaml
apiVersion: v1
data:
  clustercheck: dUt5QVlMYTVKdWxaZDA1NGI=
  monitor: NkRwbFVJcExCSFFFSHBMM3k=
  operator: OGsyMmxhaG02blh0aW9BbkFW
  proxyadmin: cVB6elZHZXUwVWNkaUV4MTJp
  root: WnV2cFNiRGU4UWhpWjNmd1Y=
  pmmserver: dGVzdDEyMw==
  xtrabackup: MmVGSGsyWTlJdk44ZUlmQXlnYQ==
kind: Secret
```

Now, verify that all required components are installed and running successfully.

```
$ kubectl get pod
NAME                                               READY   STATUS    RESTARTS   AGE
cluster1-haproxy-0                                 2/2     Running   0          4m33s
cluster1-haproxy-1                                 2/2     Running   0          2m40s
cluster1-haproxy-2                                 2/2     Running   0          2m21s
cluster1-pxc-0                                     2/2     Running   0          4m33s
cluster1-pxc-1                                     2/2     Running   0          2m46s
cluster1-pxc-2                                     2/2     Running   0          92s
monitoring-0                                       1/1     Running   0          111m
percona-xtradb-cluster-operator-749b86b678-8f4q5   1/1     Running   0          3h54m
```

```
$ kubectl get pvc
NAME                     STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS     AGE
datadir-cluster1-pxc-0   Bound    pvc-3c0ef6d5-6d04-469b-aad1-e4a6881a5176   10Gi       RWO            openebs-device   4m49s
datadir-cluster1-pxc-1   Bound    pvc-92ed74a0-ccf7-48e1-8b5a-721ea87d4282   10Gi       RWO            openebs-device   3m2s
datadir-cluster1-pxc-2   Bound    pvc-1f9f619c-998d-4f60-9ed8-7b785c5cb49e   10Gi       RWO            openebs-device   108s
pmmdata-monitoring-0     Bound    pvc-02514e16-4f3f-4123-96ae-7f609c9377f7   8Gi        RWO            gp2              3h42m
```

```
$ kubectl get pv
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                            STORAGECLASS     REASON   AGE
pvc-02514e16-4f3f-4123-96ae-7f609c9377f7   8Gi        RWO            Delete           Bound    default/pmmdata-monitoring-0     gp2                       3h42m
pvc-1f9f619c-998d-4f60-9ed8-7b785c5cb49e   10Gi       RWO            Delete           Bound    default/datadir-cluster1-pxc-2   openebs-device            117s
pvc-3c0ef6d5-6d04-469b-aad1-e4a6881a5176   10Gi       RWO            Delete           Bound    default/datadir-cluster1-pxc-0   openebs-device            4m59s
pvc-92ed74a0-ccf7-48e1-8b5a-721ea87d4282   10Gi       RWO            Delete           Bound    default/datadir-cluster1-pxc-1   openebs-device            3m11s
```
```
$ kubectl get svc
NAME                        TYPE           CLUSTER-IP       EXTERNAL-IP                                                                    PORT(S)                       AGE
cluster1-haproxy            ClusterIP      10.100.136.90    <none>                                                                         3306/TCP,3309/TCP,33062/TCP   5m18s
cluster1-haproxy-replicas   ClusterIP      10.100.244.115   <none>                                                                         3306/TCP                      5m18s
cluster1-pxc                ClusterIP      None             <none>                                                                         3306/TCP,33062/TCP            5m18s
cluster1-pxc-unready        ClusterIP      None             <none>                                                                         3306/TCP,33062/TCP            5m18s
kubernetes                  ClusterIP      10.100.0.1       <none>                                                                         443/TCP                       5h19m
monitoring-service          LoadBalancer   10.100.32.246    a543e9e1d189644f9bf4f7fdf0ba15b3-1159960729.ap-southeast-1.elb.amazonaws.com   443:30317/TCP                 112m
```
