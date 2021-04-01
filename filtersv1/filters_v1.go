/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.28.0-55613c9e-20210220-164656
 */

// Package filtersv1 : Operations and models for the FiltersV1 service
package filtersv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	common "github.ibm.com/ibmcloud/networking-go-sdk/common"
	"io"
	"net/http"
	"reflect"
	"time"
)

// FiltersV1 : Filters
//
// Version: 1.0.1
type FiltersV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "filters"

// FiltersV1Options : Service options
type FiltersV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewFiltersV1UsingExternalConfig : constructs an instance of FiltersV1 with passed in options and external configuration.
func NewFiltersV1UsingExternalConfig(options *FiltersV1Options) (filters *FiltersV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	filters, err = NewFiltersV1(options)
	if err != nil {
		return
	}

	err = filters.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = filters.Service.SetServiceURL(options.URL)
	}
	return
}

// NewFiltersV1 : constructs an instance of FiltersV1 with passed in options.
func NewFiltersV1(options *FiltersV1Options) (service *FiltersV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &FiltersV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "filters" suitable for processing requests.
func (filters *FiltersV1) Clone() *FiltersV1 {
	if core.IsNil(filters) {
		return nil
	}
	clone := *filters
	clone.Service = filters.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (filters *FiltersV1) SetServiceURL(url string) error {
	return filters.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (filters *FiltersV1) GetServiceURL() string {
	return filters.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (filters *FiltersV1) SetDefaultHeaders(headers http.Header) {
	filters.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (filters *FiltersV1) SetEnableGzipCompression(enableGzip bool) {
	filters.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (filters *FiltersV1) GetEnableGzipCompression() bool {
	return filters.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (filters *FiltersV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	filters.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (filters *FiltersV1) DisableRetries() {
	filters.Service.DisableRetries()
}

// ListAllFilters : List all filters for a zone
// List all filters for a zone.
func (filters *FiltersV1) ListAllFilters(listAllFiltersOptions *ListAllFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.ListAllFiltersWithContext(context.Background(), listAllFiltersOptions)
}

// ListAllFiltersWithContext is an alternate form of the ListAllFilters method which supports a Context parameter
func (filters *FiltersV1) ListAllFiltersWithContext(ctx context.Context, listAllFiltersOptions *ListAllFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAllFiltersOptions, "listAllFiltersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAllFiltersOptions, "listAllFiltersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *listAllFiltersOptions.Crn,
		"zone_identifier": *listAllFiltersOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAllFiltersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "ListAllFilters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if listAllFiltersOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*listAllFiltersOptions.XAuthUserToken))
	}
	if listAllFiltersOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*listAllFiltersOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// CreateFilter : Create filters for a zone
// Create new filters for a given zone under a service instance.
func (filters *FiltersV1) CreateFilter(createFilterOptions *CreateFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.CreateFilterWithContext(context.Background(), createFilterOptions)
}

// CreateFilterWithContext is an alternate form of the CreateFilter method which supports a Context parameter
func (filters *FiltersV1) CreateFilterWithContext(ctx context.Context, createFilterOptions *CreateFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createFilterOptions, "createFilterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createFilterOptions, "createFilterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *createFilterOptions.Crn,
		"zone_identifier": *createFilterOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createFilterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "CreateFilter")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if createFilterOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*createFilterOptions.XAuthUserToken))
	}
	if createFilterOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*createFilterOptions.Accept))
	}

	if createFilterOptions.FilterInput != nil {
		_, err = builder.SetBodyContentJSON(createFilterOptions.FilterInput)
		if err != nil {
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// UpdateFilters : Update filters
// Update existing filters for a given zone under a given service instance.
func (filters *FiltersV1) UpdateFilters(updateFiltersOptions *UpdateFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.UpdateFiltersWithContext(context.Background(), updateFiltersOptions)
}

// UpdateFiltersWithContext is an alternate form of the UpdateFilters method which supports a Context parameter
func (filters *FiltersV1) UpdateFiltersWithContext(ctx context.Context, updateFiltersOptions *UpdateFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateFiltersOptions, "updateFiltersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateFiltersOptions, "updateFiltersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *updateFiltersOptions.Crn,
		"zone_identifier": *updateFiltersOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateFiltersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "UpdateFilters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if updateFiltersOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*updateFiltersOptions.XAuthUserToken))
	}
	if updateFiltersOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*updateFiltersOptions.Accept))
	}

	if updateFiltersOptions.FilterUpdateInput != nil {
		_, err = builder.SetBodyContentJSON(updateFiltersOptions.FilterUpdateInput)
		if err != nil {
			return
		}
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// DeleteFilters : Delete filters
// Delete filters by filter ids.
func (filters *FiltersV1) DeleteFilters(deleteFiltersOptions *DeleteFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.DeleteFiltersWithContext(context.Background(), deleteFiltersOptions)
}

// DeleteFiltersWithContext is an alternate form of the DeleteFilters method which supports a Context parameter
func (filters *FiltersV1) DeleteFiltersWithContext(ctx context.Context, deleteFiltersOptions *DeleteFiltersOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFiltersOptions, "deleteFiltersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFiltersOptions, "deleteFiltersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *deleteFiltersOptions.Crn,
		"zone_identifier": *deleteFiltersOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFiltersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "DeleteFilters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if deleteFiltersOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*deleteFiltersOptions.XAuthUserToken))
	}
	if deleteFiltersOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*deleteFiltersOptions.Accept))
	}

	builder.AddQuery("id", fmt.Sprint(*deleteFiltersOptions.ID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// GetFilter : Get filter details by id
// Get the details of a filter for a given zone under a given service instance.
func (filters *FiltersV1) GetFilter(getFilterOptions *GetFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.GetFilterWithContext(context.Background(), getFilterOptions)
}

// GetFilterWithContext is an alternate form of the GetFilter method which supports a Context parameter
func (filters *FiltersV1) GetFilterWithContext(ctx context.Context, getFilterOptions *GetFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getFilterOptions, "getFilterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getFilterOptions, "getFilterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *getFilterOptions.Crn,
		"zone_identifier": *getFilterOptions.ZoneIdentifier,
		"filter_identifier": *getFilterOptions.FilterIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters/{filter_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFilterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "GetFilter")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if getFilterOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*getFilterOptions.XAuthUserToken))
	}
	if getFilterOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getFilterOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// UpdateFilter : Update a filter
