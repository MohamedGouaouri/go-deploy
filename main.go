package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"path/filepath"

	"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Read the deployment YAML file from the local file system
	filePath := "./tests/deployment.yaml" // Path to your deployment YAML file
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Parse the YAML into a Kubernetes Deployment object
	var deployment appsv1.Deployment
	err = yaml.Unmarshal(yamlFile, &deployment)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}

	// Set up the Kubernetes client
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config") // Path to your kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Deploy the parsed deployment object to the Kubernetes cluster
	deploymentsClient := clientset.AppsV1().Deployments(deployment.Namespace)
	result, err := deploymentsClient.Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating deployment: %v", err)
	}

	fmt.Printf("Deployment %s created successfully in namespace %s.\n", result.Name, result.Namespace)
}
