package util

import (
	//"github.com/subravi92/grafana-operator/pkg/clients/fake"
	monitorsv1alpha1 "github.com/dichque/grafana-operator/pkg/apis/monitors/v1alpha1"
	"k8s.io/kubernetes/staging/src/k8s.io/cli-runtime/pkg/kustomize/k8sdeps/kv"

	//"k8s.io/client-go/util/integer"
	"testing"
	//"k8s.io/client-go/tools/cache"
	//v1 "k8s.io/client-go/informers/core/v1"
	//apiv1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGenerateDeployment(t *testing.T) {






	instance := &monitorsv1alpha1.Grafana{}
	instance.Name = "test-grafana"
	instance.Namespace = "test-namespace"

	instance.Spec.Image = "grafana/grafana:6.0.0-beta3clear"
	instance.Spec.GrafanaAdminUser = "admin"
	instance.Spec.GrafanaAdminPassword = ""
	instance.Spec.PrometheusUrl = "http://prometheus-operated:9090"

	var err *appsv1.Deployment


	//testcase1_op:=" &Deployment{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:test-grafana-deployment,GenerateName:,Namespace:test-namespace,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:DeploymentSpec{Replicas:nil,Selector:&k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{MatchLabels:map[string]string{deployment: test-grafana-deployment,},MatchExpressions:[],},Template:k8s_io_api_core_v1.PodTemplateSpec{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{app: grafana,deployment: test-grafana-deployment,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:PodSpec{Volumes:[{grafana-config {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-config,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-data {nil &EmptyDirVolumeSource{Medium:,SizeLimit:<nil>,} nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil}} {grafana-datasources {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-datasources,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {kafka-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:kafka-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}}],Containers:[{test-grafana-grafana-deployment  grafana/grafana:6.0.0-beta3clear [] []  [{ 0 3000  }] [] [] {map[] map[]} [{grafana-config false /etc/grafana  <nil>} {grafana-data false /var/lib/grafana  <nil>} {grafana-datasources false /etc/grafana/provisioning/datasources  <nil>} {grafana-dashboards false /etc/grafana/provisioning/dashboards  <nil>} {kafka-dashboards false /grafana-dashboard-definitions/0  <nil>}] [] nil nil nil    nil false false false}],RestartPolicy:,TerminationGracePeriodSeconds:nil,ActiveDeadlineSeconds:nil,DNSPolicy:,NodeSelector:map[string]string{},ServiceAccountName:,DeprecatedServiceAccount:,NodeName:,HostNetwork:false,HostPID:false,HostIPC:false,SecurityContext:nil,ImagePullSecrets:[{cisco-cred-pull-secret}],Hostname:,Subdomain:,Affinity:nil,SchedulerName:,InitContainers:[],AutomountServiceAccountToken:nil,Tolerations:[],HostAliases:[],PriorityClassName:,Priority:nil,DNSConfig:nil,ShareProcessNamespace:nil,ReadinessGates:[],RuntimeClassName:nil,EnableServiceLinks:nil,},},Strategy:DeploymentStrategy{Type:,RollingUpdate:nil,},MinReadySeconds:0,RevisionHistoryLimit:nil,Paused:false,ProgressDeadlineSeconds:nil,},Status:DeploymentStatus{ObservedGeneration:0,Replicas:0,UpdatedReplicas:0,AvailableReplicas:0,UnavailableReplicas:0,Conditions:[],ReadyReplicas:0,CollisionCount:nil,},},errmessage false"
	i := int32(1)
	tests := []struct {
		name    string
		grafana metav1.Object

		replicas *int32
		image    string
		output   string
	}{
		// TODO: Add test cases.
		{"testcase1", instance, &i, " grafana/grafana:6.0.0-beta3clear", "&Deployment{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:test-grafana-deployment,GenerateName:,Namespace:test-namespace,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:DeploymentSpec{Replicas:nil,Selector:&k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{MatchLabels:map[string]string{deployment: test-grafana-deployment,},MatchExpressions:[],},Template:k8s_io_api_core_v1.PodTemplateSpec{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{app: grafana,deployment: test-grafana-deployment,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:PodSpec{Volumes:[{grafana-config {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-config,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-data {nil &EmptyDirVolumeSource{Medium:,SizeLimit:<nil>,} nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil}} {grafana-datasources {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-datasources,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {kafka-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:kafka-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}}],Containers:[{test-grafana-grafana-deployment  grafana/grafana:6.0.0-beta3clear [] []  [{ 0 3000  }] [] [] {map[] map[]} [{grafana-config false /etc/grafana  <nil>} {grafana-data false /var/lib/grafana  <nil>} {grafana-datasources false /etc/grafana/provisioning/datasources  <nil>} {grafana-dashboards false /etc/grafana/provisioning/dashboards  <nil>} {kafka-dashboards false /grafana-dashboard-definitions/0  <nil>}] [] nil nil nil    nil false false false}],RestartPolicy:,TerminationGracePeriodSeconds:nil,ActiveDeadlineSeconds:nil,DNSPolicy:,NodeSelector:map[string]string{},ServiceAccountName:,DeprecatedServiceAccount:,NodeName:,HostNetwork:false,HostPID:false,HostIPC:false,SecurityContext:nil,ImagePullSecrets:[{cisco-cred-pull-secret}],Hostname:,Subdomain:,Affinity:nil,SchedulerName:,InitContainers:[],AutomountServiceAccountToken:nil,Tolerations:[],HostAliases:[],PriorityClassName:,Priority:nil,DNSConfig:nil,ShareProcessNamespace:nil,ReadinessGates:[],RuntimeClassName:nil,EnableServiceLinks:nil,},},Strategy:DeploymentStrategy{Type:,RollingUpdate:nil,},MinReadySeconds:0,RevisionHistoryLimit:nil,Paused:false,ProgressDeadlineSeconds:nil,},Status:DeploymentStatus{ObservedGeneration:0,Replicas:0,UpdatedReplicas:0,AvailableReplicas:0,UnavailableReplicas:0,Conditions:[],ReadyReplicas:0,CollisionCount:nil,},}"},
		{"testcase2", &monitorsv1alpha1.Grafana{}, &i, "", "&Deployment{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:-deployment,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:DeploymentSpec{Replicas:nil,Selector:&k8s_io_apimachinery_pkg_apis_meta_v1.LabelSelector{MatchLabels:map[string]string{deployment: -deployment,},MatchExpressions:[],},Template:k8s_io_api_core_v1.PodTemplateSpec{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{app: grafana,deployment: -deployment,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:PodSpec{Volumes:[{grafana-config {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-config,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-data {nil &EmptyDirVolumeSource{Medium:,SizeLimit:<nil>,} nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil}} {grafana-datasources {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-datasources,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {grafana-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:grafana-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}} {kafka-dashboards {nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil nil &ConfigMapVolumeSource{LocalObjectReference:LocalObjectReference{Name:kafka-dashboards,},Items:[],DefaultMode:nil,Optional:nil,} nil nil nil nil nil nil nil nil}}],Containers:[{-grafana-deployment grafana/grafana:6.0 [] []  [{ 0 3000  }] [] [] {map[] map[]} [{grafana-config false /etc/grafana  <nil>} {grafana-data false /var/lib/grafana  <nil>} {grafana-datasources false /etc/grafana/provisioning/datasources  <nil>} {grafana-dashboards false /etc/grafana/provisioning/dashboards  <nil>} {kafka-dashboards false /grafana-dashboard-definitions/0  <nil>}] [] nil nil nil    nil false false false}],RestartPolicy:,TerminationGracePeriodSeconds:nil,ActiveDeadlineSeconds:nil,DNSPolicy:,NodeSelector:map[string]string{},ServiceAccountName:,DeprecatedServiceAccount:,NodeName:,HostNetwork:false,HostPID:false,HostIPC:false,SecurityContext:nil,ImagePullSecrets:[{cisco-cred-pull-secret}],Hostname:,Subdomain:,Affinity:nil,SchedulerName:,InitContainers:[],AutomountServiceAccountToken:nil,Tolerations:[],HostAliases:[],PriorityClassName:,Priority:nil,DNSConfig:nil,ShareProcessNamespace:nil,ReadinessGates:[],RuntimeClassName:nil,EnableServiceLinks:nil,},},Strategy:DeploymentStrategy{Type:,RollingUpdate:nil,},MinReadySeconds:0,RevisionHistoryLimit:nil,Paused:false,ProgressDeadlineSeconds:nil,},Status:DeploymentStatus{ObservedGeneration:0,Replicas:0,UpdatedReplicas:0,AvailableReplicas:0,UnavailableReplicas:0,Conditions:[],ReadyReplicas:0,CollisionCount:nil,},}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := GenerateDeployment(tt.grafana, tt.replicas, tt.image);  err.String() != tt.output {

				t.Errorf("GenerateDeployment() error = %v,errmessage %v", err.String(), tt.output)
			}
			print(err)
		})
	}

}

