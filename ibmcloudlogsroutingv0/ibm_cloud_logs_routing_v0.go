/**
 * (C) Copyright IBM Corp. 2023.
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
 * IBM OpenAPI SDK Code Generator Version: 3.81.0-c73a091c-20231026-215706
 */

// Package ibmcloudlogsroutingv0 : Operations and models for the IbmCloudLogsRoutingV0 service
package ibmcloudlogsroutingv0

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/logs-router-go-sdk/common"
	"github.com/go-openapi/strfmt"
	"github.com/moul/http2curl"
)

// IbmCloudLogsRoutingV0 : IBM Cloud Logs Routing is an IBM cloud platform service to collect log-events of your VPC and
// deliver them to the configured log sink.
//
// API Version: 0.0.1
// See: http://cloud.ibm.com
type IbmCloudLogsRoutingV0 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://management.us-east.logs-router.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_logs_routing"

const ParameterizedServiceURL = "https://management.{region}.logs-router.cloud.ibm.com/v1"

var defaultUrlVariables = map[string]string{
	"region": "us-east",
}

// IbmCloudLogsRoutingV0Options : Service options
type IbmCloudLogsRoutingV0Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudLogsRoutingV0UsingExternalConfig : constructs an instance of IbmCloudLogsRoutingV0 with passed in options and external configuration.
func NewIbmCloudLogsRoutingV0UsingExternalConfig(options *IbmCloudLogsRoutingV0Options) (ibmCloudLogsRouting *IbmCloudLogsRoutingV0, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudLogsRouting, err = NewIbmCloudLogsRoutingV0(options)
	if err != nil {
		return
	}

	err = ibmCloudLogsRouting.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudLogsRouting.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudLogsRoutingV0 : constructs an instance of IbmCloudLogsRoutingV0 with passed in options.
func NewIbmCloudLogsRoutingV0(options *IbmCloudLogsRoutingV0Options) (service *IbmCloudLogsRoutingV0, err error) {
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

	service = &IbmCloudLogsRoutingV0{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "ibmCloudLogsRouting" suitable for processing requests.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) Clone() *IbmCloudLogsRoutingV0 {
	if core.IsNil(ibmCloudLogsRouting) {
		return nil
	}
	clone := *ibmCloudLogsRouting
	clone.Service = ibmCloudLogsRouting.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) SetServiceURL(url string) error {
	return ibmCloudLogsRouting.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetServiceURL() string {
	return ibmCloudLogsRouting.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) SetDefaultHeaders(headers http.Header) {
	ibmCloudLogsRouting.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) SetEnableGzipCompression(enableGzip bool) {
	ibmCloudLogsRouting.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetEnableGzipCompression() bool {
	return ibmCloudLogsRouting.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmCloudLogsRouting.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DisableRetries() {
	ibmCloudLogsRouting.Service.DisableRetries()
}

// ListTenants : List of tenants
// List of tenants defined in your account.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenants(listTenantsOptions *ListTenantsOptions) (result *TenantCollection, response *core.DetailedResponse, err error) {
	return ibmCloudLogsRouting.ListTenantsWithContext(context.Background(), listTenantsOptions)
}

// ListTenantsWithContext is an alternate form of the ListTenants method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenantsWithContext(ctx context.Context, listTenantsOptions *ListTenantsOptions) (result *TenantCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTenantsOptions, "listTenantsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTenantsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "ListTenants")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenantCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTenant : Create (onboard) a new tenant
// Create (onboard) a new tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTenant(createTenantOptions *CreateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	return ibmCloudLogsRouting.CreateTenantWithContext(context.Background(), createTenantOptions)
}

// CreateTenantWithContext is an alternate form of the CreateTenant method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTenantWithContext(ctx context.Context, createTenantOptions *CreateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTenantOptions, "createTenantOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTenantOptions, "createTenantOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTenantOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "CreateTenant")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createTenantOptions.TargetType != nil {
		body["target_type"] = createTenantOptions.TargetType
	}
	if createTenantOptions.TargetHost != nil {
		body["target_host"] = createTenantOptions.TargetHost
	}
	if createTenantOptions.TargetPort != nil {
		body["target_port"] = createTenantOptions.TargetPort
	}
	if createTenantOptions.AccessCredential != nil {
		body["access_credential"] = createTenantOptions.AccessCredential
	}
	if createTenantOptions.TargetInstanceCrn != nil {
		body["target_instance_crn"] = createTenantOptions.TargetInstanceCrn
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTenantDetail : Details of a tenant
// List the details of the given tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantDetail(getTenantDetailOptions *GetTenantDetailOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	return ibmCloudLogsRouting.GetTenantDetailWithContext(context.Background(), getTenantDetailOptions)
}

// GetTenantDetailWithContext is an alternate form of the GetTenantDetail method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantDetailWithContext(ctx context.Context, getTenantDetailOptions *GetTenantDetailOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTenantDetailOptions, "getTenantDetailOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTenantDetailOptions, "getTenantDetailOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*getTenantDetailOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTenantDetailOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "GetTenantDetail")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	command, _ := http2curl.GetCurlCommand(request)
	fmt.Println("HTTP command is:")
	fmt.Println(command)

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTenant : Delete (offboard) a tenant
// Delete (offboard) a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTenant(deleteTenantOptions *DeleteTenantOptions) (result *TenantDelete, response *core.DetailedResponse, err error) {
	return ibmCloudLogsRouting.DeleteTenantWithContext(context.Background(), deleteTenantOptions)
}

// DeleteTenantWithContext is an alternate form of the DeleteTenant method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTenantWithContext(ctx context.Context, deleteTenantOptions *DeleteTenantOptions) (result *TenantDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTenantOptions, "deleteTenantOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTenantOptions, "deleteTenantOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*deleteTenantOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTenantOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "DeleteTenant")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenantDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateTarget : Update the target of a tenant
// Update the target of a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTarget(updateTargetOptions *UpdateTargetOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	return ibmCloudLogsRouting.UpdateTargetWithContext(context.Background(), updateTargetOptions)
}

