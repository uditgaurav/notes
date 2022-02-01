## Monitoring Setup on KIND Cluster

- Create a KIND Cluster - 

```
Kind create cluster --name <CLUSTER-NAME>
```

- Command

```
kubectl apply -f https://gist.githubusercontent.com/Jonsy13/f2d1c585ea32c8d23a5ddd2ebe5129ac/raw/709ada193e5da6173169d15f7291a352b9898b98/components.yaml
```

For monitoring setup, we will need a metrics-server, Which is not installed by default in KIND Cluster, so we need to install the same. 

Either you can use this modified manifest from here
https://gist.github.com/Jonsy13/f2d1c585ea32c8d23a5ddd2ebe5129ac
