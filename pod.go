package main

import (
	//"bufio"
	"flag"
	"fmt"
	//"os"

	//extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podsClient := clientset.Pods(apiv1.NamespaceAll)

	// List Pods
	fmt.Printf("Listing pods in namespace %q:\n", apiv1.NamespaceAll)
	list, err := podsClient.List(metav1.ListOptions{
})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		containerID := d.Status.ContainerStatuses[0].ContainerID[9:]
		fmt.Printf(" Name:%s-------Namespace:%s------ContainerID:%s\n", d.Name, d.Namespace,containerID)
	}
}