func TestGenerateService(t *testing.T) {

	tests := []struct {
		name    string
		grafana metav1.Object
		output string
	}{
		//TODO : Add test cases
		{name: "testcase1", grafana: &monitorsv1alpha1.Grafana{}, output: "&Service{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:grafana-operated,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{grafana-service: ,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:ServiceSpec{Ports:[{http  3000 {0 0 } 0}],Selector:map[string]string{app: grafana,},ClusterIP:,Type:,ExternalIPs:[],SessionAffinity:,LoadBalancerIP:,LoadBalancerSourceRanges:[],ExternalName:,ExternalTrafficPolicy:,HealthCheckNodePort:0,PublishNotReadyAddresses:false,SessionAffinityConfig:nil,},Status:ServiceStatus{LoadBalancer:LoadBalancerStatus{Ingress:[],},},}"},
		{name: "testcase2", grafana: &monitorsv1alpha1.Grafana{}, output: "&Service{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:grafana-operated,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{grafana-service: ,},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Spec:ServiceSpec{Ports:[{http  3000 {0 0 } 0}],Selector:map[string]string{app: grafana,},ClusterIP:,Type:,ExternalIPs:[],SessionAffinity:,LoadBalancerIP:,LoadBalancerSourceRanges:[],ExternalName:,ExternalTrafficPolicy:,HealthCheckNodePort:0,PublishNotReadyAddresses:false,SessionAffinityConfig:nil,},Status:ServiceStatus{LoadBalancer:LoadBalancerStatus{Ingress:[],},},}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateService(tt.grafana); err.String()!= tt.output {
				t.Errorf("GenerateService() error = %v,errmessage %v", err.String(), tt.output)
				//print(err.String())
			}

		})
	}

}

