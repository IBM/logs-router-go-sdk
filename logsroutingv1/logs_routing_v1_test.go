/**
 * (C) Copyright IBM Corp. 2026.
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

package logsroutingv1_test

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
	"github.com/IBM/logs-router-go-sdk/logsroutingv1/logsroutingv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`LogsRoutingV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(logsRoutingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(logsRoutingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
				URL: "https://logsroutingv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(logsRoutingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTING_URL": "https://logsroutingv1/api",
				"LOGS_ROUTING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1UsingExternalConfig(&logsroutingv1.LogsRoutingV1Options{
				})
				Expect(logsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := logsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRoutingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1UsingExternalConfig(&logsroutingv1.LogsRoutingV1Options{
					URL: "https://testService/api",
				})
				Expect(logsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRoutingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1UsingExternalConfig(&logsroutingv1.LogsRoutingV1Options{
				})
				err := logsRoutingService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logsRoutingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRoutingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRoutingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRoutingService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTING_URL": "https://logsroutingv1/api",
				"LOGS_ROUTING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1UsingExternalConfig(&logsroutingv1.LogsRoutingV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(logsRoutingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTING_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1UsingExternalConfig(&logsroutingv1.LogsRoutingV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(logsRoutingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = logsroutingv1.GetServiceURLForRegion("au-syd.private")
			Expect(url).To(Equal("https://management.private.au-syd.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://management.au-syd.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("br-sao.private")
			Expect(url).To(Equal("https://management.private.br-sao.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://management.br-sao.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("ca-tor.private")
			Expect(url).To(Equal("https://management.private.ca-tor.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://management.ca-tor.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("ca-mon.private")
			Expect(url).To(Equal("https://management.private.ca-mon.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("ca-mon")
			Expect(url).To(Equal("https://management.ca-mon.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-de.private")
			Expect(url).To(Equal("https://management.private.eu-de.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://management.eu-de.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-es.private")
			Expect(url).To(Equal("https://management.private.eu-es.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-es")
			Expect(url).To(Equal("https://management.eu-es.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-gb.private")
			Expect(url).To(Equal("https://management.private.eu-gb.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://management.eu-gb.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-fr2.private")
			Expect(url).To(Equal("https://management.private.eu-fr2.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://management.eu-fr2.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("jp-tok.private")
			Expect(url).To(Equal("https://management.private.jp-tok.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://management.jp-tok.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("jp-osa.private")
			Expect(url).To(Equal("https://management.private.jp-osa.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://management.jp-osa.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("us-south.private")
			Expect(url).To(Equal("https://management.private.us-south.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://management.us-south.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("us-east.private")
			Expect(url).To(Equal("https://management.private.us-east.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://management.us-east.logs-router.cloud.ibm.com/v1"))
			Expect(err).To(BeNil())

			url, err = logsroutingv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-logging-tenant"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTenants with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(logsroutingv1.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("my-logging-tenant")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.ListTenants(listTenantsOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-logging-tenant"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}]}`)
				}))
			})
			It(`Invoke ListTenants successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(logsroutingv1.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("my-logging-tenant")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.ListTenantsWithContext(ctx, listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.ListTenantsWithContext(ctx, listTenantsOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-logging-tenant"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}]}`)
				}))
			})
			It(`Invoke ListTenants successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.ListTenants(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(logsroutingv1.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("my-logging-tenant")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTenants with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(logsroutingv1.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("my-logging-tenant")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTenantsOptions model with no property values
				listTenantsOptionsModelNew := new(logsroutingv1.ListTenantsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.ListTenants(listTenantsOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(logsroutingv1.ListTenantsOptions)
				listTenantsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantsOptionsModel.Name = core.StringPtr("my-logging-tenant")
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.ListTenants(listTenantsOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTenant with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(logsroutingv1.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.CreateTenant(createTenantOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke CreateTenant successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(logsroutingv1.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.CreateTenantWithContext(ctx, createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.CreateTenantWithContext(ctx, createTenantOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke CreateTenant successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.CreateTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(logsroutingv1.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTenant with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(logsroutingv1.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTenantOptions model with no property values
				createTenantOptionsModelNew := new(logsroutingv1.CreateTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.CreateTenant(createTenantOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(logsroutingv1.CreateTenantOptions)
				createTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTenantOptionsModel.Name = core.StringPtr("my-logging-tenant")
				createTenantOptionsModel.Targets = []logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.CreateTenant(createTenantOptionsModel)
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
		getTenantDetailPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantDetailPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTenantDetail with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(logsroutingv1.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
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
		getTenantDetailPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantDetailPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(logsroutingv1.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.GetTenantDetail(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(logsroutingv1.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTenantDetail with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(logsroutingv1.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTenantDetailOptions model with no property values
				getTenantDetailOptionsModelNew := new(logsroutingv1.GetTenantDetailOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.GetTenantDetail(getTenantDetailOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(logsroutingv1.GetTenantDetailOptions)
				getTenantDetailOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.GetTenantDetail(getTenantDetailOptionsModel)
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
		deleteTenantPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTenantPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTenant successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := logsRoutingService.DeleteTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(logsroutingv1.DeleteTenantOptions)
				deleteTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = logsRoutingService.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTenant with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(logsroutingv1.DeleteTenantOptions)
				deleteTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := logsRoutingService.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTenantOptions model with no property values
				deleteTenantOptionsModelNew := new(logsroutingv1.DeleteTenantOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = logsRoutingService.DeleteTenant(deleteTenantOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTenant(updateTenantOptions *UpdateTenantOptions) - Operation response error`, func() {
		updateTenantPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTenantPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTenant with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(logsroutingv1.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(logsroutingv1.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTenantOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.UpdateTenant(updateTenantOptionsModel)
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
		updateTenantPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0"
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke UpdateTenant successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(logsroutingv1.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(logsroutingv1.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTenantOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.UpdateTenantWithContext(ctx, updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.UpdateTenantWithContext(ctx, updateTenantOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "8717db99-2cfb-4ba6-a033-89c994c2e9f0", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "crn": "crn:v1:bluemix:public:logs-router:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:b10d649e-5cad-4274-a223-40f8e4dbdea6::", "name": "my-logging-tenant", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}], "write_status": {"status": "success", "reason_for_last_failure": "Logs endpoint is not reachable. Received status code: 403", "last_failure": "2024-10-14T10:49:09.000Z"}}`)
				}))
			})
			It(`Invoke UpdateTenant successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.UpdateTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(logsroutingv1.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(logsroutingv1.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTenantOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTenant with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(logsroutingv1.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(logsroutingv1.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTenantOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.UpdateTenant(updateTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTenantOptions model with no property values
				updateTenantOptionsModelNew := new(logsroutingv1.UpdateTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.UpdateTenant(updateTenantOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TenantPatch model
				tenantPatchModel := new(logsroutingv1.TenantPatch)
				tenantPatchModel.Name = core.StringPtr("my-logging-tenant")
				tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTenantOptions model
				updateTenantOptionsModel := new(logsroutingv1.UpdateTenantOptions)
				updateTenantOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTenantOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTenantOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.TenantPatch = tenantPatchModelAsPatch
				updateTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.UpdateTenant(updateTenantOptionsModel)
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
		listTenantTargetsPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-log-sink"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTenantTargets with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(logsroutingv1.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel.Name = core.StringPtr("my-log-sink")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
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
		listTenantTargetsPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTenantTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-log-sink"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke ListTenantTargets successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(logsroutingv1.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel.Name = core.StringPtr("my-log-sink")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.ListTenantTargetsWithContext(ctx, listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.ListTenantTargetsWithContext(ctx, listTenantTargetsOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"my-log-sink"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}]}`)
				}))
			})
			It(`Invoke ListTenantTargets successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.ListTenantTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(logsroutingv1.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel.Name = core.StringPtr("my-log-sink")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTenantTargets with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(logsroutingv1.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel.Name = core.StringPtr("my-log-sink")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTenantTargetsOptions model with no property values
				listTenantTargetsOptionsModelNew := new(logsroutingv1.ListTenantTargetsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the ListTenantTargetsOptions model
				listTenantTargetsOptionsModel := new(logsroutingv1.ListTenantTargetsOptions)
				listTenantTargetsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				listTenantTargetsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel.Name = core.StringPtr("my-log-sink")
				listTenantTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.ListTenantTargets(listTenantTargetsOptionsModel)
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
		createTargetPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTarget with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsroutingv1.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetPrototype = targetPrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.CreateTarget(createTargetOptionsModel)
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
		createTargetPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets"
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsroutingv1.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetPrototype = targetPrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.CreateTargetWithContext(ctx, createTargetOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsroutingv1.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetPrototype = targetPrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsroutingv1.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetPrototype = targetPrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(logsroutingv1.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.CreateTarget(createTargetOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsroutingv1.CreateTargetOptions)
				createTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				createTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				createTargetOptionsModel.TargetPrototype = targetPrototypeModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.CreateTarget(createTargetOptionsModel)
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
		getTenantTargetDetailsPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets/c40e55a5-0833-4748-b032-b8e8cfe6e135"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantTargetDetailsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTenantTargetDetails with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(logsroutingv1.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
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
		getTenantTargetDetailsPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets/c40e55a5-0833-4748-b032-b8e8cfe6e135"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTenantTargetDetailsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke GetTenantTargetDetails successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(logsroutingv1.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.GetTenantTargetDetailsWithContext(ctx, getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.GetTenantTargetDetailsWithContext(ctx, getTenantTargetDetailsOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke GetTenantTargetDetails successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.GetTenantTargetDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(logsroutingv1.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTenantTargetDetails with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(logsroutingv1.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTenantTargetDetailsOptions model with no property values
				getTenantTargetDetailsOptionsModelNew := new(logsroutingv1.GetTenantTargetDetailsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the GetTenantTargetDetailsOptions model
				getTenantTargetDetailsOptionsModel := new(logsroutingv1.GetTenantTargetDetailsOptions)
				getTenantTargetDetailsOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				getTenantTargetDetailsOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantTargetDetailsOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptionsModel)
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
		updateTargetPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets/c40e55a5-0833-4748-b032-b8e8cfe6e135"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTarget with error: Operation response processing error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersPatch model
				targetParametersPatchModel := new(logsroutingv1.TargetParametersPatch)
				targetParametersPatchModel.Host = core.StringPtr("www.example.com")
				targetParametersPatchModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPatch model
				targetPatchModel := new(logsroutingv1.TargetPatch)
				targetPatchModel.Name = core.StringPtr("my-log-sink")
				targetPatchModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPatchModel.Parameters = targetParametersPatchModel
				targetPatchModelAsPatch, asPatchErr := targetPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsroutingv1.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				updateTargetOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.TargetPatch = targetPatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRoutingService.EnableRetries(0, 0)
				result, response, operationErr = logsRoutingService.UpdateTarget(updateTargetOptionsModel)
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
		updateTargetPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets/c40e55a5-0833-4748-b032-b8e8cfe6e135"
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke UpdateTarget successfully with retries`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())
				logsRoutingService.EnableRetries(0, 0)

				// Construct an instance of the TargetParametersPatch model
				targetParametersPatchModel := new(logsroutingv1.TargetParametersPatch)
				targetParametersPatchModel.Host = core.StringPtr("www.example.com")
				targetParametersPatchModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPatch model
				targetPatchModel := new(logsroutingv1.TargetPatch)
				targetPatchModel.Name = core.StringPtr("my-log-sink")
				targetPatchModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPatchModel.Parameters = targetParametersPatchModel
				targetPatchModelAsPatch, asPatchErr := targetPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsroutingv1.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				updateTargetOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.TargetPatch = targetPatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRoutingService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRoutingService.DisableRetries()
				result, response, operationErr := logsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRoutingService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
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
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c40e55a5-0833-4748-b032-b8e8cfe6e135", "name": "my-log-sink", "etag": "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"", "type": "logs", "created_at": "2024-06-20T18:30:00.143Z", "updated_at": "2024-06-20T18:30:00.143Z", "log_sink_crn": "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::", "parameters": {"host": "www.example.com", "port": 1}}`)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRoutingService.UpdateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetParametersPatch model
				targetParametersPatchModel := new(logsroutingv1.TargetParametersPatch)
				targetParametersPatchModel.Host = core.StringPtr("www.example.com")
				targetParametersPatchModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPatch model
				targetPatchModel := new(logsroutingv1.TargetPatch)
				targetPatchModel.Name = core.StringPtr("my-log-sink")
				targetPatchModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPatchModel.Parameters = targetParametersPatchModel
				targetPatchModelAsPatch, asPatchErr := targetPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsroutingv1.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				updateTargetOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.TargetPatch = targetPatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTarget with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersPatch model
				targetParametersPatchModel := new(logsroutingv1.TargetParametersPatch)
				targetParametersPatchModel.Host = core.StringPtr("www.example.com")
				targetParametersPatchModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPatch model
				targetPatchModel := new(logsroutingv1.TargetPatch)
				targetPatchModel.Name = core.StringPtr("my-log-sink")
				targetPatchModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPatchModel.Parameters = targetParametersPatchModel
				targetPatchModelAsPatch, asPatchErr := targetPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsroutingv1.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				updateTargetOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.TargetPatch = targetPatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRoutingService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTargetOptions model with no property values
				updateTargetOptionsModelNew := new(logsroutingv1.UpdateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRoutingService.UpdateTarget(updateTargetOptionsModelNew)
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
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the TargetParametersPatch model
				targetParametersPatchModel := new(logsroutingv1.TargetParametersPatch)
				targetParametersPatchModel.Host = core.StringPtr("www.example.com")
				targetParametersPatchModel.Port = core.Int64Ptr(int64(1))

				// Construct an instance of the TargetPatch model
				targetPatchModel := new(logsroutingv1.TargetPatch)
				targetPatchModel.Name = core.StringPtr("my-log-sink")
				targetPatchModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPatchModel.Parameters = targetParametersPatchModel
				targetPatchModelAsPatch, asPatchErr := targetPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsroutingv1.UpdateTargetOptions)
				updateTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				updateTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				updateTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				updateTargetOptionsModel.IfMatch = core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.TargetPatch = targetPatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRoutingService.UpdateTarget(updateTargetOptionsModel)
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
		deleteTargetPath := "/tenants/8717db99-2cfb-4ba6-a033-89c994c2e9f0/targets/c40e55a5-0833-4748-b032-b8e8cfe6e135"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Ibm-Api-Version"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Api-Version"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := logsRoutingService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(logsroutingv1.DeleteTargetOptions)
				deleteTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				deleteTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = logsRoutingService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				logsRoutingService, serviceErr := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRoutingService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(logsroutingv1.DeleteTargetOptions)
				deleteTargetOptionsModel.IBMAPIVersion = core.StringPtr("testString")
				deleteTargetOptionsModel.TenantID = CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				deleteTargetOptionsModel.TargetID = CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRoutingService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := logsRoutingService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(logsroutingv1.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = logsRoutingService.DeleteTarget(deleteTargetOptionsModelNew)
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
			logsRoutingService, _ := logsroutingv1.NewLogsRoutingV1(&logsroutingv1.LogsRoutingV1Options{
				URL:           "http://logsroutingv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				Expect(targetParametersTypeLogsPrototypeModel).ToNot(BeNil())
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))
				Expect(targetParametersTypeLogsPrototypeModel.Host).To(Equal(core.StringPtr("www.example.com")))
				Expect(targetParametersTypeLogsPrototypeModel.Port).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				Expect(targetPrototypeModel).ToNot(BeNil())
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel
				Expect(targetPrototypeModel.Name).To(Equal(core.StringPtr("my-log-sink")))
				Expect(targetPrototypeModel.LogSinkCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")))
				Expect(targetPrototypeModel.Parameters).To(Equal(targetParametersTypeLogsPrototypeModel))

				// Construct an instance of the CreateTargetOptions model
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				iBmapiVersion := "testString"
				var targetPrototype logsroutingv1.TargetPrototypeIntf = nil
				createTargetOptionsModel := logsRoutingService.NewCreateTargetOptions(tenantID, iBmapiVersion, targetPrototype)
				createTargetOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				createTargetOptionsModel.SetIBMAPIVersion("testString")
				createTargetOptionsModel.SetTargetPrototype(targetPrototypeModel)
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(createTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(createTargetOptionsModel.TargetPrototype).To(Equal(targetPrototypeModel))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTenantOptions successfully`, func() {
				// Construct an instance of the TargetParametersTypeLogsPrototype model
				targetParametersTypeLogsPrototypeModel := new(logsroutingv1.TargetParametersTypeLogsPrototype)
				Expect(targetParametersTypeLogsPrototypeModel).ToNot(BeNil())
				targetParametersTypeLogsPrototypeModel.Host = core.StringPtr("www.example.com")
				targetParametersTypeLogsPrototypeModel.Port = core.Int64Ptr(int64(1))
				Expect(targetParametersTypeLogsPrototypeModel.Host).To(Equal(core.StringPtr("www.example.com")))
				Expect(targetParametersTypeLogsPrototypeModel.Port).To(Equal(core.Int64Ptr(int64(1))))

				// Construct an instance of the TargetPrototypeTargetTypeLogsPrototype model
				targetPrototypeModel := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
				Expect(targetPrototypeModel).ToNot(BeNil())
				targetPrototypeModel.Name = core.StringPtr("my-log-sink")
				targetPrototypeModel.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
				targetPrototypeModel.Parameters = targetParametersTypeLogsPrototypeModel
				Expect(targetPrototypeModel.Name).To(Equal(core.StringPtr("my-log-sink")))
				Expect(targetPrototypeModel.LogSinkCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")))
				Expect(targetPrototypeModel.Parameters).To(Equal(targetParametersTypeLogsPrototypeModel))

				// Construct an instance of the CreateTenantOptions model
				iBmapiVersion := "testString"
				createTenantOptionsName := "my-logging-tenant"
				createTenantOptionsTargets := []logsroutingv1.TargetPrototypeIntf{}
				createTenantOptionsModel := logsRoutingService.NewCreateTenantOptions(iBmapiVersion, createTenantOptionsName, createTenantOptionsTargets)
				createTenantOptionsModel.SetIBMAPIVersion("testString")
				createTenantOptionsModel.SetName("my-logging-tenant")
				createTenantOptionsModel.SetTargets([]logsroutingv1.TargetPrototypeIntf{targetPrototypeModel})
				createTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTenantOptionsModel).ToNot(BeNil())
				Expect(createTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(createTenantOptionsModel.Name).To(Equal(core.StringPtr("my-logging-tenant")))
				Expect(createTenantOptionsModel.Targets).To(Equal([]logsroutingv1.TargetPrototypeIntf{targetPrototypeModel}))
				Expect(createTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				targetID := CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				deleteTargetOptionsModel := logsRoutingService.NewDeleteTargetOptions(iBmapiVersion, tenantID, targetID)
				deleteTargetOptionsModel.SetIBMAPIVersion("testString")
				deleteTargetOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				deleteTargetOptionsModel.SetTargetID(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135"))
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(deleteTargetOptionsModel.TargetID).To(Equal(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTenantOptions successfully`, func() {
				// Construct an instance of the DeleteTenantOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				deleteTenantOptionsModel := logsRoutingService.NewDeleteTenantOptions(iBmapiVersion, tenantID)
				deleteTenantOptionsModel.SetIBMAPIVersion("testString")
				deleteTenantOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				deleteTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTenantOptionsModel).ToNot(BeNil())
				Expect(deleteTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(deleteTenantOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(deleteTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTenantDetailOptions successfully`, func() {
				// Construct an instance of the GetTenantDetailOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				getTenantDetailOptionsModel := logsRoutingService.NewGetTenantDetailOptions(iBmapiVersion, tenantID)
				getTenantDetailOptionsModel.SetIBMAPIVersion("testString")
				getTenantDetailOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				getTenantDetailOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTenantDetailOptionsModel).ToNot(BeNil())
				Expect(getTenantDetailOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(getTenantDetailOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(getTenantDetailOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTenantTargetDetailsOptions successfully`, func() {
				// Construct an instance of the GetTenantTargetDetailsOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				targetID := CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				getTenantTargetDetailsOptionsModel := logsRoutingService.NewGetTenantTargetDetailsOptions(iBmapiVersion, tenantID, targetID)
				getTenantTargetDetailsOptionsModel.SetIBMAPIVersion("testString")
				getTenantTargetDetailsOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				getTenantTargetDetailsOptionsModel.SetTargetID(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135"))
				getTenantTargetDetailsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTenantTargetDetailsOptionsModel).ToNot(BeNil())
				Expect(getTenantTargetDetailsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(getTenantTargetDetailsOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(getTenantTargetDetailsOptionsModel.TargetID).To(Equal(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")))
				Expect(getTenantTargetDetailsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTenantTargetsOptions successfully`, func() {
				// Construct an instance of the ListTenantTargetsOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				listTenantTargetsOptionsModel := logsRoutingService.NewListTenantTargetsOptions(iBmapiVersion, tenantID)
				listTenantTargetsOptionsModel.SetIBMAPIVersion("testString")
				listTenantTargetsOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				listTenantTargetsOptionsModel.SetName("my-log-sink")
				listTenantTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTenantTargetsOptionsModel).ToNot(BeNil())
				Expect(listTenantTargetsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(listTenantTargetsOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(listTenantTargetsOptionsModel.Name).To(Equal(core.StringPtr("my-log-sink")))
				Expect(listTenantTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTenantsOptions successfully`, func() {
				// Construct an instance of the ListTenantsOptions model
				iBmapiVersion := "testString"
				listTenantsOptionsModel := logsRoutingService.NewListTenantsOptions(iBmapiVersion)
				listTenantsOptionsModel.SetIBMAPIVersion("testString")
				listTenantsOptionsModel.SetName("my-logging-tenant")
				listTenantsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTenantsOptionsModel).ToNot(BeNil())
				Expect(listTenantsOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(listTenantsOptionsModel.Name).To(Equal(core.StringPtr("my-logging-tenant")))
				Expect(listTenantsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTargetParametersTypeLogsPrototype successfully`, func() {
				host := "www.example.com"
				port := int64(1)
				_model, err := logsRoutingService.NewTargetParametersTypeLogsPrototype(host, port)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateTargetOptions successfully`, func() {
				// Construct an instance of the UpdateTargetOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				targetID := CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")
				ifMatch := "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\""
				targetPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTargetOptionsModel := logsRoutingService.NewUpdateTargetOptions(iBmapiVersion, tenantID, targetID, ifMatch, targetPatch)
				updateTargetOptionsModel.SetIBMAPIVersion("testString")
				updateTargetOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				updateTargetOptionsModel.SetTargetID(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135"))
				updateTargetOptionsModel.SetIfMatch("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTargetOptionsModel.SetTargetPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTargetOptionsModel).ToNot(BeNil())
				Expect(updateTargetOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(updateTargetOptionsModel.TargetID).To(Equal(CreateMockUUID("c40e55a5-0833-4748-b032-b8e8cfe6e135")))
				Expect(updateTargetOptionsModel.IfMatch).To(Equal(core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
				Expect(updateTargetOptionsModel.TargetPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTenantOptions successfully`, func() {
				// Construct an instance of the UpdateTenantOptions model
				iBmapiVersion := "testString"
				tenantID := CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")
				ifMatch := "\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\""
				tenantPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTenantOptionsModel := logsRoutingService.NewUpdateTenantOptions(iBmapiVersion, tenantID, ifMatch, tenantPatch)
				updateTenantOptionsModel.SetIBMAPIVersion("testString")
				updateTenantOptionsModel.SetTenantID(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0"))
				updateTenantOptionsModel.SetIfMatch("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")
				updateTenantOptionsModel.SetTenantPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTenantOptionsModel).ToNot(BeNil())
				Expect(updateTenantOptionsModel.IBMAPIVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateTenantOptionsModel.TenantID).To(Equal(CreateMockUUID("8717db99-2cfb-4ba6-a033-89c994c2e9f0")))
				Expect(updateTenantOptionsModel.IfMatch).To(Equal(core.StringPtr("\"822b4b5423e225206c1d75666595714a11925cd0f82b229839864443d6c3c049\"")))
				Expect(updateTenantOptionsModel.TenantPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTargetPrototypeTargetTypeLogsPrototype successfully`, func() {
				name := "my-log-sink"
				logSinkCrn := "crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::"
				var parameters *logsroutingv1.TargetParametersTypeLogsPrototype = nil
				_, err := logsRoutingService.NewTargetPrototypeTargetTypeLogsPrototype(name, logSinkCrn, parameters)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalTargetParametersPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TargetParametersPatch)
			model.Host = core.StringPtr("www.example.com")
			model.Port = core.Int64Ptr(int64(1))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TargetParametersPatch
			err = logsroutingv1.UnmarshalTargetParametersPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetParametersTypeLogsPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TargetParametersTypeLogsPrototype)
			model.Host = core.StringPtr("www.example.com")
			model.Port = core.Int64Ptr(int64(1))

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TargetParametersTypeLogsPrototype
			err = logsroutingv1.UnmarshalTargetParametersTypeLogsPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TargetPatch)
			model.Name = core.StringPtr("my-log-sink")
			model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TargetPatch
			err = logsroutingv1.UnmarshalTargetPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TargetPrototype)
			model.Name = core.StringPtr("my-log-sink")
			model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TargetPrototype
			err = logsroutingv1.UnmarshalTargetPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTenantPatch successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TenantPatch)
			model.Name = core.StringPtr("my-logging-tenant")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TenantPatch
			err = logsroutingv1.UnmarshalTenantPatch(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetPrototypeTargetTypeLogsPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsroutingv1.TargetPrototypeTargetTypeLogsPrototype)
			model.Name = core.StringPtr("my-log-sink")
			model.LogSinkCrn = core.StringPtr("crn:v1:bluemix:public:logs:eu-de:a/4516b8fa0a174a71899f5affa4f18d78:cfef55c6-cdfe-48c8-b882-aefc271532e4::")
			model.Parameters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsroutingv1.TargetPrototypeTargetTypeLogsPrototype
			err = logsroutingv1.UnmarshalTargetPrototypeTargetTypeLogsPrototype(raw, &result)
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
