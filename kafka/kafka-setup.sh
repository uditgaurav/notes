#!/bin/bash

echo "***********************************************"
echo "              Installing kafka                 "
echo "***********************************************"

kubectl create ns kafka

VERSION=0.12.0
OS=$(uname | tr '[:upper:]' '[:lower:]')
wget -O kubectl-kudo https://github.com/kudobuilder/kudo/releases/download/v${VERSION}/kubectl-kudo_${VERSION}_${OS}_x86_64

chmod +x kubectl-kudo

sudo mv kubectl-kudo /usr/local/bin/kubectl-kudo
kubectl kudo init
kubectl kudo install zookeeper --instance=zookeeper-instance -n kafka
kubectl kudo install kafka --instance=kafka -p ADD_SERVICE_MONITOR=true -n kafka

n=0
until [ "$n" -ge 300 ]; do
        retries=$((30-n))
	zookeeperReadyReplica=$(kubectl get sts zookeeper-instance-zookeeper -n kafka -o jsonpath='{.status.readyReplicas}')
	echo "Number Of Ready Replicas for Zookeeper pods: $zookeeperReadyReplica, Retries left: $retries"
	if [ "$zookeeperReadyReplica" == "3" ]; then
        kafkaReadyReplica=$(kubectl get sts kafka-kafka -n kafka -o jsonpath='{.status.readyReplicas}')
        echo "Number Of Ready Replicas for Kafka pods: $kafkaReadyReplica, Retries left: $retries"
        	if [ "$kafkaReadyReplica" == "3" ]; then
			break;
		fi
	fi
	n=$((n+1))
	sleep 10
done
kubectl get pods -n kafka

echo "***********************************************"
echo "              Installing Prometheus                 "
echo "***********************************************"
git clone https://github.com/litmuschaos/litmus.git
kubectl create ns monitoring
kubectl apply -f litmus/monitoring/utils/prometheus/prometheus-operator -n monitoring
sleep 5
rm -rf litmus/monitoring/utils/prometheus/prometheus-configuration/prometheus.yaml
kubectl apply -f litmus/monitoring/utils/prometheus/prometheus-configuration -n monitoring
sleep 5
kubectl get pods -n monitoring
kubectl get svc -n monitoring

echo "***********************************************"
echo "              Installing Grafana                 "
echo "***********************************************"
kubectl apply -f litmus/monitoring/utils/grafana -n monitoring
sleep 2
kubectl get svc -n monitoring

echo "***********************************************"
echo "              Kafka and Chaos exporter          "
echo "***********************************************"
git clone https://github.com/chaoscarnival/bootcamps.git
kubectl apply -f bootcamps/day1-kafkaChaos/service-monitors/chaos-exporter-service-monitor.yaml
kubectl apply -f https://raw.githubusercontent.com/chaoscarnival/bootcamps/main/day1-kafkaChaos/service-monitors/kafka-exporter-service-monitor.yaml

kubectl apply -f bootcamps/day1-kafkaChaos/chaos-exporter -n litmus

cd bootcamps/day1-kafkaChaos/kafka-exporter-helm/charts/kafka-exporter
helm install  -f values.yaml kafka-exporter --namespace=kafka .
sleep 2
kubectl label servicemonitor.monitoring.coreos.com kafka-monitor -n kafka k8s-app=kafka-exporter
cd ../../..
kubectl apply -f prometheus
cd ../..
rm -rf litmus
rm -rf bootcamps
echo "***********************************************"
echo "              Final Preview          "
echo "***********************************************"

echo "##################"
echo "## pods in kafka ns  ##"
echo "##################"
echo ""
kubectl get pods -n kafka

echo "############################"
echo "## pods in monitoring ns  ##"
echo "############################"
echo ""
kubectl get pods -n monitoring

echo "########################"
echo "## pods in litmus ns  ##"
echo "########################"
echo ""
kubectl get pods -n litmus


echo "################################"
echo "## services in monitoring ns  ##"
echo "################################"
echo ""
kubectl get svc -n monitoring

echo "Thank You Your Kafka Setup is Ready!!!"

