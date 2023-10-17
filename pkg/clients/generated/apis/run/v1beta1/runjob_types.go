// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobBinaryAuthorization struct {
	/* If present, indicates to use Breakglass using this justification. If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass. */
	// +optional
	BreakglassJustification *string `json:"breakglassJustification,omitempty"`

	/* If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled. */
	// +optional
	UseDefault *bool `json:"useDefault,omitempty"`
}

type JobContainers struct {
	/* Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell. */
	// +optional
	Args []string `json:"args,omitempty"`

	/* Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell. */
	// +optional
	Command []string `json:"command,omitempty"`

	/* List of environment variables to set in the container. */
	// +optional
	Env []JobEnv `json:"env,omitempty"`

	/* URL of the Container image in Google Container Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images. */
	Image string `json:"image"`

	/* DEPRECATED. `liveness_probe` is deprecated and will be removed in a future major release. This field is not supported by the Cloud Run API. Periodic probe of container liveness. Container will be restarted if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	This field is not supported in Cloud Run Job currently. */
	// +optional
	LivenessProbe *JobLivenessProbe `json:"livenessProbe,omitempty"`

	/* Name of the container specified as a DNS_LABEL. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.

	If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on. */
	// +optional
	Ports []JobPorts `json:"ports,omitempty"`

	/* Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources. */
	// +optional
	Resources *JobResources `json:"resources,omitempty"`

	/* DEPRECATED. `startup_probe` is deprecated and will be removed in a future major release. This field is not supported by the Cloud Run API. Startup probe of application within the container. All other probes are disabled if a startup probe is provided, until it succeeds. Container will not be added to service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	This field is not supported in Cloud Run Job currently. */
	// +optional
	StartupProbe *JobStartupProbe `json:"startupProbe,omitempty"`

	/* Volume to mount into the container's filesystem. */
	// +optional
	VolumeMounts []JobVolumeMounts `json:"volumeMounts,omitempty"`

	/* Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. */
	// +optional
	WorkingDir *string `json:"workingDir,omitempty"`
}

type JobEmptyDir struct {
	/* The different types of medium supported for EmptyDir. Default value: "MEMORY" Possible values: ["MEMORY"]. */
	// +optional
	Medium *string `json:"medium,omitempty"`

	/* Limit on the storage usable by this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir. */
	// +optional
	SizeLimit *string `json:"sizeLimit,omitempty"`
}

type JobEnv struct {
	/* Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters. */
	Name string `json:"name"`

	/* Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "", and the maximum length is 32768 bytes. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the environment variable's value. */
	// +optional
	ValueSource *JobValueSource `json:"valueSource,omitempty"`
}

type JobHttpGet struct {
	/* Custom headers to set in the request. HTTP allows repeated headers. */
	// +optional
	HttpHeaders []JobHttpHeaders `json:"httpHeaders,omitempty"`

	/* Path to access on the HTTP server. Defaults to '/'. */
	// +optional
	Path *string `json:"path,omitempty"`
}

type JobHttpHeaders struct {
	/* The header field name. */
	Name string `json:"name"`

	/* The header field value. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type JobItems struct {
	/* Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal). If 0 or not set, the Volume's default mode will be used. */
	// +optional
	Mode *int `json:"mode,omitempty"`

	/* The relative path of the secret in the container. */
	Path string `json:"path"`

	/* The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version */
	VersionRef v1alpha1.ResourceRef `json:"versionRef"`
}

