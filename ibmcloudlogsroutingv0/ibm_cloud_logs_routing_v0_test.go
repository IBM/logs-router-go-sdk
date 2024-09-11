/**
 * (C) Copyright IBM Corp. 2024.
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

package ibmcloudlogsroutingv0_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IBMCloudLogsRoutingV0`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudLogsRoutingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
				URL: "https://ibmcloudlogsroutingv0/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudLogsRoutingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_LOGS_ROUTING_URL":       "https://ibmcloudlogsroutingv0/api",
				"IBM_CLOUD_LOGS_ROUTING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{})
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudLogsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudLogsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudLogsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudLogsRoutingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudLogsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudLogsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudLogsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudLogsRoutingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{})
				err := ibmCloudLogsRoutingService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudLogsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudLogsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudLogsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudLogsRoutingService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_LOGS_ROUTING_URL":       "https://ibmcloudlogsroutingv0/api",
				"IBM_CLOUD_LOGS_ROUTING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudLogsRoutingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_LOGS_ROUTING_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudLogsRoutingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudlogsroutingv0.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := ibmcloudlogsroutingv0.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://management.eu-gb.logs-router.cloud.ibm.com/v1"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := ibmcloudlogsroutingv0.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`ListTenants(listTenantsOptions *ListTenantsOptions) - Operation response error`, func() {
		listTenantsPath := "/tenants"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTenants with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("2024-06-15")
				listTenantsOptionsModel.Name = core.StringPtr("testString")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTenants(listTenantsOptions *ListTenantsOptions)`, func() {
		listTenantsPath := "/tenants"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}]}`)
				}))
			})
			It(`Invoke ListTenants successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("testString")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.ListTenantsWithContext(ctx, listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.ListTenantsWithContext(ctx, listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}]}`)
				}))
			})
			It(`Invoke ListTenants successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenants(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("testString")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTenants with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("testString")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTenantsOptions model with no property values
				listTenantsOptionsModelNew := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTenants successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("testString")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTenant(createTenantOptions *CreateTenantOptions) - Operation response error`, func() {
		createTenantPath := "/tenants"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTenantPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTenant with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTenant(createTenantOptions *CreateTenantOptions)`, func() {
		createTenantPath := "/tenants"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTenantPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke CreateTenant successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.CreateTenantWithContext(ctx, createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.CreateTenantWithContext(ctx, createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTenantPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke CreateTenant successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTenant with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTenantOptions model with no property values
				createTenantOptionsModelNew := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTenant successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmcloudlogsroutingv0.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTenantDetail(getTenantDetailOptions *GetTenantDetailOptions) - Operation response error`, func() {
		getTenantDetailPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantDetailPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTenantDetail with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTenantDetail(getTenantDetailOptions *GetTenantDetailOptions)`, func() {
		getTenantDetailPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantDetailPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantDetailPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantDetail(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTenantDetail with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTenantDetailOptions model with no property values
				getTenantDetailOptionsModelNew := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTenantDetail successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmcloudlogsroutingv0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTenant(deleteTenantOptions *DeleteTenantOptions)`, func() {
		deleteTenantPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTenantPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTenant successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudLogsRoutingService.DeleteTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmcloudlogsroutingv0.DeleteTenantOptions)
				deleteTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudLogsRoutingService.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTenant with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmcloudlogsroutingv0.DeleteTenantOptions)
				deleteTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudLogsRoutingService.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTenantOptions model with no property values
				deleteTenantOptionsModelNew := new(ibmcloudlogsroutingv0.DeleteTenantOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudLogsRoutingService.DeleteTenant(deleteTenantOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTenant(updateTenantOptions *UpdateTenantOptions) - Operation response error`, func() {
		updateTenantPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTenantPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTenant with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(ibmcloudlogsroutingv0.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTenantOptionsModel.IfMatch = core.StringPtr("testString")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTenant(updateTenantOptions *UpdateTenantOptions)`, func() {
		updateTenantPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTenantPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke UpdateTenant successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(ibmcloudlogsroutingv0.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTenantOptionsModel.IfMatch = core.StringPtr("testString")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.UpdateTenantWithContext(ctx, updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.UpdateTenantWithContext(ctx, updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTenantPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke UpdateTenant successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(ibmcloudlogsroutingv0.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTenantOptionsModel.IfMatch = core.StringPtr("testString")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTenant with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(ibmcloudlogsroutingv0.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTenantOptionsModel.IfMatch = core.StringPtr("testString")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTenantOptions model with no property values
				updateTenantOptionsModelNew := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateTenant successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(ibmcloudlogsroutingv0.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(ibmcloudlogsroutingv0.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTenantOptionsModel.IfMatch = core.StringPtr("testString")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTenantTargets(listTenantTargetsOptions *ListTenantTargetsOptions) - Operation response error`, func() {
		listTenantTargetsPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTenantTargets with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel.Name = core.StringPtr("testString")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTenantTargets(listTenantTargetsOptions *ListTenantTargetsOptions)`, func() {
		listTenantTargetsPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke ListTenantTargets successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel.Name = core.StringPtr("testString")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.ListTenantTargetsWithContext(ctx, listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.ListTenantTargetsWithContext(ctx, listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke ListTenantTargets successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenantTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel.Name = core.StringPtr("testString")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTenantTargets with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel.Name = core.StringPtr("testString")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTenantTargetsOptions model with no property values
				listTenantTargetsOptionsModelNew := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTenantTargets successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(ibmcloudlogsroutingv0.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel.Name = core.StringPtr("testString")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions) - Operation response error`, func() {
		createTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTarget with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetTypePrototype = targetTypePrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
		createTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetTypePrototype = targetTypePrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetTypePrototype = targetTypePrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetTypePrototype = targetTypePrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(ibmcloudlogsroutingv0.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetTypePrototype = targetTypePrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTenantTargetDetails(getTenantTargetDetailsOptions *GetTenantTargetDetailsOptions) - Operation response error`, func() {
		getTenantTargetDetailsPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantTargetDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTenantTargetDetails with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTenantTargetDetails(getTenantTargetDetailsOptions *GetTenantTargetDetailsOptions)`, func() {
		getTenantTargetDetailsPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantTargetDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke GetTenantTargetDetails successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetailsWithContext(ctx, getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.GetTenantTargetDetailsWithContext(ctx, getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantTargetDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke GetTenantTargetDetails successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTenantTargetDetails with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTenantTargetDetailsOptions model with no property values
				getTenantTargetDetailsOptionsModelNew := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTenantTargetDetails successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(ibmcloudlogsroutingv0.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions) - Operation response error`, func() {
		updateTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTarget with error: Operation response processing error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePatchLogDna model
				targetTypePatchModel := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
				targetTypePatchModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePatchModel.Name = core.StringPtr("my-log-sink")
				targetTypePatchModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.IfMatch = core.StringPtr("testString")
				updateTargetOptionsModel.TargetTypePatch = targetTypePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudLogsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
		updateTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke UpdateTarget successfully with retries`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
				ibmCloudLogsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePatchLogDna model
				targetTypePatchModel := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
				targetTypePatchModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePatchModel.Name = core.StringPtr("my-log-sink")
				targetTypePatchModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.IfMatch = core.StringPtr("testString")
				updateTargetOptionsModel.TargetTypePatch = targetTypePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudLogsRoutingService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudLogsRoutingService.DisableRetries()
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudLogsRoutingService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "log_sink_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "name": "my-log-sink", "etag": "\"c3a43545a7f2675970671ac3a57b8db067a1866b2222e1b950ee8da612e347c6\"", "type": "logdna", "created_at": "2024-06-20T18:30:00.143156Z", "updated_at": "2024-06-20T18:30:00.143156Z", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePatchLogDna model
				targetTypePatchModel := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
				targetTypePatchModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePatchModel.Name = core.StringPtr("my-log-sink")
				targetTypePatchModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.IfMatch = core.StringPtr("testString")
				updateTargetOptionsModel.TargetTypePatch = targetTypePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTarget with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePatchLogDna model
				targetTypePatchModel := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
				targetTypePatchModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePatchModel.Name = core.StringPtr("my-log-sink")
				targetTypePatchModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.IfMatch = core.StringPtr("testString")
				updateTargetOptionsModel.TargetTypePatch = targetTypePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTargetOptions model with no property values
				updateTargetOptionsModelNew := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")

				// Construct an instance of the TargetTypePatchLogDna model
				targetTypePatchModel := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
				targetTypePatchModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePatchModel.Name = core.StringPtr("my-log-sink")
				targetTypePatchModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmcloudlogsroutingv0.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.IfMatch = core.StringPtr("testString")
				updateTargetOptionsModel.TargetTypePatch = targetTypePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
		deleteTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673/targets/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "2024-06-15")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudLogsRoutingService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(ibmcloudlogsroutingv0.DeleteTargetOptions)
				deleteTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudLogsRoutingService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				ibmCloudLogsRoutingService, serviceErr := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudLogsRoutingService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(ibmcloudlogsroutingv0.DeleteTargetOptions)
				deleteTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTargetOptionsModel.TargetID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudLogsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudLogsRoutingService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(ibmcloudlogsroutingv0.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudLogsRoutingService.DeleteTarget(deleteTargetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmCloudLogsRoutingService, _ := ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0(&ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{
				URL:           "http://ibmcloudlogsroutingv0modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				Expect(targetParametersTypeLogDnaPrototypeModel).ToNot(BeNil())
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")
				Expect(targetParametersTypeLogDnaPrototypeModel.Host).To(Equal(core.StringPtr("www.example.com")))
				Expect(targetParametersTypeLogDnaPrototypeModel.Port).To(Equal(core.Int64Ptr(int64(1))))
				Expect(targetParametersTypeLogDnaPrototypeModel.AccessCredential).To(Equal(core.StringPtr("ingestion-secret")))

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				Expect(targetTypePrototypeModel).ToNot(BeNil())
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				Expect(targetTypePrototypeModel.LogSinkCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")))
				Expect(targetTypePrototypeModel.Name).To(Equal(core.StringPtr("my-log-sink")))
				Expect(targetTypePrototypeModel.Parameters).To(Equal(targetParametersTypeLogDnaPrototypeModel))

				// Construct an instance of the CreateTargetOptions model
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				ibmAPIVersion := "testString"
				region := "ca-tor"
				var targetTypePrototype ibmcloudlogsroutingv0.TargetTypePrototypeIntf = nil
				createTargetOptionsModel := ibmCloudLogsRoutingService.NewCreateTargetOptions(tenantID, ibmAPIVersion, region, targetTypePrototype)
				createTargetOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				createTargetOptionsModel.SetIBMAPIVersion("testString")
				createTargetOptionsModel.SetTargetTypePrototype(targetTypePrototypeModel)
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(createTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(createTargetOptionsModel.TargetTypePrototype).To(Equal(targetTypePrototypeModel))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTenantOptions successfully`, func() {
				// Construct an instance of the TargetParametersTypeLogDnaPrototype model
				targetParametersTypeLogDnaPrototypeModel := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
				Expect(targetParametersTypeLogDnaPrototypeModel).ToNot(BeNil())
				targetParametersTypeLogDnaPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogDnaPrototypeModel.Port = core.Int64Ptr(int64(1))
				targetParametersTypeLogDnaPrototypeModel.AccessCredential = core.StringPtr("ingestion-secret")
				Expect(targetParametersTypeLogDnaPrototypeModel.Host).To(Equal(core.StringPtr("www.example.com")))
				Expect(targetParametersTypeLogDnaPrototypeModel.Port).To(Equal(core.Int64Ptr(int64(1))))
				Expect(targetParametersTypeLogDnaPrototypeModel.AccessCredential).To(Equal(core.StringPtr("ingestion-secret")))

				// Construct an instance of the TargetTypePrototypeTargetTypeLogDnaPrototype model
				targetTypePrototypeModel := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
				Expect(targetTypePrototypeModel).ToNot(BeNil())
				targetTypePrototypeModel.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				targetTypePrototypeModel.Name = core.StringPtr("my-log-sink")
				targetTypePrototypeModel.Parameters = targetParametersTypeLogDnaPrototypeModel
				Expect(targetTypePrototypeModel.LogSinkCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")))
				Expect(targetTypePrototypeModel.Name).To(Equal(core.StringPtr("my-log-sink")))
				Expect(targetTypePrototypeModel.Parameters).To(Equal(targetParametersTypeLogDnaPrototypeModel))

				// Construct an instance of the CreateTenantOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				createTenantOptionsName := "my-logging-tenant"
				createTenantOptionsTargets := []ibmcloudlogsroutingv0.TargetTypePrototypeIntf{}
				createTenantOptionsModel := ibmCloudLogsRoutingService.NewCreateTenantOptions(ibmAPIVersion, createTenantOptionsName, region, createTenantOptionsTargets)
				createTenantOptionsModel.SetIBMAPIVersion("testString")
				createTenantOptionsModel.SetName("my-logging-tenant")
				createTenantOptionsModel.SetTargets([]ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel})
				createTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTenantOptionsModel).ToNot(BeNil())
				Expect(createTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(createTenantOptionsModel.Name).To(Equal(core.StringPtr("my-logging-tenant")))
				Expect(createTenantOptionsModel.Targets).To(Equal([]ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel}))
				Expect(createTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				targetID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTargetOptionsModel := ibmCloudLogsRoutingService.NewDeleteTargetOptions(ibmAPIVersion, region, tenantID, targetID)
				deleteTargetOptionsModel.SetIBMAPIVersion("testString")
				deleteTargetOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deleteTargetOptionsModel.SetTargetID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deleteTargetOptionsModel.TargetID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTenantOptions successfully`, func() {
				// Construct an instance of the DeleteTenantOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel := ibmCloudLogsRoutingService.NewDeleteTenantOptions(ibmAPIVersion, region, tenantID)
				deleteTenantOptionsModel.SetIBMAPIVersion("testString")
				deleteTenantOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deleteTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTenantOptionsModel).ToNot(BeNil())
				Expect(deleteTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(deleteTenantOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deleteTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTenantDetailOptions successfully`, func() {
				// Construct an instance of the GetTenantDetailOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel := ibmCloudLogsRoutingService.NewGetTenantDetailOptions(ibmAPIVersion, tenantID, region)
				getTenantDetailOptionsModel.SetIBMAPIVersion("testString")
				getTenantDetailOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getTenantDetailOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTenantDetailOptionsModel).ToNot(BeNil())
				Expect(getTenantDetailOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(getTenantDetailOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getTenantDetailOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTenantTargetDetailsOptions successfully`, func() {
				// Construct an instance of the GetTenantTargetDetailsOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				targetID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantTargetDetailsOptionsModel := ibmCloudLogsRoutingService.NewGetTenantTargetDetailsOptions(ibmAPIVersion, region, tenantID, targetID)
				getTenantTargetDetailsOptionsModel.SetIBMAPIVersion("testString")
				getTenantTargetDetailsOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getTenantTargetDetailsOptionsModel.SetTargetID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getTenantTargetDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTenantTargetDetailsOptionsModel).ToNot(BeNil())
				Expect(getTenantTargetDetailsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(getTenantTargetDetailsOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getTenantTargetDetailsOptionsModel.TargetID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getTenantTargetDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTenantTargetsOptions successfully`, func() {
				// Construct an instance of the ListTenantTargetsOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listTenantTargetsOptionsModel := ibmCloudLogsRoutingService.NewListTenantTargetsOptions(ibmAPIVersion, region, tenantID)
				listTenantTargetsOptionsModel.SetIBMAPIVersion("testString")
				listTenantTargetsOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listTenantTargetsOptionsModel.SetName("testString")
				listTenantTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTenantTargetsOptionsModel).ToNot(BeNil())
				Expect(listTenantTargetsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(listTenantTargetsOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listTenantTargetsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listTenantTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTenantsOptions successfully`, func() {
				// Construct an instance of the ListTenantsOptions model
				ibmAPIVersion := "testString"
				region := "ca-tor"
				listTenantsOptionsModel := ibmCloudLogsRoutingService.NewListTenantsOptions(ibmAPIVersion, region)
				listTenantsOptionsModel.SetIBMAPIVersion("testString")
				listTenantsOptionsModel.SetName("testString")
				listTenantsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTenantsOptionsModel).ToNot(BeNil())
				Expect(listTenantsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(listTenantsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listTenantsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTargetParametersTypeLogDnaPrototype successfully`, func() {
				host := "www.example.com"
				port := int64(1)
				accessCredential := "ingestion-secret"
				_model, err := ibmCloudLogsRoutingService.NewTargetParametersTypeLogDnaPrototype(host, port, accessCredential)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetParametersTypeLogsPrototype successfully`, func() {
				host := "www.example.com"
				port := int64(1)
				_model, err := ibmCloudLogsRoutingService.NewTargetParametersTypeLogsPrototype(host, port)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateTargetOptions successfully`, func() {
				// Construct an instance of the UpdateTargetOptions model
				ibmAPIVersion := "testString"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				targetID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				ifMatch := "testString"
				region := "ca-tor"
				targetTypePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTargetOptionsModel := ibmCloudLogsRoutingService.NewUpdateTargetOptions(ibmAPIVersion, region, tenantID, targetID, ifMatch, targetTypePatch)
				updateTargetOptionsModel.SetIBMAPIVersion("testString")
				updateTargetOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateTargetOptionsModel.SetTargetID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateTargetOptionsModel.SetIfMatch("testString")
				updateTargetOptionsModel.SetTargetTypePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTargetOptionsModel).ToNot(BeNil())
				Expect(updateTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateTargetOptionsModel.TargetID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateTargetOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateTargetOptionsModel.TargetTypePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTenantOptions successfully`, func() {
				// Construct an instance of the UpdateTenantOptions model
				ibmAPIVersion := "testString"
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				ifMatch := "testString"
				region := "ca-tor"
				tenantPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTenantOptionsModel := ibmCloudLogsRoutingService.NewUpdateTenantOptions(ibmAPIVersion, region, tenantID, ifMatch, tenantPatch)
				updateTenantOptionsModel.SetIBMAPIVersion("testString")
				updateTenantOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateTenantOptionsModel.SetIfMatch("testString")
				updateTenantOptionsModel.SetTenantPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTenantOptionsModel).ToNot(BeNil())
				Expect(updateTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateTenantOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateTenantOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateTenantOptionsModel.TenantPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTargetTypePrototypeTargetTypeLogDnaPrototype successfully`, func() {
				logSinkCRN := "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
				name := "my-log-sink"
				_model, err := ibmCloudLogsRoutingService.NewTargetTypePrototypeTargetTypeLogDnaPrototype(logSinkCRN, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetTypePrototypeTargetTypeLogsPrototype successfully`, func() {
				logSinkCRN := "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
				name := "my-log-sink"
				_model, err := ibmCloudLogsRoutingService.NewTargetTypePrototypeTargetTypeLogsPrototype(logSinkCRN, name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalTargetParametersTypeLogDnaPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype)
			model.Host = core.StringPtr("www.example.com")
			model.Port = core.Int64Ptr(int64(1))
			model.AccessCredential = core.StringPtr("ingestion-secret")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetParametersTypeLogDnaPrototype
			err = ibmcloudlogsroutingv0.UnmarshalTargetParametersTypeLogDnaPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetParametersTypeLogsPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype)
			model.Host = core.StringPtr("www.example.com")
			model.Port = core.Int64Ptr(int64(1))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetParametersTypeLogsPrototype
			err = ibmcloudlogsroutingv0.UnmarshalTargetParametersTypeLogsPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePatch successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePatch)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePatch
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePrototype)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePrototype
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTenantPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TenantPatch)
			model.Name = core.StringPtr("my-logging-tenant")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TenantPatch
			err = ibmcloudlogsroutingv0.UnmarshalTenantPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePatchLogDna successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePatchLogDna)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePatchLogDna
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePatchLogDna(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePatchLogs successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePatchLogs)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePatchLogs
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePatchLogs(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePrototypeTargetTypeLogDnaPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePrototypeTargetTypeLogDnaPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetTypePrototypeTargetTypeLogsPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogsPrototype)
			model.LogSinkCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
			model.Name = core.StringPtr("my-log-sink")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogsPrototype
			err = ibmcloudlogsroutingv0.UnmarshalTargetTypePrototypeTargetTypeLogsPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