// Update an existing filter for a given zone under a given service instance.
func (filters *FiltersV1) UpdateFilter(updateFilterOptions *UpdateFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.UpdateFilterWithContext(context.Background(), updateFilterOptions)
}

// UpdateFilterWithContext is an alternate form of the UpdateFilter method which supports a Context parameter
func (filters *FiltersV1) UpdateFilterWithContext(ctx context.Context, updateFilterOptions *UpdateFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateFilterOptions, "updateFilterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateFilterOptions, "updateFilterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *updateFilterOptions.Crn,
		"zone_identifier": *updateFilterOptions.ZoneIdentifier,
		"filter_identifier": *updateFilterOptions.FilterIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters/{filter_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateFilterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "UpdateFilter")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	builder.AddHeader("Content-Type", "application/json")
	if updateFilterOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*updateFilterOptions.XAuthUserToken))
	}
	if updateFilterOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*updateFilterOptions.Accept))
	}

	body := make(map[string]interface{})
	if updateFilterOptions.ID != nil {
		body["id"] = updateFilterOptions.ID
	}
	if updateFilterOptions.Expression != nil {
		body["expression"] = updateFilterOptions.Expression
	}
	if updateFilterOptions.Description != nil {
		body["description"] = updateFilterOptions.Description
	}
	if updateFilterOptions.Paused != nil {
		body["paused"] = updateFilterOptions.Paused
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// DeleteFilter : Delete a filter
// Delete a filter given its id.
func (filters *FiltersV1) DeleteFilter(deleteFilterOptions *DeleteFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return filters.DeleteFilterWithContext(context.Background(), deleteFilterOptions)
}

// DeleteFilterWithContext is an alternate form of the DeleteFilter method which supports a Context parameter
func (filters *FiltersV1) DeleteFilterWithContext(ctx context.Context, deleteFilterOptions *DeleteFilterOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFilterOptions, "deleteFilterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFilterOptions, "deleteFilterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *deleteFilterOptions.Crn,
		"zone_identifier": *deleteFilterOptions.ZoneIdentifier,
		"filter_identifier": *deleteFilterOptions.FilterIdentifier,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = filters.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(filters.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/filters/{filter_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFilterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("filters", "V1", "DeleteFilter")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "*/*")
	if deleteFilterOptions.XAuthUserToken != nil {
		builder.AddHeader("X-Auth-User-Token", fmt.Sprint(*deleteFilterOptions.XAuthUserToken))
	}
	if deleteFilterOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*deleteFilterOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = filters.Service.Request(request, &result)

	return
}

// CreateFilterOptions : The CreateFilter options.
type CreateFilterOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required,ne="`

	// Zone identifier of the zone for which filters are created.
	ZoneIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Json objects which are used to create filters.
	FilterInput []FilterInput

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateFilterOptions : Instantiate CreateFilterOptions
func (*FiltersV1) NewCreateFilterOptions(xAuthUserToken string, crn string, zoneIdentifier string) *CreateFilterOptions {
	return &CreateFilterOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *CreateFilterOptions) SetXAuthUserToken(xAuthUserToken string) *CreateFilterOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *CreateFilterOptions) SetCrn(crn string) *CreateFilterOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *CreateFilterOptions) SetZoneIdentifier(zoneIdentifier string) *CreateFilterOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *CreateFilterOptions) SetAccept(accept string) *CreateFilterOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetFilterInput : Allow user to set FilterInput
