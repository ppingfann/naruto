package main

import (
	//"bufio"
	"flag"
	"fmt"
	//"os"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io"
	"io/ioutil"
	//extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin 
     //(only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	//"strings"
)


func main() {
	router := httprouter.New()

	router.POST("/api/v1/containerID",func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
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

		body,_ := ioutil.ReadAll(io.LimitReader(r.Body,1122334))

		for _, d := range list.Items {
			containerID := d.Status.ContainerStatuses[0].ContainerID[9:]
			if string(body)==containerID{
				fmt.Fprintf(w,"Name:%s--Namespace:%s--ContainerID:%s\n", d.Name, d.Namespace, containerID)
			}
		}
	})




	http.ListenAndServe(":8082", router)
}



