// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/logs-router-go-sdk/ibmlogsrouteropenapi30v0"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmcloudlogsroutingv0 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmCloudLogsRoutingV0 Integration Tests`, func() {
	const externalConfigFile = "../ibm_cloud_logs_routing_v0.env"

	var (
		err          error
		ibmCloudLogsRoutingService *ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmlogsrouteropenapi30v0.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			ibmCloudLogsRoutingServiceOptions := &ibmlogsrouteropenapi30v0.IbmLogsRouterOpenApi30V0Options{}

			ibmCloudLogsRoutingService, err = ibmlogsrouteropenapi30v0.NewIbmLogsRouterOpenApi30V0UsingExternalConfig(ibmCloudLogsRoutingServiceOptions)
			Expect(err).To(BeNil())
			Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
			Expect(ibmCloudLogsRoutingService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			ibmCloudLogsRoutingService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListTenants - List of tenants`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTenants(listTenantsOptions *ListTenantsOptions)`, func() {
			listTenantsOptions := &ibmlogsrouteropenapi30v0.ListTenantsOptions{
			}

			tenantDetailsResponseCollection, response, err := ibmCloudLogsRoutingService.ListTenants(listTenantsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponseCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateTenant - Create (onboard) a new tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTenant(createTenantOptions *CreateTenantOptions)`, func() {
			createTenantOptions := &ibmlogsrouteropenapi30v0.CreateTenantOptions{
				TargetType: core.StringPtr("logdna"),
				TargetHost: core.StringPtr("www.example.com"),
				TargetPort: core.Int64Ptr(int64(38)),
				AccessCredential: core.StringPtr("testString"),
				TargetInstanceCrn: core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"),
			}

			tenantDetailsResponse, response, err := ibmCloudLogsRoutingService.CreateTenant(createTenantOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
	})

	Describe(`GetTenantDetail - Details of a tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTenantDetail(getTenantDetailOptions *GetTenantDetailOptions)`, func() {
			getTenantDetailOptions := &ibmlogsrouteropenapi30v0.GetTenantDetailOptions{
				TenantID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			tenantDetailsResponse, response, err := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateTarget - Update the target of a tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
			tenantDetailsResponsePatchModel := &ibmlogsrouteropenapi30v0.TenantDetailsResponsePatch{
				TargetHost: core.StringPtr("www.example.com"),
				TargetPort: core.Int64Ptr(int64(38)),
				AccessCredential: core.StringPtr("testString"),
				TargetInstanceCrn: core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"),
			}
			tenantDetailsResponsePatchModelAsPatch, asPatchErr := tenantDetailsResponsePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTargetOptions := &ibmlogsrouteropenapi30v0.UpdateTargetOptions{
				TenantID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				TenantDetailsResponsePatch: tenantDetailsResponsePatchModelAsPatch,
			}

			tenantDetailsResponse, response, err := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDetailsResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteTenant - Delete (offboard) a tenant`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTenant(deleteTenantOptions *DeleteTenantOptions)`, func() {
			deleteTenantOptions := &ibmlogsrouteropenapi30v0.DeleteTenantOptions{
				TenantID: CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			}

			tenantDeleteResponse, response, err := ibmCloudLogsRoutingService.DeleteTenant(deleteTenantOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantDeleteResponse).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
