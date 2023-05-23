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

// checks if the OfferPanelFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferPanelFilter{}

// OfferPanelFilter struct for OfferPanelFilter
type OfferPanelFilter struct {
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	Factories []OfferPanelFilterFactory `json:"factories,omitempty"`
	DefaultFactory *string `json:"default_factory,omitempty"`
	Countries []OfferPanelFilterCountry `json:"countries,omitempty"`
	DefaultCountry *string `json:"default_country,omitempty"`
	Ports []OfferPanelFilterPort `json:"ports,omitempty"`
	DefaultPort *string `json:"default_port,omitempty"`
	Grades []OfferPanelFilterGrade `json:"grades,omitempty"`
	DefaultGrade *string `json:"default_grade,omitempty"`
	Packings []OfferPanelFilterPackings `json:"packings,omitempty"`
	DefaultPacking *string `json:"default_packing,omitempty"`
	Payments []OfferPanelFilterPayments `json:"payments,omitempty"`
	DefaultPayment *string `json:"default_payment,omitempty"`
	Shippings []OfferPanelFilterShippings `json:"shippings,omitempty"`
	DefaultShipping *string `json:"default_shipping,omitempty"`
	Destinations []OfferPanelFilterDestinations `json:"destinations,omitempty"`
	DefaultDestination *string `json:"default_destination,omitempty"`
	Months []string `json:"months,omitempty"`
}

// NewOfferPanelFilter instantiates a new OfferPanelFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferPanelFilter(name string, type_ string) *OfferPanelFilter {
	this := OfferPanelFilter{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewOfferPanelFilterWithDefaults instantiates a new OfferPanelFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferPanelFilterWithDefaults() *OfferPanelFilter {
	this := OfferPanelFilter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OfferPanelFilter) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *OfferPanelFilter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OfferPanelFilter) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *OfferPanelFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OfferPanelFilter) SetType(v string) {
	o.Type = v
}

// GetFactories returns the Factories field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetFactories() []OfferPanelFilterFactory {
	if o == nil || IsNil(o.Factories) {
		var ret []OfferPanelFilterFactory
		return ret
	}
	return o.Factories
}

// GetFactoriesOk returns a tuple with the Factories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetFactoriesOk() ([]OfferPanelFilterFactory, bool) {
	if o == nil || IsNil(o.Factories) {
		return nil, false
	}
	return o.Factories, true
}

// HasFactories returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasFactories() bool {
	if o != nil && !IsNil(o.Factories) {
		return true
	}

	return false
}

// SetFactories gets a reference to the given []OfferPanelFilterFactory and assigns it to the Factories field.
func (o *OfferPanelFilter) SetFactories(v []OfferPanelFilterFactory) {
	o.Factories = v
}

// GetDefaultFactory returns the DefaultFactory field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultFactory() string {
	if o == nil || IsNil(o.DefaultFactory) {
		var ret string
		return ret
	}
	return *o.DefaultFactory
}

// GetDefaultFactoryOk returns a tuple with the DefaultFactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultFactoryOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultFactory) {
		return nil, false
	}
	return o.DefaultFactory, true
}

// HasDefaultFactory returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultFactory() bool {
	if o != nil && !IsNil(o.DefaultFactory) {
		return true
	}

	return false
}

// SetDefaultFactory gets a reference to the given string and assigns it to the DefaultFactory field.
func (o *OfferPanelFilter) SetDefaultFactory(v string) {
	o.DefaultFactory = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetCountries() []OfferPanelFilterCountry {
	if o == nil || IsNil(o.Countries) {
		var ret []OfferPanelFilterCountry
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetCountriesOk() ([]OfferPanelFilterCountry, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []OfferPanelFilterCountry and assigns it to the Countries field.
func (o *OfferPanelFilter) SetCountries(v []OfferPanelFilterCountry) {
	o.Countries = v
}

// GetDefaultCountry returns the DefaultCountry field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultCountry() string {
	if o == nil || IsNil(o.DefaultCountry) {
		var ret string
		return ret
	}
	return *o.DefaultCountry
}

// GetDefaultCountryOk returns a tuple with the DefaultCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultCountryOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultCountry) {
		return nil, false
	}
	return o.DefaultCountry, true
}

// HasDefaultCountry returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultCountry() bool {
	if o != nil && !IsNil(o.DefaultCountry) {
		return true
	}

	return false
}

