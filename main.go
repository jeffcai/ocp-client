package main

import (
	"fmt"

	buildv1client "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// Determine the Namespace referenced by the current context in the
	// kubeconfig file.
	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		panic(err)
	}

	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}

	// Create a Kubernetes core/v1 client.
	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	// Create an OpenShift build/v1 client.
	buildclient, err := buildv1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	// List all Pods in our current Namespace.
	pods, err := coreclient.Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pods in namespace %s:\n", namespace)
	for _, pod := range pods.Items {
		fmt.Printf("  %s\n", pod.Name)
	}

	// List all ConfigMaps
	cms, err := coreclient.ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("ConfigMaps in namespace %s:\n", namespace)
	for _, cm := range cms.Items {
		fmt.Printf("  %s\n", cm.Name)
	}

	// List all Services
	svcs, err := coreclient.Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Services in namespace %s:\n", namespace)
	for _, svc := range svcs.Items {
		fmt.Printf("  %s\n", svc.Name)
	}

	// List all Endpoints
	eps, err := coreclient.Endpoints(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Endpoints in namespace %s:\n", namespace)
	for _, ep := range eps.Items {
		fmt.Printf("  %s\n", ep.Name)
	}

	// List all Builds in our current Namespace.
	builds, err := buildclient.Builds(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Builds in namespace %s:\n", namespace)
	for _, build := range builds.Items {
		fmt.Printf("  %s\n", build.Name)
	}
}
