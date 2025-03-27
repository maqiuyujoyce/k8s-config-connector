// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:types
// krm.group: colab.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.aiplatform.v1beta1
// resource: ColabRuntimeTemplate:NotebookRuntimeTemplate
// resource: ColabRuntime:NotebookRuntime
// resource: ColabSchedule:Schedule

package v1alpha1

// +kcc:proto=google.cloud.aiplatform.v1beta1.Artifact
type Artifact struct {

	// User provided display name of the Artifact.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The uniform resource identifier of the artifact file.
	//  May be empty if there is no actual artifact file.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.uri
	URI *string `json:"uri,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Artifacts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Artifact (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The state of this Artifact. This is a property of the Artifact, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines), and the system does not prescribe
	//  or check the validity of state transitions.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.state
	State *string `json:"state,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Artifact.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description of the Artifact
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Artifact.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ArtifactTypeSchema
type ArtifactTypeSchema struct {
	// The name of the type. The format of the title must be:
	//  `<namespace>.<title>`.
	//  Examples:
	//   - `aiplatform.Model`
	//   - `acme.CustomModel`
	//  When this field is set, the type must be pre-registered in the MLMD
	//  store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ArtifactTypeSchema.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// Points to a YAML file stored on Cloud Storage describing the
	//  format.
	//  Deprecated. Use [PipelineArtifactTypeSchema.schema_title][] or
	//  [PipelineArtifactTypeSchema.instance_schema][] instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ArtifactTypeSchema.schema_uri
	SchemaURI *string `json:"schemaURI,omitempty"`

	// Contains a raw YAML string, describing the format of
	//  the properties of the type.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ArtifactTypeSchema.instance_schema
	InstanceSchema *string `json:"instanceSchema,omitempty"`

	// The schema version of the artifact. If the value is not set, it defaults
	//  to the latest version in the system.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ArtifactTypeSchema.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.BatchDedicatedResources
type BatchDedicatedResources struct {
	// Required. Immutable. The specification of a single machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// Immutable. The number of machine replicas used at the start of the batch
	//  operation. If not set, Vertex AI decides starting number, not greater than
	//  [max_replica_count][google.cloud.aiplatform.v1beta1.BatchDedicatedResources.max_replica_count]
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.starting_replica_count
	StartingReplicaCount *int32 `json:"startingReplicaCount,omitempty"`

	// Immutable. The maximum number of machine replicas the batch operation may
	//  be scaled to. The default value is 10.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BatchDedicatedResources.max_replica_count
	MaxReplicaCount *int32 `json:"maxReplicaCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.BlurBaselineConfig
type BlurBaselineConfig struct {
	// The standard deviation of the blur kernel for the blurred baseline. The
	//  same blurring parameter is used for both the height and the width
	//  dimension. If not set, the method defaults to the zero (i.e. black for
	//  images) baseline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.BlurBaselineConfig.max_blur_sigma
	MaxBlurSigma *float32 `json:"maxBlurSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Context
type Context struct {
	// Immutable. The resource name of the Context.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.name
	Name *string `json:"name,omitempty"`

	// User provided display name of the Context.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Contexts.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Context (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in schema_name to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Context.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description of the Context
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Examples
type Examples struct {
	// The Cloud Storage input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.example_gcs_source
	ExampleGCSSource *Examples_ExampleGCSSource `json:"exampleGCSSource,omitempty"`

	// The full configuration for the generated index, the semantics are the
	//  same as [metadata][google.cloud.aiplatform.v1beta1.Index.metadata] and
	//  should match
	//  [NearestNeighborSearchConfig](https://cloud.google.com/vertex-ai/docs/explainable-ai/configuring-explanations-example-based#nearest-neighbor-search-config).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.nearest_neighbor_search_config
	NearestNeighborSearchConfig *Value `json:"nearestNeighborSearchConfig,omitempty"`

	// Simplified preset configuration, which automatically sets configuration
	//  values based on the desired query speed-precision trade-off and modality.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.presets
	Presets *Presets `json:"presets,omitempty"`

	// The Cloud Storage locations that contain the instances to be
	//  indexed for approximate nearest neighbor search.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.gcs_source
	GCSSource *GCSSource `json:"gcsSource,omitempty"`

	// The number of neighbors to return when querying for examples.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.neighbor_count
	NeighborCount *int32 `json:"neighborCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Examples.ExampleGcsSource
type Examples_ExampleGCSSource struct {
	// The format in which instances are given, if not specified, assume it's
	//  JSONL format. Currently only JSONL format is supported.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.ExampleGcsSource.data_format
	DataFormat *string `json:"dataFormat,omitempty"`

	// The Cloud Storage location for the input instances.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Examples.ExampleGcsSource.gcs_source
	GCSSource *GCSSource `json:"gcsSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Execution
type Execution struct {

	// User provided display name of the Execution.
	//  May be up to 128 Unicode characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The state of this Execution. This is a property of the Execution, and does
	//  not imply or capture any ongoing process. This property is managed by
	//  clients (such as Vertex AI Pipelines) and the system does not prescribe
	//  or check the validity of state transitions.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.state
	State *string `json:"state,omitempty"`

	// An eTag used to perform consistent read-modify-write updates. If not set, a
	//  blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.etag
	Etag *string `json:"etag,omitempty"`

	// The labels with user-defined metadata to organize your Executions.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//  No more than 64 user labels can be associated with one Execution (System
	//  labels are excluded).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The title of the schema describing the metadata.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.schema_title
	SchemaTitle *string `json:"schemaTitle,omitempty"`

	// The version of the schema in `schema_title` to use.
	//
	//  Schema title and version is expected to be registered in earlier Create
	//  Schema calls. And both are used together as unique identifiers to identify
	//  schemas within the local metadata store.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.schema_version
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// Properties of the Execution.
	//  Top level metadata keys' heading and trailing spaces will be trimmed.
	//  The size of this field should not exceed 200KB.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Description of the Execution
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationMetadata
type ExplanationMetadata struct {

	// TODO: unsupported map type with key string and value message

	// TODO: unsupported map type with key string and value message

	// Points to a YAML file stored on Google Cloud Storage describing the format
	//  of the [feature
	//  attributions][google.cloud.aiplatform.v1beta1.Attribution.feature_attributions].
	//  The schema is defined as an OpenAPI 3.0.2 [Schema
	//  Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	//  AutoML tabular Models always have this field populated by Vertex AI.
	//  Note: The URI given on output may be different, including the URI scheme,
	//  than the one given on input. The output URI will point to a location where
	//  the user only has a read access.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.feature_attributions_schema_uri
	FeatureAttributionsSchemaURI *string `json:"featureAttributionsSchemaURI,omitempty"`

	// Name of the source to generate embeddings for example based explanations.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.latent_space_source
	LatentSpaceSource *string `json:"latentSpaceSource,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata
type ExplanationMetadata_InputMetadata struct {
	// Baseline inputs for this feature.
	//
	//  If no baseline is specified, Vertex AI chooses the baseline for this
	//  feature. If multiple baselines are specified, Vertex AI returns the
	//  average attributions across them in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1beta1.Attribution.feature_attributions].
	//
	//  For Vertex AI-provided Tensorflow images (both 1.x and 2.x), the shape
	//  of each baseline must match the shape of the input tensor. If a scalar is
	//  provided, we broadcast to the same shape as the input tensor.
	//
	//  For custom images, the element of the baselines must be in the same
	//  format as the feature's input in the
	//  [instance][google.cloud.aiplatform.v1beta1.ExplainRequest.instances][].
	//  The schema of any single instance may be specified via Endpoint's
	//  DeployedModels'
	//  [Model's][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	//  [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	//  [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.input_baselines
	InputBaselines []Value `json:"inputBaselines,omitempty"`

	// Name of the input tensor for this feature. Required and is only
	//  applicable to Vertex AI-provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.input_tensor_name
	InputTensorName *string `json:"inputTensorName,omitempty"`

	// Defines how the feature is encoded into the input tensor. Defaults to
	//  IDENTITY.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.encoding
	Encoding *string `json:"encoding,omitempty"`

	// Modality of the feature. Valid values are: numeric, image. Defaults to
	//  numeric.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.modality
	Modality *string `json:"modality,omitempty"`

	// The domain details of the input feature value. Like min/max, original
	//  mean or standard deviation if normalized.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.feature_value_domain
	FeatureValueDomain *ExplanationMetadata_InputMetadata_FeatureValueDomain `json:"featureValueDomain,omitempty"`

	// Specifies the index of the values of the input tensor.
	//  Required when the input tensor is a sparse representation. Refer to
	//  Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.indices_tensor_name
	IndicesTensorName *string `json:"indicesTensorName,omitempty"`

	// Specifies the shape of the values of the input if the input is a sparse
	//  representation. Refer to Tensorflow documentation for more details:
	//  https://www.tensorflow.org/api_docs/python/tf/sparse/SparseTensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.dense_shape_tensor_name
	DenseShapeTensorName *string `json:"denseShapeTensorName,omitempty"`

	// A list of feature names for each index in the input tensor.
	//  Required when the input
	//  [InputMetadata.encoding][google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.encoding]
	//  is BAG_OF_FEATURES, BAG_OF_FEATURES_SPARSE, INDICATOR.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.index_feature_mapping
	IndexFeatureMapping []string `json:"indexFeatureMapping,omitempty"`

	// Encoded tensor is a transformation of the input tensor. Must be provided
	//  if choosing
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1beta1.ExplanationParameters.integrated_gradients_attribution]
	//  or [XRAI
	//  attribution][google.cloud.aiplatform.v1beta1.ExplanationParameters.xrai_attribution]
	//  and the input tensor is not differentiable.
	//
	//  An encoded tensor is generated if the input tensor is encoded by a lookup
	//  table.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.encoded_tensor_name
	EncodedTensorName *string `json:"encodedTensorName,omitempty"`

	// A list of baselines for the encoded tensor.
	//
	//  The shape of each baseline should match the shape of the encoded tensor.
	//  If a scalar is provided, Vertex AI broadcasts to the same shape as the
	//  encoded tensor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.encoded_baselines
	EncodedBaselines []Value `json:"encodedBaselines,omitempty"`

	// Visualization configurations for image explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.visualization
	Visualization *ExplanationMetadata_InputMetadata_Visualization `json:"visualization,omitempty"`

	// Name of the group that the input belongs to. Features with the same group
	//  name will be treated as one feature when computing attributions. Features
	//  grouped together can have different shapes in value. If provided, there
	//  will be one single attribution generated in
	//  [Attribution.feature_attributions][google.cloud.aiplatform.v1beta1.Attribution.feature_attributions],
	//  keyed by the group name.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.group_name
	GroupName *string `json:"groupName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.FeatureValueDomain
type ExplanationMetadata_InputMetadata_FeatureValueDomain struct {
	// The minimum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.FeatureValueDomain.min_value
	MinValue *float32 `json:"minValue,omitempty"`

	// The maximum permissible value for this feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.FeatureValueDomain.max_value
	MaxValue *float32 `json:"maxValue,omitempty"`

	// If this input feature has been normalized to a mean value of 0,
	//  the original_mean specifies the mean value of the domain prior to
	//  normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_mean
	OriginalMean *float32 `json:"originalMean,omitempty"`

	// If this input feature has been normalized to a standard deviation of
	//  1.0, the original_stddev specifies the standard deviation of the domain
	//  prior to normalization.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.FeatureValueDomain.original_stddev
	OriginalStddev *float32 `json:"originalStddev,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization
type ExplanationMetadata_InputMetadata_Visualization struct {
	// Type of the image visualization. Only applicable to
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1beta1.ExplanationParameters.integrated_gradients_attribution].
	//  OUTLINES shows regions of attribution, while PIXELS shows per-pixel
	//  attribution. Defaults to OUTLINES.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.type
	Type *string `json:"type,omitempty"`

	// Whether to only highlight pixels with positive contributions, negative
	//  or both. Defaults to POSITIVE.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.polarity
	Polarity *string `json:"polarity,omitempty"`

	// The color scheme used for the highlighted areas.
	//
	//  Defaults to PINK_GREEN for
	//  [Integrated Gradients
	//  attribution][google.cloud.aiplatform.v1beta1.ExplanationParameters.integrated_gradients_attribution],
	//  which shows positive attributions in green and negative in pink.
	//
	//  Defaults to VIRIDIS for
	//  [XRAI
	//  attribution][google.cloud.aiplatform.v1beta1.ExplanationParameters.xrai_attribution],
	//  which highlights the most influential regions in yellow and the least
	//  influential in blue.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.color_map
	ColorMap *string `json:"colorMap,omitempty"`

	// Excludes attributions above the specified percentile from the
	//  highlighted areas. Using the clip_percent_upperbound and
	//  clip_percent_lowerbound together can be useful for filtering out noise
	//  and making it easier to see areas of strong attribution. Defaults to
	//  99.9.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_upperbound
	ClipPercentUpperbound *float32 `json:"clipPercentUpperbound,omitempty"`

	// Excludes attributions below the specified percentile, from the
	//  highlighted areas. Defaults to 62.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.clip_percent_lowerbound
	ClipPercentLowerbound *float32 `json:"clipPercentLowerbound,omitempty"`

	// How the original image is displayed in the visualization.
	//  Adjusting the overlay can help increase visual clarity if the original
	//  image makes it difficult to view the visualization. Defaults to NONE.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.Visualization.overlay_type
	OverlayType *string `json:"overlayType,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata
type ExplanationMetadata_OutputMetadata struct {
	// Static mapping between the index and display name.
	//
	//  Use this if the outputs are a deterministic n-dimensional array, e.g. a
	//  list of scores of all the classes in a pre-defined order for a
	//  multi-classification Model. It's not feasible if the outputs are
	//  non-deterministic, e.g. the Model produces top-k classes or sort the
	//  outputs by their values.
	//
	//  The shape of the value must be an n-dimensional array of strings. The
	//  number of dimensions must match that of the outputs to be explained.
	//  The
	//  [Attribution.output_display_name][google.cloud.aiplatform.v1beta1.Attribution.output_display_name]
	//  is populated by locating in the mapping with
	//  [Attribution.output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata.index_display_name_mapping
	IndexDisplayNameMapping *Value `json:"indexDisplayNameMapping,omitempty"`

	// Specify a field name in the prediction to look for the display name.
	//
	//  Use this if the prediction contains the display names for the outputs.
	//
	//  The display names in the prediction must have the same shape of the
	//  outputs, so that it can be located by
	//  [Attribution.output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index]
	//  for a specific output.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata.display_name_mapping_key
	DisplayNameMappingKey *string `json:"displayNameMappingKey,omitempty"`

	// Name of the output tensor. Required and is only applicable to Vertex
	//  AI provided images for Tensorflow.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata.output_tensor_name
	OutputTensorName *string `json:"outputTensorName,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationParameters
type ExplanationParameters struct {
	// An attribution method that approximates Shapley values for features that
	//  contribute to the label being predicted. A sampling strategy is used to
	//  approximate the value rather than considering all subsets of features.
	//  Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.sampled_shapley_attribution
	SampledShapleyAttribution *SampledShapleyAttribution `json:"sampledShapleyAttribution,omitempty"`

	// An attribution method that computes Aumann-Shapley values taking
	//  advantage of the model's fully differentiable structure. Refer to this
	//  paper for more details: https://arxiv.org/abs/1703.01365
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.integrated_gradients_attribution
	IntegratedGradientsAttribution *IntegratedGradientsAttribution `json:"integratedGradientsAttribution,omitempty"`

	// An attribution method that redistributes Integrated Gradients
	//  attribution to segmented regions, taking advantage of the model's fully
	//  differentiable structure. Refer to this paper for
	//  more details: https://arxiv.org/abs/1906.02825
	//
	//  XRAI currently performs better on natural images, like a picture of a
	//  house or an animal. If the images are taken in artificial environments,
	//  like a lab or manufacturing line, or from diagnostic equipment, like
	//  x-rays or quality-control cameras, use Integrated Gradients instead.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.xrai_attribution
	XraiAttribution *XraiAttribution `json:"xraiAttribution,omitempty"`

	// Example-based explanations that returns the nearest neighbors from the
	//  provided dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.examples
	Examples *Examples `json:"examples,omitempty"`

	// If populated, returns attributions for top K indices of outputs
	//  (defaults to 1). Only applies to Models that predicts more than one outputs
	//  (e,g, multi-class Models). When set to -1, returns explanations for all
	//  outputs.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.top_k
	TopK *int32 `json:"topK,omitempty"`

	// If populated, only returns attributions that have
	//  [output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index]
	//  contained in output_indices. It must be an ndarray of integers, with the
	//  same shape of the output it's explaining.
	//
	//  If not populated, returns attributions for
	//  [top_k][google.cloud.aiplatform.v1beta1.ExplanationParameters.top_k]
	//  indices of outputs. If neither top_k nor output_indices is populated,
	//  returns the argmax index of the outputs.
	//
	//  Only applicable to Models that predict multiple outputs (e,g, multi-class
	//  Models that predict multiple classes).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationParameters.output_indices
	OutputIndices *ListValue `json:"outputIndices,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ExplanationSpec
type ExplanationSpec struct {
	// Required. Parameters that configure explaining of the Model's predictions.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationSpec.parameters
	Parameters *ExplanationParameters `json:"parameters,omitempty"`

	// Optional. Metadata describing the Model's input and output for explanation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ExplanationSpec.metadata
	Metadata *ExplanationMetadata `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FeatureNoiseSigma
type FeatureNoiseSigma struct {
	// Noise sigma per feature. No noise is added to features that are not set.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureNoiseSigma.noise_sigma
	NoiseSigma []FeatureNoiseSigma_NoiseSigmaForFeature `json:"noiseSigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.FeatureNoiseSigma.NoiseSigmaForFeature
type FeatureNoiseSigma_NoiseSigmaForFeature struct {
	// The name of the input feature for which noise sigma is provided. The
	//  features are defined in
	//  [explanation metadata
	//  inputs][google.cloud.aiplatform.v1beta1.ExplanationMetadata.inputs].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureNoiseSigma.NoiseSigmaForFeature.name
	Name *string `json:"name,omitempty"`

	// This represents the standard deviation of the Gaussian kernel that will
	//  be used to add noise to the feature prior to computing gradients. Similar
	//  to
	//  [noise_sigma][google.cloud.aiplatform.v1beta1.SmoothGradConfig.noise_sigma]
	//  but represents the noise added to the current feature. Defaults to 0.1.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureNoiseSigma.NoiseSigmaForFeature.sigma
	Sigma *float32 `json:"sigma,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.GcsDestination
type GCSDestination struct {
	// Required. Google Cloud Storage URI to output directory. If the uri doesn't
	//  end with
	//  '/', a '/' will be automatically appended. The directory is created if it
	//  doesn't exist.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GcsDestination.output_uri_prefix
	OutputURIPrefix *string `json:"outputURIPrefix,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.GcsSource
type GCSSource struct {
	// Required. Google Cloud Storage URI(-s) to the input file(s). May contain
	//  wildcards. For more information on wildcards, see
	//  https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GcsSource.uris
	Uris []string `json:"uris,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.IntegratedGradientsAttribution
type IntegratedGradientsAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.IntegratedGradientsAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.IntegratedGradientsAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for IG with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.IntegratedGradientsAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.MachineSpec
type MachineSpec struct {
	// Immutable. The type of the machine.
	//
	//  See the [list of machine types supported for
	//  prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	//  See the [list of machine types supported for custom
	//  training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	//  For [DeployedModel][google.cloud.aiplatform.v1beta1.DeployedModel] this
	//  field is optional, and the default value is `n1-standard-2`. For
	//  [BatchPredictionJob][google.cloud.aiplatform.v1beta1.BatchPredictionJob] or
	//  as part of [WorkerPoolSpec][google.cloud.aiplatform.v1beta1.WorkerPoolSpec]
	//  this field is required.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.machine_type
	MachineType *string `json:"machineType,omitempty"`

	// Immutable. The type of accelerator(s) that may be attached to the machine
	//  as per
	//  [accelerator_count][google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_type
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// The number of accelerators to attach to the machine.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count
	AcceleratorCount *int32 `json:"acceleratorCount,omitempty"`

	// Immutable. The topology of the TPUs. Corresponds to the TPU topologies
	//  available from GKE. (Example: tpu_topology: "2x2x1").
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.tpu_topology
	TpuTopology *string `json:"tpuTopology,omitempty"`

	// Optional. Immutable. Configuration controlling how this resource pool
	//  consumes reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.MachineSpec.reservation_affinity
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringAlertCondition
type ModelMonitoringAlertCondition struct {
	// A condition that compares a stats value against a threshold. Alert will
	//  be triggered if value above the threshold.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringAlertCondition.threshold
	Threshold *float64 `json:"threshold,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput
type ModelMonitoringInput struct {
	// Columnized dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.columnized_dataset
	ColumnizedDataset *ModelMonitoringInput_ModelMonitoringDataset `json:"columnizedDataset,omitempty"`

	// Vertex AI Batch prediction Job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.batch_prediction_output
	BatchPredictionOutput *ModelMonitoringInput_BatchPredictionOutput `json:"batchPredictionOutput,omitempty"`

	// Vertex AI Endpoint request & response logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.vertex_endpoint_logs
	VertexEndpointLogs *ModelMonitoringInput_VertexEndpointLogs `json:"vertexEndpointLogs,omitempty"`

	// The time interval (pair of start_time and end_time) for which results
	//  should be returned.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.time_interval
	TimeInterval *Interval `json:"timeInterval,omitempty"`

	// The time offset setting for which results should be returned.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.time_offset
	TimeOffset *ModelMonitoringInput_TimeOffset `json:"timeOffset,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.BatchPredictionOutput
type ModelMonitoringInput_BatchPredictionOutput struct {
	// Vertex AI Batch prediction job resource name. The job must match the
	//  model version specified in [ModelMonitor].[model_monitoring_target].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.BatchPredictionOutput.batch_prediction_job
	BatchPredictionJob *string `json:"batchPredictionJob,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset
type ModelMonitoringInput_ModelMonitoringDataset struct {
	// Resource name of the Vertex AI managed dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.vertex_dataset
	VertexDataset *string `json:"vertexDataset,omitempty"`

	// Google Cloud Storage data source.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.gcs_source
	GCSSource *ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringGCSSource `json:"gcsSource,omitempty"`

	// BigQuery data source.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.bigquery_source
	BigquerySource *ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringBigQuerySource `json:"bigquerySource,omitempty"`

	// The timestamp field. Usually for serving data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.timestamp_field
	TimestampField *string `json:"timestampField,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringBigQuerySource
type ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringBigQuerySource struct {
	// BigQuery URI to a table, up to 2000 characters long. All the columns
	//  in the table will be selected. Accepted forms:
	//
	//  *  BigQuery path. For example:
	//  `bq://projectId.bqDatasetId.bqTableId`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringBigQuerySource.table_uri
	TableURI *string `json:"tableURI,omitempty"`

	// Standard SQL to be used instead of the `table_uri`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringBigQuerySource.query
	Query *string `json:"query,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringGcsSource
type ModelMonitoringInput_ModelMonitoringDataset_ModelMonitoringGCSSource struct {
	// Google Cloud Storage URI to the input file(s). May contain
	//  wildcards. For more information on wildcards, see
	//  https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringGcsSource.gcs_uri
	GCSURI *string `json:"gcsURI,omitempty"`

	// Data format of the dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.ModelMonitoringDataset.ModelMonitoringGcsSource.format
	Format *string `json:"format,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.TimeOffset
type ModelMonitoringInput_TimeOffset struct {
	// [offset] is the time difference from the cut-off time.
	//  For scheduled jobs, the cut-off time is the scheduled time.
	//  For non-scheduled jobs, it's the time when the job was created.
	//  Currently we support the following format:
	//  'w|W': Week, 'd|D': Day, 'h|H': Hour
	//  E.g. '1h' stands for 1 hour, '2d' stands for 2 days.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.TimeOffset.offset
	Offset *string `json:"offset,omitempty"`

	// [window] refers to the scope of data selected for analysis.
	//  It allows you to specify the quantity of data you wish to examine.
	//  Currently we support the following format:
	//  'w|W': Week, 'd|D': Day, 'h|H': Hour
	//  E.g. '1h' stands for 1 hour, '2d' stands for 2 days.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.TimeOffset.window
	Window *string `json:"window,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.VertexEndpointLogs
type ModelMonitoringInput_VertexEndpointLogs struct {
	// List of endpoint resource names. The endpoints must enable the logging
	//  with the [Endpoint].[request_response_logging_config], and must contain
	//  the deployed model corresponding to the model version specified in
	//  [ModelMonitor].[model_monitoring_target].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringInput.VertexEndpointLogs.endpoints
	Endpoints []string `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringJob
type ModelMonitoringJob struct {

	// The display name of the ModelMonitoringJob.
	//  The name can be up to 128 characters long and can consist of any UTF-8.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Monitoring monitoring job spec. It outlines the specifications for
	//  monitoring objectives, notifications, and result exports. If left blank,
	//  the default monitoring specifications from the top-level resource
	//  'ModelMonitor' will be applied. If provided, we will use the specification
	//  defined here rather than the default one.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.model_monitoring_spec
	ModelMonitoringSpec *ModelMonitoringSpec `json:"modelMonitoringSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail
type ModelMonitoringJobExecutionDetail struct {
	// Processed baseline datasets.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.baseline_datasets
	BaselineDatasets []ModelMonitoringJobExecutionDetail_ProcessedDataset `json:"baselineDatasets,omitempty"`

	// Processed target datasets.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.target_datasets
	TargetDatasets []ModelMonitoringJobExecutionDetail_ProcessedDataset `json:"targetDatasets,omitempty"`

	// TODO: unsupported map type with key string and value message

	// Additional job error status.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.ProcessedDataset
type ModelMonitoringJobExecutionDetail_ProcessedDataset struct {
	// Actual data location of the processed dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.ProcessedDataset.location
	Location *string `json:"location,omitempty"`

	// Dataset time range information if any.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJobExecutionDetail.ProcessedDataset.time_range
	TimeRange *Interval `json:"timeRange,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec
type ModelMonitoringNotificationSpec struct {
	// Email alert config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.email_config
	EmailConfig *ModelMonitoringNotificationSpec_EmailConfig `json:"emailConfig,omitempty"`

	// Dump the anomalies to Cloud Logging. The anomalies will be put to json
	//  payload encoded from proto
	//  [google.cloud.aiplatform.logging.ModelMonitoringAnomaliesLogEntry][].
	//  This can be further sinked to Pub/Sub or any other services supported
	//  by Cloud Logging.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.enable_cloud_logging
	EnableCloudLogging *bool `json:"enableCloudLogging,omitempty"`

	// Notification channel config.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.notification_channel_configs
	NotificationChannelConfigs []ModelMonitoringNotificationSpec_NotificationChannelConfig `json:"notificationChannelConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.EmailConfig
type ModelMonitoringNotificationSpec_EmailConfig struct {
	// The email addresses to send the alerts.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.EmailConfig.user_emails
	UserEmails []string `json:"userEmails,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.NotificationChannelConfig
type ModelMonitoringNotificationSpec_NotificationChannelConfig struct {
	// Resource names of the NotificationChannels.
	//  Must be of the format
	//  `projects/<project_id_or_number>/notificationChannels/<channel_id>`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringNotificationSpec.NotificationChannelConfig.notification_channel
	NotificationChannel *string `json:"notificationChannel,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec
type ModelMonitoringObjectiveSpec struct {
	// Tabular monitoring objective.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.tabular_objective
	TabularObjective *ModelMonitoringObjectiveSpec_TabularObjective `json:"tabularObjective,omitempty"`

	// The explanation spec.
	//  This spec is required when the objectives spec includes feature attribution
	//  objectives.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// Baseline dataset.
	//  It could be the training dataset or production serving dataset from a
	//  previous period.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.baseline_dataset
	BaselineDataset *ModelMonitoringInput `json:"baselineDataset,omitempty"`

	// Target dataset.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.target_dataset
	TargetDataset *ModelMonitoringInput `json:"targetDataset,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec
type ModelMonitoringObjectiveSpec_DataDriftSpec struct {
	// Feature names / Prediction output names interested in monitoring.
	//  These should be a subset of the input feature names or prediction output
	//  names specified in the monitoring schema.
	//  If the field is not specified all features / prediction outputs outlied
	//  in the monitoring schema will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec.features
	Features []string `json:"features,omitempty"`

	// Supported metrics type:
	//   * l_infinity
	//   * jensen_shannon_divergence
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec.categorical_metric_type
	CategoricalMetricType *string `json:"categoricalMetricType,omitempty"`

	// Supported metrics type:
	//   * jensen_shannon_divergence
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec.numeric_metric_type
	NumericMetricType *string `json:"numericMetricType,omitempty"`

	// Default alert condition for all the categorical features.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec.default_categorical_alert_condition
	DefaultCategoricalAlertCondition *ModelMonitoringAlertCondition `json:"defaultCategoricalAlertCondition,omitempty"`

	// Default alert condition for all the numeric features.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.DataDriftSpec.default_numeric_alert_condition
	DefaultNumericAlertCondition *ModelMonitoringAlertCondition `json:"defaultNumericAlertCondition,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec
type ModelMonitoringObjectiveSpec_FeatureAttributionSpec struct {
	// Feature names interested in monitoring.
	//  These should be a subset of the input feature names specified in the
	//  monitoring schema. If the field is not specified all features outlied in
	//  the monitoring schema will be used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.features
	Features []string `json:"features,omitempty"`

	// Default alert condition for all the features.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.default_alert_condition
	DefaultAlertCondition *ModelMonitoringAlertCondition `json:"defaultAlertCondition,omitempty"`

	// TODO: unsupported map type with key string and value message

	// The config of resources used by the Model Monitoring during the batch
	//  explanation for non-AutoML models. If not set, `n1-standard-2` machine
	//  type will be used by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.FeatureAttributionSpec.batch_explanation_dedicated_resources
	BatchExplanationDedicatedResources *BatchDedicatedResources `json:"batchExplanationDedicatedResources,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective
type ModelMonitoringObjectiveSpec_TabularObjective struct {
	// Input feature distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.feature_drift_spec
	FeatureDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"featureDriftSpec,omitempty"`

	// Prediction output distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.prediction_output_drift_spec
	PredictionOutputDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"predictionOutputDriftSpec,omitempty"`

	// Feature attribution monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.feature_attribution_spec
	FeatureAttributionSpec *ModelMonitoringObjectiveSpec_FeatureAttributionSpec `json:"featureAttributionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringOutputSpec
type ModelMonitoringOutputSpec struct {
	// Google Cloud Storage base folder path for metrics, error logs, etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringOutputSpec.gcs_base_directory
	GCSBaseDirectory *GCSDestination `json:"gcsBaseDirectory,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringSpec
type ModelMonitoringSpec struct {
	// The monitoring objective spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringSpec.objective_spec
	ObjectiveSpec *ModelMonitoringObjectiveSpec `json:"objectiveSpec,omitempty"`

	// The model monitoring notification spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringSpec.notification_spec
	NotificationSpec *ModelMonitoringNotificationSpec `json:"notificationSpec,omitempty"`

	// The Output destination spec for metrics, error logs, etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringSpec.output_spec
	OutputSpec *ModelMonitoringOutputSpec `json:"outputSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookEucConfig
type NotebookEUCConfig struct {
	// Input only. Whether EUC is disabled in this NotebookRuntimeTemplate.
	//  In proto3, the default value of a boolean is false. In this way, by default
	//  EUC will be enabled for NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookEucConfig.euc_disabled
	EUCDisabled *bool `json:"eucDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob
type NotebookExecutionJob struct {
	// The Dataform Repository pointing to a single file notebook repository.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.dataform_repository_source
	DataformRepositorySource *NotebookExecutionJob_DataformRepositorySource `json:"dataformRepositorySource,omitempty"`

	// The Cloud Storage url pointing to the ipynb file. Format:
	//  `gs://bucket/notebook_file.ipynb`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.gcs_notebook_source
	GCSNotebookSource *NotebookExecutionJob_GCSNotebookSource `json:"gcsNotebookSource,omitempty"`

	// The contents of an input notebook file.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.direct_notebook_source
	DirectNotebookSource *NotebookExecutionJob_DirectNotebookSource `json:"directNotebookSource,omitempty"`

	// The NotebookRuntimeTemplate to source compute configuration from.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.notebook_runtime_template_resource_name
	NotebookRuntimeTemplateResourceName *string `json:"notebookRuntimeTemplateResourceName,omitempty"`

	// The custom compute configuration for an execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.custom_environment_spec
	CustomEnvironmentSpec *NotebookExecutionJob_CustomEnvironmentSpec `json:"customEnvironmentSpec,omitempty"`

	// The Cloud Storage location to upload the result to. Format:
	//  `gs://bucket-name`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.gcs_output_uri
	GCSOutputURI *string `json:"gcsOutputURI,omitempty"`

	// The user email to run the execution as. Only supported by Colab runtimes.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.execution_user
	ExecutionUser *string `json:"executionUser,omitempty"`

	// The service account to run the execution as.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The Workbench runtime configuration to use for the notebook execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.workbench_runtime
	WorkbenchRuntime *NotebookExecutionJob_WorkbenchRuntime `json:"workbenchRuntime,omitempty"`

	// The display name of the NotebookExecutionJob. The name can be up to 128
	//  characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`

	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//  System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	//  and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The name of the kernel to use during notebook execution. If unset, the
	//  default kernel is used.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.kernel_name
	KernelName *string `json:"kernelName,omitempty"`

	// Customer-managed encryption key spec for the notebook execution job.
	//  This field is auto-populated if the
	//  [NotebookRuntimeTemplate][google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate]
	//  has an encryption spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.CustomEnvironmentSpec
type NotebookExecutionJob_CustomEnvironmentSpec struct {
	// The specification of a single machine for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.CustomEnvironmentSpec.machine_spec
	MachineSpec *MachineSpec `json:"machineSpec,omitempty"`

	// The specification of a persistent disk to attach for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.CustomEnvironmentSpec.persistent_disk_spec
	PersistentDiskSpec *PersistentDiskSpec `json:"persistentDiskSpec,omitempty"`

	// The network configuration to use for the execution job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.CustomEnvironmentSpec.network_spec
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.DataformRepositorySource
type NotebookExecutionJob_DataformRepositorySource struct {
	// The resource name of the Dataform Repository. Format:
	//  `projects/{project_id}/locations/{location}/repositories/{repository_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.DataformRepositorySource.dataform_repository_resource_name
	DataformRepositoryResourceName *string `json:"dataformRepositoryResourceName,omitempty"`

	// The commit SHA to read repository with. If unset, the file will be read
	//  at HEAD.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.DataformRepositorySource.commit_sha
	CommitSha *string `json:"commitSha,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.DirectNotebookSource
type NotebookExecutionJob_DirectNotebookSource struct {
	// The base64-encoded contents of the input notebook file.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.DirectNotebookSource.content
	Content []byte `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.GcsNotebookSource
type NotebookExecutionJob_GCSNotebookSource struct {
	// The Cloud Storage uri pointing to the ipynb file. Format:
	//  `gs://bucket/notebook_file.ipynb`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.GcsNotebookSource.uri
	URI *string `json:"uri,omitempty"`

	// The version of the Cloud Storage object to read. If unset, the current
	//  version of the object is read. See
	//  https://cloud.google.com/storage/docs/metadata#generation-number.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.GcsNotebookSource.generation
	Generation *string `json:"generation,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.WorkbenchRuntime
type NotebookExecutionJob_WorkbenchRuntime struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig
type NotebookIdleShutdownConfig struct {
	// Required. Duration is accurate to the second. In Notebook, Idle Timeout is
	//  accurate to minute so the range of idle_timeout (second) is: 10 * 60 ~ 1440
	//  * 60.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_timeout
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Whether Idle Shutdown is disabled in this NotebookRuntimeTemplate.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookIdleShutdownConfig.idle_shutdown_disabled
	IdleShutdownDisabled *bool `json:"idleShutdownDisabled,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PersistentDiskSpec
type PersistentDiskSpec struct {
	// Type of the disk (default is "pd-standard").
	//  Valid values: "pd-ssd" (Persistent Disk Solid State Drive)
	//  "pd-standard" (Persistent Disk Hard Disk Drive)
	//  "pd-balanced" (Balanced Persistent Disk)
	//  "pd-extreme" (Extreme Persistent Disk)
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PersistentDiskSpec.disk_type
	DiskType *string `json:"diskType,omitempty"`

	// Size in GB of the disk (default is 100GB).
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PersistentDiskSpec.disk_size_gb
	DiskSizeGB *int64 `json:"diskSizeGB,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob
type PipelineJob struct {

	// The display name of the Pipeline.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The spec of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec
	PipelineSpec map[string]string `json:"pipelineSpec,omitempty"`

	// The labels with user-defined metadata to organize PipelineJob.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	//
	//  Note there is some reserved label key for Vertex AI Pipelines.
	//  - `vertex-ai-pipelines-run-billing-id`, user set value will get overrided.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Runtime config of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.runtime_config
	RuntimeConfig *PipelineJob_RuntimeConfig `json:"runtimeConfig,omitempty"`

	// Customer-managed encryption key spec for a pipelineJob. If set, this
	//  PipelineJob and all of its sub-resources will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the pipeline workload runs as.
	//  If not specified, the Compute Engine default service account in the project
	//  will be used.
	//  See
	//  https://cloud.google.com/compute/docs/access/service-accounts#default_service_account
	//
	//  Users starting the pipeline must have the `iam.serviceAccounts.actAs`
	//  permission on this service account.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The full name of the Compute Engine
	//  [network](/compute/docs/networks-and-firewalls#networks) to which the
	//  Pipeline Job's workload should be peered. For example,
	//  `projects/12345/global/networks/myVPC`.
	//  [Format](/compute/docs/reference/rest/v1/networks/insert)
	//  is of the form `projects/{project}/global/networks/{network}`.
	//  Where {project} is a project number, as in `12345`, and {network} is a
	//  network name.
	//
	//  Private services access must already be configured for the network.
	//  Pipeline job will apply the network configuration to the Google Cloud
	//  resources being launched, if applied, such as Vertex AI
	//  Training or Dataflow job. If left unspecified, the workload is not peered
	//  with any network.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.network
	Network *string `json:"network,omitempty"`

	// A list of names for the reserved ip ranges under the VPC network
	//  that can be used for this Pipeline Job's workload.
	//
	//  If set, we will deploy the Pipeline Job's workload within the provided ip
	//  ranges. Otherwise, the job will be deployed to any ip ranges under the
	//  provided VPC network.
	//
	//  Example: ['vertex-ai-ip-range'].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`

	// Optional. Configuration for PSC-I for PipelineJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// A template uri from where the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec],
	//  if empty, will be downloaded. Currently, only uri from Vertex Template
	//  Registry & Gallery is supported. Reference to
	//  https://cloud.google.com/vertex-ai/docs/pipelines/create-pipeline-template.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.template_uri
	TemplateURI *string `json:"templateURI,omitempty"`

	// Optional. Whether to do component level validations before job creation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.preflight_validations
	PreflightValidations *bool `json:"preflightValidations,omitempty"`

	// Optional. The original pipeline job id if this pipeline job is a rerun of a
	//  previous pipeline job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.original_pipeline_job_id
	OriginalPipelineJobID *int64 `json:"originalPipelineJobID,omitempty"`

	// Optional. The rerun configs for each task in the pipeline job.
	//  By default, the rerun will:
	//  1. Use the same input artifacts as the original run.
	//  2. Use the same input parameters as the original run.
	//  3. Skip all the tasks that are already succeeded in the original run.
	//  4. Rerun all the tasks that are not succeeded in the original run.
	//  By providing this field, users can override the default behavior and
	//  specify the rerun config for each task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_task_rerun_configs
	PipelineTaskRerunConfigs []PipelineTaskRerunConfig `json:"pipelineTaskRerunConfigs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.DefaultRuntime
type PipelineJob_RuntimeConfig_DefaultRuntime struct {
	// Persistent resource based runtime detail.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.DefaultRuntime.persistent_resource_runtime_detail
	PersistentResourceRuntimeDetail *PipelineJob_RuntimeConfig_PersistentResourceRuntimeDetail `json:"persistentResourceRuntimeDetail,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.InputArtifact
type PipelineJob_RuntimeConfig_InputArtifact struct {
	// Artifact resource id from MLMD. Which is the last portion of an
	//  artifact resource name:
	//  `projects/{project}/locations/{location}/metadataStores/default/artifacts/{artifact_id}`.
	//  The artifact must stay within the same project, location and default
	//  metadatastore as the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.InputArtifact.artifact_id
	ArtifactID *string `json:"artifactID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.PersistentResourceRuntimeDetail
type PipelineJob_RuntimeConfig_PersistentResourceRuntimeDetail struct {
	// Persistent resource name.
	//  Format:
	//  `projects/{project}/locations/{location}/persistentResources/{persistent_resource}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.PersistentResourceRuntimeDetail.persistent_resource_name
	PersistentResourceName *string `json:"persistentResourceName,omitempty"`

	// The max time a pipeline task waits for the required CPU, memory, or
	//  accelerator resource to become available from the specified persistent
	//  resource. Default wait time is 0.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.PersistentResourceRuntimeDetail.task_resource_unavailable_wait_time_ms
	TaskResourceUnavailableWaitTimeMs *int64 `json:"taskResourceUnavailableWaitTimeMs,omitempty"`

	// Specifies the behavior to take if the timeout is reached.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.RuntimeConfig.PersistentResourceRuntimeDetail.task_resource_unavailable_timeout_behavior
	TaskResourceUnavailableTimeoutBehavior *string `json:"taskResourceUnavailableTimeoutBehavior,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJobDetail
type PipelineJobDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskDetail
type PipelineTaskDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.ArtifactList
type PipelineTaskDetail_ArtifactList struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.PipelineTaskStatus
type PipelineTaskDetail_PipelineTaskStatus struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail
type PipelineTaskExecutorDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail
type PipelineTaskExecutorDetail_ContainerDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.CustomJobDetail
type PipelineTaskExecutorDetail_CustomJobDetail struct {
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig
type PipelineTaskRerunConfig struct {
	// Optional. The system generated ID of the task. Retrieved from original run.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.task_id
	TaskID *int64 `json:"taskID,omitempty"`

	// Optional. The name of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.task_name
	TaskName *string `json:"taskName,omitempty"`

	// Optional. The runtime input of the task overridden by the user.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.inputs
	Inputs *PipelineTaskRerunConfig_Inputs `json:"inputs,omitempty"`

	// Optional. Whether to skip this task. Default value is False.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.skip_task
	SkipTask *bool `json:"skipTask,omitempty"`

	// Optional. Whether to skip downstream tasks. Default value is False.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.skip_downstream_tasks
	SkipDownstreamTasks *bool `json:"skipDownstreamTasks,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.ArtifactList
type PipelineTaskRerunConfig_ArtifactList struct {
	// Optional. A list of artifact metadata.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.ArtifactList.artifacts
	Artifacts []RuntimeArtifact `json:"artifacts,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskRerunConfig.Inputs
type PipelineTaskRerunConfig_Inputs struct {

	// TODO: unsupported map type with key string and value message

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTemplateMetadata
type PipelineTemplateMetadata struct {
	// The version_name in artifact registry.
	//
	//  Will always be presented in output if the
	//  [PipelineJob.template_uri][google.cloud.aiplatform.v1beta1.PipelineJob.template_uri]
	//  is from supported template registry.
	//
	//  Format is "sha256:abcdef123456...".
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTemplateMetadata.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Presets
type Presets struct {
	// Preset option controlling parameters for speed-precision trade-off when
	//  querying for examples. If omitted, defaults to `PRECISE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Presets.query
	Query *string `json:"query,omitempty"`

	// The modality of the uploaded model, which automatically configures the
	//  distance measurement and feature normalization for the underlying example
	//  index and queries. If your model does not precisely fit one of these types,
	//  it is okay to choose the closest type.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Presets.modality
	Modality *string `json:"modality,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PscInterfaceConfig
type PSCInterfaceConfig struct {
	// Optional. The full name of the Compute Engine
	//  [network
	//  attachment](https://cloud.google.com/vpc/docs/about-network-attachments) to
	//  attach to the resource.
	//  For example, `projects/12345/regions/us-central1/networkAttachments/myNA`.
	//  is of the form
	//  `projects/{project}/regions/{region}/networkAttachments/{networkAttachment}`.
	//  Where {project} is a project number, as in `12345`, and {networkAttachment}
	//  is a network attachment name.
	//  To specify this field, you must have already [created a network attachment]
	//  (https://cloud.google.com/vpc/docs/create-manage-network-attachments#create-network-attachments).
	//  This field is only used for resources using PSC-I.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PscInterfaceConfig.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ReservationAffinity
type ReservationAffinity struct {
	// Required. Specifies the reservation affinity type.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.reservation_affinity_type
	ReservationAffinityType *string `json:"reservationAffinityType,omitempty"`

	// Optional. Corresponds to the label key of a reservation resource. To target
	//  a SPECIFIC_RESERVATION by name, use
	//  `compute.googleapis.com/reservation-name` as the key and specify the name
	//  of your reservation as its value.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.key
	Key *string `json:"key,omitempty"`

	// Optional. Corresponds to the label values of a reservation resource. This
	//  must be the full resource name of the reservation.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ReservationAffinity.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.RuntimeArtifact
type RuntimeArtifact struct {
	// The name of an artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeArtifact.name
	Name *string `json:"name,omitempty"`

	// The type of the artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeArtifact.type
	Type *ArtifactTypeSchema `json:"type,omitempty"`

	// The URI of the artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeArtifact.uri
	URI *string `json:"uri,omitempty"`

	// TODO: unsupported map type with key string and value message

	// TODO: unsupported map type with key string and value message

	// Properties of the Artifact.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.RuntimeArtifact.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.SampledShapleyAttribution
type SampledShapleyAttribution struct {
	// Required. The number of feature permutations to consider when approximating
	//  the Shapley values.
	//
	//  Valid range of its value is [1, 50], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SampledShapleyAttribution.path_count
	PathCount *int32 `json:"pathCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Schedule.RunResponse
type Schedule_RunResponse struct {
	// The scheduled run time based on the user-specified schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.RunResponse.scheduled_run_time
	ScheduledRunTime *string `json:"scheduledRunTime,omitempty"`

	// The response of the scheduled run.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Schedule.RunResponse.run_response
	RunResponse *string `json:"runResponse,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ShieldedVmConfig
type ShieldedVMConfig struct {
	// Defines whether the instance has [Secure
	//  Boot](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#secure-boot)
	//  enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ShieldedVmConfig.enable_secure_boot
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.SmoothGradConfig
type SmoothGradConfig struct {
	// This is a single float value and will be used to add noise to all the
	//  features. Use this field when all features are normalized to have the
	//  same distribution: scale to range [0, 1], [-1, 1] or z-scoring, where
	//  features are normalized to have 0-mean and 1-variance. Learn more about
	//  [normalization](https://developers.google.com/machine-learning/data-prep/transform/normalization).
	//
	//  For best results the recommended value is about 10% - 20% of the standard
	//  deviation of the input feature. Refer to section 3.2 of the SmoothGrad
	//  paper: https://arxiv.org/pdf/1706.03825.pdf. Defaults to 0.1.
	//
	//  If the distribution is different per feature, set
	//  [feature_noise_sigma][google.cloud.aiplatform.v1beta1.SmoothGradConfig.feature_noise_sigma]
	//  instead for each feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SmoothGradConfig.noise_sigma
	NoiseSigma *float32 `json:"noiseSigma,omitempty"`

	// This is similar to
	//  [noise_sigma][google.cloud.aiplatform.v1beta1.SmoothGradConfig.noise_sigma],
	//  but provides additional flexibility. A separate noise sigma can be
	//  provided for each feature, which is useful if their distributions are
	//  different. No noise is added to features that are not set. If this field
	//  is unset,
	//  [noise_sigma][google.cloud.aiplatform.v1beta1.SmoothGradConfig.noise_sigma]
	//  will be used for all features.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SmoothGradConfig.feature_noise_sigma
	FeatureNoiseSigma *FeatureNoiseSigma `json:"featureNoiseSigma,omitempty"`

	// The number of gradient samples to use for
	//  approximation. The higher this number, the more accurate the gradient
	//  is, but the runtime complexity increases by this factor as well.
	//  Valid range of its value is [1, 50]. Defaults to 3.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.SmoothGradConfig.noisy_sample_count
	NoisySampleCount *int32 `json:"noisySampleCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.XraiAttribution
type XraiAttribution struct {
	// Required. The number of steps for approximating the path integral.
	//  A good value to start is 50 and gradually increase until the
	//  sum to diff property is met within the desired error range.
	//
	//  Valid range of its value is [1, 100], inclusively.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.XraiAttribution.step_count
	StepCount *int32 `json:"stepCount,omitempty"`

	// Config for SmoothGrad approximation of gradients.
	//
	//  When enabled, the gradients are approximated by averaging the gradients
	//  from noisy samples in the vicinity of the inputs. Adding
	//  noise can help improve the computed gradients. Refer to this paper for more
	//  details: https://arxiv.org/pdf/1706.03825.pdf
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.XraiAttribution.smooth_grad_config
	SmoothGradConfig *SmoothGradConfig `json:"smoothGradConfig,omitempty"`

	// Config for XRAI with blur baseline.
	//
	//  When enabled, a linear path from the maximally blurred image to the input
	//  image is created. Using a blurred baseline instead of zero (black image) is
	//  motivated by the BlurIG approach explained here:
	//  https://arxiv.org/abs/2004.03383
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.XraiAttribution.blur_baseline_config
	BlurBaselineConfig *BlurBaselineConfig `json:"blurBaselineConfig,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.Interval
type Interval struct {
	// Optional. Inclusive start of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be the same
	//  or after the start.
	// +kcc:proto:field=google.type.Interval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Exclusive end of the interval.
	//
	//  If specified, a Timestamp matching this interval will have to be before the
	//  end.
	// +kcc:proto:field=google.type.Interval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Context
type ContextObservedState struct {
	// Output only. Timestamp when this Context was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Context was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of resource names of Contexts that are parents of this
	//  Context. A Context may have at most 10 parent_contexts.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Context.parent_contexts
	ParentContexts []string `json:"parentContexts,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.Execution
type ExecutionObservedState struct {
	// Output only. The resource name of the Execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this Execution was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Execution was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Execution.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringJob
type ModelMonitoringJobObservedState struct {
	// Output only. Resource name of a ModelMonitoringJob. Format:
	//  `projects/{project_id}/locations/{location_id}/modelMonitors/{model_monitor_id}/modelMonitoringJobs/{model_monitoring_job_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was updated most
	//  recently.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the monitoring job.
	//   * When the job is still creating, the state will be 'JOB_STATE_PENDING'.
	//   * Once the job is successfully created, the state will be
	//     'JOB_STATE_RUNNING'.
	//   * Once the job is finished, the state will be one of
	//     'JOB_STATE_FAILED', 'JOB_STATE_SUCCEEDED',
	//     'JOB_STATE_PARTIALLY_SUCCEEDED'.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.state
	State *string `json:"state,omitempty"`

	// Output only. Schedule resource name. It will only appear when this job is
	//  triggered by a schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.schedule
	Schedule *string `json:"schedule,omitempty"`

	// Output only. Execution results for all the monitoring objectives.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.job_execution_detail
	JobExecutionDetail *ModelMonitoringJobExecutionDetail `json:"jobExecutionDetail,omitempty"`

	// Output only. Timestamp when this ModelMonitoringJob was scheduled. It will
	//  only appear when this job is triggered by a schedule.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringJob.schedule_time
	ScheduleTime *string `json:"scheduleTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookEucConfig
type NotebookEUCConfigObservedState struct {
	// Output only. Whether ActAs check is bypassed for service account attached
	//  to the VM. If false, we need ActAs check for the default Compute Engine
	//  Service account. When a Runtime is created, a VM is allocated using Default
	//  Compute Engine Service Account. Any user requesting to use this Runtime
	//  requires Service Account User (ActAs) permission over this SA. If true,
	//  Runtime owner is using EUC and does not require the above permission as VM
	//  no longer use default Compute Engine SA, but a P4SA.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookEucConfig.bypass_actas_check
	BypassActasCheck *bool `json:"bypassActasCheck,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.NotebookExecutionJob
type NotebookExecutionJobObservedState struct {
	// Output only. The resource name of this NotebookExecutionJob. Format:
	//  `projects/{project_id}/locations/{location}/notebookExecutionJobs/{job_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.name
	Name *string `json:"name,omitempty"`

	// Output only. The Schedule resource name if this job is triggered by one.
	//  Format:
	//  `projects/{project_id}/locations/{location}/schedules/{schedule_id}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.schedule_resource_name
	ScheduleResourceName *string `json:"scheduleResourceName,omitempty"`

	// Output only. The state of the NotebookExecutionJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.job_state
	JobState *string `json:"jobState,omitempty"`

	// Output only. Populated when the NotebookExecutionJob is completed. When
	//  there is an error during notebook execution, the error details are
	//  populated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.status
	Status *Status `json:"status,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this NotebookExecutionJob was most recently
	//  updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.NotebookExecutionJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJob
type PipelineJobObservedState struct {
	// Output only. The resource name of the PipelineJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.name
	Name *string `json:"name,omitempty"`

	// Output only. Pipeline creation time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Pipeline start time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Pipeline end time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. Timestamp when this PipelineJob was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The detailed state of the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.state
	State *string `json:"state,omitempty"`

	// Output only. The details of pipeline run. Not available in the list view.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.job_detail
	JobDetail *PipelineJobDetail `json:"jobDetail,omitempty"`

	// Output only. The error that occurred during pipeline execution.
	//  Only populated when the pipeline's state is FAILED or CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.error
	Error *Status `json:"error,omitempty"`

	// Output only. Pipeline template metadata. Will fill up fields if
	//  [PipelineJob.template_uri][google.cloud.aiplatform.v1beta1.PipelineJob.template_uri]
	//  is from supported template registry.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.template_metadata
	TemplateMetadata *PipelineTemplateMetadata `json:"templateMetadata,omitempty"`

	// Output only. The schedule resource name.
	//  Only returned if the Pipeline is created by Schedule API.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.schedule_name
	ScheduleName *string `json:"scheduleName,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJob.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineJobDetail
type PipelineJobDetailObservedState struct {
	// Output only. The context of the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJobDetail.pipeline_context
	PipelineContext *Context `json:"pipelineContext,omitempty"`

	// Output only. The context of the current pipeline run.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJobDetail.pipeline_run_context
	PipelineRunContext *Context `json:"pipelineRunContext,omitempty"`

	// Output only. The runtime details of the tasks under the pipeline.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineJobDetail.task_details
	TaskDetails []PipelineTaskDetail `json:"taskDetails,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskDetail
type PipelineTaskDetailObservedState struct {
	// Output only. The system generated ID of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.task_id
	TaskID *int64 `json:"taskID,omitempty"`

	// Output only. The id of the parent task if the task is within a component
	//  scope. Empty if the task is at the root level.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.parent_task_id
	ParentTaskID *int64 `json:"parentTaskID,omitempty"`

	// Output only. The user specified name of the task that is defined in
	//  [pipeline_spec][google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.task_name
	TaskName *string `json:"taskName,omitempty"`

	// Output only. Task create time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Task start time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Task end time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.end_time
	EndTime *string `json:"endTime,omitempty"`

	// Output only. The detailed execution info.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.executor_detail
	ExecutorDetail *PipelineTaskExecutorDetail `json:"executorDetail,omitempty"`

	// Output only. State of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.state
	State *string `json:"state,omitempty"`

	// Output only. The execution metadata of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.execution
	Execution *Execution `json:"execution,omitempty"`

	// Output only. The error that occurred during task execution.
	//  Only populated when the task's state is FAILED or CANCELLED.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.error
	Error *Status `json:"error,omitempty"`

	// Output only. A list of task status. This field keeps a record of task
	//  status evolving over time.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.pipeline_task_status
	PipelineTaskStatus []PipelineTaskDetail_PipelineTaskStatus `json:"pipelineTaskStatus,omitempty"`

	// TODO: unsupported map type with key string and value message

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.PipelineTaskStatus
type PipelineTaskDetail_PipelineTaskStatusObservedState struct {
	// Output only. Update time of this status.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.PipelineTaskStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the task.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.PipelineTaskStatus.state
	State *string `json:"state,omitempty"`

	// Output only. The error that occurred during the state. May be set when
	//  the state is any of the non-final state (PENDING/RUNNING/CANCELLING) or
	//  FAILED state. If the state is FAILED, the error here is final and not
	//  going to be retried. If the state is a non-final state, the error
	//  indicates a system-error being retried.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskDetail.PipelineTaskStatus.error
	Error *Status `json:"error,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail
type PipelineTaskExecutorDetailObservedState struct {
	// Output only. The detailed info for a container executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.container_detail
	ContainerDetail *PipelineTaskExecutorDetail_ContainerDetail `json:"containerDetail,omitempty"`

	// Output only. The detailed info for a custom job executor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.custom_job_detail
	CustomJobDetail *PipelineTaskExecutorDetail_CustomJobDetail `json:"customJobDetail,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail
type PipelineTaskExecutorDetail_ContainerDetailObservedState struct {
	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob] for the main
	//  container execution.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail.main_job
	MainJob *string `json:"mainJob,omitempty"`

	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob] for the
	//  pre-caching-check container execution. This job will be available if the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec]
	//  specifies the `pre_caching_check` hook in the lifecycle events.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail.pre_caching_check_job
	PreCachingCheckJob *string `json:"preCachingCheckJob,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob] for the main
	//  container executions. The list includes the all attempts in chronological
	//  order.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail.failed_main_jobs
	FailedMainJobs []string `json:"failedMainJobs,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob] for the
	//  pre-caching-check container executions. This job will be available if the
	//  [PipelineJob.pipeline_spec][google.cloud.aiplatform.v1beta1.PipelineJob.pipeline_spec]
	//  specifies the `pre_caching_check` hook in the lifecycle events. The list
	//  includes the all attempts in chronological order.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.ContainerDetail.failed_pre_caching_check_jobs
	FailedPreCachingCheckJobs []string `json:"failedPreCachingCheckJobs,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.CustomJobDetail
type PipelineTaskExecutorDetail_CustomJobDetailObservedState struct {
	// Output only. The name of the
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob].
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.CustomJobDetail.job
	Job *string `json:"job,omitempty"`

	// Output only. The names of the previously failed
	//  [CustomJob][google.cloud.aiplatform.v1beta1.CustomJob]. The list includes
	//  the all attempts in chronological order.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.PipelineTaskExecutorDetail.CustomJobDetail.failed_jobs
	FailedJobs []string `json:"failedJobs,omitempty"`
}
