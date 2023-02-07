@echo off
Set namespace=litmus
Echo Caution - This script will delete the namespace %namespace%
kubectl config current-context
echo DO YOU WANT TO CONTINUE WITH THE CLEANUP (y/n)
set /p Input=Enter y or n:
if /I "%Input%"=="y" goto yes
goto no
:yes
echo Cleanup Started
echo Deleting ChaosEngines
kubectl delete chaosengines.litmuschaos.io --all -n %namespace%
echo Deleting ChaosResults
kubectl delete chaosresults.litmuschaos.io --all -n %namespace%
echo Deleting ChaosExperiments
kubectl delete chaosexperiments.litmuschaos.io --all -n %namespace%
echo Deleting %namespace% Namespace
kubectl delete ns %namespace%
echo Deleting Clusterroles
for /f "tokens=1" %%a in ('kubectl get clusterroles ^| findstr "argo litmus subscriber-cluster chaos-cluster-role"') do kubectl delete clusterroles %%a
echo Deleting Cluster Role Binding
for /f "tokens=1" %%a in ('kubectl get clusterrolebinding ^| findstr "argo litmus subscriber-cluster chaos-cluster-role-binding"') do kubectl delete clusterrolebinding %%a
echo Cleanup Done
pause
:no
echo Done !!!
