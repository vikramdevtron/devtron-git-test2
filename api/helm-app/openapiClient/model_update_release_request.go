/*
Devtron Labs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateReleaseRequest struct for UpdateReleaseRequest
type UpdateReleaseRequest struct {
	// helm app id
	AppId *string `json:"appId,omitempty"`
	// updated values yaml string
	ValuesYaml *string `json:"valuesYaml,omitempty"`
}

// NewUpdateReleaseRequest instantiates a new UpdateReleaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReleaseRequest() *UpdateReleaseRequest {
	this := UpdateReleaseRequest{}
	return &this
}

// NewUpdateReleaseRequestWithDefaults instantiates a new UpdateReleaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReleaseRequestWithDefaults() *UpdateReleaseRequest {
	this := UpdateReleaseRequest{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UpdateReleaseRequest) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseRequest) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UpdateReleaseRequest) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UpdateReleaseRequest) SetAppId(v string) {
	o.AppId = &v
}

// GetValuesYaml returns the ValuesYaml field value if set, zero value otherwise.
func (o *UpdateReleaseRequest) GetValuesYaml() string {
	if o == nil || o.ValuesYaml == nil {
		var ret string
		return ret
	}
	return *o.ValuesYaml
}

// GetValuesYamlOk returns a tuple with the ValuesYaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseRequest) GetValuesYamlOk() (*string, bool) {
	if o == nil || o.ValuesYaml == nil {
		return nil, false
	}
	return o.ValuesYaml, true
}

// HasValuesYaml returns a boolean if a field has been set.
func (o *UpdateReleaseRequest) HasValuesYaml() bool {
	if o != nil && o.ValuesYaml != nil {
		return true
	}

	return false
}

// SetValuesYaml gets a reference to the given string and assigns it to the ValuesYaml field.
func (o *UpdateReleaseRequest) SetValuesYaml(v string) {
	o.ValuesYaml = &v
}

func (o UpdateReleaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.ValuesYaml != nil {
		toSerialize["valuesYaml"] = o.ValuesYaml
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReleaseRequest struct {
	value *UpdateReleaseRequest
	isSet bool
}

func (v NullableUpdateReleaseRequest) Get() *UpdateReleaseRequest {
	return v.value
}

func (v *NullableUpdateReleaseRequest) Set(val *UpdateReleaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReleaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReleaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReleaseRequest(val *UpdateReleaseRequest) *NullableUpdateReleaseRequest {
	return &NullableUpdateReleaseRequest{value: val, isSet: true}
}

func (v NullableUpdateReleaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReleaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


