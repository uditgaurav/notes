## Create OpenEBS [Storage Class](https://docs.openebs.io/docs/next/uglocalpv-device.html#create-storageclass)

- To create your own StorageClass to customize how Local PV with devices are created. For instance, if you would like to run MongoDB stateful applications with Local PV device, you would want to set the default filesystem as xfs and/or also dedicate some devices on node that you want to use for Local PV. Save the following StorageClass definition as `local-device-sc.yaml`.

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-device
  annotations:
    openebs.io/cas-type: local
    cas.openebs.io/config: |
      - name: StorageType
        value: device
      - name: FSType
        value: xfs
      - name: BlockDeviceTag
        value: "mongo"
provisioner: openebs.io/local
reclaimPolicy: Delete
volumeBindingMode: WaitForFirstConsumer
```

**NOTE**
```
The volumeBindingMode MUST ALWAYS be set to WaitForFirstConsumer. volumeBindingMode: WaitForFirstConsumer instructs Kubernetes to initiate
the creation of PV only after Pod using PVC is scheduled to the node.
```

- Create OpenEBS Local PV Device Storage Class.

```bash
kubectl apply -f local-device-sc.yaml
```

- Verify that the StorageClass is successfully created.

```yaml
kubectl get sc local-device -o yaml
```

