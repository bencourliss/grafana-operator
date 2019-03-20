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
	"context"
	"github.com/dichque/grafana-operator/pkg/util"
	"reflect"

	monitorsv1alpha1 "github.com/dichque/grafana-operator/pkg/apis/monitors/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Grafana Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGrafana{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("grafana-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Grafana
	err = c.Watch(&source.Kind{Type: &monitorsv1alpha1.Grafana{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by Grafana - change this for objects you create
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &monitorsv1alpha1.Grafana{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &monitorsv1alpha1.Grafana{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &monitorsv1alpha1.Grafana{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileGrafana{}

// ReconcileGrafana reconciles a Grafana object
type ReconcileGrafana struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Grafana object and makes changes based on the state read
// and what is in the Grafana.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitors.aims.cisco.com,resources=grafanas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitors.aims.cisco.com,resources=grafanas/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitors.aims.cisco.com,resources=grafanas/finalizers,verbs=get;update;patch;delete;list;watch
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;watch;list
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;watch;list;create;update;delete
func (r *ReconcileGrafana) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the Grafana instance
	instance := &monitorsv1alpha1.Grafana{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define desired state of configMaps
	//grafanaConfig := util.GenerateConfigMap(instance, "grafana-config", []string{"config/grafana.ini"})

	cfg := new(util.GrafanaConfig)
	if instance.Spec.GrafanaAdminUser != "" {
		cfg.AdminUser = instance.Spec.GrafanaAdminUser
	}
	if instance.Spec.GrafanaAdminPassword != "" {
		cfg.AdminPassword = instance.Spec.GrafanaAdminPassword
	}

	grafanaConfig := util.GenerateConfigMapFromTemplate(instance, "grafana-config", "/config/grafana.ini.tmpl", cfg)
	if err := controllerutil.SetControllerReference(instance, grafanaConfig, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundGrafanaConfig := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: grafanaConfig.Name, Namespace: grafanaConfig.Namespace}, foundGrafanaConfig)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating Grafana config", "namespace", grafanaConfig.Namespace, "name", grafanaConfig.Name)
		err = r.Create(context.TODO(), grafanaConfig)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if !reflect.DeepEqual(grafanaConfig, foundGrafanaConfig) {
		foundGrafanaConfig = grafanaConfig
		log.Info("Updating grafana config", "namespace", grafanaConfig.Namespace, "name", grafanaConfig.Name)
		err = r.Update(context.TODO(), foundGrafanaConfig)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Grafana Data Source
	//grafanaDataSrc := util.GenerateConfigMap(instance, "grafana-datasources", []string{"/config/datasources.yaml"})
	dataSrc := new(util.GrafanaConfig)
	if instance.Spec.PrometheusUrl != "" {
		dataSrc.PrometheusUrl = instance.Spec.PrometheusUrl
	} else {
		dataSrc.PrometheusUrl = "http://prometheus-operated:9090"
	}

	grafanaDataSrc := util.GenerateConfigMapFromTemplate(instance, "grafana-datasources", "/config/datasources.yaml.tmpl", dataSrc)
	if err := controllerutil.SetControllerReference(instance, grafanaDataSrc, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundGrafanaDataSrc := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: grafanaDataSrc.Name, Namespace: grafanaDataSrc.Namespace}, foundGrafanaDataSrc)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating grafana data source", "namespace", grafanaDataSrc.Namespace, "name", grafanaDataSrc.Name)
		err = r.Create(context.TODO(), grafanaDataSrc)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if !reflect.DeepEqual(grafanaDataSrc, foundGrafanaDataSrc) {
		foundGrafanaDataSrc = grafanaDataSrc
		log.Info("Updating grafana data source", "namespace", grafanaDataSrc.Namespace, "name", grafanaDataSrc.Name)
		err = r.Update(context.TODO(), foundGrafanaDataSrc)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Grafana Dashboard Source
	grafanaDash := util.GenerateConfigMap(instance, "grafana-dashboards", []string{"/config/dashboards.yaml"})
	if err := controllerutil.SetControllerReference(instance, grafanaDash, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundGrafanaDash := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: grafanaDash.Name, Namespace: grafanaDash.Namespace}, foundGrafanaDash)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating grafana dashboard source", "namespace", grafanaDash.Namespace, "name", grafanaDash.Name)
		err = r.Create(context.TODO(), grafanaDash)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if !reflect.DeepEqual(grafanaDash, foundGrafanaDash) {
		foundGrafanaDash = grafanaDash
		log.Info("Updating grafana dashboard source", "namespace", grafanaDash.Namespace, "name", grafanaDash.Name)
		err = r.Update(context.TODO(), foundGrafanaDash)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Grafana Kafka Dashboards
	kafkaDashFilePath := []string{"/config/dashboards/strimzi-zookeeper.json", "/config/dashboards/dichque-burrow.json", "/config/dashboards/strimzi-kafka.json"}
	kafkaDash := util.GenerateConfigMap(instance, "kafka-dashboards", kafkaDashFilePath)
	if err := controllerutil.SetControllerReference(instance, kafkaDash, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundKafkaDash := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: kafkaDash.Name, Namespace: grafanaDash.Namespace}, foundKafkaDash)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating kafka dashboard source", "namespace", kafkaDash.Namespace, "name", kafkaDash.Name)
		err = r.Create(context.TODO(), kafkaDash)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if !reflect.DeepEqual(kafkaDash, foundKafkaDash) {
		foundKafkaDash = kafkaDash
		log.Info("Updating kafka dashboard source", "namespace", kafkaDash.Namespace, "name", kafkaDash.Name)
		err = r.Update(context.TODO(), foundKafkaDash)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Define desired state of the deployment
	deploy := util.GenerateDeployment(instance, instance.Spec.Replicas, instance.Spec.Image)
	if err := controllerutil.SetControllerReference(instance, deploy, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Deployment already exists
	found := &appsv1.Deployment{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
		err = r.Create(context.TODO(), deploy)
		return reconcile.Result{}, err
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Updating status field of Grafana
	if found.Status.AvailableReplicas != instance.Status.ActiveGrafanaCount {
		instance.Status.ActiveGrafanaCount = found.Status.AvailableReplicas
		log.Info("Updating Status", "namespace", instance.Namespace, "name", instance.Name)
		err = r.Update(context.TODO(), instance)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Update the found object and write the result back if there are any changes
	if !reflect.DeepEqual(deploy.Spec, found.Spec) {
		found.Spec = deploy.Spec
		log.Info("Updating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
		err = r.Update(context.TODO(), found)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	// Define desired state of the service
	service := util.GenerateService(instance)
	if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundService := &corev1.Service{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating service ", "namespace", service.Namespace, "name", service.Name)
		err = r.Create(context.TODO(), service)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Update the found object and write the result back if there are any changes
	if !reflect.DeepEqual(service.Spec, foundService.Spec) {
		foundService.Spec = service.Spec
		log.Info("Updating Service", "namespace", service.Namespace, "name", service.Name)
		err = r.Update(context.TODO(), found)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}