func (options *CreateFilterOptions) SetFilterInput(filterInput []FilterInput) *CreateFilterOptions {
	options.FilterInput = filterInput
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateFilterOptions) SetHeaders(param map[string]string) *CreateFilterOptions {
	options.Headers = param
	return options
}

// DeleteFilterOptions : The DeleteFilter options.
type DeleteFilterOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full crn of the service instance.
	Crn *string `validate:"required,ne="`

	// Identifier of zone whose filter is to be deleted.
	ZoneIdentifier *string `validate:"required,ne="`

	// Identifier of the filter to be deleted.
	FilterIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteFilterOptions : Instantiate DeleteFilterOptions
func (*FiltersV1) NewDeleteFilterOptions(xAuthUserToken string, crn string, zoneIdentifier string, filterIdentifier string) *DeleteFilterOptions {
	return &DeleteFilterOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
		FilterIdentifier: core.StringPtr(filterIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *DeleteFilterOptions) SetXAuthUserToken(xAuthUserToken string) *DeleteFilterOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *DeleteFilterOptions) SetCrn(crn string) *DeleteFilterOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *DeleteFilterOptions) SetZoneIdentifier(zoneIdentifier string) *DeleteFilterOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetFilterIdentifier : Allow user to set FilterIdentifier
func (options *DeleteFilterOptions) SetFilterIdentifier(filterIdentifier string) *DeleteFilterOptions {
	options.FilterIdentifier = core.StringPtr(filterIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *DeleteFilterOptions) SetAccept(accept string) *DeleteFilterOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFilterOptions) SetHeaders(param map[string]string) *DeleteFilterOptions {
	options.Headers = param
	return options
}

// DeleteFiltersOptions : The DeleteFilters options.
type DeleteFiltersOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full crn of the service instance.
	Crn *string `validate:"required,ne="`

	// Identifier of zone whose filters are to be deleted.
	ZoneIdentifier *string `validate:"required,ne="`

	// ids of filters which will be deleted.
	ID *string `validate:"required"`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteFiltersOptions : Instantiate DeleteFiltersOptions
func (*FiltersV1) NewDeleteFiltersOptions(xAuthUserToken string, crn string, zoneIdentifier string, id string) *DeleteFiltersOptions {
	return &DeleteFiltersOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
		ID: core.StringPtr(id),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *DeleteFiltersOptions) SetXAuthUserToken(xAuthUserToken string) *DeleteFiltersOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *DeleteFiltersOptions) SetCrn(crn string) *DeleteFiltersOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *DeleteFiltersOptions) SetZoneIdentifier(zoneIdentifier string) *DeleteFiltersOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetID : Allow user to set ID