type JobLivenessProbe struct {
	/* Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. */
	// +optional
	FailureThreshold *int `json:"failureThreshold,omitempty"`

	/* HTTPGet specifies the http request to perform. Exactly one of HTTPGet or TCPSocket must be specified. */
	// +optional
	HttpGet *JobHttpGet `json:"httpGet,omitempty"`

	/* Number of seconds after the container has started before the probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes. */
	// +optional
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty"`

	/* How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. Must be greater or equal than timeoutSeconds. */
	// +optional
	PeriodSeconds *int `json:"periodSeconds,omitempty"`

	/* TCPSocket specifies an action involving a TCP port. Exactly one of HTTPGet or TCPSocket must be specified. */
	// +optional
	TcpSocket *JobTcpSocket `json:"tcpSocket,omitempty"`

	/* Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. Maximum value is 3600. Must be smaller than periodSeconds. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes. */
	// +optional
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

type JobNetworkInterfaces struct {
	/* The VPC network that the Cloud Run resource will be able to send traffic to. At least one of network or subnetwork must be specified. If both
	network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be
	looked up from the subnetwork. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* The VPC subnetwork that the Cloud Run resource will get IPs from. At least one of network or subnetwork must be specified. If both
	network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the
	subnetwork with the same name with the network will be used. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`

	/* Network tags applied to this Cloud Run job. */
	// +optional
	Tags []string `json:"tags,omitempty"`
}

type JobPorts struct {
	/* Port number the container listens on. This must be a valid TCP port number, 0 < containerPort < 65536. */
	// +optional
	ContainerPort *int `json:"containerPort,omitempty"`

	/* If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c". */
	// +optional
	Name *string `json:"name,omitempty"`
}

type JobResources struct {
	/* Only memory and CPU are supported. Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go. */
	// +optional
	Limits map[string]string `json:"limits,omitempty"`
}

type JobSecret struct {
	/* Integer representation of mode bits to use on created files by default. Must be a value between 0000 and 0777 (octal), defaulting to 0444. Directories within the path are not affected by this setting. */
	// +optional
	DefaultMode *int `json:"defaultMode,omitempty"`

	/* If unspecified, the volume will expose a file whose name is the secret, relative to VolumeMount.mount_path. If specified, the key will be used as the version to fetch from Cloud Secret Manager and the path will be the name of the file exposed in the volume. When items are defined, they must specify a path and a version. */
	// +optional
	Items []JobItems `json:"items,omitempty"`

	/* The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project. */
	SecretRef v1alpha1.ResourceRef `json:"secretRef"`
}

type JobStartupProbe struct {
	/* Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1. */
	// +optional
	FailureThreshold *int `json:"failureThreshold,omitempty"`

	/* HTTPGet specifies the http request to perform. Exactly one of HTTPGet or TCPSocket must be specified. */
	// +optional
	HttpGet *JobHttpGet `json:"httpGet,omitempty"`

	/* Number of seconds after the container has started before the probe is initiated. Defaults to 0 seconds. Minimum value is 0. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes. */
	// +optional
	InitialDelaySeconds *int `json:"initialDelaySeconds,omitempty"`

	/* How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1. Maximum value for liveness probe is 3600. Maximum value for startup probe is 240. Must be greater or equal than timeoutSeconds. */
	// +optional
	PeriodSeconds *int `json:"periodSeconds,omitempty"`

	/* TCPSocket specifies an action involving a TCP port. Exactly one of HTTPGet or TCPSocket must be specified. */
	// +optional
	TcpSocket *JobTcpSocket `json:"tcpSocket,omitempty"`

	/* Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. Maximum value is 3600. Must be smaller than periodSeconds. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes. */
	// +optional
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

type JobTcpSocket struct {
	/* Port number to access on the container. Must be in the range 1 to 65535. If not specified, defaults to 8080. */
	// +optional
	Port *int `json:"port,omitempty"`
}

type JobTemplate struct {
	/* Holds the single container that defines the unit of execution for this task. */
	// +optional
	Containers []JobContainers `json:"containers,omitempty"`

	/* A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek */
	// +optional
	EncryptionKeyRef *v1alpha1.ResourceRef `json:"encryptionKeyRef,omitempty"`

	/* The execution environment being used to host this Task. Possible values: ["EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"]. */
	// +optional
	ExecutionEnvironment *string `json:"executionEnvironment,omitempty"`

	/* Number of retries allowed per Task, before marking this Task failed. */
	// +optional
	MaxRetries *int `json:"maxRetries,omitempty"`

	/* Email address of the IAM service account associated with the revision of the service. The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.

	A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". */
	// +optional
	Timeout *string `json:"timeout,omitempty"`

	/* A list of Volumes to make available to containers. */
	// +optional
	Volumes []JobVolumes `json:"volumes,omitempty"`

	/* VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc. */
	// +optional
	VpcAccess *JobVpcAccess `json:"vpcAccess,omitempty"`
}

type JobValueSource struct {
	/* Selects a secret and a specific version from Cloud Secret Manager. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type JobVolumeMounts struct {
	/* Path within the container at which the volume should be mounted. Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must otherwise be /cloudsql. All instances defined in the Volume will be available as /cloudsql/[instance]. For more information on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run. */
	MountPath string `json:"mountPath"`

	/* This must match the Name of a Volume. */
	Name string `json:"name"`
}

type JobVolumes struct {
	/* Ephemeral storage used as a shared volume. */
	// +optional
	EmptyDir *JobEmptyDir `json:"emptyDir,omitempty"`

	/* Volume's name. */
	Name string `json:"name"`

	/* Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret. */
	// +optional
	Secret *JobSecret `json:"secret,omitempty"`
}

type JobVpcAccess struct {
	/* Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"]. */
	// +optional
	Egress *string `json:"egress,omitempty"`

	/* Direct VPC egress settings. Currently only single network interface is supported. */
	// +optional
	NetworkInterfaces []JobNetworkInterfaces `json:"networkInterfaces,omitempty"`
}

type RunJobSpec struct {
	/* Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.

	Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected on new resources.
	All system annotations in v1 now have a corresponding field in v2 Job.

	This field follows Kubernetes annotations' namespacing, limits, and rules. */
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	/* Settings for the Binary Authorization feature. */
	// +optional
	BinaryAuthorization *JobBinaryAuthorization `json:"binaryAuthorization,omitempty"`

	/* Arbitrary identifier for the API client. */
	// +optional
	Client *string `json:"client,omitempty"`

	/* Arbitrary version identifier for the API client. */
	// +optional
	ClientVersion *string `json:"clientVersion,omitempty"`

	/* The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.
	If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.

	For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]. */
	// +optional
	LaunchStage *string `json:"launchStage,omitempty"`

	/* Immutable. The location of the cloud run job. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The template used to create executions for this Job. */
	Template JobTemplate `json:"template"`
}

type JobLatestCreatedExecutionStatus struct {
	/* Completion timestamp of the execution.

	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	CompletionTime *string `json:"completionTime,omitempty"`

	/* Creation timestamp of the execution.

	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Name of the execution. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type JobTerminalConditionStatus struct {
	/* A reason for the execution condition. */
	// +optional
	ExecutionReason *string `json:"executionReason,omitempty"`

	/* Last time the condition transitioned from one status to another.

	A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z". */
	// +optional
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`

	/* Human readable message indicating details about the current status. */
	// +optional
	Message *string `json:"message,omitempty"`

	/* A common (service-level) reason for this condition. */
	// +optional
	Reason *string `json:"reason,omitempty"`

	/* A reason for the revision condition. */
	// +optional
	RevisionReason *string `json:"revisionReason,omitempty"`

	/* How to interpret failures of this condition, one of Error, Warning, Info. */
	// +optional
	Severity *string `json:"severity,omitempty"`

	/* State of the condition. */
	// +optional
	State *string `json:"state,omitempty"`

	/* type is used to communicate the status of the reconciliation process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting Types common to all resources include: * "Ready": True when the Resource is ready. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type RunJobStatus struct {
	/* Conditions represent the latest available observations of the
	   RunJob's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The creation time. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Email address of the authenticated creator. */
	// +optional
	Creator *string `json:"creator,omitempty"`

	/* The deletion time. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* Number of executions created for this job. */
	// +optional
	ExecutionCount *int `json:"executionCount,omitempty"`

	/* For a deleted resource, the time after which it will be permamently deleted. */
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`

	/* Email address of the last authenticated modifier. */
	// +optional
	LastModifier *string `json:"lastModifier,omitempty"`

	/* Name of the last created execution. */
	// +optional
	LatestCreatedExecution []JobLatestCreatedExecutionStatus `json:"latestCreatedExecution,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Returns true if the Job is currently being acted upon by the system to bring it into the desired state.

	When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.

	If reconciliation succeeded, the following fields will match: observedGeneration and generation, latest_succeeded_execution and latestCreatedExecution.

	If reconciliation failed, observedGeneration and latest_succeeded_execution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions. */
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	/* The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state. */
	// +optional
	TerminalCondition []JobTerminalConditionStatus `json:"terminalCondition,omitempty"`

	/* Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* The last-modified time. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RunJob is the Schema for the run API
// +k8s:openapi-gen=true
type RunJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunJobSpec   `json:"spec,omitempty"`
	Status RunJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RunJobList contains a list of RunJob
type RunJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RunJob{}, &RunJobList{})
}
