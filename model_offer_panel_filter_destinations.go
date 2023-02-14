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

// checks if the OfferPanelFilterDestinations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferPanelFilterDestinations{}

// OfferPanelFilterDestinations struct for OfferPanelFilterDestinations
type OfferPanelFilterDestinations struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DestinationCountry *string `json:"destination_country,omitempty"`
}

// NewOfferPanelFilterDestinations instantiates a new OfferPanelFilterDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferPanelFilterDestinations() *OfferPanelFilterDestinations {
	this := OfferPanelFilterDestinations{}
	return &this
}

// NewOfferPanelFilterDestinationsWithDefaults instantiates a new OfferPanelFilterDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferPanelFilterDestinationsWithDefaults() *OfferPanelFilterDestinations {
	this := OfferPanelFilterDestinations{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfferPanelFilterDestinations) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterDestinations) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfferPanelFilterDestinations) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OfferPanelFilterDestinations) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OfferPanelFilterDestinations) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterDestinations) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OfferPanelFilterDestinations) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OfferPanelFilterDestinations) SetName(v string) {
	o.Name = &v
}

// GetDestinationCountry returns the DestinationCountry field value if set, zero value otherwise.
func (o *OfferPanelFilterDestinations) GetDestinationCountry() string {
	if o == nil || isNil(o.DestinationCountry) {
		var ret string
		return ret
	}
	return *o.DestinationCountry
}

// GetDestinationCountryOk returns a tuple with the DestinationCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterDestinations) GetDestinationCountryOk() (*string, bool) {
	if o == nil || isNil(o.DestinationCountry) {
		return nil, false
	}
	return o.DestinationCountry, true
}

// HasDestinationCountry returns a boolean if a field has been set.
func (o *OfferPanelFilterDestinations) HasDestinationCountry() bool {
	if o != nil && !isNil(o.DestinationCountry) {
		return true
	}

	return false
}

// SetDestinationCountry gets a reference to the given string and assigns it to the DestinationCountry field.
func (o *OfferPanelFilterDestinations) SetDestinationCountry(v string) {
	o.DestinationCountry = &v
}

func (o OfferPanelFilterDestinations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferPanelFilterDestinations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: name is readOnly
	// skip: destination_country is readOnly
	return toSerialize, nil
}

type NullableOfferPanelFilterDestinations struct {
	value *OfferPanelFilterDestinations
	isSet bool
}

func (v NullableOfferPanelFilterDestinations) Get() *OfferPanelFilterDestinations {
	return v.value
}

func (v *NullableOfferPanelFilterDestinations) Set(val *OfferPanelFilterDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferPanelFilterDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferPanelFilterDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferPanelFilterDestinations(val *OfferPanelFilterDestinations) *NullableOfferPanelFilterDestinations {
	return &NullableOfferPanelFilterDestinations{value: val, isSet: true}
}

func (v NullableOfferPanelFilterDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferPanelFilterDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

