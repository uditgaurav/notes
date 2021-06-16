# Podtato Head

It contains the steps to setup podtato head application. Credit to [ispeakc0de/datastore](https://github.com/ispeakc0de/datastore/) repository.

## Podtato pod delete with probes 

- Clone the repository and create the hello appliation. It will create the hello app in monitoring namespace by default.

_Create Target Namespace_
```bash
kubectl create ns <app-ns>(default monitoring) 
```

```bash
git clone https://github.com/uditgaurav/notes.git
cd notes/podtato-head/hello-service
kubectl apply -f app
```
_This will setup podtato along with black box exporter_

If you want to setup the application in different namespace then can the namespace in the files of `app` directory before creating it.

- Setup Prometheus for the hellow service

(_When using default namespace that is monitoring_)
```bash
kubectl apply -f notes/podtato-head/hello-service/prometheus-scrape-configuration/01-prometheus-rbac.yaml
```

(_For other namesapces_)

Modify the configmap file:
```bash
        static_configs:
        - targets:
          - helloservice.monitoring.svc.cluster.local:9000
        relabel_configs:
          - source_labels: [__address__]
            target_label: __param_target
```

to pick the correct namespace and port.
