## Docker Images Used:

#### LitmusPortal Components:

```bash
    Image:         argoproj/argocli:v2.9.3
    Image:         litmuschaos/litmusportal-event-tracker:2.0.0-Beta4
    Image:         litmuschaos/litmusportal-frontend:2.0.0-Beta4
    Image:         mongo:4.2.8
    Image:         litmuschaos/litmusportal-server:2.0.0-Beta4
    Image:         litmuschaos/litmusportal-auth-server:2.0.0-Beta4
    Image:         litmuschaos/litmusportal-subscriber:2.0.0-Beta4
    Image:         argoproj/workflow-controller:v2.9.3
```    

#### Litmus Infra Components:
```bash
    Image:         litmuschaos/chaos-runner:1.13.3
    Image:         litmuschaos/chaos-operator:1.13.3
```    

#### Experiment Image
```bash
    Image:         litmuschaos/go-runner:1.13.3
```

#### Monitoring
```bash
    Image:         litmuschaos/chaos-exporter:1.13.3
```
#### (Optional)
```bash
    Image:         grafana/grafana:latest
    Image:         quay.io/prometheus/prometheus:v2.19.2
    Image:         quay.io/coreos/prometheus-config-reloader:v0.42.1
    Image:         jimmidyson/configmap-reload:v0.5.0
    Image:         quay.io/coreos/prometheus-operator:v0.40.0
    Image:         quay.io/coreos/kube-rbac-proxy:v0.5.0
```
