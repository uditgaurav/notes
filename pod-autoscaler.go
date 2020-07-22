package pod_autoscaler

import (
	"fmt"
	"strconv"
	"time"

	clients "github.com/litmuschaos/litmus-go/pkg/clients"
	"github.com/litmuschaos/litmus-go/pkg/events"
	experimentTypes "github.com/litmuschaos/litmus-go/pkg/generic/pod-autoscaler/types"
	"github.com/litmuschaos/litmus-go/pkg/log"
	"github.com/litmuschaos/litmus-go/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	retries "k8s.io/client-go/util/retry"
	"k8s.io/klog"

	"github.com/pkg/errors"
)

var err error

//PreparePodAutoscaler contains the prepration steps before chaos injection
func PreparePodAutoscaler(experimentsDetails *experimentTypes.ExperimentDetails, clients clients.ClientSets, resultDetails *types.ResultDetails, eventsDetails *types.EventDetails, chaosDetails *types.ChaosDetails) error {

	//Waiting for the ramp time before chaos injection
	if experimentsDetails.RampTime != 0 {
		log.Infof("[Ramp]: Waiting for the %vs ramp time before injecting chaos", strconv.Itoa(experimentsDetails.RampTime))
		waitForRampTime(experimentsDetails)
	}
	if err != nil {
		return errors.Errorf("Unable to get the serviceAccountName, err: %v", err)
	}

	err = PodAutoscalerChaos(experimentsDetails, clients, eventsDetails, chaosDetails)

	if err != nil {
		return errors.Errorf("Unable to perform autoscaling, due to %v", err)
	}

	//Waiting for the ramp time after chaos injection
	if experimentsDetails.RampTime != 0 {
		log.Infof("[Ramp]: Waiting for the %vs ramp time after injecting chaos", strconv.Itoa(experimentsDetails.RampTime))
		waitForRampTime(experimentsDetails)
	}
	return nil
}

//waitForRampTime waits for the given ramp time duration (in seconds)
func waitForRampTime(experimentsDetails *experimentTypes.ExperimentDetails) {
	time.Sleep(time.Duration(experimentsDetails.RampTime) * time.Second)
}

//PodAutoscalerChaos scales up the pod replicas
func PodAutoscalerChaos(experimentsDetails *experimentTypes.ExperimentDetails, clients clients.ClientSets, eventsDetails *types.EventDetails, chaosDetails *types.ChaosDetails) error {

	deploymentsClient := clients.KubeClient.AppsV1().Deployments(experimentsDetails.ChaosNamespace)

	// List Deployments
	var deploymentName string

	log.Infof("Listing deployments in namespace %q:\n", experimentsDetails.ChaosNamespace)
	deploymentList, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		return errors.Errorf("Unable to get the list of deployments, err: %v", err)
	}
	for _, deployment := range deploymentList.Items {
		deploymentName = deployment.Name
		fmt.Printf(" * %s (%d replicas)\n", deployment.Name, *deployment.Spec.Replicas)

	}
	replicas := int32(experimentsDetails.Replicas)

	// Scale Deployments
	retryErr := retries.RetryOnConflict(retries.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
		if err != nil {
			return errors.Errorf("Failed to get latest version of Deployment: %v", err)
		}

		result.Spec.Replicas = int32Ptr(replicas) // modify replica count
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		return errors.Errorf("Unable to scale the deployment, due to: %v", retryErr)
	}
	log.Info("Deployment Started Scaling")

	if experimentsDetails.EngineName != "" {
		types.SetEngineEventAttributes(eventsDetails, types.PreChaosCheck, "Injecting "+experimentsDetails.ExperimentName+" chaos on "+deploymentName+" deployment", chaosDetails)
		events.GenerateEvents(eventsDetails, clients, chaosDetails, "ChaosEngine")
	}

	err = ApplicationPodStatusCheck(experimentsDetails.ChaosDuration, experimentsDetails.AppNS, experimentsDetails.AppLabel, clients)
	if err != nil {
		return errors.Errorf("Status Check failed, err: %v", err)
	}

	return nil
}

// ApplicationPodStatusCheck checks the status of the application pod
func ApplicationPodStatusCheck(ChaosDuration int, appNs string, deploymentName string, clients clients.ClientSets) error {

	//ChaosStartTimeStamp contains the start timestamp, when the chaos injection begin
	ChaosStartTimeStamp := time.Now().Unix()
	failFlag := false

	// err := retry.
	// 	Times(uint(ChaosDuration / 2)).
	// 	Wait( * time.Second).
	// 	Try(func(attempt uint) error {
	// 		podSpec, err := clients.KubeClient.CoreV1().Pods(appNs).List(metav1.ListOptions{LabelSelector: appLabel})
	// 		if err != nil || len(podSpec.Items) == 0 {
	// 			return errors.Errorf("Unable to get the pod, err: %v", err)
	// 		}
	// 		err = nil
	// 		for _, pod := range podSpec.Items {
	// 			if string(pod.Status.Phase) != "Running" {
	// 				return errors.Errorf("Pod is not yet in running state")
	// 			}
	// 			log.InfoWithValues("The running status of Pods are as follows", logrus.Fields{
	// 				"Pod": pod.Name, "Status": pod.Status.Phase})

	// 			//ChaosCurrentTimeStamp contains the current timestamp
	// 			ChaosCurrentTimeStamp := time.Now().Unix()

	// 			//ChaosDiffTimeStamp contains the difference of current timestamp and start timestamp
	// 			//It will helpful to track the total chaos duration
	// 			chaosDiffTimeStamp := ChaosCurrentTimeStamp - ChaosStartTimeStamp

	// 			if int(chaosDiffTimeStamp) >= ChaosDuration {
	// 				failFlag = true
	// 				break
	// 			}
	// 		}
	// 		if failFlag == true {
	// 			log.Info("[Info]: Application pods fail to come in running state after Chaos Duration of %i sec")
	// 			return errors.Errorf("Application pods fail to come in running state after Chaos Duration of %d sec", ChaosDuration)
	// 		}
	// 		return nil
	// 	})
	// if err != nil {
	// 	return err
	// }

	sampleApp, _ := clients.KubeClient.AppsV1().Deployments(appNs).Get(deploymentName, metav1.GetOptions{})
	count := 0
	for sampleApp.Status.UnavailableReplicas != 0 {
		if count < 20 {
			klog.Infof("Application is Creating, Currently Unavaliable Count is: %v \n", sampleApp.Status.UnavailableReplicas)
			sampleApp, _ = clients.KubeClient.AppsV1().Deployments(appNs).Get(deploymentName, metav1.GetOptions{})
			time.Sleep(5 * time.Second)
			count++
		} else {
			return errors.Wrapf(err, "%v deployment fail to get in Running state, due to:%v", deploymentName, err)
		}
	}

	return nil
}

func int32Ptr(i int32) *int32 { return &i }
