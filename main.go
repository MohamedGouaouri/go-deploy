package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Read the deployment YAML file
	filePath := "./tests/deployment.yaml"
	yamlFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer yamlFile.Close()

	// Decode the YAML into a Kubernetes Deployment object
	decoder := yaml.NewYAMLOrJSONDecoder(yamlFile, 100)
	var deployment appsv1.Deployment
	if err = decoder.Decode(&deployment); err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}

	// Set up Kubernetes client using kubeconfig
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
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