// UpdateTargetWithContext is an alternate form of the UpdateTarget method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTargetWithContext(ctx context.Context, updateTargetOptions *UpdateTargetOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTargetOptions, "updateTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTargetOptions, "updateTargetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*updateTargetOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "UpdateTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateTargetOptions.TenantPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	command, _ := http2curl.GetCurlCommand(request)
	fmt.Println("HTTP command is:")
	fmt.Println(command)

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTenantOptions : The CreateTenant options.
type CreateTenantOptions struct {
	// Type of log-sink to connect to. Supported: [logdna].
	TargetType *string `json:"target_type" validate:"required"`

	// Full qualified host name of log-sink.
	TargetHost *string `json:"target_host" validate:"required"`

	// Port number at target_host of log-sink.
	TargetPort *int64 `json:"target_port" validate:"required"`

	// Secret to connect to the log-sink.
	AccessCredential *string `json:"access_credential" validate:"required"`

	// Cloud resource name of the log-sink target instance.
	TargetInstanceCrn *string `json:"target_instance_crn" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTenantOptions : Instantiate CreateTenantOptions
func (*IbmCloudLogsRoutingV0) NewCreateTenantOptions(targetType string, targetHost string, targetPort int64, accessCredential string, targetInstanceCrn string) *CreateTenantOptions {
	return &CreateTenantOptions{
		TargetType:        core.StringPtr(targetType),
		TargetHost:        core.StringPtr(targetHost),
		TargetPort:        core.Int64Ptr(targetPort),
		AccessCredential:  core.StringPtr(accessCredential),
		TargetInstanceCrn: core.StringPtr(targetInstanceCrn),
	}
}

// SetTargetType : Allow user to set TargetType
func (_options *CreateTenantOptions) SetTargetType(targetType string) *CreateTenantOptions {
	_options.TargetType = core.StringPtr(targetType)
	return _options
}

// SetTargetHost : Allow user to set TargetHost
func (_options *CreateTenantOptions) SetTargetHost(targetHost string) *CreateTenantOptions {
	_options.TargetHost = core.StringPtr(targetHost)
	return _options
}

// SetTargetPort : Allow user to set TargetPort
func (_options *CreateTenantOptions) SetTargetPort(targetPort int64) *CreateTenantOptions {
	_options.TargetPort = core.Int64Ptr(targetPort)
	return _options
}

// SetAccessCredential : Allow user to set AccessCredential
func (_options *CreateTenantOptions) SetAccessCredential(accessCredential string) *CreateTenantOptions {
	_options.AccessCredential = core.StringPtr(accessCredential)
	return _options
}

// SetTargetInstanceCrn : Allow user to set TargetInstanceCrn
func (_options *CreateTenantOptions) SetTargetInstanceCrn(targetInstanceCrn string) *CreateTenantOptions {
	_options.TargetInstanceCrn = core.StringPtr(targetInstanceCrn)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTenantOptions) SetHeaders(param map[string]string) *CreateTenantOptions {
	options.Headers = param
	return options
}

// DeleteTenantOptions : The DeleteTenant options.
type DeleteTenantOptions struct {
	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTenantOptions : Instantiate DeleteTenantOptions
func (*IbmCloudLogsRoutingV0) NewDeleteTenantOptions(tenantID *strfmt.UUID) *DeleteTenantOptions {
	return &DeleteTenantOptions{
		TenantID: tenantID,
	}
}

// SetTenantID : Allow user to set TenantID
func (_options *DeleteTenantOptions) SetTenantID(tenantID *strfmt.UUID) *DeleteTenantOptions {
	_options.TenantID = tenantID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTenantOptions) SetHeaders(param map[string]string) *DeleteTenantOptions {
	options.Headers = param
	return options
}

// GetTenantDetailOptions : The GetTenantDetail options.
type GetTenantDetailOptions struct {
	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTenantDetailOptions : Instantiate GetTenantDetailOptions
func (*IbmCloudLogsRoutingV0) NewGetTenantDetailOptions(tenantID *strfmt.UUID) *GetTenantDetailOptions {
	return &GetTenantDetailOptions{
		TenantID: tenantID,
	}
}

// SetTenantID : Allow user to set TenantID
func (_options *GetTenantDetailOptions) SetTenantID(tenantID *strfmt.UUID) *GetTenantDetailOptions {
	_options.TenantID = tenantID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTenantDetailOptions) SetHeaders(param map[string]string) *GetTenantDetailOptions {
	options.Headers = param
	return options
}

// ListTenantsOptions : The ListTenants options.
type ListTenantsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTenantsOptions : Instantiate ListTenantsOptions
func (*IbmCloudLogsRoutingV0) NewListTenantsOptions() *ListTenantsOptions {
	return &ListTenantsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListTenantsOptions) SetHeaders(param map[string]string) *ListTenantsOptions {
	options.Headers = param
	return options
}

// Tenant : Response body from a successful list tenant details operation.
type Tenant struct {
	// Unique ID of the created instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The account ID the tenant belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// Type of log-sink.
	TargetType *string `json:"target_type" validate:"required"`

	// Host name of log-sink.
	TargetHost *string `json:"target_host" validate:"required"`

	// Network port of log sink.
	TargetPort *int64 `json:"target_port" validate:"required"`

	// Cloud resource name of the log-sink target instance.
	TargetInstanceCrn *string `json:"target_instance_crn" validate:"required"`

	// Time stamp the tenant was originally created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// time stamp the tenant was last updated.
	UpdatedAt *string `json:"updated_at" validate:"required"`
}

// UnmarshalTenant unmarshals an instance of Tenant from the specified map of raw messages.
func UnmarshalTenant(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tenant)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_host", &obj.TargetHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_port", &obj.TargetPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_instance_crn", &obj.TargetInstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TenantCollection : Response body from a successful list tenants operation.
type TenantCollection struct {
	// List of tenants in the account.
	Tenants []Tenant `json:"tenants,omitempty"`
}

// UnmarshalTenantCollection unmarshals an instance of TenantCollection from the specified map of raw messages.
func UnmarshalTenantCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TenantCollection)
	err = core.UnmarshalModel(m, "tenants", &obj.Tenants, UnmarshalTenant)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TenantDelete : Response body from successful delete tenant operation.
type TenantDelete struct {
	// HTTP status code.
	Status *int64 `json:"status" validate:"required"`

	// Status message.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalTenantDelete unmarshals an instance of TenantDelete from the specified map of raw messages.
func UnmarshalTenantDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TenantDelete)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TenantPatch : The request body used when updating the target of a tenant. At least one other value is required.
type TenantPatch struct {
	// Fully-qualified host name of log-sink.
	TargetHost *string `json:"target_host,omitempty"`

	// Port number at target_host of log-sink.
	TargetPort *int64 `json:"target_port,omitempty"`

	// Secret to connect to the log-sink.
	AccessCredential *string `json:"access_credential,omitempty"`

	// Cloud resource name of the log-sink target instance.
	TargetInstanceCrn *string `json:"target_instance_crn,omitempty"`
}

// UnmarshalTenantPatch unmarshals an instance of TenantPatch from the specified map of raw messages.
func UnmarshalTenantPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TenantPatch)
	err = core.UnmarshalPrimitive(m, "target_host", &obj.TargetHost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_port", &obj.TargetPort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access_credential", &obj.AccessCredential)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_instance_crn", &obj.TargetInstanceCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the TenantPatch
func (tenantPatch *TenantPatch) AsPatch() (_patch map[string]interface{}, err error) {
	var jsonData []byte
	jsonData, err = json.Marshal(tenantPatch)
	if err == nil {
		err = json.Unmarshal(jsonData, &_patch)
	}
	return
}

// UpdateTargetOptions : The UpdateTarget options.
type UpdateTargetOptions struct {
	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Updates the target setup. Only the listed fields can be updated and only the fields that need to be changed are
	// required in the body. At least one field must be specified for the update.
	TenantPatch map[string]interface{} `json:"Tenant_patch" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTargetOptions : Instantiate UpdateTargetOptions
func (*IbmCloudLogsRoutingV0) NewUpdateTargetOptions(tenantID *strfmt.UUID, tenantPatch map[string]interface{}) *UpdateTargetOptions {
	return &UpdateTargetOptions{
		TenantID:    tenantID,
		TenantPatch: tenantPatch,
	}
}

// SetTenantID : Allow user to set TenantID
func (_options *UpdateTargetOptions) SetTenantID(tenantID *strfmt.UUID) *UpdateTargetOptions {
	_options.TenantID = tenantID
	return _options
}

// SetTenantPatch : Allow user to set TenantPatch
func (_options *UpdateTargetOptions) SetTenantPatch(tenantPatch map[string]interface{}) *UpdateTargetOptions {
	_options.TenantPatch = tenantPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTargetOptions) SetHeaders(param map[string]string) *UpdateTargetOptions {
	options.Headers = param
	return options
}
