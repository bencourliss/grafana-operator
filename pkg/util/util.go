/*
Copyright 2018 The Kubernetes Authors.

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

package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/kubernetes/staging/src/k8s.io/cli-runtime/pkg/kustomize/k8sdeps/kv"
)

type GrafanaConfig struct {
	AdminUser     string
	AdminPassword string
	PrometheusUrl string
}

func GenerateDeployment(grafana metav1.Object, replicas *int32, image string) *appsv1.Deployment {

	if replicas == nil {
		r := int32(1)
		replicas = &r
	}

	if image == "" {
		image = "grafana/grafana:6.0"
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      grafana.GetName() + "-deployment",
			Namespace: grafana.GetNamespace(),
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"deployment": grafana.GetName() + "-deployment"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"deployment": grafana.GetName() + "-deployment", "app": "grafana"}},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  grafana.GetName() + "-grafana-deployment",
							Image: image,
							Ports: []corev1.ContainerPort{{ContainerPort: 3000}},
							VolumeMounts: []corev1.VolumeMount{
								{Name: "grafana-config", MountPath: "/etc/grafana"},
								{Name: "grafana-data", MountPath: "/var/lib/grafana"},
								{Name: "grafana-datasources", MountPath: "/etc/grafana/provisioning/datasources"},
								{Name: "grafana-dashboards", MountPath: "/etc/grafana/provisioning/dashboards"},
								{Name: "kafka-dashboards", MountPath: "/grafana-dashboard-definitions/0"},
							},
						},
					},
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: "cisco-cred-pull-secret",
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "grafana-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "grafana-config",
									},
								},
							},
						},
						{
							Name: "grafana-data",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "grafana-datasources",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "grafana-datasources",
									},
								},
							},
						},
						{
							Name: "grafana-dashboards",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "grafana-dashboards",
									},
								},
							},
						},
						{
							Name: "kafka-dashboards",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "kafka-dashboards",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return deploy
}

// GenerateService returns a new corev1.Service pointer generated for the MongoDB instance
// grafana: Grafana instance
func GenerateService(grafana metav1.Object) *corev1.Service {
	// TODO: Default and Validate these with Webhooks
	copyLabels := grafana.GetLabels()
	if copyLabels == nil {
		copyLabels = map[string]string{}
	}
	labels := map[string]string{}
	for k, v := range copyLabels {
		labels[k] = v
	}
	labels["grafana-service"] = grafana.GetName()

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      grafana.GetName() + "-grafana-service",
			Namespace: grafana.GetNamespace(),
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{Name: "http", Port: 3000},
			},
			Selector: map[string]string{"app": "grafana"},
		},
	}
	return service
}

func GenerateConfigMap(grafana metav1.Object, name string, filePath []string) *corev1.ConfigMap {

	configMap := &corev1.ConfigMap{}
	configMap.Name = name
	configMap.Namespace = grafana.GetNamespace()
	configMap.Data = map[string]string{}

	var all []kv.Pair
	pairs, err := keyValuesFromFileSources(filePath)
	all = append(all, pairs...)

	for _, kvPair := range all {
		err = addKvToConfigMap(configMap, kvPair.Key, kvPair.Value)
		if err != nil {
			return nil
		}
	}
	return configMap
}

func GenerateConfigMapFromTemplate(grafana metav1.Object, name string, tmplPath string, cfg *GrafanaConfig) *corev1.ConfigMap {

	// Tread cautiously fragile code ahead !!

	configMap := &corev1.ConfigMap{}
	configMap.Name = name
	configMap.Namespace = grafana.GetNamespace()
	configMap.Data = map[string]string{}
	var cm []byte
	var all []kv.Pair

	tmpl, err := template.ParseFiles(tmplPath)
	wr, err := ioutil.TempFile("/tmp", ".grafana")

	err = tmpl.Execute(wr, cfg)

	wr.Close()
	defer os.Remove(wr.Name())

	cm, err = ioutil.ReadFile(wr.Name())
	all = append(all, kv.Pair{Key: strings.ReplaceAll(path.Base(tmplPath), ".tmpl", ""), Value: string(cm)})

	for _, kvPair := range all {
		err = addKvToConfigMap(configMap, kvPair.Key, kvPair.Value)
		if err != nil {
			return nil
		}
	}

	return configMap
}

func keyValuesFromFileSources(sources []string) ([]kv.Pair, error) {
	var kvs []kv.Pair
	for _, s := range sources {
		k, fPath, err := parseFileSource(s)
		if err != nil {
			return nil, err
		}
		content, err := ioutil.ReadFile(fPath)
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, kv.Pair{Key: k, Value: string(content)})
	}
	return kvs, nil
}

// parseFileSource parses the source given.
//
//  Acceptable formats include:
//   1.  source-path: the basename will become the key name
//   2.  source-name=source-path: the source-name will become the key name and
//       source-path is the path to the key file.
//
// Key names cannot include '='.
func parseFileSource(source string) (keyName, filePath string, err error) {
	numSeparators := strings.Count(source, "=")
	switch {
	case numSeparators == 0:
		return path.Base(source), source, nil
	case numSeparators == 1 && strings.HasPrefix(source, "="):
		return "", "", fmt.Errorf("key name for file path %v missing", strings.TrimPrefix(source, "="))
	case numSeparators == 1 && strings.HasSuffix(source, "="):
		return "", "", fmt.Errorf("file path for key name %v missing", strings.TrimSuffix(source, "="))
	case numSeparators > 1:
		return "", "", errors.New("key names or file paths cannot contain '='")
	default:
		components := strings.Split(source, "=")
		return components[0], components[1], nil
	}
}

// addKvToConfigMap adds the given key and data to the given config map.
// Error if key invalid, or already exists.
func addKvToConfigMap(configMap *corev1.ConfigMap, keyName, data string) error {
	// Note, the rules for ConfigMap keys are the exact same as the ones for SecretKeys.
	if errs := validation.IsConfigMapKey(keyName); len(errs) != 0 {
		return fmt.Errorf("%q is not a valid key name for a ConfigMap: %s", keyName, strings.Join(errs, ";"))
	}

	keyExistsErrorMsg := "cannot add key %s, another key by that name already exists: %v"

	// If the configmap data contains byte sequences that are all in the UTF-8
	// range, we will write it to .Data
	if utf8.Valid([]byte(data)) {
		if _, entryExists := configMap.Data[keyName]; entryExists {
			return fmt.Errorf(keyExistsErrorMsg, keyName, configMap.Data)
		}
		configMap.Data[keyName] = data
		return nil
	}

	// otherwise, it's BinaryData
	if configMap.BinaryData == nil {
		configMap.BinaryData = map[string][]byte{}
	}
	if _, entryExists := configMap.BinaryData[keyName]; entryExists {
		return fmt.Errorf(keyExistsErrorMsg, keyName, configMap.BinaryData)
	}
	configMap.BinaryData[keyName] = []byte(data)
	return nil
}