// SetDefaultCountry gets a reference to the given string and assigns it to the DefaultCountry field.
func (o *OfferPanelFilter) SetDefaultCountry(v string) {
	o.DefaultCountry = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetPorts() []OfferPanelFilterPort {
	if o == nil || IsNil(o.Ports) {
		var ret []OfferPanelFilterPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetPortsOk() ([]OfferPanelFilterPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []OfferPanelFilterPort and assigns it to the Ports field.
func (o *OfferPanelFilter) SetPorts(v []OfferPanelFilterPort) {
	o.Ports = v
}

// GetDefaultPort returns the DefaultPort field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultPort() string {
	if o == nil || IsNil(o.DefaultPort) {
		var ret string
		return ret
	}
	return *o.DefaultPort
}

// GetDefaultPortOk returns a tuple with the DefaultPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultPortOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPort) {
		return nil, false
	}
	return o.DefaultPort, true
}

// HasDefaultPort returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultPort() bool {
	if o != nil && !IsNil(o.DefaultPort) {
		return true
	}

	return false
}

// SetDefaultPort gets a reference to the given string and assigns it to the DefaultPort field.
func (o *OfferPanelFilter) SetDefaultPort(v string) {
	o.DefaultPort = &v
}

// GetGrades returns the Grades field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetGrades() []OfferPanelFilterGrade {
	if o == nil || IsNil(o.Grades) {
		var ret []OfferPanelFilterGrade
		return ret
	}
	return o.Grades
}

// GetGradesOk returns a tuple with the Grades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetGradesOk() ([]OfferPanelFilterGrade, bool) {
	if o == nil || IsNil(o.Grades) {
		return nil, false
	}
	return o.Grades, true
}

// HasGrades returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasGrades() bool {
	if o != nil && !IsNil(o.Grades) {
		return true
	}

	return false
}

// SetGrades gets a reference to the given []OfferPanelFilterGrade and assigns it to the Grades field.
func (o *OfferPanelFilter) SetGrades(v []OfferPanelFilterGrade) {
	o.Grades = v
}

// GetDefaultGrade returns the DefaultGrade field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultGrade() string {
	if o == nil || IsNil(o.DefaultGrade) {
		var ret string
		return ret
	}
	return *o.DefaultGrade
}

// GetDefaultGradeOk returns a tuple with the DefaultGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultGradeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultGrade) {
		return nil, false
	}
	return o.DefaultGrade, true
}

// HasDefaultGrade returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultGrade() bool {
	if o != nil && !IsNil(o.DefaultGrade) {
		return true
	}

	return false
}

// SetDefaultGrade gets a reference to the given string and assigns it to the DefaultGrade field.
func (o *OfferPanelFilter) SetDefaultGrade(v string) {
	o.DefaultGrade = &v
}

// GetPackings returns the Packings field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetPackings() []OfferPanelFilterPackings {
	if o == nil || IsNil(o.Packings) {
		var ret []OfferPanelFilterPackings
		return ret
	}
	return o.Packings
}

// GetPackingsOk returns a tuple with the Packings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetPackingsOk() ([]OfferPanelFilterPackings, bool) {
	if o == nil || IsNil(o.Packings) {
		return nil, false
	}
	return o.Packings, true
}

// HasPackings returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasPackings() bool {
	if o != nil && !IsNil(o.Packings) {
		return true
	}

	return false
}

// SetPackings gets a reference to the given []OfferPanelFilterPackings and assigns it to the Packings field.
func (o *OfferPanelFilter) SetPackings(v []OfferPanelFilterPackings) {
	o.Packings = v
}

// GetDefaultPacking returns the DefaultPacking field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultPacking() string {
	if o == nil || IsNil(o.DefaultPacking) {
		var ret string
		return ret
	}
	return *o.DefaultPacking
}

// GetDefaultPackingOk returns a tuple with the DefaultPacking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultPackingOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPacking) {
		return nil, false
	}
	return o.DefaultPacking, true
}

// HasDefaultPacking returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultPacking() bool {
	if o != nil && !IsNil(o.DefaultPacking) {
		return true
	}

	return false
}

// SetDefaultPacking gets a reference to the given string and assigns it to the DefaultPacking field.
func (o *OfferPanelFilter) SetDefaultPacking(v string) {
	o.DefaultPacking = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetPayments() []OfferPanelFilterPayments {
	if o == nil || IsNil(o.Payments) {
		var ret []OfferPanelFilterPayments
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetPaymentsOk() ([]OfferPanelFilterPayments, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []OfferPanelFilterPayments and assigns it to the Payments field.
func (o *OfferPanelFilter) SetPayments(v []OfferPanelFilterPayments) {
	o.Payments = v
}

// GetDefaultPayment returns the DefaultPayment field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultPayment() string {
	if o == nil || IsNil(o.DefaultPayment) {
		var ret string
		return ret
	}
	return *o.DefaultPayment
}

// GetDefaultPaymentOk returns a tuple with the DefaultPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPayment) {
		return nil, false
	}
	return o.DefaultPayment, true
}

// HasDefaultPayment returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultPayment() bool {
	if o != nil && !IsNil(o.DefaultPayment) {
		return true
	}

	return false
}

