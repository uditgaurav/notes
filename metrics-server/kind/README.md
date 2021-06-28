## Monitoring Setup on KIND Cluster

- Create a KIND Cluster - 

```
Kind create cluster --name <CLUSTER-NAME>
```

For monitoring setup, we will need a metrics-server, Which is not installed by default in KIND Cluster, so we need to install the same. 

Either you can use this modified manifest from here
https://gist.github.com/Jonsy13/f2d1c585ea32c8d23a5ddd2ebe5129ac
