/*
 * madana-api
 *
 * <h1>Using the madana-api</h1>        <p>This documentation contains a Quickstart Guide, relating client functionality and information about the available         endpoints and used datamodels.   </p>       <p> The madana-api and its implementations are still in heavy development. This means that there may be problems in our protocols, or there may be mistakes in our implementations. We take security vulnerabilities very seriously. If you discover a security issue, please bring it to our attention right away! If you find a vulnerability that may affect live deployments -- for example, by exposing a remote execution exploit -- please send your report privately to info@madana.io. Please DO NOT file a public issue. If the issue is a protocol weakness that cannot be immediately exploited or something not yet deployed, just discuss it openly   </p>   <br>   <p> Note: Not all functionality might be acessible without having accquired and api-license token. For more information visit <a href=\"https://www.madana.io\">www.madana.io</a> </p>       <br>
 *
 * API version: 0.5.0-master.29
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package madanaapiclient

import (
	"encoding/json"
)

// JsonV1ObjectMeta 
type JsonV1ObjectMeta struct {
	// 
	Name *string `json:"name,omitempty"`
	// 
	GenerateName *string `json:"generateName,omitempty"`
	// 
	Finalizers *[]string `json:"finalizers,omitempty"`
	// 
	Namespace *string `json:"namespace,omitempty"`
	// 
	SelfLink *string `json:"selfLink,omitempty"`
	// 
	ManagedFields *[]JsonV1ManagedFieldsEntry `json:"managedFields,omitempty"`
	// 
	Generation *float32 `json:"generation,omitempty"`
	// 
	Annotations *map[string]string `json:"annotations,omitempty"`
	// 
	DeletionTimestamp *float32 `json:"deletionTimestamp,omitempty"`
	// 
	DeletionGracePeriodSeconds *float32 `json:"deletionGracePeriodSeconds,omitempty"`
	// 
	OwnerReferences *[]JsonV1OwnerReference `json:"ownerReferences,omitempty"`
	// 
	Uid *string `json:"uid,omitempty"`
	// 
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// 
	CreationTimestamp *float32 `json:"creationTimestamp,omitempty"`
	// 
	ClusterName *string `json:"clusterName,omitempty"`
	// 
	Labels *map[string]string `json:"labels,omitempty"`
}

// NewJsonV1ObjectMeta instantiates a new JsonV1ObjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonV1ObjectMeta() *JsonV1ObjectMeta {
	this := JsonV1ObjectMeta{}
	return &this
}

// NewJsonV1ObjectMetaWithDefaults instantiates a new JsonV1ObjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonV1ObjectMetaWithDefaults() *JsonV1ObjectMeta {
	this := JsonV1ObjectMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JsonV1ObjectMeta) SetName(v string) {
	o.Name = &v
}

// GetGenerateName returns the GenerateName field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetGenerateName() string {
	if o == nil || o.GenerateName == nil {
		var ret string
		return ret
	}
	return *o.GenerateName
}

// GetGenerateNameOk returns a tuple with the GenerateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetGenerateNameOk() (*string, bool) {
	if o == nil || o.GenerateName == nil {
		return nil, false
	}
	return o.GenerateName, true
}

// HasGenerateName returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasGenerateName() bool {
	if o != nil && o.GenerateName != nil {
		return true
	}

	return false
}

// SetGenerateName gets a reference to the given string and assigns it to the GenerateName field.
func (o *JsonV1ObjectMeta) SetGenerateName(v string) {
	o.GenerateName = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetFinalizers() []string {
	if o == nil || o.Finalizers == nil {
		var ret []string
		return ret
	}
	return *o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetFinalizersOk() (*[]string, bool) {
	if o == nil || o.Finalizers == nil {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasFinalizers() bool {
	if o != nil && o.Finalizers != nil {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []string and assigns it to the Finalizers field.
func (o *JsonV1ObjectMeta) SetFinalizers(v []string) {
	o.Finalizers = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *JsonV1ObjectMeta) SetNamespace(v string) {
	o.Namespace = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetSelfLink() string {
	if o == nil || o.SelfLink == nil {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetSelfLinkOk() (*string, bool) {
	if o == nil || o.SelfLink == nil {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasSelfLink() bool {
	if o != nil && o.SelfLink != nil {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *JsonV1ObjectMeta) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetManagedFields returns the ManagedFields field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetManagedFields() []JsonV1ManagedFieldsEntry {
	if o == nil || o.ManagedFields == nil {
		var ret []JsonV1ManagedFieldsEntry
		return ret
	}
	return *o.ManagedFields
}

// GetManagedFieldsOk returns a tuple with the ManagedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetManagedFieldsOk() (*[]JsonV1ManagedFieldsEntry, bool) {
	if o == nil || o.ManagedFields == nil {
		return nil, false
	}
	return o.ManagedFields, true
}

// HasManagedFields returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasManagedFields() bool {
	if o != nil && o.ManagedFields != nil {
		return true
	}

	return false
}

// SetManagedFields gets a reference to the given []JsonV1ManagedFieldsEntry and assigns it to the ManagedFields field.
func (o *JsonV1ObjectMeta) SetManagedFields(v []JsonV1ManagedFieldsEntry) {
	o.ManagedFields = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetGeneration() float32 {
	if o == nil || o.Generation == nil {
		var ret float32
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetGenerationOk() (*float32, bool) {
	if o == nil || o.Generation == nil {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasGeneration() bool {
	if o != nil && o.Generation != nil {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given float32 and assigns it to the Generation field.
func (o *JsonV1ObjectMeta) SetGeneration(v float32) {
	o.Generation = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *JsonV1ObjectMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetDeletionTimestamp returns the DeletionTimestamp field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetDeletionTimestamp() float32 {
	if o == nil || o.DeletionTimestamp == nil {
		var ret float32
		return ret
	}
	return *o.DeletionTimestamp
}

// GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetDeletionTimestampOk() (*float32, bool) {
	if o == nil || o.DeletionTimestamp == nil {
		return nil, false
	}
	return o.DeletionTimestamp, true
}

// HasDeletionTimestamp returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasDeletionTimestamp() bool {
	if o != nil && o.DeletionTimestamp != nil {
		return true
	}

	return false
}

// SetDeletionTimestamp gets a reference to the given float32 and assigns it to the DeletionTimestamp field.
func (o *JsonV1ObjectMeta) SetDeletionTimestamp(v float32) {
	o.DeletionTimestamp = &v
}

// GetDeletionGracePeriodSeconds returns the DeletionGracePeriodSeconds field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetDeletionGracePeriodSeconds() float32 {
	if o == nil || o.DeletionGracePeriodSeconds == nil {
		var ret float32
		return ret
	}
	return *o.DeletionGracePeriodSeconds
}

// GetDeletionGracePeriodSecondsOk returns a tuple with the DeletionGracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetDeletionGracePeriodSecondsOk() (*float32, bool) {
	if o == nil || o.DeletionGracePeriodSeconds == nil {
		return nil, false
	}
	return o.DeletionGracePeriodSeconds, true
}

// HasDeletionGracePeriodSeconds returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasDeletionGracePeriodSeconds() bool {
	if o != nil && o.DeletionGracePeriodSeconds != nil {
		return true
	}

	return false
}

// SetDeletionGracePeriodSeconds gets a reference to the given float32 and assigns it to the DeletionGracePeriodSeconds field.
func (o *JsonV1ObjectMeta) SetDeletionGracePeriodSeconds(v float32) {
	o.DeletionGracePeriodSeconds = &v
}

// GetOwnerReferences returns the OwnerReferences field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetOwnerReferences() []JsonV1OwnerReference {
	if o == nil || o.OwnerReferences == nil {
		var ret []JsonV1OwnerReference
		return ret
	}
	return *o.OwnerReferences
}

// GetOwnerReferencesOk returns a tuple with the OwnerReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetOwnerReferencesOk() (*[]JsonV1OwnerReference, bool) {
	if o == nil || o.OwnerReferences == nil {
		return nil, false
	}
	return o.OwnerReferences, true
}

// HasOwnerReferences returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasOwnerReferences() bool {
	if o != nil && o.OwnerReferences != nil {
		return true
	}

	return false
}

// SetOwnerReferences gets a reference to the given []JsonV1OwnerReference and assigns it to the OwnerReferences field.
func (o *JsonV1ObjectMeta) SetOwnerReferences(v []JsonV1OwnerReference) {
	o.OwnerReferences = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *JsonV1ObjectMeta) SetUid(v string) {
	o.Uid = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *JsonV1ObjectMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetCreationTimestamp() float32 {
	if o == nil || o.CreationTimestamp == nil {
		var ret float32
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetCreationTimestampOk() (*float32, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasCreationTimestamp() bool {
	if o != nil && o.CreationTimestamp != nil {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given float32 and assigns it to the CreationTimestamp field.
func (o *JsonV1ObjectMeta) SetCreationTimestamp(v float32) {
	o.CreationTimestamp = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *JsonV1ObjectMeta) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *JsonV1ObjectMeta) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonV1ObjectMeta) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *JsonV1ObjectMeta) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *JsonV1ObjectMeta) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o JsonV1ObjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GenerateName != nil {
		toSerialize["generateName"] = o.GenerateName
	}
	if o.Finalizers != nil {
		toSerialize["finalizers"] = o.Finalizers
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.SelfLink != nil {
		toSerialize["selfLink"] = o.SelfLink
	}
	if o.ManagedFields != nil {
		toSerialize["managedFields"] = o.ManagedFields
	}
	if o.Generation != nil {
		toSerialize["generation"] = o.Generation
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.DeletionTimestamp != nil {
		toSerialize["deletionTimestamp"] = o.DeletionTimestamp
	}
	if o.DeletionGracePeriodSeconds != nil {
		toSerialize["deletionGracePeriodSeconds"] = o.DeletionGracePeriodSeconds
	}
	if o.OwnerReferences != nil {
		toSerialize["ownerReferences"] = o.OwnerReferences
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.ResourceVersion != nil {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if o.CreationTimestamp != nil {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableJsonV1ObjectMeta struct {
	value *JsonV1ObjectMeta
	isSet bool
}

func (v NullableJsonV1ObjectMeta) Get() *JsonV1ObjectMeta {
	return v.value
}

func (v *NullableJsonV1ObjectMeta) Set(val *JsonV1ObjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonV1ObjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonV1ObjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonV1ObjectMeta(val *JsonV1ObjectMeta) *NullableJsonV1ObjectMeta {
	return &NullableJsonV1ObjectMeta{value: val, isSet: true}
}

func (v NullableJsonV1ObjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonV1ObjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


