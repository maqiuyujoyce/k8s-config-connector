// Copyright 2025 Google LLC
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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ColabScheduleGVK = GroupVersion.WithKind("ColabSchedule")

// ColabScheduleSpec defines the desired state of ColabSchedule
// +kcc:proto=google.cloud.aiplatform.v1beta1.Schedule
type ColabScheduleSpec struct {
	// The ColabSchedule name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	//  runs. To explicitly set a timezone to the cron tab, apply a prefix in the
	//  cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	//  The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	//  database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	//  "TZ=America/New_York 1 * * * *".
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.cron
	Cron *string `json:"cron,omitempty"`

	// Request for
	//  [PipelineService.CreatePipelineJob][google.cloud.aiplatform.v1beta1.PipelineService.CreatePipelineJob].
	//  CreatePipelineJobRequest.parent field is required (format:
	//  projects/{project}/locations/{location}).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_pipeline_job_request
	CreatePipelineJobRequest *CreatePipelineJobRequest `json:"createPipelineJobRequest,omitempty"`

	// Request for
	//  [ModelMonitoringService.CreateModelMonitoringJob][google.cloud.aiplatform.v1beta1.ModelMonitoringService.CreateModelMonitoringJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_model_monitoring_job_request
	CreateModelMonitoringJobRequest *CreateModelMonitoringJobRequest `json:"createModelMonitoringJobRequest,omitempty"`

	// Request for
	//  [NotebookService.CreateNotebookExecutionJob][google.cloud.aiplatform.v1beta1.NotebookService.CreateNotebookExecutionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_notebook_execution_job_request
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequest `json:"createNotebookExecutionJobRequest,omitempty"`

	// Required. User provided name of the Schedule.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.display_name
	DisplayName *string `json:"displayName"`

	// Optional. Timestamp after which the first run can be scheduled.
	//  Default to Schedule create time if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Timestamp after which no new runs can be scheduled.
	//  If specified, The schedule will be completed when either
	//  end_time is reached or when scheduled_run_count >= max_run_count.
	//  If not specified, new runs will keep getting scheduled until this Schedule
	//  is paused or deleted. Already scheduled runs will be allowed to complete.
	//  Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Optional. Maximum run count of the schedule.
	//  If specified, The schedule will be completed when either
	//  started_run_count >= max_run_count or when end_time is reached.
	//  If not specified, new runs will keep getting scheduled until this Schedule
	//  is paused or deleted. Already scheduled runs will be allowed to complete.
	//  Unset if not specified.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.max_run_count
	MaxRunCount *int64 `json:"maxRunCount,omitempty"`

	// Required. Maximum number of runs that can be started concurrently for this
	//  Schedule. This is the limit for starting the scheduled requests and not the
	//  execution of the operations/jobs created by the requests (if applicable).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.max_concurrent_run_count
	MaxConcurrentRunCount *int64 `json:"maxConcurrentRunCount"`

	// Optional. Whether new scheduled runs can be queued when max_concurrent_runs
	//  limit is reached. If set to true, new runs will be queued instead of
	//  skipped. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.allow_queueing
	AllowQueueing *bool `json:"allowQueueing,omitempty"`
}

