## Setup Kube-Slack

#### Create an incoming webhook:

- Setup [Slack Webhook](https://api.slack.com/messaging/webhooks) and add integrate it with your slack as shown below:

1. click on the gears button (Channel Settings) near the search box.
2. Select "Add an app or integration":- Create a new app if you don't have one.
3. Search for "Incoming WebHooks"
4. Click on "Add configuration"
5. Select the channel you want the bot to post to and submit.
6. You can customize the icon and name if you want.

**Note:** Take note of the "Webhook URL". This will be something like `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX`

- If your kubernetes uses RBAC, you should apply the following manifest as well:


```yaml
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kube-slack
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-slack
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kube-slack
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: kube-slack
subjects:
  - kind: ServiceAccount
    name: kube-slack
    namespace: kube-system
```
- Load this Deployment into your Kubernetes. Make sure you set `SLACK_URL` to the Webhook URL and uncomment serviceAccountName if you use RBAC also setup username in`SLACK_USERNAME` env.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-slack
  namespace: kube-system
spec:
  replicas: 1
  revisionHistoryLimit: 3
  selector:
    matchLabels:
      app: kube-slack
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
      name: kube-slack
      labels:
        app: kube-slack
    spec:
      serviceAccountName: kube-slack
      containers:
      - name: kube-slack
        image: willwill/kube-slack:v4.2.0
        env:
        - name: SLACK_URL
          value: https://hooks.slack.com/services/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
        - name: SLACK_USERNAME
          value: 'LitmusCI'
        resources:
          requests:
            memory: 30M
            cpu: 5m
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      - key: CriticalAddonsOnly
        operator: Exists
```
