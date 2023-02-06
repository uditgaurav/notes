#!/bin/bash
echo "Caution - This script will delete the namespace litmus "
kubectl config current-context
echo "Do you want to cleanup Litmus for this cluster"
echo  'Press "0" & Enter to continue Or Press "any other key" & Enter to Abort'
read line
if [ "$line" -eq "0" ]
then
	echo " Cleanup Started" 
	echo " Deleting ChaosEngines"
	# kubectl delete chaosengines.litmuschaos.io --all -n litmus
	echo " Deleting ChaosResults"
	kubectl delete chaosresults.litmuschaos.io --all -n litmus
	echo " Deleting ChaosExperiments"
	kubectl delete chaosexperiments.litmuschaos.io --all -n litmus
	echo " Deleting Litmus Namespace"
	kubectl delete ns litmus
	echo " Deleting Clusterroles"
	kubectl get clusterroles | grep -i "argo\|litmus\|subscriber-cluster\|chaos-cluster-role" | awk '{print $1}' | xargs -n 1 kubectl delete clusterroles
	echo " Deleting Cluster Role Binding"
	kubectl get clusterrolebinding | grep -i "argo\|litmus\|subscriber-cluster\|chaos-cluster-role-binding" | awk '{print $1}' | xargs -n 1 kubectl delete clusterrolebinding
	echo " Cleanup Done" 
else
	echo " Aborting Cleanup" 
fi
exit
