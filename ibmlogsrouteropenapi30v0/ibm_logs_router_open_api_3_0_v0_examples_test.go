//go:build examples
// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM logs-router - OpenAPI 3.0 service.
//
// The following configuration properties are assumed to be defined:
// IBM_LOGS_ROUTER_OPEN_API_3_0_URL=<service base url>
// IBM_LOGS_ROUTER_OPEN_API_3_0_AUTH_TYPE=iam
// IBM_LOGS_ROUTER_OPEN_API_3_0_APIKEY=<IAM apikey>
// IBM_LOGS_ROUTER_OPEN_API_3_0_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`IbmLogsRouterOpenApi30V0 Examples Tests`, func() {

	const externalConfigFile = "../ibm_logs_router_open_api30_v0.env"

	var (
		ibmLogsRouterOpenApi30Service *ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0
		config                        map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmlogsrouteropenapi30v0.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			ibmLogsRouterOpenApi30ServiceOptions := &ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{}

			ibmLogsRouterOpenApi30Service, err = ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(ibmLogsRouterOpenApi30ServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmLogsRouterOpenApi30Service).ToNot(BeNil())
		})
	})

	Describe(`IbmLogsRouterOpenApi30V0 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTenants request example`, func() {
			fmt.Println("\nListTenants() result:")
			// begin-list_tenants

			listTenantsOptions := ibmLogsRouterOpenApi30Service.NewListTenantsOptions()

			tenantDetailsResponseCollection, response, err := ibmLogsRouterOpenApi30Service.ListTenants(listTenantsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantDetailsResponseCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_tenants

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponseCollection).ToNot(BeNil())
		})
		It(`CreateTenant request example`, func() {
			fmt.Println("\nCreateTenant() result:")
			// begin-create_tenant

			createTenantOptions := ibmLogsRouterOpenApi30Service.NewCreateTenantOptions(
				"logdna",
				"www.example.com",
				int64(38),
				"testString",
				"crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::",
			)

			tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Service.CreateTenant(createTenantOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantDetailsResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_tenant

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
		It(`GetTenantDetail request example`, func() {
			fmt.Println("\nGetTenantDetail() result:")
			// begin-get_tenant_detail

			getTenantDetailOptions := ibmLogsRouterOpenApi30Service.NewGetTenantDetailOptions(
				CreateMockUUID("f3a466c9-c4db-4eee-95cc-ba82db58e2b5"),
			)

			tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Service.GetTenantDetail(getTenantDetailOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantDetailsResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_tenant_detail

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
		It(`UpdateTarget request example`, func() {
			fmt.Println("\nUpdateTarget() result:")
			// begin-update_target

			tenantDetailsResponsePatchModel := &ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch{}
			tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTargetOptions := ibmLogsRouterOpenApi30Service.NewUpdateTargetOptions(
				CreateMockUUID("f3a466c9-c4db-4eee-95cc-ba82db58e2b5"),
				tenantDetailsResponsePatchModelAsPatch,
			)

			tenantDetailsResponse, response, err := ibmLogsRouterOpenApi30Service.UpdateTarget(updateTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantDetailsResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
		It(`DeleteTenant request example`, func() {
			fmt.Println("\nDeleteTenant() result:")
			// begin-delete_tenant

			deleteTenantOptions := ibmLogsRouterOpenApi30Service.NewDeleteTenantOptions(
				CreateMockUUID("f3a466c9-c4db-4eee-95cc-ba82db58e2b5"),
			)

			tenantDeleteResponse, response, err := ibmLogsRouterOpenApi30Service.DeleteTenant(deleteTenantOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_tenant

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDeleteResponse).ToNot(BeNil())
		})
	})
})
