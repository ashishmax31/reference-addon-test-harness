package tests

import (
	"context"
	"fmt"
	"time"

	addonmeta "github.com/mt-sre/reference-addon-test-harness/pkg/metadata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	addonsv1alpha1 "github.com/openshift/addon-operator/apis/addons/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Reference addon test harness", func() {
	const (
		ExpectedDeploymentName  = "reference-addon"
		ExpectedAddonCrName     = "reference-addon"
		ReferenceAddonNamespace = "redhat-reference-addon"

		timeout  = time.Second * 5 * 60
		interval = time.Second * 1
	)
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	err = addonsv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).ShouldNot(HaveOccurred())
	apiClient, err := client.New(config, client.Options{Scheme: scheme.Scheme})
	Expect(err).ShouldNot(HaveOccurred())

	Context("After installation", func() {
		It("Should have an Addon CR in the available state", func() {
			addon := addonsv1alpha1.Addon{}
			Eventually(func() bool {
				err := apiClient.Get(context.TODO(), types.NamespacedName{Namespace: ReferenceAddonNamespace, Name: ExpectedAddonCrName}, &addon)
				if err != nil {
					fmt.Printf("Errored while Getting, msg: %s\n", err.Error())
					return false
				}
				if meta.IsStatusConditionTrue(addon.Status.Conditions,
					addonsv1alpha1.Available) {
					return true
				} else {
					return false
				}
			}, timeout, interval).Should(BeTrue())
		})
		It("Should have a reference addon deployment", func() {
			deploymentObj := appsv1.Deployment{}
			// Assert that the reference-addon deployment is created.
			Eventually(func() bool {
				err := apiClient.Get(context.TODO(), types.NamespacedName{Namespace: ReferenceAddonNamespace, Name: ExpectedDeploymentName}, &deploymentObj)
				if err != nil {
					fmt.Printf("Errored while Getting, msg: %s\n", err.Error())
					return false
				}
				addonmeta.Instance.AddonInstalled = true
				return true
			}, timeout, interval).Should(BeTrue())
		})
	})
})