func (options *DeleteFiltersOptions) SetID(id string) *DeleteFiltersOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccept : Allow user to set Accept
func (options *DeleteFiltersOptions) SetAccept(accept string) *DeleteFiltersOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFiltersOptions) SetHeaders(param map[string]string) *DeleteFiltersOptions {
	options.Headers = param
	return options
}

// GetFilterOptions : The GetFilter options.
type GetFilterOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full crn of the service instance.
	Crn *string `validate:"required,ne="`

	// Zone identifier (zone id).
	ZoneIdentifier *string `validate:"required,ne="`

	// Identifier of filter for the given zone.
	FilterIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFilterOptions : Instantiate GetFilterOptions
func (*FiltersV1) NewGetFilterOptions(xAuthUserToken string, crn string, zoneIdentifier string, filterIdentifier string) *GetFilterOptions {
	return &GetFilterOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
		FilterIdentifier: core.StringPtr(filterIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *GetFilterOptions) SetXAuthUserToken(xAuthUserToken string) *GetFilterOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *GetFilterOptions) SetCrn(crn string) *GetFilterOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *GetFilterOptions) SetZoneIdentifier(zoneIdentifier string) *GetFilterOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetFilterIdentifier : Allow user to set FilterIdentifier
func (options *GetFilterOptions) SetFilterIdentifier(filterIdentifier string) *GetFilterOptions {
	options.FilterIdentifier = core.StringPtr(filterIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *GetFilterOptions) SetAccept(accept string) *GetFilterOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFilterOptions) SetHeaders(param map[string]string) *GetFilterOptions {
	options.Headers = param
	return options
}

// ListAllFiltersOptions : The ListAllFilters options.
type ListAllFiltersOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full url-encoded cloud resource name (CRN) of resource instance.
	Crn *string `validate:"required,ne="`

	// Zone identifier of the zone for which filters are listed.
	ZoneIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAllFiltersOptions : Instantiate ListAllFiltersOptions
func (*FiltersV1) NewListAllFiltersOptions(xAuthUserToken string, crn string, zoneIdentifier string) *ListAllFiltersOptions {
	return &ListAllFiltersOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *ListAllFiltersOptions) SetXAuthUserToken(xAuthUserToken string) *ListAllFiltersOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *ListAllFiltersOptions) SetCrn(crn string) *ListAllFiltersOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *ListAllFiltersOptions) SetZoneIdentifier(zoneIdentifier string) *ListAllFiltersOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *ListAllFiltersOptions) SetAccept(accept string) *ListAllFiltersOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllFiltersOptions) SetHeaders(param map[string]string) *ListAllFiltersOptions {
	options.Headers = param
	return options
}

// UpdateFilterOptions : The UpdateFilter options.
type UpdateFilterOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full crn of the service instance.
	Crn *string `validate:"required,ne="`

	// Zone identifier (zone id).
	ZoneIdentifier *string `validate:"required,ne="`

	// Identifier of filter.
	FilterIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	// Identifier of the filter.
	ID *string

	// A filter expression.
	Expression *string

	// To briefly describe the filter.
	Description *string

	// Indicates if the filter is active.
	Paused *bool

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateFilterOptions : Instantiate UpdateFilterOptions
func (*FiltersV1) NewUpdateFilterOptions(xAuthUserToken string, crn string, zoneIdentifier string, filterIdentifier string) *UpdateFilterOptions {
	return &UpdateFilterOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
		FilterIdentifier: core.StringPtr(filterIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *UpdateFilterOptions) SetXAuthUserToken(xAuthUserToken string) *UpdateFilterOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *UpdateFilterOptions) SetCrn(crn string) *UpdateFilterOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *UpdateFilterOptions) SetZoneIdentifier(zoneIdentifier string) *UpdateFilterOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetFilterIdentifier : Allow user to set FilterIdentifier
