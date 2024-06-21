//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM Cloud Logs Routing service.
//
// The following configuration properties are assumed to be defined:
// IBM_CLOUD_LOGS_ROUTING_URL=<service base url>
// IBM_CLOUD_LOGS_ROUTING_AUTH_TYPE=iam
// IBM_CLOUD_LOGS_ROUTING_APIKEY=<IAM apikey>
// IBM_CLOUD_LOGS_ROUTING_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`IBMCloudLogsRoutingV0 Examples Tests`, func() {

	const externalConfigFile = "../ibm_cloud_logs_routing_v0.env"

	var (
		ibmCloudLogsRoutingService *ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0
		config       map[string]string
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
			config, err = core.GetServiceProperties(ibmcloudlogsroutingv0.DefaultServiceName)
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

			ibmCloudLogsRoutingServiceOptions := &ibmcloudlogsroutingv0.IBMCloudLogsRoutingV0Options{}

			ibmCloudLogsRoutingService, err = ibmcloudlogsroutingv0.NewIBMCloudLogsRoutingV0UsingExternalConfig(ibmCloudLogsRoutingServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmCloudLogsRoutingService).ToNot(BeNil())
		})
	})

	Describe(`IBMCloudLogsRoutingV0 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTenants request example`, func() {
			fmt.Println("\nListTenants() result:")
			// begin-list_tenants

			listTenantsOptions := ibmCloudLogsRoutingService.NewListTenantsOptions(
				"testString",
			)

			tenantCollection, response, err := ibmCloudLogsRoutingService.ListTenants(listTenantsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenantCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_tenants

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenantCollection).ToNot(BeNil())
		})
		It(`CreateTenant request example`, func() {
			fmt.Println("\nCreateTenant() result:")
			// begin-create_tenant

			targetTypePrototypeModel := &ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype{
				LogSinkCRN: core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"),
				Name: core.StringPtr("my-log-sink"),
			}

			createTenantOptions := ibmCloudLogsRoutingService.NewCreateTenantOptions(
				"testString",
				"my-logging-tenant",
				[]ibmcloudlogsroutingv0.TargetTypePrototypeIntf{targetTypePrototypeModel},
			)

			tenant, response, err := ibmCloudLogsRoutingService.CreateTenant(createTenantOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenant, "", "  ")
			fmt.Println(string(b))

			// end-create_tenant

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(tenant).ToNot(BeNil())
		})
		It(`GetTenantDetail request example`, func() {
			fmt.Println("\nGetTenantDetail() result:")
			// begin-get_tenant_detail

			getTenantDetailOptions := ibmCloudLogsRoutingService.NewGetTenantDetailOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			tenant, response, err := ibmCloudLogsRoutingService.GetTenantDetail(getTenantDetailOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenant, "", "  ")
			fmt.Println(string(b))

			// end-get_tenant_detail

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenant).ToNot(BeNil())
		})
		It(`UpdateTenant request example`, func() {
			fmt.Println("\nUpdateTenant() result:")
			// begin-update_tenant

			tenantPatchModel := &ibmcloudlogsroutingv0.TenantPatch{
			}
			tenantPatchModelAsPatch, asPatchErr := tenantPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTenantOptions := ibmCloudLogsRoutingService.NewUpdateTenantOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				"testString",
				tenantPatchModelAsPatch,
			)

			tenant, response, err := ibmCloudLogsRoutingService.UpdateTenant(updateTenantOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tenant, "", "  ")
			fmt.Println(string(b))

			// end-update_tenant

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tenant).ToNot(BeNil())
		})
		It(`ListTenantTargets request example`, func() {
			fmt.Println("\nListTenantTargets() result:")
			// begin-list_tenant_targets

			listTenantTargetsOptions := ibmCloudLogsRoutingService.NewListTenantTargetsOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			targetTypeCollection, response, err := ibmCloudLogsRoutingService.ListTenantTargets(listTenantTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetTypeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_tenant_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetTypeCollection).ToNot(BeNil())
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			targetTypePrototypeModel := &ibmcloudlogsroutingv0.TargetTypePrototypeTargetTypeLogDnaPrototype{
				LogSinkCRN: core.StringPtr("crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"),
				Name: core.StringPtr("my-log-sink"),
			}

			createTargetOptions := ibmCloudLogsRoutingService.NewCreateTargetOptions(
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				"testString",
				targetTypePrototypeModel,
			)

			targetType, response, err := ibmCloudLogsRoutingService.CreateTarget(createTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetType, "", "  ")
			fmt.Println(string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(targetType).ToNot(BeNil())
		})
		It(`GetTenantTargetDetails request example`, func() {
			fmt.Println("\nGetTenantTargetDetails() result:")
			// begin-get_tenant_target_details

			getTenantTargetDetailsOptions := ibmCloudLogsRoutingService.NewGetTenantTargetDetailsOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			targetType, response, err := ibmCloudLogsRoutingService.GetTenantTargetDetails(getTenantTargetDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetType, "", "  ")
			fmt.Println(string(b))

			// end-get_tenant_target_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetType).ToNot(BeNil())
		})
		It(`UpdateTarget request example`, func() {
			fmt.Println("\nUpdateTarget() result:")
			// begin-update_target

			targetTypePatchModel := &ibmcloudlogsroutingv0.TargetTypePatchLogDna{
			}
			targetTypePatchModelAsPatch, asPatchErr := targetTypePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateTargetOptions := ibmCloudLogsRoutingService.NewUpdateTargetOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				"testString",
				targetTypePatchModelAsPatch,
			)

			targetType, response, err := ibmCloudLogsRoutingService.UpdateTarget(updateTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetType, "", "  ")
			fmt.Println(string(b))

			// end-update_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetType).ToNot(BeNil())
		})
		It(`DeleteTenant request example`, func() {
			// begin-delete_tenant

			deleteTenantOptions := ibmCloudLogsRoutingService.NewDeleteTenantOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			response, err := ibmCloudLogsRoutingService.DeleteTenant(deleteTenantOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTenant(): %d\n", response.StatusCode)
			}

			// end-delete_tenant

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTarget request example`, func() {
			// begin-delete_target

			deleteTargetOptions := ibmCloudLogsRoutingService.NewDeleteTargetOptions(
				"testString",
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
				CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"),
			)

			response, err := ibmCloudLogsRoutingService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTarget(): %d\n", response.StatusCode)
			}

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
