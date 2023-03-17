/*
Agridence API

Agridence

API version: v1
Contact: it@agridence.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arp_sdk_go_v1

import (
	"encoding/json"
)

// checks if the GraphDataForSymbolCategoryValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphDataForSymbolCategoryValueInner{}

// GraphDataForSymbolCategoryValueInner struct for GraphDataForSymbolCategoryValueInner
type GraphDataForSymbolCategoryValueInner struct {
	X *string `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
}

// NewGraphDataForSymbolCategoryValueInner instantiates a new GraphDataForSymbolCategoryValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphDataForSymbolCategoryValueInner() *GraphDataForSymbolCategoryValueInner {
	this := GraphDataForSymbolCategoryValueInner{}
	return &this
}

// NewGraphDataForSymbolCategoryValueInnerWithDefaults instantiates a new GraphDataForSymbolCategoryValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphDataForSymbolCategoryValueInnerWithDefaults() *GraphDataForSymbolCategoryValueInner {
	this := GraphDataForSymbolCategoryValueInner{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *GraphDataForSymbolCategoryValueInner) GetX() string {
	if o == nil || isNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDataForSymbolCategoryValueInner) GetXOk() (*string, bool) {
	if o == nil || isNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *GraphDataForSymbolCategoryValueInner) HasX() bool {
	if o != nil && !isNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *GraphDataForSymbolCategoryValueInner) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *GraphDataForSymbolCategoryValueInner) GetY() float32 {
	if o == nil || isNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphDataForSymbolCategoryValueInner) GetYOk() (*float32, bool) {
	if o == nil || isNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *GraphDataForSymbolCategoryValueInner) HasY() bool {
	if o != nil && !isNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *GraphDataForSymbolCategoryValueInner) SetY(v float32) {
	o.Y = &v
}

func (o GraphDataForSymbolCategoryValueInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphDataForSymbolCategoryValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !isNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableGraphDataForSymbolCategoryValueInner struct {
	value *GraphDataForSymbolCategoryValueInner
	isSet bool
}

func (v NullableGraphDataForSymbolCategoryValueInner) Get() *GraphDataForSymbolCategoryValueInner {
	return v.value
}

func (v *NullableGraphDataForSymbolCategoryValueInner) Set(val *GraphDataForSymbolCategoryValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphDataForSymbolCategoryValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphDataForSymbolCategoryValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphDataForSymbolCategoryValueInner(val *GraphDataForSymbolCategoryValueInner) *NullableGraphDataForSymbolCategoryValueInner {
	return &NullableGraphDataForSymbolCategoryValueInner{value: val, isSet: true}
}

func (v NullableGraphDataForSymbolCategoryValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphDataForSymbolCategoryValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