/*func TestGenerateConfigMapFromTemplate(t *testing.T) {

	type GrafanaConfigtest struct {
		AdminUser     string
		AdminPassword string
		PrometheusUrl string
	}

	tests := []struct {
		name          string
		grafana       metav1.Object
		description   string
		tmplpath      string
		grafanaconfig *GrafanaConfig
		errmessage    bool
	}{
		{name: "grafana-datasources", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase1", tmplpath: "/config/templates/datasources.yaml.tmpl", grafanaconfig: &GrafanaConfig{"admin", "changeit", "http://prometheus-operated:9090"}, errmessage: false},
		{name: "grafana-datasources", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase2", tmplpath: "", grafanaconfig: &GrafanaConfig{"", "", ""}, errmessage: false},
		{name: "grafana-datasources", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase3", tmplpath: "/config/templates/datasources.yaml.tmpl", grafanaconfig: &GrafanaConfig{"", "changeit", "http://prometheus-operated:9090"}, errmessage: false},
		{name: "grafana-datasources", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase4", tmplpath: "/config/templates/datasources.yaml.tmpl", grafanaconfig: &GrafanaConfig{"admin", "", "http://prometheus-operated:9090"}, errmessage: false},
		{name: "grafana-datasources", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase5", tmplpath: "/config/templates/datasources.yaml.tmpl", grafanaconfig: &GrafanaConfig{"admin", "changeit", ""}, errmessage: false},
		{name: "grafana-config", grafana: &monitorsv1alpha1.Grafana{}, description: "testcase6", tmplpath: "/config/templates/grafana.ini.tmpl", grafanaconfig: &GrafanaConfig{"admin", "changeit", "http://prometheus-operated:9090"}, errmessage: false},
	}

	print(tests)
	/*for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if err := GenerateConfigMapFromTemplate(tt.grafana, tt.name, tt.tmplpath,tt.grafanaconfig); (err == nil)  {
				t.Errorf("GenerateConfigMapFromTemplate() error = %v,errmessage %v", err, tt.errmessage)
			}
		})
	}

}*/
func TestKeyValuesFromFileSources(t *testing.T) {

	tests := []struct {
		description string
		sources     []string
		output      string
	}{

		//Add Test cases
		{description: "testcase1", sources: []string{"/config/dashboards/strimzi-zookeeper.json", "/config/dashboards/dichque-grafana.json", "/config/dashboards/strimzi-kafka.json"}, output: ""},
		{description: "testcase2", sources: []string{"/config/dashboards/strimzi-zookeeper.json", "/config/dashboards/dichque-grafana.json", "/config/dashboards/strimzi-kafka.json"}, output: ""},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

			if _, err := keyValuesFromFileSources(tt.sources); err == nil {
				t.Errorf("keyValuesFromFileSources() error = %v,errmessage %v", err, tt.output)
			}
		})
	}

}
func TestParseFileSource(t *testing.T) {

	tests := []struct {
		description string
		sources     string
		output      string
	}{
		// TODO: Add test cases

		{description: "testcase1", sources: "/config/dashboards/strimzi-zookeeper.json", output: "strimzi-zookeeper.json"},
		{description: "testcase2", sources: "/config/dashboards/dichque-grafana.json", output: "dichque-grafana.json"},
		{description: "testcase3", sources: "/config/dashboards/strimzi-kafka.json", output: "strimzi-kafka.json"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if key, file, err := parseFileSource(tt.sources); (err != nil) && key == tt.output {
				t.Errorf("parseFileSource() error = %v,errmessage %v%v%v", err, tt.output, key, file)
			}
		})
	}

}

