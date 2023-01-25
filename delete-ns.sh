#!/bin/bash

kubectl get namespace <YOUR_NAMESPACE> -o json > <YOUR_NAMESPACE>.json
remove kubernetes from finalizers array which is under spec
kubectl replace --raw "/api/v1/namespaces/<YOUR_NAMESPACE>/finalize" -f ./<YOUR_NAMESPACE>.json
