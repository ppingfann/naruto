/*
opyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License
.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
    Unless required by applicable law or agreed to in writing, 
software
    distributed under the License is distributed on an "AS IS" 
BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expres
s or implied.
    See the License for the specific language governing permiss
ions and
    limitations under the License.
*/

// Note: the example only works with the code within the same r
//elease/branch.
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
	// Uncomment the following line to load the gcp plugin 
     //(only required to authenticate against GKE clusters).
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