func TestGenerateConfigMap(t *testing.T) {

	tests := []struct {
		name        string
		description string
		grafana     metav1.Object
		filePath    []string
		output string
	}{
		//TODO : Add test cases
		{name: "grafana", description: "testcase1", grafana: &monitorsv1alpha1.Grafana{}, filePath: []string{"/config/dashboards.yaml"}, output: "&ConfigMap{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:grafana,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Data:map[string]string{},BinaryData:map[string][]byte{},}"},
		{name: "grafana", description: "testcase2", grafana: &monitorsv1alpha1.Grafana{}, filePath: []string{"/Config/datasources.yaml"}, output: "&ConfigMap{ObjectMeta:k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta{Name:grafana,GenerateName:,Namespace:,SelfLink:,UID:,ResourceVersion:,Generation:0,CreationTimestamp:0001-01-01 00:00:00 +0000 UTC,DeletionTimestamp:<nil>,DeletionGracePeriodSeconds:nil,Labels:map[string]string{},Annotations:map[string]string{},OwnerReferences:[],Finalizers:[],ClusterName:,Initializers:nil,},Data:map[string]string{},BinaryData:map[string][]byte{},}"},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if err := GenerateConfigMap(tt.grafana, tt.name, tt.filePath); err.String()!=tt.output {
				t.Errorf("GenerateConfigMap() error = %v,errmessage %v", err, tt.output)
			}
		})
	}

}

