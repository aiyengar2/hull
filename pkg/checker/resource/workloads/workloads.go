package workloads

import (
	"github.com/aiyengar2/hull/pkg/checker"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
)

func init() {
	if err := appsv1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
	if err := autoscalingv1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
	if err := batchv1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
	if err := corev1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
	if err := policyv1beta1.AddToScheme(checker.Scheme); err != nil {
		panic(err)
	}
}

type CronJobs []*batchv1.CronJob
type DaemonSets []*appsv1.DaemonSet
type Deployments []*appsv1.Deployment
type Jobs []*batchv1.Job
type StatefulSets []*appsv1.StatefulSet

type ConfigMaps []*corev1.ConfigMap
type Secrets []*corev1.Secret

type PodSecurityPolicies []*policyv1beta1.PodSecurityPolicy
type HorizontalPodAutoscalers []*autoscalingv1.HorizontalPodAutoscaler
