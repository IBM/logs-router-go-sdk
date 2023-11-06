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

package ibmlogsrouteropenapi30v0_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IbmLogsRouterOpenApi30V0`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmLogsRouterOpenApi30Service).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				URL: "https://ibmlogsrouteropenapi30v0/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmLogsRouterOpenApi30Service).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_LOGS_ROUTER_OPEN_API_3_0_URL": "https://ibmlogsrouteropenapi30v0/api",
				"IBM_LOGS_ROUTER_OPEN_API_3_0_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				})
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmLogsRouterOpenApi30Service.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmLogsRouterOpenApi30Service.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmLogsRouterOpenApi30Service.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmLogsRouterOpenApi30Service.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL: "https://testService/api",
				})
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmLogsRouterOpenApi30Service.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmLogsRouterOpenApi30Service.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmLogsRouterOpenApi30Service.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmLogsRouterOpenApi30Service.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				})
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmLogsRouterOpenApi30Service.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmLogsRouterOpenApi30Service.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmLogsRouterOpenApi30Service.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmLogsRouterOpenApi30Service.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_LOGS_ROUTER_OPEN_API_3_0_URL": "https://ibmlogsrouteropenapi30v0/api",
				"IBM_LOGS_ROUTER_OPEN_API_3_0_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmLogsRouterOpenApi30Service).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_LOGS_ROUTER_OPEN_API_3_0_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmLogsRouterOpenApi30Service).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmlogsrouteropenapi30v0.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := ibmlogsrouteropenapi30v0.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://management.private.eu-gb.logs-router.test.cloud.ibm.com/v1"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := ibmlogsrouteropenapi30v0.ConstructServiceURL(providedUrlVariables)
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTenants with error: Operation response processing error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmlogsrouteropenapi30v0.ListTenantsOptions)
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}]}`)
				}))
			})
			It(`Invoke ListTenants successfully with retries`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmlogsrouteropenapi30v0.ListTenantsOptions)
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmLogsRouterOpenApi30Service.ListTenantsWithContext(ctx, listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmLogsRouterOpenApi30Service.DisableRetries()
				result, response, operationErr := ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmLogsRouterOpenApi30Service.ListTenantsWithContext(ctx, listTenantsOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tenants": [{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}]}`)
				}))
			})
			It(`Invoke ListTenants successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmLogsRouterOpenApi30Service.ListTenants(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmlogsrouteropenapi30v0.ListTenantsOptions)
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTenants with error: Operation request error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmlogsrouteropenapi30v0.ListTenantsOptions)
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := new(ibmlogsrouteropenapi30v0.ListTenantsOptions)
				listTenantsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptionsModel)
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTenant with error: Operation response processing error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				createTenantOptionsModel.TargetType = core.StringPtr("logdna")
				createTenantOptionsModel.TargetHost = core.StringPtr("www.example.com")
				createTenantOptionsModel.TargetPort = core.Int64Ptr(int64(38))
				createTenantOptionsModel.AccessCredential = core.StringPtr("testString")
				createTenantOptionsModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke CreateTenant successfully with retries`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				createTenantOptionsModel.TargetType = core.StringPtr("logdna")
				createTenantOptionsModel.TargetHost = core.StringPtr("www.example.com")
				createTenantOptionsModel.TargetPort = core.Int64Ptr(int64(38))
				createTenantOptionsModel.AccessCredential = core.StringPtr("testString")
				createTenantOptionsModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmLogsRouterOpenApi30Service.CreateTenantWithContext(ctx, createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmLogsRouterOpenApi30Service.DisableRetries()
				result, response, operationErr := ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmLogsRouterOpenApi30Service.CreateTenantWithContext(ctx, createTenantOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke CreateTenant successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmLogsRouterOpenApi30Service.CreateTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				createTenantOptionsModel.TargetType = core.StringPtr("logdna")
				createTenantOptionsModel.TargetHost = core.StringPtr("www.example.com")
				createTenantOptionsModel.TargetPort = core.Int64Ptr(int64(38))
				createTenantOptionsModel.AccessCredential = core.StringPtr("testString")
				createTenantOptionsModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTenant with error: Operation validation and request error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				createTenantOptionsModel.TargetType = core.StringPtr("logdna")
				createTenantOptionsModel.TargetHost = core.StringPtr("www.example.com")
				createTenantOptionsModel.TargetPort = core.Int64Ptr(int64(38))
				createTenantOptionsModel.AccessCredential = core.StringPtr("testString")
				createTenantOptionsModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTenantOptions model with no property values
				createTenantOptionsModelNew := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModelNew)
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
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsModel := new(ibmlogsrouteropenapi30v0.CreateTenantOptions)
				createTenantOptionsModel.TargetType = core.StringPtr("logdna")
				createTenantOptionsModel.TargetHost = core.StringPtr("www.example.com")
				createTenantOptionsModel.TargetPort = core.Int64Ptr(int64(38))
				createTenantOptionsModel.AccessCredential = core.StringPtr("testString")
				createTenantOptionsModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptionsModel)
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
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTenantDetail with error: Operation response processing error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully with retries`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmLogsRouterOpenApi30Service.DisableRetries()
				result, response, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmLogsRouterOpenApi30Service.GetTenantDetailWithContext(ctx, getTenantDetailOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke GetTenantDetail successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetail(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTenantDetail with error: Operation validation and request error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTenantDetailOptions model with no property values
				getTenantDetailOptionsModelNew := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModelNew)
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
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the GetTenantDetailOptions model
				getTenantDetailOptionsModel := new(ibmlogsrouteropenapi30v0.GetTenantDetailOptions)
				getTenantDetailOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptionsModel)
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
	Describe(`DeleteTenant(deleteTenantOptions *DeleteTenantOptions) - Operation response error`, func() {
		deleteTenantPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTenantPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTenant with error: Operation response processing error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTenant(deleteTenantOptions *DeleteTenantOptions)`, func() {
		deleteTenantPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTenantPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": 6, "message": "Message"}`)
				}))
			})
			It(`Invoke DeleteTenant successfully with retries`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenantWithContext(ctx, deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmLogsRouterOpenApi30Service.DisableRetries()
				result, response, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmLogsRouterOpenApi30Service.DeleteTenantWithContext(ctx, deleteTenantOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTenantPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status": 6, "message": "Message"}`)
				}))
			})
			It(`Invoke DeleteTenant successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenant(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTenant with error: Operation validation and request error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteTenantOptions model with no property values
				deleteTenantOptionsModelNew := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModelNew)
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
			It(`Invoke DeleteTenant successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the DeleteTenantOptions model
				deleteTenantOptionsModel := new(ibmlogsrouteropenapi30v0.DeleteTenantOptions)
				deleteTenantOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptionsModel)
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
		updateTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTarget with error: Operation response processing error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the TenantDetailsResponsePatch model
				tenantDetailsResponsePatchModel := new(ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch)
				tenantDetailsResponsePatchModel.TargetHost = core.StringPtr("www.example.com")
				tenantDetailsResponsePatchModel.TargetPort = core.Int64Ptr(int64(38))
				tenantDetailsResponsePatchModel.AccessCredential = core.StringPtr("testString")
				tenantDetailsResponsePatchModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TenantDetailsResponsePatch = tenantDetailsResponsePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
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
		updateTargetPath := "/tenants/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully with retries`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
				ibmLogsRouterOpenApi30Service.EnableRetries(0, 0)

				// Construct an instance of the TenantDetailsResponsePatch model
				tenantDetailsResponsePatchModel := new(ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch)
				tenantDetailsResponsePatchModel.TargetHost = core.StringPtr("www.example.com")
				tenantDetailsResponsePatchModel.TargetPort = core.Int64Ptr(int64(38))
				tenantDetailsResponsePatchModel.AccessCredential = core.StringPtr("testString")
				tenantDetailsResponsePatchModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TenantDetailsResponsePatch = tenantDetailsResponsePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmLogsRouterOpenApi30Service.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmLogsRouterOpenApi30Service.DisableRetries()
				result, response, operationErr := ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmLogsRouterOpenApi30Service.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", "account_id": "AccountID", "target_type": "logdna", "target_host": "www.example.com", "target_port": 10, "target_instance_crn": "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::", "created_at": "2023-10-20T18:30:00.143156Z", "updated_at": "2023-10-20T18:30:00.143156Z"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmLogsRouterOpenApi30Service.UpdateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TenantDetailsResponsePatch model
				tenantDetailsResponsePatchModel := new(ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch)
				tenantDetailsResponsePatchModel.TargetHost = core.StringPtr("www.example.com")
				tenantDetailsResponsePatchModel.TargetPort = core.Int64Ptr(int64(38))
				tenantDetailsResponsePatchModel.AccessCredential = core.StringPtr("testString")
				tenantDetailsResponsePatchModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TenantDetailsResponsePatch = tenantDetailsResponsePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTarget with error: Operation validation and request error`, func() {
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the TenantDetailsResponsePatch model
				tenantDetailsResponsePatchModel := new(ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch)
				tenantDetailsResponsePatchModel.TargetHost = core.StringPtr("www.example.com")
				tenantDetailsResponsePatchModel.TargetPort = core.Int64Ptr(int64(38))
				tenantDetailsResponsePatchModel.AccessCredential = core.StringPtr("testString")
				tenantDetailsResponsePatchModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TenantDetailsResponsePatch = tenantDetailsResponsePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmLogsRouterOpenApi30Service.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTargetOptions model with no property values
				updateTargetOptionsModelNew := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModelNew)
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
				ibmLogsRouterOpenApi30Service, serviceErr := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())

				// Construct an instance of the TenantDetailsResponsePatch model
				tenantDetailsResponsePatchModel := new(ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch)
				tenantDetailsResponsePatchModel.TargetHost = core.StringPtr("www.example.com")
				tenantDetailsResponsePatchModel.TargetPort = core.Int64Ptr(int64(38))
				tenantDetailsResponsePatchModel.AccessCredential = core.StringPtr("testString")
				tenantDetailsResponsePatchModel.TargetInstanceCrn = core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(ibmlogsrouteropenapi30v0.UpdateTargetOptions)
				updateTargetOptionsModel.TenantID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateTargetOptionsModel.TenantDetailsResponsePatch = tenantDetailsResponsePatchModelAsPatch
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptionsModel)
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
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmLogsRouterOpenApi30Service, _ := ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0(&ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{
				URL:           "http://ibmlogsrouteropenapi30v0modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateTenantOptions successfully`, func() {
				// Construct an instance of the CreateTenantOptions model
				createTenantOptionsTargetType := "logdna"
				createTenantOptionsTargetHost := "www.example.com"
				createTenantOptionsTargetPort := int64(38)
				createTenantOptionsAccessCredential := "testString"
				createTenantOptionsTargetInstanceCrn := "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
				createTenantOptionsModel := ibmLogsRouterOpenApi30Service.NewCreateTenantOptions(createTenantOptionsTargetType, createTenantOptionsTargetHost, createTenantOptionsTargetPort, createTenantOptionsAccessCredential, createTenantOptionsTargetInstanceCrn)
				createTenantOptionsModel.SetTargetType("logdna")
				createTenantOptionsModel.SetTargetHost("www.example.com")
				createTenantOptionsModel.SetTargetPort(int64(38))
				createTenantOptionsModel.SetAccessCredential("testString")
				createTenantOptionsModel.SetTargetInstanceCrn("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")
				createTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTenantOptionsModel).ToNot(BeNil())
				Expect(createTenantOptionsModel.TargetType).To(Equal(core.StringPtr("logdna")))
				Expect(createTenantOptionsModel.TargetHost).To(Equal(core.StringPtr("www.example.com")))
				Expect(createTenantOptionsModel.TargetPort).To(Equal(core.Int64Ptr(int64(38))))
				Expect(createTenantOptionsModel.AccessCredential).To(Equal(core.StringPtr("testString")))
				Expect(createTenantOptionsModel.TargetInstanceCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::")))
				Expect(createTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTenantOptions successfully`, func() {
				// Construct an instance of the DeleteTenantOptions model
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteTenantOptionsModel := ibmLogsRouterOpenApi30Service.NewDeleteTenantOptions(tenantID)
				deleteTenantOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deleteTenantOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTenantOptionsModel).ToNot(BeNil())
				Expect(deleteTenantOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deleteTenantOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTenantDetailOptions successfully`, func() {
				// Construct an instance of the GetTenantDetailOptions model
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getTenantDetailOptionsModel := ibmLogsRouterOpenApi30Service.NewGetTenantDetailOptions(tenantID)
				getTenantDetailOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getTenantDetailOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTenantDetailOptionsModel).ToNot(BeNil())
				Expect(getTenantDetailOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getTenantDetailOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTenantsOptions successfully`, func() {
				// Construct an instance of the ListTenantsOptions model
				listTenantsOptionsModel := ibmLogsRouterOpenApi30Service.NewListTenantsOptions()
				listTenantsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTenantsOptionsModel).ToNot(BeNil())
				Expect(listTenantsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTargetOptions successfully`, func() {
				// Construct an instance of the UpdateTargetOptions model
				tenantID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				tenantDetailsResponsePatch := map[string]interface{}{"anyKey": "anyValue"}
				updateTargetOptionsModel := ibmLogsRouterOpenApi30Service.NewUpdateTargetOptions(tenantID, tenantDetailsResponsePatch)
				updateTargetOptionsModel.SetTenantID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateTargetOptionsModel.SetTenantDetailsResponsePatch(map[string]interface{}{"anyKey": "anyValue"})
				updateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTargetOptionsModel).ToNot(BeNil())
				Expect(updateTargetOptionsModel.TenantID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateTargetOptionsModel.TenantDetailsResponsePatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
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

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
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