func (options *UpdateFilterOptions) SetFilterIdentifier(filterIdentifier string) *UpdateFilterOptions {
	options.FilterIdentifier = core.StringPtr(filterIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *UpdateFilterOptions) SetAccept(accept string) *UpdateFilterOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetID : Allow user to set ID
func (options *UpdateFilterOptions) SetID(id string) *UpdateFilterOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetExpression : Allow user to set Expression
func (options *UpdateFilterOptions) SetExpression(expression string) *UpdateFilterOptions {
	options.Expression = core.StringPtr(expression)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateFilterOptions) SetDescription(description string) *UpdateFilterOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetPaused : Allow user to set Paused
func (options *UpdateFilterOptions) SetPaused(paused bool) *UpdateFilterOptions {
	options.Paused = core.BoolPtr(paused)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateFilterOptions) SetHeaders(param map[string]string) *UpdateFilterOptions {
	options.Headers = param
	return options
}

// UpdateFiltersOptions : The UpdateFilters options.
type UpdateFiltersOptions struct {
	// IBM Cloud user IAM token.
	XAuthUserToken *string `validate:"required"`

	// Full crn of the service instance.
	Crn *string `validate:"required,ne="`

	// Zone identifier (zone id).
	ZoneIdentifier *string `validate:"required,ne="`

	// The type of the response: *_/_* or application/json.
	Accept *string

	FilterUpdateInput []FilterUpdateInput

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateFiltersOptions : Instantiate UpdateFiltersOptions
func (*FiltersV1) NewUpdateFiltersOptions(xAuthUserToken string, crn string, zoneIdentifier string) *UpdateFiltersOptions {
	return &UpdateFiltersOptions{
		XAuthUserToken: core.StringPtr(xAuthUserToken),
		Crn: core.StringPtr(crn),
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetXAuthUserToken : Allow user to set XAuthUserToken
func (options *UpdateFiltersOptions) SetXAuthUserToken(xAuthUserToken string) *UpdateFiltersOptions {
	options.XAuthUserToken = core.StringPtr(xAuthUserToken)
	return options
}

// SetCrn : Allow user to set Crn
func (options *UpdateFiltersOptions) SetCrn(crn string) *UpdateFiltersOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *UpdateFiltersOptions) SetZoneIdentifier(zoneIdentifier string) *UpdateFiltersOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetAccept : Allow user to set Accept
func (options *UpdateFiltersOptions) SetAccept(accept string) *UpdateFiltersOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetFilterUpdateInput : Allow user to set FilterUpdateInput
func (options *UpdateFiltersOptions) SetFilterUpdateInput(filterUpdateInput []FilterUpdateInput) *UpdateFiltersOptions {
	options.FilterUpdateInput = filterUpdateInput
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateFiltersOptions) SetHeaders(param map[string]string) *UpdateFiltersOptions {
	options.Headers = param
	return options
}

// FilterInput : Json objects which are used to create filters.
type FilterInput struct {
	// A filter expression.
	Expression *string `json:"expression" validate:"required"`

	// Indicates if the filter is active.
	Paused *bool `json:"paused,omitempty"`

	// To briefly describe the filter, omitted from object if empty.
	Description *string `json:"description,omitempty"`
}

// NewFilterInput : Instantiate FilterInput (Generic Model Constructor)
func (*FiltersV1) NewFilterInput(expression string) (model *FilterInput, err error) {
	model = &FilterInput{
		Expression: core.StringPtr(expression),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFilterInput unmarshals an instance of FilterInput from the specified map of raw messages.
func UnmarshalFilterInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FilterInput)
	err = core.UnmarshalPrimitive(m, "expression", &obj.Expression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "paused", &obj.Paused)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FilterUpdateInput : FilterUpdateInput struct
type FilterUpdateInput struct {
	// Identifier of the filter.
	ID *string `json:"id" validate:"required"`

	// A filter expression.
	Expression *string `json:"expression" validate:"required"`

	// To briefly describe the filter.
	Description *string `json:"description,omitempty"`

	// Indicates if the filter is active.
	Paused *bool `json:"paused,omitempty"`
}

// NewFilterUpdateInput : Instantiate FilterUpdateInput (Generic Model Constructor)
func (*FiltersV1) NewFilterUpdateInput(id string, expression string) (model *FilterUpdateInput, err error) {
	model = &FilterUpdateInput{
		ID: core.StringPtr(id),
		Expression: core.StringPtr(expression),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFilterUpdateInput unmarshals an instance of FilterUpdateInput from the specified map of raw messages.
func UnmarshalFilterUpdateInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FilterUpdateInput)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expression", &obj.Expression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "paused", &obj.Paused)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