// SetDefaultPayment gets a reference to the given string and assigns it to the DefaultPayment field.
func (o *OfferPanelFilter) SetDefaultPayment(v string) {
	o.DefaultPayment = &v
}

// GetShippings returns the Shippings field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetShippings() []OfferPanelFilterShippings {
	if o == nil || IsNil(o.Shippings) {
		var ret []OfferPanelFilterShippings
		return ret
	}
	return o.Shippings
}

// GetShippingsOk returns a tuple with the Shippings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetShippingsOk() ([]OfferPanelFilterShippings, bool) {
	if o == nil || IsNil(o.Shippings) {
		return nil, false
	}
	return o.Shippings, true
}

// HasShippings returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasShippings() bool {
	if o != nil && !IsNil(o.Shippings) {
		return true
	}

	return false
}

// SetShippings gets a reference to the given []OfferPanelFilterShippings and assigns it to the Shippings field.
func (o *OfferPanelFilter) SetShippings(v []OfferPanelFilterShippings) {
	o.Shippings = v
}

// GetDefaultShipping returns the DefaultShipping field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultShipping() string {
	if o == nil || IsNil(o.DefaultShipping) {
		var ret string
		return ret
	}
	return *o.DefaultShipping
}

// GetDefaultShippingOk returns a tuple with the DefaultShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultShippingOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultShipping) {
		return nil, false
	}
	return o.DefaultShipping, true
}

// HasDefaultShipping returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultShipping() bool {
	if o != nil && !IsNil(o.DefaultShipping) {
		return true
	}

	return false
}

// SetDefaultShipping gets a reference to the given string and assigns it to the DefaultShipping field.
func (o *OfferPanelFilter) SetDefaultShipping(v string) {
	o.DefaultShipping = &v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDestinations() []OfferPanelFilterDestinations {
	if o == nil || IsNil(o.Destinations) {
		var ret []OfferPanelFilterDestinations
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDestinationsOk() ([]OfferPanelFilterDestinations, bool) {
	if o == nil || IsNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []OfferPanelFilterDestinations and assigns it to the Destinations field.
func (o *OfferPanelFilter) SetDestinations(v []OfferPanelFilterDestinations) {
	o.Destinations = v
}

// GetDefaultDestination returns the DefaultDestination field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetDefaultDestination() string {
	if o == nil || IsNil(o.DefaultDestination) {
		var ret string
		return ret
	}
	return *o.DefaultDestination
}

// GetDefaultDestinationOk returns a tuple with the DefaultDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetDefaultDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultDestination) {
		return nil, false
	}
	return o.DefaultDestination, true
}

// HasDefaultDestination returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasDefaultDestination() bool {
	if o != nil && !IsNil(o.DefaultDestination) {
		return true
	}

	return false
}

// SetDefaultDestination gets a reference to the given string and assigns it to the DefaultDestination field.
func (o *OfferPanelFilter) SetDefaultDestination(v string) {
	o.DefaultDestination = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *OfferPanelFilter) GetMonths() []string {
	if o == nil || IsNil(o.Months) {
		var ret []string
		return ret
	}
	return o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilter) GetMonthsOk() ([]string, bool) {
	if o == nil || IsNil(o.Months) {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *OfferPanelFilter) HasMonths() bool {
	if o != nil && !IsNil(o.Months) {
		return true
	}

	return false
}

// SetMonths gets a reference to the given []string and assigns it to the Months field.
func (o *OfferPanelFilter) SetMonths(v []string) {
	o.Months = v
}

func (o OfferPanelFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferPanelFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	// skip: factories is readOnly
	// skip: default_factory is readOnly
	// skip: countries is readOnly
	// skip: default_country is readOnly
	// skip: ports is readOnly
	// skip: default_port is readOnly
	// skip: grades is readOnly
	// skip: default_grade is readOnly
	// skip: packings is readOnly
	// skip: default_packing is readOnly
	// skip: payments is readOnly
	// skip: default_payment is readOnly
	// skip: shippings is readOnly
	// skip: default_shipping is readOnly
	// skip: destinations is readOnly
	// skip: default_destination is readOnly
	// skip: months is readOnly
	return toSerialize, nil
}

type NullableOfferPanelFilter struct {
	value *OfferPanelFilter
	isSet bool
}

func (v NullableOfferPanelFilter) Get() *OfferPanelFilter {
	return v.value
}

func (v *NullableOfferPanelFilter) Set(val *OfferPanelFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferPanelFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferPanelFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferPanelFilter(val *OfferPanelFilter) *NullableOfferPanelFilter {
	return &NullableOfferPanelFilter{value: val, isSet: true}
}

func (v NullableOfferPanelFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferPanelFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


