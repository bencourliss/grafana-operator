/*
Copyright 2019 Jagadish Nagarajaiah.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package grafana

import (
	"testing"
	"time"

	monitorsv1alpha1 "github.com/dichque/grafana-operator/pkg/apis/monitors/v1alpha1"
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedDeployRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "grafana-test-deployment", Namespace: "default"}}
var depKey = types.NamespacedName{Name: "grafana-test-deployment", Namespace: "default"}
var confKey = types.NamespacedName{Name: "grafana-test-config", Namespace: "default"}
var serKey = types.NamespacedName{Name: "grafana-test-service", Namespace: "default"}
var expectedConfigRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "grafana-test-config", Namespace: "default"}}
var expectedServiceRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "grafana-test-service", Namespace: "default"}}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &monitorsv1alpha1.Grafana{ObjectMeta: metav1.ObjectMeta{Name: "grafana-test-deployment", Namespace: "default"}}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Grafana object and expect the Reconcile and Deployment to be created
	err = c.Create(context.TODO(), instance)
	// The instance object may not be a valid object because it might be missing some required fields.
	// Please modify the instance object by adding required fields and then remove the following if statement.
	if apierrors.IsInvalid(err) {
		t.Logf("failed to create object, got an invalid object error: %v", err)
		return
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), instance)
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedDeployRequest)))

	//Deployment

	deploy := &appsv1.Deployment{}
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Delete the Deployment and expect Reconcile to be called for Deployment deletion
	g.Expect(c.Delete(context.TODO(), deploy)).NotTo(gomega.HaveOccurred())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedDeployRequest)))
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Manually delete Deployment since GC isn't enabled in the test control plane
	g.Eventually(func() error { return c.Delete(context.TODO(), deploy) }, timeout).
		Should(gomega.MatchError("deployments.apps \"grafana-test-deployment\" not found"))

	//ConfigMap
	//Create grafana configMap and expect the configMap to be created

	configmap_test := &corev1.ConfigMap{}
	g.Eventually(func() error { return c.Get(context.TODO(), confKey, configmap_test) }, timeout).Should(gomega.Succeed())

	g.Expect(c.Delete(context.TODO(), configmap_test)).NotTo(gomega.HaveOccurred())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedConfigRequest)))
	g.Eventually(func() error { return c.Get(context.TODO(), confKey, configmap_test) }, timeout).Should(gomega.Succeed())

	//Manually delete configmap since GC isn't enabled in the test controller plane
	g.Eventually(func() error { return c.Delete(context.TODO(), configmap_test) }, timeout).
		Should(gomega.MatchError("configmap.corev1 \"grafana-test-configmap\" not found"))

	//Service
	service_test := &corev1.Service{}

	g.Eventually(func() error { return c.Get(context.TODO(), serKey, service_test) }, timeout).
		Should(gomega.Succeed())

	// Delete the Deployment and expect Reconcile to be called for Deployment deletion
	g.Expect(c.Delete(context.TODO(), deploy)).NotTo(gomega.HaveOccurred())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedServiceRequest)))
	g.Eventually(func() error { return c.Get(context.TODO(), serKey, service_test) }, timeout).
		Should(gomega.Succeed())

	// Manually delete Deployment since GC isn't enabled in the test control plane
	g.Eventually(func() error { return c.Delete(context.TODO(), deploy) }, timeout).
		Should(gomega.MatchError("service.corev1 \"grafana-test-service\" not found"))

}