func TestAddKvToConfigMap(t *testing.T) {

	tests := []struct {
		description string
		configMap   *corev1.ConfigMap
		keyvalue  kv.Pair
		errmessage  bool
	}{
		//TODO: Add test cases
		{description: "testcase1", configMap: &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace"}}, keyvalue: kv.Pair{"key1","{\r\n  \"aliasColors\": {\r\n    \"localhost:7071\": \"#890F02\"\r\n  },\r\n  \"bars\": false,\r\n  \"dashLength\": 10,\r\n  \"dashes\": false,\r\n  \"datasource\": \"prometheus\",\r\n  \"editable\": true,\r\n  \"error\": false,\r\n  \"fill\": 1,\r\n  \"grid\": {},\r\n  \"gridPos\": {\r\n    \"h\": 6,\r\n    \"w\": 8,\r\n    \"x\": 16,\r\n    \"y\": 6\r\n  },\r\n  \"id\": 3,\r\n  \"isNew\": true,\r\n  \"legend\": {\r\n    \"avg\": false,\r\n    \"current\": false,\r\n    \"max\": false,\r\n    \"min\": false,\r\n    \"show\": true,\r\n    \"total\": false,\r\n    \"values\": false\r\n  },\r\n  \"lines\": true,\r\n  \"linewidth\": 2,\r\n  \"links\": [],\r\n  \"nullPointMode\": \"connected\",\r\n  \"percentage\": false,\r\n  \"pointradius\": 5,\r\n  \"points\": false,\r\n  \"renderer\": \"flot\",\r\n  \"seriesOverrides\": [],\r\n  \"spaceLength\": 10,\r\n  \"stack\": false,\r\n  \"steppedLine\": false,\r\n  \"targets\": [\r\n    {\r\n      \"expr\": \"sum without(gc)(rate(jvm_gc_collection_seconds_sum{job=\\\"my-cluster-kafka-bootstrap\\\"}[5m]))\",\r\n      \"format\": \"time_series\",\r\n      \"intervalFactor\": 10,\r\n      \"legendFormat\": \"{{pod}}\",\r\n      \"metric\": \"jvm_gc_collection_seconds_sum\",\r\n      \"refId\": \"A\",\r\n      \"step\": 4\r\n    }\r\n  ],\r\n  \"thresholds\": [],\r\n  \"timeFrom\": null,\r\n  \"timeRegions\": [],\r\n  \"timeShift\": null,\r\n  \"title\": \"Time spent in GC\",\r\n  \"tooltip\": {\r\n    \"msResolution\": false,\r\n    \"shared\": true,\r\n    \"sort\": 0,\r\n    \"value_type\": \"cumulative\"\r\n  },\r\n  \"type\": \"graph\",\r\n  \"xaxis\": {\r\n    \"buckets\": null,\r\n    \"mode\": \"time\",\r\n    \"name\": null,\r\n    \"show\": true,\r\n    \"values\": []\r\n  }\r\n}"},  errmessage: false},
		{description: "testcase2", configMap: &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace"}}, keyvalue: kv.Pair{Key:"",Value:""}, errmessage: false},
		{description: "testcase3", configMap: &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace"}}, keyvalue: kv.Pair{Key:"",Value:"without key"},  errmessage: false},
		{description: "testcase4", configMap: &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace"}}, keyvalue: kv.Pair{Key:"without value",Value:""},  errmessage: false},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if err := addKvToConfigMap(tt.configMap, tt.keyvalue.Key, tt.keyvalue.Value); err != nil {
				t.Errorf("addKvToConfigMap() error = %v,errmessage %v", err, tt.errmessage)
			}
		})
	}

}

