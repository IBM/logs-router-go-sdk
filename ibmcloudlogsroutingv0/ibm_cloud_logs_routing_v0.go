/**
 * (C) Copyright IBM Corp. 2025.
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
 * IBM OpenAPI SDK Code Generator Version: 3.102.0-615ec964-20250307-203034
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
)

// IbmCloudLogsRoutingV0 : IBM Cloud Logs Routing is an IBM cloud platform service to collect log-events of your VPC and
// deliver them to the configured log-sink.
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
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	ibmCloudLogsRouting, err = NewIbmCloudLogsRoutingV0(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = ibmCloudLogsRouting.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = ibmCloudLogsRouting.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
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
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
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
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
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
	err := ibmCloudLogsRouting.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
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

// ListTenants : List tenants
// List tenants.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenants(listTenantsOptions *ListTenantsOptions) (result *TenantCollection, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.ListTenantsWithContext(context.Background(), listTenantsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListTenantsWithContext is an alternate form of the ListTenants method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenantsWithContext(ctx context.Context, listTenantsOptions *ListTenantsOptions) (result *TenantCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTenantsOptions, "listTenantsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listTenantsOptions, "listTenantsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
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
	if listTenantsOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*listTenantsOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	if listTenantsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listTenantsOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_tenants", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenantCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateTenant : Create (onboard) a tenant
// Create (onboard) a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTenant(createTenantOptions *CreateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.CreateTenantWithContext(context.Background(), createTenantOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateTenantWithContext is an alternate form of the CreateTenant method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTenantWithContext(ctx context.Context, createTenantOptions *CreateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTenantOptions, "createTenantOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createTenantOptions, "createTenantOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
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
	if createTenantOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*createTenantOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	body := make(map[string]interface{})
	if createTenantOptions.Name != nil {
		body["name"] = createTenantOptions.Name
	}
	if createTenantOptions.Targets != nil {
		body["targets"] = createTenantOptions.Targets
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_tenant", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTenantDetail : Retrieve a tenant
// Retrieve a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantDetail(getTenantDetailOptions *GetTenantDetailOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.GetTenantDetailWithContext(context.Background(), getTenantDetailOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTenantDetailWithContext is an alternate form of the GetTenantDetail method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantDetailWithContext(ctx context.Context, getTenantDetailOptions *GetTenantDetailOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTenantDetailOptions, "getTenantDetailOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getTenantDetailOptions, "getTenantDetailOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
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
	if getTenantDetailOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*getTenantDetailOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_tenant_detail", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteTenant : Delete (offboard) a tenant
// Delete (offboard) a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTenant(deleteTenantOptions *DeleteTenantOptions) (response *core.DetailedResponse, err error) {
	response, err = ibmCloudLogsRouting.DeleteTenantWithContext(context.Background(), deleteTenantOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteTenantWithContext is an alternate form of the DeleteTenant method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTenantWithContext(ctx context.Context, deleteTenantOptions *DeleteTenantOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTenantOptions, "deleteTenantOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteTenantOptions, "deleteTenantOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteTenantOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "DeleteTenant")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTenantOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*deleteTenantOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = ibmCloudLogsRouting.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_tenant", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateTenant : Update a tenant
// Update a tenant. To update a target please refer to the <b>targets</b> section.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTenant(updateTenantOptions *UpdateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.UpdateTenantWithContext(context.Background(), updateTenantOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateTenantWithContext is an alternate form of the UpdateTenant method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTenantWithContext(ctx context.Context, updateTenantOptions *UpdateTenantOptions) (result *Tenant, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTenantOptions, "updateTenantOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateTenantOptions, "updateTenantOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*updateTenantOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateTenantOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "UpdateTenant")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateTenantOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*updateTenantOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}
	if updateTenantOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateTenantOptions.IfMatch))
	}

	_, err = builder.SetBodyContentJSON(updateTenantOptions.TenantPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_tenant", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTenant)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListTenantTargets : List targets for a tenant
// List targets for a tenant.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenantTargets(listTenantTargetsOptions *ListTenantTargetsOptions) (result *TargetCollection, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.ListTenantTargetsWithContext(context.Background(), listTenantTargetsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListTenantTargetsWithContext is an alternate form of the ListTenantTargets method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) ListTenantTargetsWithContext(ctx context.Context, listTenantTargetsOptions *ListTenantTargetsOptions) (result *TargetCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTenantTargetsOptions, "listTenantTargetsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listTenantTargetsOptions, "listTenantTargetsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*listTenantTargetsOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}/targets`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listTenantTargetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "ListTenantTargets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTenantTargetsOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*listTenantTargetsOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	if listTenantTargetsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listTenantTargetsOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_tenant_targets", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTargetCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateTarget : Create a target
// Create a new target for a tenant.<br><b>Note:</b> The tenant must not already have a target of the specified type.
// The supported target type is <b>logs</b> (IBM Cloud Logs).
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTarget(createTargetOptions *CreateTargetOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.CreateTargetWithContext(context.Background(), createTargetOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateTargetWithContext is an alternate form of the CreateTarget method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) CreateTargetWithContext(ctx context.Context, createTargetOptions *CreateTargetOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTargetOptions, "createTargetOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createTargetOptions, "createTargetOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*createTargetOptions.TenantID),
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}/targets`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "CreateTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createTargetOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*createTargetOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	_, err = builder.SetBodyContentJSON(createTargetOptions.TargetPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_target", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTenantTargetDetails : Retrieve a target
// Retrieve a target.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantTargetDetails(getTenantTargetDetailsOptions *GetTenantTargetDetailsOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.GetTenantTargetDetailsWithContext(context.Background(), getTenantTargetDetailsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTenantTargetDetailsWithContext is an alternate form of the GetTenantTargetDetails method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) GetTenantTargetDetailsWithContext(ctx context.Context, getTenantTargetDetailsOptions *GetTenantTargetDetailsOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTenantTargetDetailsOptions, "getTenantTargetDetailsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getTenantTargetDetailsOptions, "getTenantTargetDetailsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*getTenantTargetDetailsOptions.TenantID),
		"target_id": fmt.Sprint(*getTenantTargetDetailsOptions.TargetID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}/targets/{target_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getTenantTargetDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "GetTenantTargetDetails")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTenantTargetDetailsOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*getTenantTargetDetailsOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_tenant_target_details", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateTarget : Update a target
// Update a target.<br><b>Note:</b>A change of the target type is only supported for tenants with a single target.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTarget(updateTargetOptions *UpdateTargetOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	result, response, err = ibmCloudLogsRouting.UpdateTargetWithContext(context.Background(), updateTargetOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateTargetWithContext is an alternate form of the UpdateTarget method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) UpdateTargetWithContext(ctx context.Context, updateTargetOptions *UpdateTargetOptions) (result TargetIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTargetOptions, "updateTargetOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateTargetOptions, "updateTargetOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*updateTargetOptions.TenantID),
		"target_id": fmt.Sprint(*updateTargetOptions.TargetID),
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}/targets/{target_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
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
	builder.AddHeader("Content-Type", "application/merge-patch+json")
	if updateTargetOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*updateTargetOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}
	if updateTargetOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateTargetOptions.IfMatch))
	}

	_, err = builder.SetBodyContentJSON(updateTargetOptions.TargetPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudLogsRouting.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_target", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteTarget : Delete a target
// Delete a target.<br><b>Note:</b> The last target is not allowed to be deleted.
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTarget(deleteTargetOptions *DeleteTargetOptions) (response *core.DetailedResponse, err error) {
	response, err = ibmCloudLogsRouting.DeleteTargetWithContext(context.Background(), deleteTargetOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteTargetWithContext is an alternate form of the DeleteTarget method which supports a Context parameter
func (ibmCloudLogsRouting *IbmCloudLogsRoutingV0) DeleteTargetWithContext(ctx context.Context, deleteTargetOptions *DeleteTargetOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTargetOptions, "deleteTargetOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteTargetOptions, "deleteTargetOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"tenant_id": fmt.Sprint(*deleteTargetOptions.TenantID),
		"target_id": fmt.Sprint(*deleteTargetOptions.TargetID),
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudLogsRouting.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudLogsRouting.Service.Options.URL, `/tenants/{tenant_id}/targets/{target_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_logs_routing", "V0", "DeleteTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTargetOptions.IBMAPIVersion != nil {
		builder.AddHeader("IBM-API-Version", fmt.Sprint(*deleteTargetOptions.IBMAPIVersion))
	} else {
		builder.AddHeader("IBM-API-Version", fmt.Sprint("2024-06-15"))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = ibmCloudLogsRouting.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_target", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "0.0.1")
}

// CreateTargetOptions : The CreateTarget options.
type CreateTargetOptions struct {
	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// log-sink endpoint and authentication.
	TargetPrototype TargetPrototypeIntf `json:"TargetPrototype" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateTargetOptions : Instantiate CreateTargetOptions
func (*IbmCloudLogsRoutingV0) NewCreateTargetOptions(tenantID *strfmt.UUID, iBMAPIVersion string, targetPrototype TargetPrototypeIntf) *CreateTargetOptions {
	return &CreateTargetOptions{
		TenantID:        tenantID,
		IBMAPIVersion:   core.StringPtr(iBMAPIVersion),
		TargetPrototype: targetPrototype,
	}
}

// SetTenantID : Allow user to set TenantID
func (_options *CreateTargetOptions) SetTenantID(tenantID *strfmt.UUID) *CreateTargetOptions {
	_options.TenantID = tenantID
	return _options
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *CreateTargetOptions) SetIBMAPIVersion(iBMAPIVersion string) *CreateTargetOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTargetPrototype : Allow user to set TargetPrototype
func (_options *CreateTargetOptions) SetTargetPrototype(targetPrototype TargetPrototypeIntf) *CreateTargetOptions {
	_options.TargetPrototype = targetPrototype
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTargetOptions) SetHeaders(param map[string]string) *CreateTargetOptions {
	options.Headers = param
	return options
}

// CreateTenantOptions : The CreateTenant options.
type CreateTenantOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The name for this tenant. The name is regionally unique across all tenants in the account.
	Name *string `json:"name" validate:"required"`

	// List of targets.<br>Two targets must not be of the same type. The supported target type is <b>logs</b> (IBM Cloud
	// Logs).
	Targets []TargetPrototypeIntf `json:"targets" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateTenantOptions : Instantiate CreateTenantOptions
func (*IbmCloudLogsRoutingV0) NewCreateTenantOptions(iBMAPIVersion string, name string, targets []TargetPrototypeIntf) *CreateTenantOptions {
	return &CreateTenantOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		Name:          core.StringPtr(name),
		Targets:       targets,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *CreateTenantOptions) SetIBMAPIVersion(iBMAPIVersion string) *CreateTenantOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateTenantOptions) SetName(name string) *CreateTenantOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTargets : Allow user to set Targets
func (_options *CreateTenantOptions) SetTargets(targets []TargetPrototypeIntf) *CreateTenantOptions {
	_options.Targets = targets
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTenantOptions) SetHeaders(param map[string]string) *CreateTenantOptions {
	options.Headers = param
	return options
}

// DeleteTargetOptions : The DeleteTarget options.
type DeleteTargetOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// The instance ID of the tenant target.
	TargetID *strfmt.UUID `json:"target_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteTargetOptions : Instantiate DeleteTargetOptions
func (*IbmCloudLogsRoutingV0) NewDeleteTargetOptions(iBMAPIVersion string, tenantID *strfmt.UUID, targetID *strfmt.UUID) *DeleteTargetOptions {
	return &DeleteTargetOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
		TargetID:      targetID,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *DeleteTargetOptions) SetIBMAPIVersion(iBMAPIVersion string) *DeleteTargetOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTenantID : Allow user to set TenantID
func (_options *DeleteTargetOptions) SetTenantID(tenantID *strfmt.UUID) *DeleteTargetOptions {
	_options.TenantID = tenantID
	return _options
}

// SetTargetID : Allow user to set TargetID
func (_options *DeleteTargetOptions) SetTargetID(targetID *strfmt.UUID) *DeleteTargetOptions {
	_options.TargetID = targetID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTargetOptions) SetHeaders(param map[string]string) *DeleteTargetOptions {
	options.Headers = param
	return options
}

// DeleteTenantOptions : The DeleteTenant options.
type DeleteTenantOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteTenantOptions : Instantiate DeleteTenantOptions
func (*IbmCloudLogsRoutingV0) NewDeleteTenantOptions(iBMAPIVersion string, tenantID *strfmt.UUID) *DeleteTenantOptions {
	return &DeleteTenantOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *DeleteTenantOptions) SetIBMAPIVersion(iBMAPIVersion string) *DeleteTenantOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
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
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTenantDetailOptions : Instantiate GetTenantDetailOptions
func (*IbmCloudLogsRoutingV0) NewGetTenantDetailOptions(iBMAPIVersion string, tenantID *strfmt.UUID) *GetTenantDetailOptions {
	return &GetTenantDetailOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *GetTenantDetailOptions) SetIBMAPIVersion(iBMAPIVersion string) *GetTenantDetailOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
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

// GetTenantTargetDetailsOptions : The GetTenantTargetDetails options.
type GetTenantTargetDetailsOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// The instance ID of the tenant target.
	TargetID *strfmt.UUID `json:"target_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTenantTargetDetailsOptions : Instantiate GetTenantTargetDetailsOptions
func (*IbmCloudLogsRoutingV0) NewGetTenantTargetDetailsOptions(iBMAPIVersion string, tenantID *strfmt.UUID, targetID *strfmt.UUID) *GetTenantTargetDetailsOptions {
	return &GetTenantTargetDetailsOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
		TargetID:      targetID,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *GetTenantTargetDetailsOptions) SetIBMAPIVersion(iBMAPIVersion string) *GetTenantTargetDetailsOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTenantID : Allow user to set TenantID
func (_options *GetTenantTargetDetailsOptions) SetTenantID(tenantID *strfmt.UUID) *GetTenantTargetDetailsOptions {
	_options.TenantID = tenantID
	return _options
}

// SetTargetID : Allow user to set TargetID
func (_options *GetTenantTargetDetailsOptions) SetTargetID(targetID *strfmt.UUID) *GetTenantTargetDetailsOptions {
	_options.TargetID = targetID
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTenantTargetDetailsOptions) SetHeaders(param map[string]string) *GetTenantTargetDetailsOptions {
	options.Headers = param
	return options
}

// ListTenantTargetsOptions : The ListTenantTargets options.
type ListTenantTargetsOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// Optional: Name of the tenant target.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListTenantTargetsOptions : Instantiate ListTenantTargetsOptions
func (*IbmCloudLogsRoutingV0) NewListTenantTargetsOptions(iBMAPIVersion string, tenantID *strfmt.UUID) *ListTenantTargetsOptions {
	return &ListTenantTargetsOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *ListTenantTargetsOptions) SetIBMAPIVersion(iBMAPIVersion string) *ListTenantTargetsOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTenantID : Allow user to set TenantID
func (_options *ListTenantTargetsOptions) SetTenantID(tenantID *strfmt.UUID) *ListTenantTargetsOptions {
	_options.TenantID = tenantID
	return _options
}

// SetName : Allow user to set Name
func (_options *ListTenantTargetsOptions) SetName(name string) *ListTenantTargetsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTenantTargetsOptions) SetHeaders(param map[string]string) *ListTenantTargetsOptions {
	options.Headers = param
	return options
}

// ListTenantsOptions : The ListTenants options.
type ListTenantsOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// Optional: The name of a tenant.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListTenantsOptions : Instantiate ListTenantsOptions
func (*IbmCloudLogsRoutingV0) NewListTenantsOptions(iBMAPIVersion string) *ListTenantsOptions {
	return &ListTenantsOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *ListTenantsOptions) SetIBMAPIVersion(iBMAPIVersion string) *ListTenantsOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetName : Allow user to set Name
func (_options *ListTenantsOptions) SetName(name string) *ListTenantsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTenantsOptions) SetHeaders(param map[string]string) *ListTenantsOptions {
	options.Headers = param
	return options
}

// Target : Parameter set describing a particular type of target.
// Models which "extend" this model:
// - TargetTypeLogs
type Target struct {
	// Unique ID of the target.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The name for this tenant target. The name must be unique across all targets for this tenant.
	Name *string `json:"name" validate:"required"`

	// Resource version identifier.
	Etag *string `json:"etag" validate:"required"`

	// Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
	Type *string `json:"type" validate:"required"`

	// Timestamp the target was originally created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// Timestamp the target was last updated.
	UpdatedAt *string `json:"updated_at" validate:"required"`

	// Cloud resource name of the log-sink target instance. Must be a valid CRN.
	LogSinkCrn *string `json:"log_sink_crn,omitempty"`

	// List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Parameters *TargetParametersTypeLogs `json:"parameters,omitempty"`
}

// Constants associated with the Target.Type property.
// Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
const (
	Target_Type_Logs = "logs"
)

func (*Target) isaTarget() bool {
	return true
}

type TargetIntf interface {
	isaTarget() bool
}

// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "etag", &obj.Etag)
	if err != nil {
		err = core.SDKErrorf(err, "", "etag-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "log_sink_crn", &obj.LogSinkCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "log_sink_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalTargetParametersTypeLogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "parameters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetCollection : Response body from a successful list targets operation.
type TargetCollection struct {
	// List of targets of a tenant.
	Targets []TargetIntf `json:"targets,omitempty"`
}

// UnmarshalTargetCollection unmarshals an instance of TargetCollection from the specified map of raw messages.
func UnmarshalTargetCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetCollection)
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTarget)
	if err != nil {
		err = core.SDKErrorf(err, "", "targets-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetParametersPatch : List of properties to update a log-sink of type <b>logs</b> (IBM Cloud Logs).
type TargetParametersPatch struct {
	// Host name of the log-sink.
	Host *string `json:"host,omitempty"`

	// Network port of the log-sink.
	Port *int64 `json:"port,omitempty"`
}

// UnmarshalTargetParametersPatch unmarshals an instance of TargetParametersPatch from the specified map of raw messages.
func UnmarshalTargetParametersPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetParametersPatch)
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// asPatch returns a generic map representation of the TargetParametersPatch
func (targetParametersPatch *TargetParametersPatch) asPatch() (_patch map[string]interface{}) {
	_patch = map[string]interface{}{}
	if !core.IsNil(targetParametersPatch.Host) {
		_patch["host"] = targetParametersPatch.Host
	}
	if !core.IsNil(targetParametersPatch.Port) {
		_patch["port"] = targetParametersPatch.Port
	}

	return
}

// TargetParametersTypeLogs : List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).
type TargetParametersTypeLogs struct {
	// Host name of the log-sink.
	Host *string `json:"host" validate:"required"`

	// Network port of the log-sink.
	Port *int64 `json:"port" validate:"required"`
}

// UnmarshalTargetParametersTypeLogs unmarshals an instance of TargetParametersTypeLogs from the specified map of raw messages.
func UnmarshalTargetParametersTypeLogs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetParametersTypeLogs)
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetParametersTypeLogsPrototype : List of properties to create or update a log-sink of type <b>logs</b> (IBM Cloud Logs).
type TargetParametersTypeLogsPrototype struct {
	// Host name of the log-sink.
	Host *string `json:"host" validate:"required"`

	// Network port of the log-sink.
	Port *int64 `json:"port" validate:"required"`
}

// NewTargetParametersTypeLogsPrototype : Instantiate TargetParametersTypeLogsPrototype (Generic Model Constructor)
func (*IbmCloudLogsRoutingV0) NewTargetParametersTypeLogsPrototype(host string, port int64) (_model *TargetParametersTypeLogsPrototype, err error) {
	_model = &TargetParametersTypeLogsPrototype{
		Host: core.StringPtr(host),
		Port: core.Int64Ptr(port),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalTargetParametersTypeLogsPrototype unmarshals an instance of TargetParametersTypeLogsPrototype from the specified map of raw messages.
func UnmarshalTargetParametersTypeLogsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetParametersTypeLogsPrototype)
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetPatch : Parameter set describing a particular type of target.
type TargetPatch struct {
	// The name for this tenant target. The name must be unique across all targets for this tenant.
	Name *string `json:"name,omitempty"`

	// Cloud resource name of the log-sink target instance. Must be a valid CRN.
	LogSinkCrn *string `json:"log_sink_crn,omitempty"`

	// List of properties to update a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Parameters *TargetParametersPatch `json:"parameters,omitempty"`
}

// UnmarshalTargetPatch unmarshals an instance of TargetPatch from the specified map of raw messages.
func UnmarshalTargetPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "log_sink_crn", &obj.LogSinkCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "log_sink_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalTargetParametersPatch)
	if err != nil {
		err = core.SDKErrorf(err, "", "parameters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the TargetPatch
func (targetPatch *TargetPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(targetPatch.Name) {
		_patch["name"] = targetPatch.Name
	}
	if !core.IsNil(targetPatch.LogSinkCrn) {
		_patch["log_sink_crn"] = targetPatch.LogSinkCrn
	}
	if !core.IsNil(targetPatch.Parameters) {
		_patch["parameters"] = targetPatch.Parameters.asPatch()
	}

	return
}

// TargetPrototype : Parameter set describing a particular type of target.
// Models which "extend" this model:
// - TargetPrototypeTargetTypeLogsPrototype
type TargetPrototype struct {
	// The name for this tenant target. The name must be unique across all targets for this tenant.
	Name *string `json:"name" validate:"required"`

	// Cloud resource name of the log-sink target instance. Must be a valid CRN.
	LogSinkCrn *string `json:"log_sink_crn,omitempty"`

	// List of properties to create or update a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Parameters *TargetParametersTypeLogsPrototype `json:"parameters,omitempty"`
}

func (*TargetPrototype) isaTargetPrototype() bool {
	return true
}

type TargetPrototypeIntf interface {
	isaTargetPrototype() bool
}

// UnmarshalTargetPrototype unmarshals an instance of TargetPrototype from the specified map of raw messages.
func UnmarshalTargetPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetPrototype)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "log_sink_crn", &obj.LogSinkCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "log_sink_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalTargetParametersTypeLogsPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "parameters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tenant : Full tenant description.
type Tenant struct {
	// Unique ID of the tenant.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// Timestamp the tenant was originally created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// Timestamp the tenant was last updated.
	UpdatedAt *string `json:"updated_at" validate:"required"`

	// Cloud resource name of the tenant. Must be a valid CRN.
	Crn *string `json:"crn" validate:"required"`

	// The name for this tenant. The name is regionally unique across all tenants in the account.
	Name *string `json:"name" validate:"required"`

	// Resource version identifier.
	Etag *string `json:"etag" validate:"required"`

	// List of targets.
	Targets []TargetIntf `json:"targets" validate:"required"`

	// The status of the write attempt to the target with the provided endpoint parameters.
	WriteStatus *WriteStatus `json:"write_status" validate:"required"`
}

// UnmarshalTenant unmarshals an instance of Tenant from the specified map of raw messages.
func UnmarshalTenant(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tenant)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "etag", &obj.Etag)
	if err != nil {
		err = core.SDKErrorf(err, "", "etag-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTarget)
	if err != nil {
		err = core.SDKErrorf(err, "", "targets-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "write_status", &obj.WriteStatus, UnmarshalWriteStatus)
	if err != nil {
		err = core.SDKErrorf(err, "", "write_status-error", common.GetComponentInfo())
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
		err = core.SDKErrorf(err, "", "tenants-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TenantPatch : The request body used when updating a tenant.
type TenantPatch struct {
	// The name for this tenant. The name is regionally unique across all tenants in the account.
	Name *string `json:"name,omitempty"`
}

// UnmarshalTenantPatch unmarshals an instance of TenantPatch from the specified map of raw messages.
func UnmarshalTenantPatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TenantPatch)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AsPatch returns a generic map representation of the TenantPatch
func (tenantPatch *TenantPatch) AsPatch() (_patch map[string]interface{}, err error) {
	_patch = map[string]interface{}{}
	if !core.IsNil(tenantPatch.Name) {
		_patch["name"] = tenantPatch.Name
	}

	return
}

// UpdateTargetOptions : The UpdateTarget options.
type UpdateTargetOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// The instance ID of the tenant target.
	TargetID *strfmt.UUID `json:"target_id" validate:"required"`

	// ETag provided by GET, POST or PATCH.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Updates a target of a tenant.
	TargetPatch map[string]interface{} `json:"TargetPatch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateTargetOptions : Instantiate UpdateTargetOptions
func (*IbmCloudLogsRoutingV0) NewUpdateTargetOptions(iBMAPIVersion string, tenantID *strfmt.UUID, targetID *strfmt.UUID, ifMatch string, targetPatch map[string]interface{}) *UpdateTargetOptions {
	return &UpdateTargetOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
		TargetID:      targetID,
		IfMatch:       core.StringPtr(ifMatch),
		TargetPatch:   targetPatch,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *UpdateTargetOptions) SetIBMAPIVersion(iBMAPIVersion string) *UpdateTargetOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTenantID : Allow user to set TenantID
func (_options *UpdateTargetOptions) SetTenantID(tenantID *strfmt.UUID) *UpdateTargetOptions {
	_options.TenantID = tenantID
	return _options
}

// SetTargetID : Allow user to set TargetID
func (_options *UpdateTargetOptions) SetTargetID(targetID *strfmt.UUID) *UpdateTargetOptions {
	_options.TargetID = targetID
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateTargetOptions) SetIfMatch(ifMatch string) *UpdateTargetOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetTargetPatch : Allow user to set TargetPatch
func (_options *UpdateTargetOptions) SetTargetPatch(targetPatch map[string]interface{}) *UpdateTargetOptions {
	_options.TargetPatch = targetPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTargetOptions) SetHeaders(param map[string]string) *UpdateTargetOptions {
	options.Headers = param
	return options
}

// UpdateTenantOptions : The UpdateTenant options.
type UpdateTenantOptions struct {
	// Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be
	// provided. Specify the current date to request the latest version.
	IBMAPIVersion *string `json:"IBM-API-Version"`

	// The instance ID of the tenant.
	TenantID *strfmt.UUID `json:"tenant_id" validate:"required"`

	// ETag provided by GET, POST or PATCH.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Update of a tenant.
	TenantPatch map[string]interface{} `json:"Tenant_patch" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateTenantOptions : Instantiate UpdateTenantOptions
func (*IbmCloudLogsRoutingV0) NewUpdateTenantOptions(iBMAPIVersion string, tenantID *strfmt.UUID, ifMatch string, tenantPatch map[string]interface{}) *UpdateTenantOptions {
	return &UpdateTenantOptions{
		IBMAPIVersion: core.StringPtr(iBMAPIVersion),
		TenantID:      tenantID,
		IfMatch:       core.StringPtr(ifMatch),
		TenantPatch:   tenantPatch,
	}
}

// SetIBMAPIVersion : Allow user to set IBMAPIVersion
func (_options *UpdateTenantOptions) SetIBMAPIVersion(iBMAPIVersion string) *UpdateTenantOptions {
	_options.IBMAPIVersion = core.StringPtr(iBMAPIVersion)
	return _options
}

// SetTenantID : Allow user to set TenantID
func (_options *UpdateTenantOptions) SetTenantID(tenantID *strfmt.UUID) *UpdateTenantOptions {
	_options.TenantID = tenantID
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateTenantOptions) SetIfMatch(ifMatch string) *UpdateTenantOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetTenantPatch : Allow user to set TenantPatch
func (_options *UpdateTenantOptions) SetTenantPatch(tenantPatch map[string]interface{}) *UpdateTenantOptions {
	_options.TenantPatch = tenantPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTenantOptions) SetHeaders(param map[string]string) *UpdateTenantOptions {
	options.Headers = param
	return options
}

// WriteStatus : The status of the write attempt to the target with the provided endpoint parameters.
type WriteStatus struct {
	// The status such as failed or success.
	Status *string `json:"status" validate:"required"`

	// Detailed description of the cause of the failure.
	ReasonForLastFailure *string `json:"reason_for_last_failure,omitempty"`

	// The timestamp of the failure.
	LastFailure *string `json:"last_failure,omitempty"`
}

// UnmarshalWriteStatus unmarshals an instance of WriteStatus from the specified map of raw messages.
func UnmarshalWriteStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WriteStatus)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reason_for_last_failure", &obj.ReasonForLastFailure)
	if err != nil {
		err = core.SDKErrorf(err, "", "reason_for_last_failure-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_failure", &obj.LastFailure)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_failure-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetPrototypeTargetTypeLogsPrototype : Full target description.
// This model "extends" TargetPrototype
type TargetPrototypeTargetTypeLogsPrototype struct {
	// The name for this tenant target. The name must be unique across all targets for this tenant.
	Name *string `json:"name" validate:"required"`

	// Cloud resource name of the log-sink target instance. Must be a valid CRN.
	LogSinkCrn *string `json:"log_sink_crn" validate:"required"`

	// List of properties to create or update a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Parameters *TargetParametersTypeLogsPrototype `json:"parameters" validate:"required"`
}

// NewTargetPrototypeTargetTypeLogsPrototype : Instantiate TargetPrototypeTargetTypeLogsPrototype (Generic Model Constructor)
func (*IbmCloudLogsRoutingV0) NewTargetPrototypeTargetTypeLogsPrototype(name string, logSinkCrn string, parameters *TargetParametersTypeLogsPrototype) (_model *TargetPrototypeTargetTypeLogsPrototype, err error) {
	_model = &TargetPrototypeTargetTypeLogsPrototype{
		Name:       core.StringPtr(name),
		LogSinkCrn: core.StringPtr(logSinkCrn),
		Parameters: parameters,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*TargetPrototypeTargetTypeLogsPrototype) isaTargetPrototype() bool {
	return true
}

// UnmarshalTargetPrototypeTargetTypeLogsPrototype unmarshals an instance of TargetPrototypeTargetTypeLogsPrototype from the specified map of raw messages.
func UnmarshalTargetPrototypeTargetTypeLogsPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetPrototypeTargetTypeLogsPrototype)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "log_sink_crn", &obj.LogSinkCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "log_sink_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalTargetParametersTypeLogsPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "parameters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetTypeLogs : Full target description.
// This model "extends" Target
type TargetTypeLogs struct {
	// Unique ID of the target.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The name for this tenant target. The name must be unique across all targets for this tenant.
	Name *string `json:"name" validate:"required"`

	// Resource version identifier.
	Etag *string `json:"etag" validate:"required"`

	// Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
	Type *string `json:"type" validate:"required"`

	// Timestamp the target was originally created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// Timestamp the target was last updated.
	UpdatedAt *string `json:"updated_at" validate:"required"`

	// Cloud resource name of the log-sink target instance. Must be a valid CRN.
	LogSinkCrn *string `json:"log_sink_crn" validate:"required"`

	// List of properties returned from a successful list operation for a log-sink of type <b>logs</b> (IBM Cloud Logs).
	Parameters *TargetParametersTypeLogs `json:"parameters" validate:"required"`
}

// Constants associated with the TargetTypeLogs.Type property.
// Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
const (
	TargetTypeLogs_Type_Logs = "logs"
)

func (*TargetTypeLogs) isaTarget() bool {
	return true
}

// UnmarshalTargetTypeLogs unmarshals an instance of TargetTypeLogs from the specified map of raw messages.
func UnmarshalTargetTypeLogs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetTypeLogs)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "etag", &obj.Etag)
	if err != nil {
		err = core.SDKErrorf(err, "", "etag-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "created_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		err = core.SDKErrorf(err, "", "updated_at-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "log_sink_crn", &obj.LogSinkCrn)
	if err != nil {
		err = core.SDKErrorf(err, "", "log_sink_crn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalTargetParametersTypeLogs)
	if err != nil {
		err = core.SDKErrorf(err, "", "parameters-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
