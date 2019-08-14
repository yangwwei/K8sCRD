package main

import (
	"github.com/yangwwei/K8sCRD/pkg/apis/example/v1alpha1"
	"github.com/yangwwei/K8sCRD/pkg/client/clientset/versioned"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

// FIXME
// This code is moved from my another repo, it's not tested so it may not work : )
// This is just to demostrate how to write a golang program to create CRDs.
func main() {
	// create CRDs
	crdVersions := make([]apiextensions.CustomResourceDefinitionVersion, 0)
	crdVersions = append(crdVersions, apiextensions.CustomResourceDefinitionVersion{
		Name:    "v1",
		Served:  true,
		Storage: true,
	})

	client := clientset.NewForConfigOrDie(ctx.kubeClient.GetConfig())
	crd := &apiextensions.CustomResourceDefinition{
		ObjectMeta: v12.ObjectMeta{
			Name: "applications.cloudera.com",
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: cloudera_com.GroupName,
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:   "applications",
				Singular: "application",
				Kind:     "Application",
			},
			Scope:    apiextensions.NamespaceScoped,
			Versions: crdVersions,
		},
	}

	_, err := client.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil {
		fmt.println(err.Error())
	}

	// create CRD object
	crdclientset := versioned.NewForConfigOrDie(ctx.kubeClient.GetConfig())
	appCRD, err := crdclientset.ClouderaV1alpha1().Applications("default").Create(&v1alpha1.Application{
		TypeMeta: v12.TypeMeta{
			Kind:       "Application",
			APIVersion: "applications.example/v1alpha1",
		},
		ObjectMeta: v12.ObjectMeta{
			Name:      "app-00001",
			Namespace: "default",
			//ResourceVersion:            "v1alpha1",
		},
		Spec: v1alpha1.ApplicationSpec{
			ApplicationId: "app-00001",
			QueueName:     "root.sandbox",
		},
		Status: v1alpha1.ApplicationStatus{
			Status:  newApp.GetApplicationState(),
			Message: "application is initiated in the scheduler",
		},
	})
	if err != nil {
		panic(err.Error())
	}
}
