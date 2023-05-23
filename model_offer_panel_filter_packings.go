/*
Agridence API

Agridence

API version: v1.1
Contact: it@agridence.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arp_sdk_go_v1

import (
	"encoding/json"
)

// checks if the OfferPanelFilterPackings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferPanelFilterPackings{}

// OfferPanelFilterPackings struct for OfferPanelFilterPackings
type OfferPanelFilterPackings struct {
	Id *int32 `json:"id,omitempty"`
	PackingCode *string `json:"packing_code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewOfferPanelFilterPackings instantiates a new OfferPanelFilterPackings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferPanelFilterPackings() *OfferPanelFilterPackings {
	this := OfferPanelFilterPackings{}
	return &this
}

// NewOfferPanelFilterPackingsWithDefaults instantiates a new OfferPanelFilterPackings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferPanelFilterPackingsWithDefaults() *OfferPanelFilterPackings {
	this := OfferPanelFilterPackings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfferPanelFilterPackings) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPackings) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfferPanelFilterPackings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OfferPanelFilterPackings) SetId(v int32) {
	o.Id = &v
}

// GetPackingCode returns the PackingCode field value if set, zero value otherwise.
func (o *OfferPanelFilterPackings) GetPackingCode() string {
	if o == nil || IsNil(o.PackingCode) {
		var ret string
		return ret
	}
	return *o.PackingCode
}

// GetPackingCodeOk returns a tuple with the PackingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPackings) GetPackingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackingCode) {
		return nil, false
	}
	return o.PackingCode, true
}

// HasPackingCode returns a boolean if a field has been set.
func (o *OfferPanelFilterPackings) HasPackingCode() bool {
	if o != nil && !IsNil(o.PackingCode) {
		return true
	}

	return false
}

// SetPackingCode gets a reference to the given string and assigns it to the PackingCode field.
func (o *OfferPanelFilterPackings) SetPackingCode(v string) {
	o.PackingCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OfferPanelFilterPackings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPackings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OfferPanelFilterPackings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OfferPanelFilterPackings) SetName(v string) {
	o.Name = &v
}

func (o OfferPanelFilterPackings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferPanelFilterPackings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: packing_code is readOnly
	// skip: name is readOnly
	return toSerialize, nil
}

type NullableOfferPanelFilterPackings struct {
	value *OfferPanelFilterPackings
	isSet bool
}

func (v NullableOfferPanelFilterPackings) Get() *OfferPanelFilterPackings {
	return v.value
}

func (v *NullableOfferPanelFilterPackings) Set(val *OfferPanelFilterPackings) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferPanelFilterPackings) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferPanelFilterPackings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferPanelFilterPackings(val *OfferPanelFilterPackings) *NullableOfferPanelFilterPackings {
	return &NullableOfferPanelFilterPackings{value: val, isSet: true}
}

func (v NullableOfferPanelFilterPackings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferPanelFilterPackings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


