package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to kubeconfig")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig")
	}
	//deploymentName := flag.String("deployment", "", "deployment name")
	//imageName := flag.String("image", "", "image name")
	//appName := flag.String("app", "app", "application name")
	//flag.Parse()
	//if *deploymentName == "" {
	//	log.Print("[ERROR] you must specify the deployment name")
	//	os.Exit(1)
	//}
	//if *imageName == "" {
	//	log.Print("[ERROR] you must specify the image name")
	//	os.Exit(1)
	//}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deployment, err := clientset.AppsV1().Deployments("kube-system").Get(context.TODO(), "coredns", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(deployment.GetCreationTimestamp())
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