type CreatePipelineJobRequest_Parent struct {
	/* The Project to create the PipelineJob in. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The name of the location to create the PipelineJob in.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable."
	// Required.
	Location string `json:"location"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest
type CreatePipelineJobRequest struct {
	// PipelineJob is a separate resource. I think we should not embed the resource
	// In ColabSchedule.
	//
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest.parent
	//CreatePipelineJobRequest_Parent `json:",inline"`
	//
	//// Required. The PipelineJob to create.
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest.pipeline_job
	//PipelineJob *PipelineJob `json:"pipelineJob"`

	// The ID to use for the PipelineJob, which will become the final component of
	//  the PipelineJob name. If not provided, an ID will be automatically
	//  generated.
	//
	//  This value should be less than 128 characters, and valid characters
	//  are `/[a-z][0-9]-/`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest.pipeline_job_id
	PipelineJobID *string `json:"pipelineJobID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest
type CreateModelMonitoringJobRequest struct {
	// ModelMonitoringJob is a separate resource. I think we should not embed the resource
	// In ColabSchedule.
	//
	//// Required. The parent of the ModelMonitoringJob.
	////  Format:
	////  `projects/{project}/locations/{location}/modelMoniitors/{model_monitor}`
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest.parent
	//Parent *string `json:"parent,omitempty"`
	//
	//// Required. The ModelMonitoringJob to create
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest.model_monitoring_job
	//ModelMonitoringJob *ModelMonitoringJob `json:"modelMonitoringJob,omitempty"`

	// Optional. The ID to use for the Model Monitoring Job, which will become the
	//  final component of the model monitoring job resource name.
	//
	//  The maximum length is 63 characters, and valid characters are
	//  `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest.model_monitoring_job_id
	ModelMonitoringJobID *string `json:"modelMonitoringJobID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest
type CreateNotebookExecutionJobRequest struct {
	// NotebookExecutionJob is a separate resource. I think we should not embed the resource
	// In ColabSchedule.
	//
	//// Required. The resource name of the Location to create the
	////  NotebookExecutionJob. Format: `projects/{project}/locations/{location}`
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest.parent
	//Parent *string `json:"parent,omitempty"`
	//
	//// Required. The NotebookExecutionJob to create.
	//// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest.notebook_execution_job
	//NotebookExecutionJob *NotebookExecutionJob `json:"notebookExecutionJob,omitempty"`

	// Optional. User specified ID for the NotebookExecutionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest.notebook_execution_job_id
	NotebookExecutionJobID *string `json:"notebookExecutionJobID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig
type PipelineJob_RuntimeConfig struct {

	// Required. A path in a Cloud Storage bucket, which will be treated as the
	//  root output directory of the pipeline. It is used by the system to
	//  generate the paths of output artifacts. The artifact paths are generated
	//  with a sub-path pattern `{job_id}/{task_id}/{output_key}` under the
	//  specified output directory. The service account specified in this
	//  pipeline must have the `storage.objects.get` and `storage.objects.create`
	//  permissions for this bucket.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.gcs_output_directory
	GCSOutputDirectory *string `json:"gcsOutputDirectory,omitempty"`

	// TODO: unsupported map type with key string and value message
	//// The runtime parameters of the PipelineJob. The parameters will be
	//// passed into
	//// [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec]
	//// to replace the placeholders at runtime. This field is used by pipelines
	//// built using `PipelineJob.pipeline_spec.schema_version` 2.1.0, such as
	//// pipelines built using Kubeflow Pipelines SDK 1.9 or higher and the v2
	//// DSL.
	//map<string, google.protobuf.Value> parameter_values = 3;

	// Represents the failure policy of a pipeline. Currently, the default of a
	//  pipeline is that the pipeline will continue to run until no more tasks
	//  can be executed, also known as PIPELINE_FAILURE_POLICY_FAIL_SLOW.
	//  However, if a pipeline is set to PIPELINE_FAILURE_POLICY_FAIL_FAST, it
	//  will stop scheduling any new tasks when a task has failed. Any scheduled
	//  tasks will continue to completion.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.failure_policy
	FailurePolicy *string `json:"failurePolicy,omitempty"`

	// TODO: unsupported map type with key string and value message
	//// The runtime artifacts of the PipelineJob. The key will be the input
	//// artifact name and the value would be one of the InputArtifact.
	//map<string, InputArtifact> input_artifacts = 5;

	// Optional. The default runtime for the PipelineJob. If not provided,
	//  Vertex Custom Job(on demand) is used as the runtime. For Vertex Custom
	//  Job, please refer to
	//  https://cloud.google.com/vertex-ai/docs/training/overview.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.default_runtime
	DefaultRuntime *PipelineJob_RuntimeConfig_DefaultRuntime `json:"defaultRuntime,omitempty"`
}

// ColabScheduleStatus defines the config connector machine state of ColabSchedule
type ColabScheduleStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ColabSchedule resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ColabScheduleObservedState `json:"observedState,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Value
type Colab_Value struct {
	// An integer value.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Value.int_value
	IntValue *int64 `json:"intValue,omitempty"`

	// A double value.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Value.double_value
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// A string value.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`
}

// ColabScheduleObservedState is the state of the ColabSchedule resource as most recently observed in GCP.
// +kcc:proto=google.cloud.aiplatform.v1beta1.Schedule
type ColabScheduleObservedState struct {
	// Request for
	//  [PipelineService.CreatePipelineJob][google.cloud.aiplatform.v1beta1.PipelineService.CreatePipelineJob].
	//  CreatePipelineJobRequest.parent field is required (format:
	//  projects/{project}/locations/{location}).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_pipeline_job_request
	CreatePipelineJobRequest *CreatePipelineJobRequestObservedState `json:"createPipelineJobRequest,omitempty"`

	// Request for
	//  [ModelMonitoringService.CreateModelMonitoringJob][google.cloud.aiplatform.v1beta1.ModelMonitoringService.CreateModelMonitoringJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_model_monitoring_job_request
	CreateModelMonitoringJobRequest *CreateModelMonitoringJobRequestObservedState `json:"createModelMonitoringJobRequest,omitempty"`

	// Request for
	//  [NotebookService.CreateNotebookExecutionJob][google.cloud.aiplatform.v1beta1.NotebookService.CreateNotebookExecutionJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_notebook_execution_job_request
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequestObservedState `json:"createNotebookExecutionJobRequest,omitempty"`

	// Output only. The number of runs started by this schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.started_run_count
	StartedRunCount *int64 `json:"startedRunCount,omitempty"`

	// Output only. The state of this Schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.state
	State *string `json:"state,omitempty"`

	// Output only. Timestamp when this Schedule was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Schedule was updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Timestamp when this Schedule should schedule the next run.
	//  Having a next_run_time in the past means the runs are being started
	//  behind schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.next_run_time
	NextRunTime *string `json:"nextRunTime,omitempty"`

	// Output only. Timestamp when this Schedule was last paused.
	//  Unset if never paused.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.last_pause_time
	LastPauseTime *string `json:"lastPauseTime,omitempty"`

	// Output only. Timestamp when this Schedule was last resumed.
	//  Unset if never resumed from pause.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.last_resume_time
	LastResumeTime *string `json:"lastResumeTime,omitempty"`

	// Output only. Whether to backfill missed runs when the schedule is resumed
	//  from PAUSED state. If set to true, all missed runs will be scheduled. New
	//  runs will be scheduled after the backfill is complete. Default to false.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.catch_up
	CatchUp *bool `json:"catchUp,omitempty"`

	// Output only. Response of the last scheduled run.
	//  This is the response for starting the scheduled requests and not the
	//  execution of the operations/jobs created by the requests (if applicable).
	//  Unset if no run has been scheduled yet.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.last_scheduled_run_response
	LastScheduledRunResponse *Schedule_RunResponse `json:"lastScheduledRunResponse,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest
type CreateModelMonitoringJobRequestObservedState struct {
	// Required. The ModelMonitoringJob to create
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateModelMonitoringJobRequest.model_monitoring_job
	ModelMonitoringJob *ModelMonitoringJobObservedState `json:"modelMonitoringJob,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest
type CreateNotebookExecutionJobRequestObservedState struct {
	// Required. The NotebookExecutionJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreateNotebookExecutionJobRequest.notebook_execution_job
	NotebookExecutionJob *NotebookExecutionJobObservedState `json:"notebookExecutionJob,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest
type CreatePipelineJobRequestObservedState struct {
	// Required. The PipelineJob to create.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.CreatePipelineJobRequest.pipeline_job
	PipelineJob *PipelineJobObservedState `json:"pipelineJob,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcolabschedule;gcpcolabschedules
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ColabSchedule is the Schema for the ColabSchedule API
// +k8s:openapi-gen=true
type ColabSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ColabScheduleSpec   `json:"spec,omitempty"`
	Status ColabScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ColabScheduleList contains a list of ColabSchedule
type ColabScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ColabSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ColabSchedule{}, &ColabScheduleList{})
}
