/**
 * (C) Copyright IBM Corp. 2020.
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

package zonessettingsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/networking-go-sdk/zonessettingsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`ZonesSettingsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(zonesSettingsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(zonesSettingsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
				URL: "https://zonessettingsv1/api",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(zonesSettingsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{})
			Expect(zonesSettingsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_SETTINGS_URL": "https://zonessettingsv1/api",
				"ZONES_SETTINGS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1UsingExternalConfig(&zonessettingsv1.ZonesSettingsV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(zonesSettingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := zonesSettingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesSettingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesSettingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesSettingsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1UsingExternalConfig(&zonessettingsv1.ZonesSettingsV1Options{
					URL: "https://testService/api",
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(zonesSettingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zonesSettingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesSettingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesSettingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesSettingsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1UsingExternalConfig(&zonessettingsv1.ZonesSettingsV1Options{
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := zonesSettingsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zonesSettingsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zonesSettingsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zonesSettingsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zonesSettingsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_SETTINGS_URL": "https://zonessettingsv1/api",
				"ZONES_SETTINGS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1UsingExternalConfig(&zonessettingsv1.ZonesSettingsV1Options{
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zonesSettingsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONES_SETTINGS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1UsingExternalConfig(&zonessettingsv1.ZonesSettingsV1Options{
				URL: "{BAD_URL_STRING",
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zonesSettingsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = zonessettingsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetZoneDnssec(getZoneDnssecOptions *GetZoneDnssecOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneDnssecPath := "/v1/testString/zones/testString/dnssec"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneDnssecPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneDnssec with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetZoneDnssecOptions model
				getZoneDnssecOptionsModel := new(zonessettingsv1.GetZoneDnssecOptions)
				getZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetZoneDnssec(getZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetZoneDnssec(getZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetZoneDnssec(getZoneDnssecOptions *GetZoneDnssecOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneDnssecPath := "/v1/testString/zones/testString/dnssec"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneDnssecPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"status": "active", "flags": 257, "algorithm": "13", "key_type": "ECDSAP256SHA256", "digest_type": "2", "digest_algorithm": "SHA256", "digest": "48E939042E82C22542CB377B580DFDC52A361CEFDC72E7F9107E2B6BD9306A45", "ds": "example.com. 3600 IN DS 16953 13 2 248E939042E82C22542CB377B580DFDC52A361CEFDC72E7F9107E2B6BD9306A45", "key_tag": 42, "public_key": "oXiGYrSTO+LSCJ3mohc8EP+CzF9KxBj8/ydXJ22pKuZP3VAC3/Md/k7xZfz470CoRyZJ6gV6vml07IC3d8xqhA=="}}`)
				}))
			})
			It(`Invoke GetZoneDnssec successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetZoneDnssec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneDnssecOptions model
				getZoneDnssecOptionsModel := new(zonessettingsv1.GetZoneDnssecOptions)
				getZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetZoneDnssec(getZoneDnssecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetZoneDnssecWithContext(ctx, getZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetZoneDnssec(getZoneDnssecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetZoneDnssecWithContext(ctx, getZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetZoneDnssec with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetZoneDnssecOptions model
				getZoneDnssecOptionsModel := new(zonessettingsv1.GetZoneDnssecOptions)
				getZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetZoneDnssec(getZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneDnssec(updateZoneDnssecOptions *UpdateZoneDnssecOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneDnssecPath := "/v1/testString/zones/testString/dnssec"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneDnssecPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneDnssec with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneDnssecOptions model
				updateZoneDnssecOptionsModel := new(zonessettingsv1.UpdateZoneDnssecOptions)
				updateZoneDnssecOptionsModel.Status = core.StringPtr("active")
				updateZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateZoneDnssec(updateZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateZoneDnssec(updateZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZoneDnssec(updateZoneDnssecOptions *UpdateZoneDnssecOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneDnssecPath := "/v1/testString/zones/testString/dnssec"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneDnssecPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"status": "active", "flags": 257, "algorithm": "13", "key_type": "ECDSAP256SHA256", "digest_type": "2", "digest_algorithm": "SHA256", "digest": "48E939042E82C22542CB377B580DFDC52A361CEFDC72E7F9107E2B6BD9306A45", "ds": "example.com. 3600 IN DS 16953 13 2 248E939042E82C22542CB377B580DFDC52A361CEFDC72E7F9107E2B6BD9306A45", "key_tag": 42, "public_key": "oXiGYrSTO+LSCJ3mohc8EP+CzF9KxBj8/ydXJ22pKuZP3VAC3/Md/k7xZfz470CoRyZJ6gV6vml07IC3d8xqhA=="}}`)
				}))
			})
			It(`Invoke UpdateZoneDnssec successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateZoneDnssec(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneDnssecOptions model
				updateZoneDnssecOptionsModel := new(zonessettingsv1.UpdateZoneDnssecOptions)
				updateZoneDnssecOptionsModel.Status = core.StringPtr("active")
				updateZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateZoneDnssec(updateZoneDnssecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateZoneDnssecWithContext(ctx, updateZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateZoneDnssec(updateZoneDnssecOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateZoneDnssecWithContext(ctx, updateZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateZoneDnssec with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneDnssecOptions model
				updateZoneDnssecOptionsModel := new(zonessettingsv1.UpdateZoneDnssecOptions)
				updateZoneDnssecOptionsModel.Status = core.StringPtr("active")
				updateZoneDnssecOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateZoneDnssec(updateZoneDnssecOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetZoneCnameFlattening(getZoneCnameFlatteningOptions *GetZoneCnameFlatteningOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCnameFlatteningPath := "/v1/testString/zones/testString/settings/cname_flattening"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneCnameFlatteningPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetZoneCnameFlattening with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetZoneCnameFlatteningOptions model
				getZoneCnameFlatteningOptionsModel := new(zonessettingsv1.GetZoneCnameFlatteningOptions)
				getZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetZoneCnameFlattening(getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetZoneCnameFlattening(getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetZoneCnameFlattening(getZoneCnameFlatteningOptions *GetZoneCnameFlatteningOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getZoneCnameFlatteningPath := "/v1/testString/zones/testString/settings/cname_flattening"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getZoneCnameFlatteningPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "cname_flattening", "value": "flatten_all", "modified_on": "2019-01-01T12:00:00", "editable": true}}`)
				}))
			})
			It(`Invoke GetZoneCnameFlattening successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetZoneCnameFlattening(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetZoneCnameFlatteningOptions model
				getZoneCnameFlatteningOptionsModel := new(zonessettingsv1.GetZoneCnameFlatteningOptions)
				getZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetZoneCnameFlattening(getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetZoneCnameFlatteningWithContext(ctx, getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetZoneCnameFlattening(getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetZoneCnameFlatteningWithContext(ctx, getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetZoneCnameFlattening with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetZoneCnameFlatteningOptions model
				getZoneCnameFlatteningOptionsModel := new(zonessettingsv1.GetZoneCnameFlatteningOptions)
				getZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetZoneCnameFlattening(getZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptions *UpdateZoneCnameFlatteningOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCnameFlatteningPath := "/v1/testString/zones/testString/settings/cname_flattening"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneCnameFlatteningPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateZoneCnameFlattening with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCnameFlatteningOptions model
				updateZoneCnameFlatteningOptionsModel := new(zonessettingsv1.UpdateZoneCnameFlatteningOptions)
				updateZoneCnameFlatteningOptionsModel.Value = core.StringPtr("flatten_all")
				updateZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptions *UpdateZoneCnameFlatteningOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateZoneCnameFlatteningPath := "/v1/testString/zones/testString/settings/cname_flattening"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateZoneCnameFlatteningPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "cname_flattening", "value": "flatten_all", "modified_on": "2019-01-01T12:00:00", "editable": true}}`)
				}))
			})
			It(`Invoke UpdateZoneCnameFlattening successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateZoneCnameFlattening(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateZoneCnameFlatteningOptions model
				updateZoneCnameFlatteningOptionsModel := new(zonessettingsv1.UpdateZoneCnameFlatteningOptions)
				updateZoneCnameFlatteningOptionsModel.Value = core.StringPtr("flatten_all")
				updateZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateZoneCnameFlatteningWithContext(ctx, updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateZoneCnameFlatteningWithContext(ctx, updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateZoneCnameFlattening with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateZoneCnameFlatteningOptions model
				updateZoneCnameFlatteningOptionsModel := new(zonessettingsv1.UpdateZoneCnameFlatteningOptions)
				updateZoneCnameFlatteningOptionsModel.Value = core.StringPtr("flatten_all")
				updateZoneCnameFlatteningOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateZoneCnameFlattening(updateZoneCnameFlatteningOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOpportunisticEncryption(getOpportunisticEncryptionOptions *GetOpportunisticEncryptionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getOpportunisticEncryptionPath := "/v1/testString/zones/testString/settings/opportunistic_encryption"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOpportunisticEncryptionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOpportunisticEncryption with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetOpportunisticEncryptionOptions model
				getOpportunisticEncryptionOptionsModel := new(zonessettingsv1.GetOpportunisticEncryptionOptions)
				getOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetOpportunisticEncryption(getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetOpportunisticEncryption(getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetOpportunisticEncryption(getOpportunisticEncryptionOptions *GetOpportunisticEncryptionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getOpportunisticEncryptionPath := "/v1/testString/zones/testString/settings/opportunistic_encryption"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOpportunisticEncryptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "opportunistic_encryption", "value": "off", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetOpportunisticEncryption successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetOpportunisticEncryption(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOpportunisticEncryptionOptions model
				getOpportunisticEncryptionOptionsModel := new(zonessettingsv1.GetOpportunisticEncryptionOptions)
				getOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetOpportunisticEncryption(getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetOpportunisticEncryptionWithContext(ctx, getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetOpportunisticEncryption(getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetOpportunisticEncryptionWithContext(ctx, getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetOpportunisticEncryption with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetOpportunisticEncryptionOptions model
				getOpportunisticEncryptionOptionsModel := new(zonessettingsv1.GetOpportunisticEncryptionOptions)
				getOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetOpportunisticEncryption(getOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptions *UpdateOpportunisticEncryptionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateOpportunisticEncryptionPath := "/v1/testString/zones/testString/settings/opportunistic_encryption"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOpportunisticEncryptionPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOpportunisticEncryption with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateOpportunisticEncryptionOptions model
				updateOpportunisticEncryptionOptionsModel := new(zonessettingsv1.UpdateOpportunisticEncryptionOptions)
				updateOpportunisticEncryptionOptionsModel.Value = core.StringPtr("false")
				updateOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptions *UpdateOpportunisticEncryptionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateOpportunisticEncryptionPath := "/v1/testString/zones/testString/settings/opportunistic_encryption"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateOpportunisticEncryptionPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "opportunistic_encryption", "value": "off", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateOpportunisticEncryption successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateOpportunisticEncryption(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateOpportunisticEncryptionOptions model
				updateOpportunisticEncryptionOptionsModel := new(zonessettingsv1.UpdateOpportunisticEncryptionOptions)
				updateOpportunisticEncryptionOptionsModel.Value = core.StringPtr("false")
				updateOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateOpportunisticEncryptionWithContext(ctx, updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateOpportunisticEncryptionWithContext(ctx, updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateOpportunisticEncryption with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateOpportunisticEncryptionOptions model
				updateOpportunisticEncryptionOptionsModel := new(zonessettingsv1.UpdateOpportunisticEncryptionOptions)
				updateOpportunisticEncryptionOptionsModel.Value = core.StringPtr("false")
				updateOpportunisticEncryptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateOpportunisticEncryption(updateOpportunisticEncryptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChallengeTTL(getChallengeTtlOptions *GetChallengeTtlOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getChallengeTTLPath := "/v1/testString/zones/testString/settings/challenge_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChallengeTTLPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChallengeTTL with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetChallengeTtlOptions model
				getChallengeTtlOptionsModel := new(zonessettingsv1.GetChallengeTtlOptions)
				getChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetChallengeTTL(getChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetChallengeTTL(getChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChallengeTTL(getChallengeTtlOptions *GetChallengeTtlOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getChallengeTTLPath := "/v1/testString/zones/testString/settings/challenge_ttl"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChallengeTTLPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "challenge_ttl", "value": 1800, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetChallengeTTL successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetChallengeTTL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChallengeTtlOptions model
				getChallengeTtlOptionsModel := new(zonessettingsv1.GetChallengeTtlOptions)
				getChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetChallengeTTL(getChallengeTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetChallengeTTLWithContext(ctx, getChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetChallengeTTL(getChallengeTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetChallengeTTLWithContext(ctx, getChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetChallengeTTL with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetChallengeTtlOptions model
				getChallengeTtlOptionsModel := new(zonessettingsv1.GetChallengeTtlOptions)
				getChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetChallengeTTL(getChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateChallengeTTL(updateChallengeTtlOptions *UpdateChallengeTtlOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateChallengeTTLPath := "/v1/testString/zones/testString/settings/challenge_ttl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateChallengeTTLPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateChallengeTTL with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateChallengeTtlOptions model
				updateChallengeTtlOptionsModel := new(zonessettingsv1.UpdateChallengeTtlOptions)
				updateChallengeTtlOptionsModel.Value = core.Int64Ptr(int64(1800))
				updateChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateChallengeTTL(updateChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateChallengeTTL(updateChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateChallengeTTL(updateChallengeTtlOptions *UpdateChallengeTtlOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateChallengeTTLPath := "/v1/testString/zones/testString/settings/challenge_ttl"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateChallengeTTLPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "challenge_ttl", "value": 1800, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateChallengeTTL successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateChallengeTTL(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateChallengeTtlOptions model
				updateChallengeTtlOptionsModel := new(zonessettingsv1.UpdateChallengeTtlOptions)
				updateChallengeTtlOptionsModel.Value = core.Int64Ptr(int64(1800))
				updateChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateChallengeTTL(updateChallengeTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateChallengeTTLWithContext(ctx, updateChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateChallengeTTL(updateChallengeTtlOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateChallengeTTLWithContext(ctx, updateChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateChallengeTTL with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateChallengeTtlOptions model
				updateChallengeTtlOptionsModel := new(zonessettingsv1.UpdateChallengeTtlOptions)
				updateChallengeTtlOptionsModel.Value = core.Int64Ptr(int64(1800))
				updateChallengeTtlOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateChallengeTTL(updateChallengeTtlOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptions *GetAutomaticHttpsRewritesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAutomaticHttpsRewritesPath := "/v1/testString/zones/testString/settings/automatic_https_rewrites"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutomaticHttpsRewritesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAutomaticHttpsRewrites with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetAutomaticHttpsRewritesOptions model
				getAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.GetAutomaticHttpsRewritesOptions)
				getAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptions *GetAutomaticHttpsRewritesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAutomaticHttpsRewritesPath := "/v1/testString/zones/testString/settings/automatic_https_rewrites"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutomaticHttpsRewritesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "automatic_https_rewrites", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetAutomaticHttpsRewrites successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetAutomaticHttpsRewrites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAutomaticHttpsRewritesOptions model
				getAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.GetAutomaticHttpsRewritesOptions)
				getAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetAutomaticHttpsRewritesWithContext(ctx, getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetAutomaticHttpsRewritesWithContext(ctx, getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAutomaticHttpsRewrites with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetAutomaticHttpsRewritesOptions model
				getAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.GetAutomaticHttpsRewritesOptions)
				getAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetAutomaticHttpsRewrites(getAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptions *UpdateAutomaticHttpsRewritesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAutomaticHttpsRewritesPath := "/v1/testString/zones/testString/settings/automatic_https_rewrites"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAutomaticHttpsRewritesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAutomaticHttpsRewrites with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateAutomaticHttpsRewritesOptions model
				updateAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.UpdateAutomaticHttpsRewritesOptions)
				updateAutomaticHttpsRewritesOptionsModel.Value = core.StringPtr("false")
				updateAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptions *UpdateAutomaticHttpsRewritesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAutomaticHttpsRewritesPath := "/v1/testString/zones/testString/settings/automatic_https_rewrites"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAutomaticHttpsRewritesPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "automatic_https_rewrites", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateAutomaticHttpsRewrites successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateAutomaticHttpsRewrites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAutomaticHttpsRewritesOptions model
				updateAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.UpdateAutomaticHttpsRewritesOptions)
				updateAutomaticHttpsRewritesOptionsModel.Value = core.StringPtr("false")
				updateAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateAutomaticHttpsRewritesWithContext(ctx, updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateAutomaticHttpsRewritesWithContext(ctx, updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAutomaticHttpsRewrites with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateAutomaticHttpsRewritesOptions model
				updateAutomaticHttpsRewritesOptionsModel := new(zonessettingsv1.UpdateAutomaticHttpsRewritesOptions)
				updateAutomaticHttpsRewritesOptionsModel.Value = core.StringPtr("false")
				updateAutomaticHttpsRewritesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateAutomaticHttpsRewrites(updateAutomaticHttpsRewritesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTrueClientIp(getTrueClientIpOptions *GetTrueClientIpOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTrueClientIpPath := "/v1/testString/zones/testString/settings/true_client_ip_header"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrueClientIpPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTrueClientIp with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetTrueClientIpOptions model
				getTrueClientIpOptionsModel := new(zonessettingsv1.GetTrueClientIpOptions)
				getTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetTrueClientIp(getTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetTrueClientIp(getTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTrueClientIp(getTrueClientIpOptions *GetTrueClientIpOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTrueClientIpPath := "/v1/testString/zones/testString/settings/true_client_ip_header"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTrueClientIpPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "true_client_ip_header", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetTrueClientIp successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetTrueClientIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTrueClientIpOptions model
				getTrueClientIpOptionsModel := new(zonessettingsv1.GetTrueClientIpOptions)
				getTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetTrueClientIp(getTrueClientIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetTrueClientIpWithContext(ctx, getTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetTrueClientIp(getTrueClientIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetTrueClientIpWithContext(ctx, getTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTrueClientIp with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetTrueClientIpOptions model
				getTrueClientIpOptionsModel := new(zonessettingsv1.GetTrueClientIpOptions)
				getTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetTrueClientIp(getTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTrueClientIp(updateTrueClientIpOptions *UpdateTrueClientIpOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateTrueClientIpPath := "/v1/testString/zones/testString/settings/true_client_ip_header"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrueClientIpPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTrueClientIp with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateTrueClientIpOptions model
				updateTrueClientIpOptionsModel := new(zonessettingsv1.UpdateTrueClientIpOptions)
				updateTrueClientIpOptionsModel.Value = core.StringPtr("true")
				updateTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateTrueClientIp(updateTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateTrueClientIp(updateTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTrueClientIp(updateTrueClientIpOptions *UpdateTrueClientIpOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateTrueClientIpPath := "/v1/testString/zones/testString/settings/true_client_ip_header"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTrueClientIpPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "true_client_ip_header", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateTrueClientIp successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateTrueClientIp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTrueClientIpOptions model
				updateTrueClientIpOptionsModel := new(zonessettingsv1.UpdateTrueClientIpOptions)
				updateTrueClientIpOptionsModel.Value = core.StringPtr("true")
				updateTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateTrueClientIp(updateTrueClientIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateTrueClientIpWithContext(ctx, updateTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateTrueClientIp(updateTrueClientIpOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateTrueClientIpWithContext(ctx, updateTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateTrueClientIp with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateTrueClientIpOptions model
				updateTrueClientIpOptionsModel := new(zonessettingsv1.UpdateTrueClientIpOptions)
				updateTrueClientIpOptionsModel.Value = core.StringPtr("true")
				updateTrueClientIpOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateTrueClientIp(updateTrueClientIpOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAlwaysUseHttps(getAlwaysUseHttpsOptions *GetAlwaysUseHttpsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAlwaysUseHttpsPath := "/v1/testString/zones/testString/settings/always_use_https"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlwaysUseHttpsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAlwaysUseHttps with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetAlwaysUseHttpsOptions model
				getAlwaysUseHttpsOptionsModel := new(zonessettingsv1.GetAlwaysUseHttpsOptions)
				getAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetAlwaysUseHttps(getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetAlwaysUseHttps(getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAlwaysUseHttps(getAlwaysUseHttpsOptions *GetAlwaysUseHttpsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getAlwaysUseHttpsPath := "/v1/testString/zones/testString/settings/always_use_https"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAlwaysUseHttpsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "always_use_https", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetAlwaysUseHttps successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetAlwaysUseHttps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAlwaysUseHttpsOptions model
				getAlwaysUseHttpsOptionsModel := new(zonessettingsv1.GetAlwaysUseHttpsOptions)
				getAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetAlwaysUseHttps(getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetAlwaysUseHttpsWithContext(ctx, getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetAlwaysUseHttps(getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetAlwaysUseHttpsWithContext(ctx, getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAlwaysUseHttps with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetAlwaysUseHttpsOptions model
				getAlwaysUseHttpsOptionsModel := new(zonessettingsv1.GetAlwaysUseHttpsOptions)
				getAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetAlwaysUseHttps(getAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptions *UpdateAlwaysUseHttpsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAlwaysUseHttpsPath := "/v1/testString/zones/testString/settings/always_use_https"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlwaysUseHttpsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAlwaysUseHttps with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateAlwaysUseHttpsOptions model
				updateAlwaysUseHttpsOptionsModel := new(zonessettingsv1.UpdateAlwaysUseHttpsOptions)
				updateAlwaysUseHttpsOptionsModel.Value = core.StringPtr("true")
				updateAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptions *UpdateAlwaysUseHttpsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateAlwaysUseHttpsPath := "/v1/testString/zones/testString/settings/always_use_https"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAlwaysUseHttpsPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "always_use_https", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateAlwaysUseHttps successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateAlwaysUseHttps(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAlwaysUseHttpsOptions model
				updateAlwaysUseHttpsOptionsModel := new(zonessettingsv1.UpdateAlwaysUseHttpsOptions)
				updateAlwaysUseHttpsOptionsModel.Value = core.StringPtr("true")
				updateAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateAlwaysUseHttpsWithContext(ctx, updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateAlwaysUseHttpsWithContext(ctx, updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAlwaysUseHttps with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateAlwaysUseHttpsOptions model
				updateAlwaysUseHttpsOptionsModel := new(zonessettingsv1.UpdateAlwaysUseHttpsOptions)
				updateAlwaysUseHttpsOptionsModel.Value = core.StringPtr("true")
				updateAlwaysUseHttpsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateAlwaysUseHttps(updateAlwaysUseHttpsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetImageSizeOptimization(getImageSizeOptimizationOptions *GetImageSizeOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getImageSizeOptimizationPath := "/v1/testString/zones/testString/settings/image_size_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageSizeOptimizationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetImageSizeOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetImageSizeOptimizationOptions model
				getImageSizeOptimizationOptionsModel := new(zonessettingsv1.GetImageSizeOptimizationOptions)
				getImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetImageSizeOptimization(getImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetImageSizeOptimization(getImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetImageSizeOptimization(getImageSizeOptimizationOptions *GetImageSizeOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getImageSizeOptimizationPath := "/v1/testString/zones/testString/settings/image_size_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageSizeOptimizationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "image_size_optimization", "value": "lossless", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetImageSizeOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetImageSizeOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImageSizeOptimizationOptions model
				getImageSizeOptimizationOptionsModel := new(zonessettingsv1.GetImageSizeOptimizationOptions)
				getImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetImageSizeOptimization(getImageSizeOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetImageSizeOptimizationWithContext(ctx, getImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetImageSizeOptimization(getImageSizeOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetImageSizeOptimizationWithContext(ctx, getImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetImageSizeOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetImageSizeOptimizationOptions model
				getImageSizeOptimizationOptionsModel := new(zonessettingsv1.GetImageSizeOptimizationOptions)
				getImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetImageSizeOptimization(getImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateImageSizeOptimization(updateImageSizeOptimizationOptions *UpdateImageSizeOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateImageSizeOptimizationPath := "/v1/testString/zones/testString/settings/image_size_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateImageSizeOptimizationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateImageSizeOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateImageSizeOptimizationOptions model
				updateImageSizeOptimizationOptionsModel := new(zonessettingsv1.UpdateImageSizeOptimizationOptions)
				updateImageSizeOptimizationOptionsModel.Value = core.StringPtr("lossless")
				updateImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateImageSizeOptimization(updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateImageSizeOptimization(updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateImageSizeOptimization(updateImageSizeOptimizationOptions *UpdateImageSizeOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateImageSizeOptimizationPath := "/v1/testString/zones/testString/settings/image_size_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateImageSizeOptimizationPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "image_size_optimization", "value": "lossless", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateImageSizeOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateImageSizeOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateImageSizeOptimizationOptions model
				updateImageSizeOptimizationOptionsModel := new(zonessettingsv1.UpdateImageSizeOptimizationOptions)
				updateImageSizeOptimizationOptionsModel.Value = core.StringPtr("lossless")
				updateImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateImageSizeOptimization(updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateImageSizeOptimizationWithContext(ctx, updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateImageSizeOptimization(updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateImageSizeOptimizationWithContext(ctx, updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateImageSizeOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateImageSizeOptimizationOptions model
				updateImageSizeOptimizationOptionsModel := new(zonessettingsv1.UpdateImageSizeOptimizationOptions)
				updateImageSizeOptimizationOptionsModel.Value = core.StringPtr("lossless")
				updateImageSizeOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateImageSizeOptimization(updateImageSizeOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetScriptLoadOptimization(getScriptLoadOptimizationOptions *GetScriptLoadOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getScriptLoadOptimizationPath := "/v1/testString/zones/testString/settings/script_load_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScriptLoadOptimizationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetScriptLoadOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetScriptLoadOptimizationOptions model
				getScriptLoadOptimizationOptionsModel := new(zonessettingsv1.GetScriptLoadOptimizationOptions)
				getScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetScriptLoadOptimization(getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetScriptLoadOptimization(getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetScriptLoadOptimization(getScriptLoadOptimizationOptions *GetScriptLoadOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getScriptLoadOptimizationPath := "/v1/testString/zones/testString/settings/script_load_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getScriptLoadOptimizationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "script_load_optimization", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetScriptLoadOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetScriptLoadOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetScriptLoadOptimizationOptions model
				getScriptLoadOptimizationOptionsModel := new(zonessettingsv1.GetScriptLoadOptimizationOptions)
				getScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetScriptLoadOptimization(getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetScriptLoadOptimizationWithContext(ctx, getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetScriptLoadOptimization(getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetScriptLoadOptimizationWithContext(ctx, getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetScriptLoadOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetScriptLoadOptimizationOptions model
				getScriptLoadOptimizationOptionsModel := new(zonessettingsv1.GetScriptLoadOptimizationOptions)
				getScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetScriptLoadOptimization(getScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptions *UpdateScriptLoadOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateScriptLoadOptimizationPath := "/v1/testString/zones/testString/settings/script_load_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateScriptLoadOptimizationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateScriptLoadOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateScriptLoadOptimizationOptions model
				updateScriptLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateScriptLoadOptimizationOptions)
				updateScriptLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptions *UpdateScriptLoadOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateScriptLoadOptimizationPath := "/v1/testString/zones/testString/settings/script_load_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateScriptLoadOptimizationPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "script_load_optimization", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateScriptLoadOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateScriptLoadOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateScriptLoadOptimizationOptions model
				updateScriptLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateScriptLoadOptimizationOptions)
				updateScriptLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateScriptLoadOptimizationWithContext(ctx, updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateScriptLoadOptimizationWithContext(ctx, updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateScriptLoadOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateScriptLoadOptimizationOptions model
				updateScriptLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateScriptLoadOptimizationOptions)
				updateScriptLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateScriptLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateScriptLoadOptimization(updateScriptLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetImageLoadOptimization(getImageLoadOptimizationOptions *GetImageLoadOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getImageLoadOptimizationPath := "/v1/testString/zones/testString/settings/image_load_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageLoadOptimizationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetImageLoadOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetImageLoadOptimizationOptions model
				getImageLoadOptimizationOptionsModel := new(zonessettingsv1.GetImageLoadOptimizationOptions)
				getImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetImageLoadOptimization(getImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetImageLoadOptimization(getImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetImageLoadOptimization(getImageLoadOptimizationOptions *GetImageLoadOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getImageLoadOptimizationPath := "/v1/testString/zones/testString/settings/image_load_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getImageLoadOptimizationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "image_load_optimization", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetImageLoadOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetImageLoadOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImageLoadOptimizationOptions model
				getImageLoadOptimizationOptionsModel := new(zonessettingsv1.GetImageLoadOptimizationOptions)
				getImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetImageLoadOptimization(getImageLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetImageLoadOptimizationWithContext(ctx, getImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetImageLoadOptimization(getImageLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetImageLoadOptimizationWithContext(ctx, getImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetImageLoadOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetImageLoadOptimizationOptions model
				getImageLoadOptimizationOptionsModel := new(zonessettingsv1.GetImageLoadOptimizationOptions)
				getImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetImageLoadOptimization(getImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateImageLoadOptimization(updateImageLoadOptimizationOptions *UpdateImageLoadOptimizationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateImageLoadOptimizationPath := "/v1/testString/zones/testString/settings/image_load_optimization"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateImageLoadOptimizationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateImageLoadOptimization with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateImageLoadOptimizationOptions model
				updateImageLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateImageLoadOptimizationOptions)
				updateImageLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateImageLoadOptimization(updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateImageLoadOptimization(updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateImageLoadOptimization(updateImageLoadOptimizationOptions *UpdateImageLoadOptimizationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateImageLoadOptimizationPath := "/v1/testString/zones/testString/settings/image_load_optimization"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateImageLoadOptimizationPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "image_load_optimization", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateImageLoadOptimization successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateImageLoadOptimization(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateImageLoadOptimizationOptions model
				updateImageLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateImageLoadOptimizationOptions)
				updateImageLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateImageLoadOptimization(updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateImageLoadOptimizationWithContext(ctx, updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateImageLoadOptimization(updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateImageLoadOptimizationWithContext(ctx, updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateImageLoadOptimization with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateImageLoadOptimizationOptions model
				updateImageLoadOptimizationOptionsModel := new(zonessettingsv1.UpdateImageLoadOptimizationOptions)
				updateImageLoadOptimizationOptionsModel.Value = core.StringPtr("true")
				updateImageLoadOptimizationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateImageLoadOptimization(updateImageLoadOptimizationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMinify(getMinifyOptions *GetMinifyOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMinifyPath := "/v1/testString/zones/testString/settings/minify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMinifyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMinify with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMinifyOptions model
				getMinifyOptionsModel := new(zonessettingsv1.GetMinifyOptions)
				getMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetMinify(getMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetMinify(getMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMinify(getMinifyOptions *GetMinifyOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMinifyPath := "/v1/testString/zones/testString/settings/minify"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMinifyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "minify", "value": {"css": "true", "html": "true", "js": "true"}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetMinify successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetMinify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMinifyOptions model
				getMinifyOptionsModel := new(zonessettingsv1.GetMinifyOptions)
				getMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetMinify(getMinifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMinifyWithContext(ctx, getMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetMinify(getMinifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMinifyWithContext(ctx, getMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMinify with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMinifyOptions model
				getMinifyOptionsModel := new(zonessettingsv1.GetMinifyOptions)
				getMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetMinify(getMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMinify(updateMinifyOptions *UpdateMinifyOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMinifyPath := "/v1/testString/zones/testString/settings/minify"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMinifyPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMinify with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the MinifySettingValue model
				minifySettingValueModel := new(zonessettingsv1.MinifySettingValue)
				minifySettingValueModel.Css = core.StringPtr("false")
				minifySettingValueModel.HTML = core.StringPtr("false")
				minifySettingValueModel.Js = core.StringPtr("false")

				// Construct an instance of the UpdateMinifyOptions model
				updateMinifyOptionsModel := new(zonessettingsv1.UpdateMinifyOptions)
				updateMinifyOptionsModel.Value = minifySettingValueModel
				updateMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateMinify(updateMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateMinify(updateMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMinify(updateMinifyOptions *UpdateMinifyOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMinifyPath := "/v1/testString/zones/testString/settings/minify"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMinifyPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "minify", "value": {"css": "true", "html": "true", "js": "true"}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateMinify successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateMinify(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MinifySettingValue model
				minifySettingValueModel := new(zonessettingsv1.MinifySettingValue)
				minifySettingValueModel.Css = core.StringPtr("false")
				minifySettingValueModel.HTML = core.StringPtr("false")
				minifySettingValueModel.Js = core.StringPtr("false")

				// Construct an instance of the UpdateMinifyOptions model
				updateMinifyOptionsModel := new(zonessettingsv1.UpdateMinifyOptions)
				updateMinifyOptionsModel.Value = minifySettingValueModel
				updateMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateMinify(updateMinifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMinifyWithContext(ctx, updateMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateMinify(updateMinifyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMinifyWithContext(ctx, updateMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateMinify with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the MinifySettingValue model
				minifySettingValueModel := new(zonessettingsv1.MinifySettingValue)
				minifySettingValueModel.Css = core.StringPtr("false")
				minifySettingValueModel.HTML = core.StringPtr("false")
				minifySettingValueModel.Js = core.StringPtr("false")

				// Construct an instance of the UpdateMinifyOptions model
				updateMinifyOptionsModel := new(zonessettingsv1.UpdateMinifyOptions)
				updateMinifyOptionsModel.Value = minifySettingValueModel
				updateMinifyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateMinify(updateMinifyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMinTlsVersion(getMinTlsVersionOptions *GetMinTlsVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMinTlsVersionPath := "/v1/testString/zones/testString/settings/min_tls_version"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMinTlsVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMinTlsVersion with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMinTlsVersionOptions model
				getMinTlsVersionOptionsModel := new(zonessettingsv1.GetMinTlsVersionOptions)
				getMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetMinTlsVersion(getMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetMinTlsVersion(getMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMinTlsVersion(getMinTlsVersionOptions *GetMinTlsVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMinTlsVersionPath := "/v1/testString/zones/testString/settings/min_tls_version"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMinTlsVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "min_tls_version", "value": "1.2", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetMinTlsVersion successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetMinTlsVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMinTlsVersionOptions model
				getMinTlsVersionOptionsModel := new(zonessettingsv1.GetMinTlsVersionOptions)
				getMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetMinTlsVersion(getMinTlsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMinTlsVersionWithContext(ctx, getMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetMinTlsVersion(getMinTlsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMinTlsVersionWithContext(ctx, getMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMinTlsVersion with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMinTlsVersionOptions model
				getMinTlsVersionOptionsModel := new(zonessettingsv1.GetMinTlsVersionOptions)
				getMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetMinTlsVersion(getMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMinTlsVersion(updateMinTlsVersionOptions *UpdateMinTlsVersionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMinTlsVersionPath := "/v1/testString/zones/testString/settings/min_tls_version"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMinTlsVersionPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMinTlsVersion with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateMinTlsVersionOptions model
				updateMinTlsVersionOptionsModel := new(zonessettingsv1.UpdateMinTlsVersionOptions)
				updateMinTlsVersionOptionsModel.Value = core.StringPtr("1.2")
				updateMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateMinTlsVersion(updateMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateMinTlsVersion(updateMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMinTlsVersion(updateMinTlsVersionOptions *UpdateMinTlsVersionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMinTlsVersionPath := "/v1/testString/zones/testString/settings/min_tls_version"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMinTlsVersionPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "min_tls_version", "value": "1.2", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateMinTlsVersion successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateMinTlsVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateMinTlsVersionOptions model
				updateMinTlsVersionOptionsModel := new(zonessettingsv1.UpdateMinTlsVersionOptions)
				updateMinTlsVersionOptionsModel.Value = core.StringPtr("1.2")
				updateMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateMinTlsVersion(updateMinTlsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMinTlsVersionWithContext(ctx, updateMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateMinTlsVersion(updateMinTlsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMinTlsVersionWithContext(ctx, updateMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateMinTlsVersion with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateMinTlsVersionOptions model
				updateMinTlsVersionOptionsModel := new(zonessettingsv1.UpdateMinTlsVersionOptions)
				updateMinTlsVersionOptionsModel.Value = core.StringPtr("1.2")
				updateMinTlsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateMinTlsVersion(updateMinTlsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIpGeolocation(getIpGeolocationOptions *GetIpGeolocationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getIpGeolocationPath := "/v1/testString/zones/testString/settings/ip_geolocation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIpGeolocationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIpGeolocation with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetIpGeolocationOptions model
				getIpGeolocationOptionsModel := new(zonessettingsv1.GetIpGeolocationOptions)
				getIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetIpGeolocation(getIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetIpGeolocation(getIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetIpGeolocation(getIpGeolocationOptions *GetIpGeolocationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getIpGeolocationPath := "/v1/testString/zones/testString/settings/ip_geolocation"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIpGeolocationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ip_geolocation", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetIpGeolocation successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetIpGeolocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIpGeolocationOptions model
				getIpGeolocationOptionsModel := new(zonessettingsv1.GetIpGeolocationOptions)
				getIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetIpGeolocation(getIpGeolocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetIpGeolocationWithContext(ctx, getIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetIpGeolocation(getIpGeolocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetIpGeolocationWithContext(ctx, getIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetIpGeolocation with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetIpGeolocationOptions model
				getIpGeolocationOptionsModel := new(zonessettingsv1.GetIpGeolocationOptions)
				getIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetIpGeolocation(getIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIpGeolocation(updateIpGeolocationOptions *UpdateIpGeolocationOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateIpGeolocationPath := "/v1/testString/zones/testString/settings/ip_geolocation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIpGeolocationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateIpGeolocation with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateIpGeolocationOptions model
				updateIpGeolocationOptionsModel := new(zonessettingsv1.UpdateIpGeolocationOptions)
				updateIpGeolocationOptionsModel.Value = core.StringPtr("true")
				updateIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateIpGeolocation(updateIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateIpGeolocation(updateIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateIpGeolocation(updateIpGeolocationOptions *UpdateIpGeolocationOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateIpGeolocationPath := "/v1/testString/zones/testString/settings/ip_geolocation"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIpGeolocationPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ip_geolocation", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateIpGeolocation successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateIpGeolocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateIpGeolocationOptions model
				updateIpGeolocationOptionsModel := new(zonessettingsv1.UpdateIpGeolocationOptions)
				updateIpGeolocationOptionsModel.Value = core.StringPtr("true")
				updateIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateIpGeolocation(updateIpGeolocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateIpGeolocationWithContext(ctx, updateIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateIpGeolocation(updateIpGeolocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateIpGeolocationWithContext(ctx, updateIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateIpGeolocation with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateIpGeolocationOptions model
				updateIpGeolocationOptionsModel := new(zonessettingsv1.UpdateIpGeolocationOptions)
				updateIpGeolocationOptionsModel.Value = core.StringPtr("true")
				updateIpGeolocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateIpGeolocation(updateIpGeolocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetServerSideExclude(getServerSideExcludeOptions *GetServerSideExcludeOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getServerSideExcludePath := "/v1/testString/zones/testString/settings/server_side_exclude"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServerSideExcludePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServerSideExclude with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetServerSideExcludeOptions model
				getServerSideExcludeOptionsModel := new(zonessettingsv1.GetServerSideExcludeOptions)
				getServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetServerSideExclude(getServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetServerSideExclude(getServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetServerSideExclude(getServerSideExcludeOptions *GetServerSideExcludeOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getServerSideExcludePath := "/v1/testString/zones/testString/settings/server_side_exclude"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getServerSideExcludePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "server_side_exclude", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetServerSideExclude successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetServerSideExclude(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServerSideExcludeOptions model
				getServerSideExcludeOptionsModel := new(zonessettingsv1.GetServerSideExcludeOptions)
				getServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetServerSideExclude(getServerSideExcludeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetServerSideExcludeWithContext(ctx, getServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetServerSideExclude(getServerSideExcludeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetServerSideExcludeWithContext(ctx, getServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetServerSideExclude with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetServerSideExcludeOptions model
				getServerSideExcludeOptionsModel := new(zonessettingsv1.GetServerSideExcludeOptions)
				getServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetServerSideExclude(getServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateServerSideExclude(updateServerSideExcludeOptions *UpdateServerSideExcludeOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateServerSideExcludePath := "/v1/testString/zones/testString/settings/server_side_exclude"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServerSideExcludePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateServerSideExclude with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateServerSideExcludeOptions model
				updateServerSideExcludeOptionsModel := new(zonessettingsv1.UpdateServerSideExcludeOptions)
				updateServerSideExcludeOptionsModel.Value = core.StringPtr("true")
				updateServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateServerSideExclude(updateServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateServerSideExclude(updateServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateServerSideExclude(updateServerSideExcludeOptions *UpdateServerSideExcludeOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateServerSideExcludePath := "/v1/testString/zones/testString/settings/server_side_exclude"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateServerSideExcludePath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "server_side_exclude", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateServerSideExclude successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateServerSideExclude(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateServerSideExcludeOptions model
				updateServerSideExcludeOptionsModel := new(zonessettingsv1.UpdateServerSideExcludeOptions)
				updateServerSideExcludeOptionsModel.Value = core.StringPtr("true")
				updateServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateServerSideExclude(updateServerSideExcludeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateServerSideExcludeWithContext(ctx, updateServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateServerSideExclude(updateServerSideExcludeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateServerSideExcludeWithContext(ctx, updateServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateServerSideExclude with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateServerSideExcludeOptions model
				updateServerSideExcludeOptionsModel := new(zonessettingsv1.UpdateServerSideExcludeOptions)
				updateServerSideExcludeOptionsModel.Value = core.StringPtr("true")
				updateServerSideExcludeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateServerSideExclude(updateServerSideExcludeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSecurityHeader(getSecurityHeaderOptions *GetSecurityHeaderOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSecurityHeaderPath := "/v1/testString/zones/testString/settings/security_header"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityHeaderPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSecurityHeader with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetSecurityHeaderOptions model
				getSecurityHeaderOptionsModel := new(zonessettingsv1.GetSecurityHeaderOptions)
				getSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetSecurityHeader(getSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetSecurityHeader(getSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSecurityHeader(getSecurityHeaderOptions *GetSecurityHeaderOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getSecurityHeaderPath := "/v1/testString/zones/testString/settings/security_header"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSecurityHeaderPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "security_header", "value": {"strict_transport_security": {"enabled": true, "max_age": 86400, "include_subdomains": true, "nosniff": true}}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetSecurityHeader successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetSecurityHeader(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSecurityHeaderOptions model
				getSecurityHeaderOptionsModel := new(zonessettingsv1.GetSecurityHeaderOptions)
				getSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetSecurityHeader(getSecurityHeaderOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetSecurityHeaderWithContext(ctx, getSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetSecurityHeader(getSecurityHeaderOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetSecurityHeaderWithContext(ctx, getSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSecurityHeader with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetSecurityHeaderOptions model
				getSecurityHeaderOptionsModel := new(zonessettingsv1.GetSecurityHeaderOptions)
				getSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetSecurityHeader(getSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSecurityHeader(updateSecurityHeaderOptions *UpdateSecurityHeaderOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateSecurityHeaderPath := "/v1/testString/zones/testString/settings/security_header"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSecurityHeaderPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSecurityHeader with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the SecurityHeaderSettingValueStrictTransportSecurity model
				securityHeaderSettingValueStrictTransportSecurityModel := new(zonessettingsv1.SecurityHeaderSettingValueStrictTransportSecurity)
				securityHeaderSettingValueStrictTransportSecurityModel.Enabled = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.MaxAge = core.Int64Ptr(int64(86400))
				securityHeaderSettingValueStrictTransportSecurityModel.IncludeSubdomains = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.Nosniff = core.BoolPtr(true)

				// Construct an instance of the SecurityHeaderSettingValue model
				securityHeaderSettingValueModel := new(zonessettingsv1.SecurityHeaderSettingValue)
				securityHeaderSettingValueModel.StrictTransportSecurity = securityHeaderSettingValueStrictTransportSecurityModel

				// Construct an instance of the UpdateSecurityHeaderOptions model
				updateSecurityHeaderOptionsModel := new(zonessettingsv1.UpdateSecurityHeaderOptions)
				updateSecurityHeaderOptionsModel.Value = securityHeaderSettingValueModel
				updateSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateSecurityHeader(updateSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateSecurityHeader(updateSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSecurityHeader(updateSecurityHeaderOptions *UpdateSecurityHeaderOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateSecurityHeaderPath := "/v1/testString/zones/testString/settings/security_header"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSecurityHeaderPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "security_header", "value": {"strict_transport_security": {"enabled": true, "max_age": 86400, "include_subdomains": true, "nosniff": true}}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateSecurityHeader successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateSecurityHeader(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SecurityHeaderSettingValueStrictTransportSecurity model
				securityHeaderSettingValueStrictTransportSecurityModel := new(zonessettingsv1.SecurityHeaderSettingValueStrictTransportSecurity)
				securityHeaderSettingValueStrictTransportSecurityModel.Enabled = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.MaxAge = core.Int64Ptr(int64(86400))
				securityHeaderSettingValueStrictTransportSecurityModel.IncludeSubdomains = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.Nosniff = core.BoolPtr(true)

				// Construct an instance of the SecurityHeaderSettingValue model
				securityHeaderSettingValueModel := new(zonessettingsv1.SecurityHeaderSettingValue)
				securityHeaderSettingValueModel.StrictTransportSecurity = securityHeaderSettingValueStrictTransportSecurityModel

				// Construct an instance of the UpdateSecurityHeaderOptions model
				updateSecurityHeaderOptionsModel := new(zonessettingsv1.UpdateSecurityHeaderOptions)
				updateSecurityHeaderOptionsModel.Value = securityHeaderSettingValueModel
				updateSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateSecurityHeader(updateSecurityHeaderOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateSecurityHeaderWithContext(ctx, updateSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateSecurityHeader(updateSecurityHeaderOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateSecurityHeaderWithContext(ctx, updateSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateSecurityHeader with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the SecurityHeaderSettingValueStrictTransportSecurity model
				securityHeaderSettingValueStrictTransportSecurityModel := new(zonessettingsv1.SecurityHeaderSettingValueStrictTransportSecurity)
				securityHeaderSettingValueStrictTransportSecurityModel.Enabled = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.MaxAge = core.Int64Ptr(int64(86400))
				securityHeaderSettingValueStrictTransportSecurityModel.IncludeSubdomains = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.Nosniff = core.BoolPtr(true)

				// Construct an instance of the SecurityHeaderSettingValue model
				securityHeaderSettingValueModel := new(zonessettingsv1.SecurityHeaderSettingValue)
				securityHeaderSettingValueModel.StrictTransportSecurity = securityHeaderSettingValueStrictTransportSecurityModel

				// Construct an instance of the UpdateSecurityHeaderOptions model
				updateSecurityHeaderOptionsModel := new(zonessettingsv1.UpdateSecurityHeaderOptions)
				updateSecurityHeaderOptionsModel.Value = securityHeaderSettingValueModel
				updateSecurityHeaderOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateSecurityHeader(updateSecurityHeaderOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMobileRedirect(getMobileRedirectOptions *GetMobileRedirectOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMobileRedirectPath := "/v1/testString/zones/testString/settings/mobile_redirect"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMobileRedirectPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMobileRedirect with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMobileRedirectOptions model
				getMobileRedirectOptionsModel := new(zonessettingsv1.GetMobileRedirectOptions)
				getMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetMobileRedirect(getMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetMobileRedirect(getMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMobileRedirect(getMobileRedirectOptions *GetMobileRedirectOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMobileRedirectPath := "/v1/testString/zones/testString/settings/mobile_redirect"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMobileRedirectPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "mobile_redirect", "value": {"status": "true", "mobile_subdomain": "m", "strip_uri": false}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetMobileRedirect successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetMobileRedirect(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMobileRedirectOptions model
				getMobileRedirectOptionsModel := new(zonessettingsv1.GetMobileRedirectOptions)
				getMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetMobileRedirect(getMobileRedirectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMobileRedirectWithContext(ctx, getMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetMobileRedirect(getMobileRedirectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMobileRedirectWithContext(ctx, getMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMobileRedirect with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMobileRedirectOptions model
				getMobileRedirectOptionsModel := new(zonessettingsv1.GetMobileRedirectOptions)
				getMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetMobileRedirect(getMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMobileRedirect(updateMobileRedirectOptions *UpdateMobileRedirectOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMobileRedirectPath := "/v1/testString/zones/testString/settings/mobile_redirect"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMobileRedirectPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMobileRedirect with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the MobileRedirecSettingValue model
				mobileRedirecSettingValueModel := new(zonessettingsv1.MobileRedirecSettingValue)
				mobileRedirecSettingValueModel.Status = core.StringPtr("true")
				mobileRedirecSettingValueModel.MobileSubdomain = core.StringPtr("m")
				mobileRedirecSettingValueModel.StripURI = core.BoolPtr(false)

				// Construct an instance of the UpdateMobileRedirectOptions model
				updateMobileRedirectOptionsModel := new(zonessettingsv1.UpdateMobileRedirectOptions)
				updateMobileRedirectOptionsModel.Value = mobileRedirecSettingValueModel
				updateMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateMobileRedirect(updateMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateMobileRedirect(updateMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMobileRedirect(updateMobileRedirectOptions *UpdateMobileRedirectOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMobileRedirectPath := "/v1/testString/zones/testString/settings/mobile_redirect"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMobileRedirectPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "mobile_redirect", "value": {"status": "true", "mobile_subdomain": "m", "strip_uri": false}, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateMobileRedirect successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateMobileRedirect(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MobileRedirecSettingValue model
				mobileRedirecSettingValueModel := new(zonessettingsv1.MobileRedirecSettingValue)
				mobileRedirecSettingValueModel.Status = core.StringPtr("true")
				mobileRedirecSettingValueModel.MobileSubdomain = core.StringPtr("m")
				mobileRedirecSettingValueModel.StripURI = core.BoolPtr(false)

				// Construct an instance of the UpdateMobileRedirectOptions model
				updateMobileRedirectOptionsModel := new(zonessettingsv1.UpdateMobileRedirectOptions)
				updateMobileRedirectOptionsModel.Value = mobileRedirecSettingValueModel
				updateMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateMobileRedirect(updateMobileRedirectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMobileRedirectWithContext(ctx, updateMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateMobileRedirect(updateMobileRedirectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMobileRedirectWithContext(ctx, updateMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateMobileRedirect with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the MobileRedirecSettingValue model
				mobileRedirecSettingValueModel := new(zonessettingsv1.MobileRedirecSettingValue)
				mobileRedirecSettingValueModel.Status = core.StringPtr("true")
				mobileRedirecSettingValueModel.MobileSubdomain = core.StringPtr("m")
				mobileRedirecSettingValueModel.StripURI = core.BoolPtr(false)

				// Construct an instance of the UpdateMobileRedirectOptions model
				updateMobileRedirectOptionsModel := new(zonessettingsv1.UpdateMobileRedirectOptions)
				updateMobileRedirectOptionsModel.Value = mobileRedirecSettingValueModel
				updateMobileRedirectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateMobileRedirect(updateMobileRedirectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPrefetchPreload(getPrefetchPreloadOptions *GetPrefetchPreloadOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getPrefetchPreloadPath := "/v1/testString/zones/testString/settings/prefetch_preload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrefetchPreloadPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPrefetchPreload with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetPrefetchPreloadOptions model
				getPrefetchPreloadOptionsModel := new(zonessettingsv1.GetPrefetchPreloadOptions)
				getPrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetPrefetchPreload(getPrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetPrefetchPreload(getPrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPrefetchPreload(getPrefetchPreloadOptions *GetPrefetchPreloadOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getPrefetchPreloadPath := "/v1/testString/zones/testString/settings/prefetch_preload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPrefetchPreloadPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "prefetch_preload", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetPrefetchPreload successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetPrefetchPreload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPrefetchPreloadOptions model
				getPrefetchPreloadOptionsModel := new(zonessettingsv1.GetPrefetchPreloadOptions)
				getPrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetPrefetchPreload(getPrefetchPreloadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetPrefetchPreloadWithContext(ctx, getPrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetPrefetchPreload(getPrefetchPreloadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetPrefetchPreloadWithContext(ctx, getPrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetPrefetchPreload with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetPrefetchPreloadOptions model
				getPrefetchPreloadOptionsModel := new(zonessettingsv1.GetPrefetchPreloadOptions)
				getPrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetPrefetchPreload(getPrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePrefetchPreload(updatePrefetchPreloadOptions *UpdatePrefetchPreloadOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updatePrefetchPreloadPath := "/v1/testString/zones/testString/settings/prefetch_preload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePrefetchPreloadPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePrefetchPreload with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdatePrefetchPreloadOptions model
				updatePrefetchPreloadOptionsModel := new(zonessettingsv1.UpdatePrefetchPreloadOptions)
				updatePrefetchPreloadOptionsModel.Value = core.StringPtr("true")
				updatePrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdatePrefetchPreload(updatePrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdatePrefetchPreload(updatePrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePrefetchPreload(updatePrefetchPreloadOptions *UpdatePrefetchPreloadOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updatePrefetchPreloadPath := "/v1/testString/zones/testString/settings/prefetch_preload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePrefetchPreloadPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "prefetch_preload", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdatePrefetchPreload successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdatePrefetchPreload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePrefetchPreloadOptions model
				updatePrefetchPreloadOptionsModel := new(zonessettingsv1.UpdatePrefetchPreloadOptions)
				updatePrefetchPreloadOptionsModel.Value = core.StringPtr("true")
				updatePrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdatePrefetchPreload(updatePrefetchPreloadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdatePrefetchPreloadWithContext(ctx, updatePrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdatePrefetchPreload(updatePrefetchPreloadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdatePrefetchPreloadWithContext(ctx, updatePrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdatePrefetchPreload with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdatePrefetchPreloadOptions model
				updatePrefetchPreloadOptionsModel := new(zonessettingsv1.UpdatePrefetchPreloadOptions)
				updatePrefetchPreloadOptionsModel.Value = core.StringPtr("true")
				updatePrefetchPreloadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdatePrefetchPreload(updatePrefetchPreloadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHttp2(getHttp2Options *GetHttp2Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHttp2Path := "/v1/testString/zones/testString/settings/http2"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHttp2Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHttp2 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetHttp2Options model
				getHttp2OptionsModel := new(zonessettingsv1.GetHttp2Options)
				getHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetHttp2(getHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetHttp2(getHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHttp2(getHttp2Options *GetHttp2Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHttp2Path := "/v1/testString/zones/testString/settings/http2"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHttp2Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "http2", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetHttp2 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetHttp2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHttp2Options model
				getHttp2OptionsModel := new(zonessettingsv1.GetHttp2Options)
				getHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetHttp2(getHttp2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetHttp2WithContext(ctx, getHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetHttp2(getHttp2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetHttp2WithContext(ctx, getHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetHttp2 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetHttp2Options model
				getHttp2OptionsModel := new(zonessettingsv1.GetHttp2Options)
				getHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetHttp2(getHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateHttp2(updateHttp2Options *UpdateHttp2Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateHttp2Path := "/v1/testString/zones/testString/settings/http2"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHttp2Path))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateHttp2 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateHttp2Options model
				updateHttp2OptionsModel := new(zonessettingsv1.UpdateHttp2Options)
				updateHttp2OptionsModel.Value = core.StringPtr("true")
				updateHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateHttp2(updateHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateHttp2(updateHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateHttp2(updateHttp2Options *UpdateHttp2Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateHttp2Path := "/v1/testString/zones/testString/settings/http2"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHttp2Path))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "http2", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateHttp2 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateHttp2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateHttp2Options model
				updateHttp2OptionsModel := new(zonessettingsv1.UpdateHttp2Options)
				updateHttp2OptionsModel.Value = core.StringPtr("true")
				updateHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateHttp2(updateHttp2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateHttp2WithContext(ctx, updateHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateHttp2(updateHttp2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateHttp2WithContext(ctx, updateHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateHttp2 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateHttp2Options model
				updateHttp2OptionsModel := new(zonessettingsv1.UpdateHttp2Options)
				updateHttp2OptionsModel.Value = core.StringPtr("true")
				updateHttp2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateHttp2(updateHttp2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIpv6(getIpv6Options *GetIpv6Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getIpv6Path := "/v1/testString/zones/testString/settings/ipv6"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIpv6Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIpv6 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetIpv6Options model
				getIpv6OptionsModel := new(zonessettingsv1.GetIpv6Options)
				getIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetIpv6(getIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetIpv6(getIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetIpv6(getIpv6Options *GetIpv6Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getIpv6Path := "/v1/testString/zones/testString/settings/ipv6"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIpv6Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ipv6", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetIpv6 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetIpv6(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIpv6Options model
				getIpv6OptionsModel := new(zonessettingsv1.GetIpv6Options)
				getIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetIpv6(getIpv6OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetIpv6WithContext(ctx, getIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetIpv6(getIpv6OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetIpv6WithContext(ctx, getIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetIpv6 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetIpv6Options model
				getIpv6OptionsModel := new(zonessettingsv1.GetIpv6Options)
				getIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetIpv6(getIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateIpv6(updateIpv6Options *UpdateIpv6Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateIpv6Path := "/v1/testString/zones/testString/settings/ipv6"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIpv6Path))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateIpv6 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateIpv6Options model
				updateIpv6OptionsModel := new(zonessettingsv1.UpdateIpv6Options)
				updateIpv6OptionsModel.Value = core.StringPtr("true")
				updateIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateIpv6(updateIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateIpv6(updateIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateIpv6(updateIpv6Options *UpdateIpv6Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateIpv6Path := "/v1/testString/zones/testString/settings/ipv6"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateIpv6Path))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ipv6", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateIpv6 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateIpv6(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateIpv6Options model
				updateIpv6OptionsModel := new(zonessettingsv1.UpdateIpv6Options)
				updateIpv6OptionsModel.Value = core.StringPtr("true")
				updateIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateIpv6(updateIpv6OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateIpv6WithContext(ctx, updateIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateIpv6(updateIpv6OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateIpv6WithContext(ctx, updateIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateIpv6 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateIpv6Options model
				updateIpv6OptionsModel := new(zonessettingsv1.UpdateIpv6Options)
				updateIpv6OptionsModel.Value = core.StringPtr("true")
				updateIpv6OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateIpv6(updateIpv6OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWebSockets(getWebSocketsOptions *GetWebSocketsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getWebSocketsPath := "/v1/testString/zones/testString/settings/websockets"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebSocketsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWebSockets with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetWebSocketsOptions model
				getWebSocketsOptionsModel := new(zonessettingsv1.GetWebSocketsOptions)
				getWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetWebSockets(getWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetWebSockets(getWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWebSockets(getWebSocketsOptions *GetWebSocketsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getWebSocketsPath := "/v1/testString/zones/testString/settings/websockets"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebSocketsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "websockets", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetWebSockets successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetWebSockets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWebSocketsOptions model
				getWebSocketsOptionsModel := new(zonessettingsv1.GetWebSocketsOptions)
				getWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetWebSockets(getWebSocketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetWebSocketsWithContext(ctx, getWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetWebSockets(getWebSocketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetWebSocketsWithContext(ctx, getWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWebSockets with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetWebSocketsOptions model
				getWebSocketsOptionsModel := new(zonessettingsv1.GetWebSocketsOptions)
				getWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetWebSockets(getWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWebSockets(updateWebSocketsOptions *UpdateWebSocketsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateWebSocketsPath := "/v1/testString/zones/testString/settings/websockets"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWebSocketsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWebSockets with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateWebSocketsOptions model
				updateWebSocketsOptionsModel := new(zonessettingsv1.UpdateWebSocketsOptions)
				updateWebSocketsOptionsModel.Value = core.StringPtr("true")
				updateWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateWebSockets(updateWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateWebSockets(updateWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWebSockets(updateWebSocketsOptions *UpdateWebSocketsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateWebSocketsPath := "/v1/testString/zones/testString/settings/websockets"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWebSocketsPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "websockets", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateWebSockets successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateWebSockets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateWebSocketsOptions model
				updateWebSocketsOptionsModel := new(zonessettingsv1.UpdateWebSocketsOptions)
				updateWebSocketsOptionsModel.Value = core.StringPtr("true")
				updateWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateWebSockets(updateWebSocketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateWebSocketsWithContext(ctx, updateWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateWebSockets(updateWebSocketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateWebSocketsWithContext(ctx, updateWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateWebSockets with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateWebSocketsOptions model
				updateWebSocketsOptionsModel := new(zonessettingsv1.UpdateWebSocketsOptions)
				updateWebSocketsOptionsModel.Value = core.StringPtr("true")
				updateWebSocketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateWebSockets(updateWebSocketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPseudoIpv4(getPseudoIpv4Options *GetPseudoIpv4Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getPseudoIpv4Path := "/v1/testString/zones/testString/settings/pseudo_ipv4"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPseudoIpv4Path))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPseudoIpv4 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetPseudoIpv4Options model
				getPseudoIpv4OptionsModel := new(zonessettingsv1.GetPseudoIpv4Options)
				getPseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetPseudoIpv4(getPseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetPseudoIpv4(getPseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPseudoIpv4(getPseudoIpv4Options *GetPseudoIpv4Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getPseudoIpv4Path := "/v1/testString/zones/testString/settings/pseudo_ipv4"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPseudoIpv4Path))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "pseudo_ipv4", "value": "add_header", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetPseudoIpv4 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetPseudoIpv4(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPseudoIpv4Options model
				getPseudoIpv4OptionsModel := new(zonessettingsv1.GetPseudoIpv4Options)
				getPseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetPseudoIpv4(getPseudoIpv4OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetPseudoIpv4WithContext(ctx, getPseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetPseudoIpv4(getPseudoIpv4OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetPseudoIpv4WithContext(ctx, getPseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetPseudoIpv4 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetPseudoIpv4Options model
				getPseudoIpv4OptionsModel := new(zonessettingsv1.GetPseudoIpv4Options)
				getPseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetPseudoIpv4(getPseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePseudoIpv4(updatePseudoIpv4Options *UpdatePseudoIpv4Options) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updatePseudoIpv4Path := "/v1/testString/zones/testString/settings/pseudo_ipv4"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePseudoIpv4Path))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePseudoIpv4 with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdatePseudoIpv4Options model
				updatePseudoIpv4OptionsModel := new(zonessettingsv1.UpdatePseudoIpv4Options)
				updatePseudoIpv4OptionsModel.Value = core.StringPtr("add_header")
				updatePseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdatePseudoIpv4(updatePseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdatePseudoIpv4(updatePseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePseudoIpv4(updatePseudoIpv4Options *UpdatePseudoIpv4Options)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updatePseudoIpv4Path := "/v1/testString/zones/testString/settings/pseudo_ipv4"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePseudoIpv4Path))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "pseudo_ipv4", "value": "add_header", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdatePseudoIpv4 successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdatePseudoIpv4(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePseudoIpv4Options model
				updatePseudoIpv4OptionsModel := new(zonessettingsv1.UpdatePseudoIpv4Options)
				updatePseudoIpv4OptionsModel.Value = core.StringPtr("add_header")
				updatePseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdatePseudoIpv4(updatePseudoIpv4OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdatePseudoIpv4WithContext(ctx, updatePseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdatePseudoIpv4(updatePseudoIpv4OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdatePseudoIpv4WithContext(ctx, updatePseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdatePseudoIpv4 with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdatePseudoIpv4Options model
				updatePseudoIpv4OptionsModel := new(zonessettingsv1.UpdatePseudoIpv4Options)
				updatePseudoIpv4OptionsModel.Value = core.StringPtr("add_header")
				updatePseudoIpv4OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdatePseudoIpv4(updatePseudoIpv4OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResponseBuffering(getResponseBufferingOptions *GetResponseBufferingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getResponseBufferingPath := "/v1/testString/zones/testString/settings/response_buffering"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResponseBufferingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResponseBuffering with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetResponseBufferingOptions model
				getResponseBufferingOptionsModel := new(zonessettingsv1.GetResponseBufferingOptions)
				getResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetResponseBuffering(getResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetResponseBuffering(getResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResponseBuffering(getResponseBufferingOptions *GetResponseBufferingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getResponseBufferingPath := "/v1/testString/zones/testString/settings/response_buffering"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResponseBufferingPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "response_buffering", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetResponseBuffering successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetResponseBuffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResponseBufferingOptions model
				getResponseBufferingOptionsModel := new(zonessettingsv1.GetResponseBufferingOptions)
				getResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetResponseBuffering(getResponseBufferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetResponseBufferingWithContext(ctx, getResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetResponseBuffering(getResponseBufferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetResponseBufferingWithContext(ctx, getResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetResponseBuffering with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetResponseBufferingOptions model
				getResponseBufferingOptionsModel := new(zonessettingsv1.GetResponseBufferingOptions)
				getResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetResponseBuffering(getResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResponseBuffering(updateResponseBufferingOptions *UpdateResponseBufferingOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateResponseBufferingPath := "/v1/testString/zones/testString/settings/response_buffering"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResponseBufferingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResponseBuffering with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateResponseBufferingOptions model
				updateResponseBufferingOptionsModel := new(zonessettingsv1.UpdateResponseBufferingOptions)
				updateResponseBufferingOptionsModel.Value = core.StringPtr("true")
				updateResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateResponseBuffering(updateResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateResponseBuffering(updateResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResponseBuffering(updateResponseBufferingOptions *UpdateResponseBufferingOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateResponseBufferingPath := "/v1/testString/zones/testString/settings/response_buffering"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResponseBufferingPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "response_buffering", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateResponseBuffering successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateResponseBuffering(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResponseBufferingOptions model
				updateResponseBufferingOptionsModel := new(zonessettingsv1.UpdateResponseBufferingOptions)
				updateResponseBufferingOptionsModel.Value = core.StringPtr("true")
				updateResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateResponseBuffering(updateResponseBufferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateResponseBufferingWithContext(ctx, updateResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateResponseBuffering(updateResponseBufferingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateResponseBufferingWithContext(ctx, updateResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateResponseBuffering with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateResponseBufferingOptions model
				updateResponseBufferingOptionsModel := new(zonessettingsv1.UpdateResponseBufferingOptions)
				updateResponseBufferingOptionsModel.Value = core.StringPtr("true")
				updateResponseBufferingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateResponseBuffering(updateResponseBufferingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHotlinkProtection(getHotlinkProtectionOptions *GetHotlinkProtectionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHotlinkProtectionPath := "/v1/testString/zones/testString/settings/hotlink_protection"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHotlinkProtectionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHotlinkProtection with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetHotlinkProtectionOptions model
				getHotlinkProtectionOptionsModel := new(zonessettingsv1.GetHotlinkProtectionOptions)
				getHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetHotlinkProtection(getHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetHotlinkProtection(getHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHotlinkProtection(getHotlinkProtectionOptions *GetHotlinkProtectionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getHotlinkProtectionPath := "/v1/testString/zones/testString/settings/hotlink_protection"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getHotlinkProtectionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "hotlink_protection", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetHotlinkProtection successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetHotlinkProtection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHotlinkProtectionOptions model
				getHotlinkProtectionOptionsModel := new(zonessettingsv1.GetHotlinkProtectionOptions)
				getHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetHotlinkProtection(getHotlinkProtectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetHotlinkProtectionWithContext(ctx, getHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetHotlinkProtection(getHotlinkProtectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetHotlinkProtectionWithContext(ctx, getHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetHotlinkProtection with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetHotlinkProtectionOptions model
				getHotlinkProtectionOptionsModel := new(zonessettingsv1.GetHotlinkProtectionOptions)
				getHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetHotlinkProtection(getHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateHotlinkProtection(updateHotlinkProtectionOptions *UpdateHotlinkProtectionOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateHotlinkProtectionPath := "/v1/testString/zones/testString/settings/hotlink_protection"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHotlinkProtectionPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateHotlinkProtection with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateHotlinkProtectionOptions model
				updateHotlinkProtectionOptionsModel := new(zonessettingsv1.UpdateHotlinkProtectionOptions)
				updateHotlinkProtectionOptionsModel.Value = core.StringPtr("true")
				updateHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateHotlinkProtection(updateHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateHotlinkProtection(updateHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateHotlinkProtection(updateHotlinkProtectionOptions *UpdateHotlinkProtectionOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateHotlinkProtectionPath := "/v1/testString/zones/testString/settings/hotlink_protection"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateHotlinkProtectionPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "hotlink_protection", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateHotlinkProtection successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateHotlinkProtection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateHotlinkProtectionOptions model
				updateHotlinkProtectionOptionsModel := new(zonessettingsv1.UpdateHotlinkProtectionOptions)
				updateHotlinkProtectionOptionsModel.Value = core.StringPtr("true")
				updateHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateHotlinkProtection(updateHotlinkProtectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateHotlinkProtectionWithContext(ctx, updateHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateHotlinkProtection(updateHotlinkProtectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateHotlinkProtectionWithContext(ctx, updateHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateHotlinkProtection with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateHotlinkProtectionOptions model
				updateHotlinkProtectionOptionsModel := new(zonessettingsv1.UpdateHotlinkProtectionOptions)
				updateHotlinkProtectionOptionsModel.Value = core.StringPtr("true")
				updateHotlinkProtectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateHotlinkProtection(updateHotlinkProtectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMaxUpload(getMaxUploadOptions *GetMaxUploadOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMaxUploadPath := "/v1/testString/zones/testString/settings/max_upload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMaxUploadPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMaxUpload with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMaxUploadOptions model
				getMaxUploadOptionsModel := new(zonessettingsv1.GetMaxUploadOptions)
				getMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetMaxUpload(getMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetMaxUpload(getMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMaxUpload(getMaxUploadOptions *GetMaxUploadOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getMaxUploadPath := "/v1/testString/zones/testString/settings/max_upload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMaxUploadPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "max_upload", "value": 300, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetMaxUpload successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetMaxUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMaxUploadOptions model
				getMaxUploadOptionsModel := new(zonessettingsv1.GetMaxUploadOptions)
				getMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetMaxUpload(getMaxUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMaxUploadWithContext(ctx, getMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetMaxUpload(getMaxUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetMaxUploadWithContext(ctx, getMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetMaxUpload with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetMaxUploadOptions model
				getMaxUploadOptionsModel := new(zonessettingsv1.GetMaxUploadOptions)
				getMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetMaxUpload(getMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateMaxUpload(updateMaxUploadOptions *UpdateMaxUploadOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMaxUploadPath := "/v1/testString/zones/testString/settings/max_upload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMaxUploadPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateMaxUpload with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateMaxUploadOptions model
				updateMaxUploadOptionsModel := new(zonessettingsv1.UpdateMaxUploadOptions)
				updateMaxUploadOptionsModel.Value = core.Int64Ptr(int64(300))
				updateMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateMaxUpload(updateMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateMaxUpload(updateMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateMaxUpload(updateMaxUploadOptions *UpdateMaxUploadOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateMaxUploadPath := "/v1/testString/zones/testString/settings/max_upload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateMaxUploadPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "max_upload", "value": 300, "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateMaxUpload successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateMaxUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateMaxUploadOptions model
				updateMaxUploadOptionsModel := new(zonessettingsv1.UpdateMaxUploadOptions)
				updateMaxUploadOptionsModel.Value = core.Int64Ptr(int64(300))
				updateMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateMaxUpload(updateMaxUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMaxUploadWithContext(ctx, updateMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateMaxUpload(updateMaxUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateMaxUploadWithContext(ctx, updateMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateMaxUpload with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateMaxUploadOptions model
				updateMaxUploadOptionsModel := new(zonessettingsv1.UpdateMaxUploadOptions)
				updateMaxUploadOptionsModel.Value = core.Int64Ptr(int64(300))
				updateMaxUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateMaxUpload(updateMaxUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTlsClientAuth(getTlsClientAuthOptions *GetTlsClientAuthOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTlsClientAuthPath := "/v1/testString/zones/testString/settings/tls_client_auth"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTlsClientAuthPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTlsClientAuth with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetTlsClientAuthOptions model
				getTlsClientAuthOptionsModel := new(zonessettingsv1.GetTlsClientAuthOptions)
				getTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetTlsClientAuth(getTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetTlsClientAuth(getTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTlsClientAuth(getTlsClientAuthOptions *GetTlsClientAuthOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getTlsClientAuthPath := "/v1/testString/zones/testString/settings/tls_client_auth"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTlsClientAuthPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "tls_client_auth", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetTlsClientAuth successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetTlsClientAuth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTlsClientAuthOptions model
				getTlsClientAuthOptionsModel := new(zonessettingsv1.GetTlsClientAuthOptions)
				getTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetTlsClientAuth(getTlsClientAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetTlsClientAuthWithContext(ctx, getTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetTlsClientAuth(getTlsClientAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetTlsClientAuthWithContext(ctx, getTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTlsClientAuth with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetTlsClientAuthOptions model
				getTlsClientAuthOptionsModel := new(zonessettingsv1.GetTlsClientAuthOptions)
				getTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetTlsClientAuth(getTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTlsClientAuth(updateTlsClientAuthOptions *UpdateTlsClientAuthOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateTlsClientAuthPath := "/v1/testString/zones/testString/settings/tls_client_auth"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTlsClientAuthPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTlsClientAuth with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateTlsClientAuthOptions model
				updateTlsClientAuthOptionsModel := new(zonessettingsv1.UpdateTlsClientAuthOptions)
				updateTlsClientAuthOptionsModel.Value = core.StringPtr("true")
				updateTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateTlsClientAuth(updateTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateTlsClientAuth(updateTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateTlsClientAuth(updateTlsClientAuthOptions *UpdateTlsClientAuthOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateTlsClientAuthPath := "/v1/testString/zones/testString/settings/tls_client_auth"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTlsClientAuthPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "tls_client_auth", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateTlsClientAuth successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateTlsClientAuth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTlsClientAuthOptions model
				updateTlsClientAuthOptionsModel := new(zonessettingsv1.UpdateTlsClientAuthOptions)
				updateTlsClientAuthOptionsModel.Value = core.StringPtr("true")
				updateTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateTlsClientAuth(updateTlsClientAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateTlsClientAuthWithContext(ctx, updateTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateTlsClientAuth(updateTlsClientAuthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateTlsClientAuthWithContext(ctx, updateTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateTlsClientAuth with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateTlsClientAuthOptions model
				updateTlsClientAuthOptionsModel := new(zonessettingsv1.UpdateTlsClientAuthOptions)
				updateTlsClientAuthOptionsModel.Value = core.StringPtr("true")
				updateTlsClientAuthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateTlsClientAuth(updateTlsClientAuthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBrowserCheck(getBrowserCheckOptions *GetBrowserCheckOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBrowserCheckPath := "/v1/testString/zones/testString/settings/browser_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBrowserCheckPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBrowserCheck with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetBrowserCheckOptions model
				getBrowserCheckOptionsModel := new(zonessettingsv1.GetBrowserCheckOptions)
				getBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetBrowserCheck(getBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetBrowserCheck(getBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBrowserCheck(getBrowserCheckOptions *GetBrowserCheckOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getBrowserCheckPath := "/v1/testString/zones/testString/settings/browser_check"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBrowserCheckPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "browser_check", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetBrowserCheck successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetBrowserCheck(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBrowserCheckOptions model
				getBrowserCheckOptionsModel := new(zonessettingsv1.GetBrowserCheckOptions)
				getBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetBrowserCheck(getBrowserCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetBrowserCheckWithContext(ctx, getBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetBrowserCheck(getBrowserCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetBrowserCheckWithContext(ctx, getBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetBrowserCheck with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetBrowserCheckOptions model
				getBrowserCheckOptionsModel := new(zonessettingsv1.GetBrowserCheckOptions)
				getBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetBrowserCheck(getBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBrowserCheck(updateBrowserCheckOptions *UpdateBrowserCheckOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateBrowserCheckPath := "/v1/testString/zones/testString/settings/browser_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBrowserCheckPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBrowserCheck with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateBrowserCheckOptions model
				updateBrowserCheckOptionsModel := new(zonessettingsv1.UpdateBrowserCheckOptions)
				updateBrowserCheckOptionsModel.Value = core.StringPtr("true")
				updateBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateBrowserCheck(updateBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateBrowserCheck(updateBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateBrowserCheck(updateBrowserCheckOptions *UpdateBrowserCheckOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateBrowserCheckPath := "/v1/testString/zones/testString/settings/browser_check"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBrowserCheckPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "browser_check", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateBrowserCheck successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateBrowserCheck(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateBrowserCheckOptions model
				updateBrowserCheckOptionsModel := new(zonessettingsv1.UpdateBrowserCheckOptions)
				updateBrowserCheckOptionsModel.Value = core.StringPtr("true")
				updateBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateBrowserCheck(updateBrowserCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateBrowserCheckWithContext(ctx, updateBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateBrowserCheck(updateBrowserCheckOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateBrowserCheckWithContext(ctx, updateBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateBrowserCheck with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateBrowserCheckOptions model
				updateBrowserCheckOptionsModel := new(zonessettingsv1.UpdateBrowserCheckOptions)
				updateBrowserCheckOptionsModel.Value = core.StringPtr("true")
				updateBrowserCheckOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateBrowserCheck(updateBrowserCheckOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEnableErrorPagesOn(getEnableErrorPagesOnOptions *GetEnableErrorPagesOnOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getEnableErrorPagesOnPath := "/v1/testString/zones/testString/settings/origin_error_page_pass_thru"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnableErrorPagesOnPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEnableErrorPagesOn with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetEnableErrorPagesOnOptions model
				getEnableErrorPagesOnOptionsModel := new(zonessettingsv1.GetEnableErrorPagesOnOptions)
				getEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetEnableErrorPagesOn(getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetEnableErrorPagesOn(getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEnableErrorPagesOn(getEnableErrorPagesOnOptions *GetEnableErrorPagesOnOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getEnableErrorPagesOnPath := "/v1/testString/zones/testString/settings/origin_error_page_pass_thru"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEnableErrorPagesOnPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "origin_error_page_pass_thru", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetEnableErrorPagesOn successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetEnableErrorPagesOn(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnableErrorPagesOnOptions model
				getEnableErrorPagesOnOptionsModel := new(zonessettingsv1.GetEnableErrorPagesOnOptions)
				getEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetEnableErrorPagesOn(getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetEnableErrorPagesOnWithContext(ctx, getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetEnableErrorPagesOn(getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetEnableErrorPagesOnWithContext(ctx, getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetEnableErrorPagesOn with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetEnableErrorPagesOnOptions model
				getEnableErrorPagesOnOptionsModel := new(zonessettingsv1.GetEnableErrorPagesOnOptions)
				getEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetEnableErrorPagesOn(getEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptions *UpdateEnableErrorPagesOnOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEnableErrorPagesOnPath := "/v1/testString/zones/testString/settings/origin_error_page_pass_thru"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnableErrorPagesOnPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEnableErrorPagesOn with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateEnableErrorPagesOnOptions model
				updateEnableErrorPagesOnOptionsModel := new(zonessettingsv1.UpdateEnableErrorPagesOnOptions)
				updateEnableErrorPagesOnOptionsModel.Value = core.StringPtr("true")
				updateEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptions *UpdateEnableErrorPagesOnOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateEnableErrorPagesOnPath := "/v1/testString/zones/testString/settings/origin_error_page_pass_thru"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEnableErrorPagesOnPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "origin_error_page_pass_thru", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateEnableErrorPagesOn successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateEnableErrorPagesOn(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateEnableErrorPagesOnOptions model
				updateEnableErrorPagesOnOptionsModel := new(zonessettingsv1.UpdateEnableErrorPagesOnOptions)
				updateEnableErrorPagesOnOptionsModel.Value = core.StringPtr("true")
				updateEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateEnableErrorPagesOnWithContext(ctx, updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateEnableErrorPagesOnWithContext(ctx, updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateEnableErrorPagesOn with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateEnableErrorPagesOnOptions model
				updateEnableErrorPagesOnOptionsModel := new(zonessettingsv1.UpdateEnableErrorPagesOnOptions)
				updateEnableErrorPagesOnOptionsModel.Value = core.StringPtr("true")
				updateEnableErrorPagesOnOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateEnableErrorPagesOn(updateEnableErrorPagesOnOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWebApplicationFirewall(getWebApplicationFirewallOptions *GetWebApplicationFirewallOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getWebApplicationFirewallPath := "/v1/testString/zones/testString/settings/waf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebApplicationFirewallPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWebApplicationFirewall with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetWebApplicationFirewallOptions model
				getWebApplicationFirewallOptionsModel := new(zonessettingsv1.GetWebApplicationFirewallOptions)
				getWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetWebApplicationFirewall(getWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetWebApplicationFirewall(getWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWebApplicationFirewall(getWebApplicationFirewallOptions *GetWebApplicationFirewallOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getWebApplicationFirewallPath := "/v1/testString/zones/testString/settings/waf"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebApplicationFirewallPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "waf", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetWebApplicationFirewall successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetWebApplicationFirewall(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWebApplicationFirewallOptions model
				getWebApplicationFirewallOptionsModel := new(zonessettingsv1.GetWebApplicationFirewallOptions)
				getWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetWebApplicationFirewall(getWebApplicationFirewallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetWebApplicationFirewallWithContext(ctx, getWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetWebApplicationFirewall(getWebApplicationFirewallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetWebApplicationFirewallWithContext(ctx, getWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWebApplicationFirewall with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetWebApplicationFirewallOptions model
				getWebApplicationFirewallOptionsModel := new(zonessettingsv1.GetWebApplicationFirewallOptions)
				getWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetWebApplicationFirewall(getWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWebApplicationFirewall(updateWebApplicationFirewallOptions *UpdateWebApplicationFirewallOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateWebApplicationFirewallPath := "/v1/testString/zones/testString/settings/waf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWebApplicationFirewallPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWebApplicationFirewall with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateWebApplicationFirewallOptions model
				updateWebApplicationFirewallOptionsModel := new(zonessettingsv1.UpdateWebApplicationFirewallOptions)
				updateWebApplicationFirewallOptionsModel.Value = core.StringPtr("true")
				updateWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateWebApplicationFirewall(updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateWebApplicationFirewall(updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWebApplicationFirewall(updateWebApplicationFirewallOptions *UpdateWebApplicationFirewallOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateWebApplicationFirewallPath := "/v1/testString/zones/testString/settings/waf"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWebApplicationFirewallPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "waf", "value": "false", "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateWebApplicationFirewall successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateWebApplicationFirewall(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateWebApplicationFirewallOptions model
				updateWebApplicationFirewallOptionsModel := new(zonessettingsv1.UpdateWebApplicationFirewallOptions)
				updateWebApplicationFirewallOptionsModel.Value = core.StringPtr("true")
				updateWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateWebApplicationFirewall(updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateWebApplicationFirewallWithContext(ctx, updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateWebApplicationFirewall(updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateWebApplicationFirewallWithContext(ctx, updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateWebApplicationFirewall with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateWebApplicationFirewallOptions model
				updateWebApplicationFirewallOptionsModel := new(zonessettingsv1.UpdateWebApplicationFirewallOptions)
				updateWebApplicationFirewallOptionsModel.Value = core.StringPtr("true")
				updateWebApplicationFirewallOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateWebApplicationFirewall(updateWebApplicationFirewallOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCiphers(getCiphersOptions *GetCiphersOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getCiphersPath := "/v1/testString/zones/testString/settings/ciphers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCiphersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCiphers with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetCiphersOptions model
				getCiphersOptionsModel := new(zonessettingsv1.GetCiphersOptions)
				getCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.GetCiphers(getCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.GetCiphers(getCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCiphers(getCiphersOptions *GetCiphersOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getCiphersPath := "/v1/testString/zones/testString/settings/ciphers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCiphersPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ciphers", "value": ["Value"], "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke GetCiphers successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.GetCiphers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCiphersOptions model
				getCiphersOptionsModel := new(zonessettingsv1.GetCiphersOptions)
				getCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.GetCiphers(getCiphersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetCiphersWithContext(ctx, getCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.GetCiphers(getCiphersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.GetCiphersWithContext(ctx, getCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetCiphers with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the GetCiphersOptions model
				getCiphersOptionsModel := new(zonessettingsv1.GetCiphersOptions)
				getCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.GetCiphers(getCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCiphers(updateCiphersOptions *UpdateCiphersOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateCiphersPath := "/v1/testString/zones/testString/settings/ciphers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCiphersPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCiphers with error: Operation response processing error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateCiphersOptions model
				updateCiphersOptionsModel := new(zonessettingsv1.UpdateCiphersOptions)
				updateCiphersOptionsModel.Value = []string{"ECDHE-ECDSA-AES128-GCM-SHA256"}
				updateCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zonesSettingsService.UpdateCiphers(updateCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zonesSettingsService.EnableRetries(0, 0)
				result, response, operationErr = zonesSettingsService.UpdateCiphers(updateCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCiphers(updateCiphersOptions *UpdateCiphersOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateCiphersPath := "/v1/testString/zones/testString/settings/ciphers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCiphersPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": {"id": "ciphers", "value": ["Value"], "editable": true, "modified_on": "2019-01-01T12:00:00"}, "success": true, "errors": [["Errors"]], "messages": [["Messages"]]}`)
				}))
			})
			It(`Invoke UpdateCiphers successfully`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())
				zonesSettingsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zonesSettingsService.UpdateCiphers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCiphersOptions model
				updateCiphersOptionsModel := new(zonessettingsv1.UpdateCiphersOptions)
				updateCiphersOptionsModel.Value = []string{"ECDHE-ECDSA-AES128-GCM-SHA256"}
				updateCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zonesSettingsService.UpdateCiphers(updateCiphersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateCiphersWithContext(ctx, updateCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zonesSettingsService.DisableRetries()
				result, response, operationErr = zonesSettingsService.UpdateCiphers(updateCiphersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zonesSettingsService.UpdateCiphersWithContext(ctx, updateCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateCiphers with error: Operation request error`, func() {
				zonesSettingsService, serviceErr := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Crn: core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zonesSettingsService).ToNot(BeNil())

				// Construct an instance of the UpdateCiphersOptions model
				updateCiphersOptionsModel := new(zonessettingsv1.UpdateCiphersOptions)
				updateCiphersOptionsModel.Value = []string{"ECDHE-ECDSA-AES128-GCM-SHA256"}
				updateCiphersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zonesSettingsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zonesSettingsService.UpdateCiphers(updateCiphersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneIdentifier := "testString"
			zonesSettingsService, _ := zonessettingsv1.NewZonesSettingsV1(&zonessettingsv1.ZonesSettingsV1Options{
				URL:           "http://zonessettingsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Crn: core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewGetAlwaysUseHttpsOptions successfully`, func() {
				// Construct an instance of the GetAlwaysUseHttpsOptions model
				getAlwaysUseHttpsOptionsModel := zonesSettingsService.NewGetAlwaysUseHttpsOptions()
				getAlwaysUseHttpsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAlwaysUseHttpsOptionsModel).ToNot(BeNil())
				Expect(getAlwaysUseHttpsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAutomaticHttpsRewritesOptions successfully`, func() {
				// Construct an instance of the GetAutomaticHttpsRewritesOptions model
				getAutomaticHttpsRewritesOptionsModel := zonesSettingsService.NewGetAutomaticHttpsRewritesOptions()
				getAutomaticHttpsRewritesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAutomaticHttpsRewritesOptionsModel).ToNot(BeNil())
				Expect(getAutomaticHttpsRewritesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBrowserCheckOptions successfully`, func() {
				// Construct an instance of the GetBrowserCheckOptions model
				getBrowserCheckOptionsModel := zonesSettingsService.NewGetBrowserCheckOptions()
				getBrowserCheckOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBrowserCheckOptionsModel).ToNot(BeNil())
				Expect(getBrowserCheckOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChallengeTtlOptions successfully`, func() {
				// Construct an instance of the GetChallengeTtlOptions model
				getChallengeTtlOptionsModel := zonesSettingsService.NewGetChallengeTtlOptions()
				getChallengeTtlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChallengeTtlOptionsModel).ToNot(BeNil())
				Expect(getChallengeTtlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCiphersOptions successfully`, func() {
				// Construct an instance of the GetCiphersOptions model
				getCiphersOptionsModel := zonesSettingsService.NewGetCiphersOptions()
				getCiphersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCiphersOptionsModel).ToNot(BeNil())
				Expect(getCiphersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEnableErrorPagesOnOptions successfully`, func() {
				// Construct an instance of the GetEnableErrorPagesOnOptions model
				getEnableErrorPagesOnOptionsModel := zonesSettingsService.NewGetEnableErrorPagesOnOptions()
				getEnableErrorPagesOnOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEnableErrorPagesOnOptionsModel).ToNot(BeNil())
				Expect(getEnableErrorPagesOnOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHotlinkProtectionOptions successfully`, func() {
				// Construct an instance of the GetHotlinkProtectionOptions model
				getHotlinkProtectionOptionsModel := zonesSettingsService.NewGetHotlinkProtectionOptions()
				getHotlinkProtectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHotlinkProtectionOptionsModel).ToNot(BeNil())
				Expect(getHotlinkProtectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHttp2Options successfully`, func() {
				// Construct an instance of the GetHttp2Options model
				getHttp2OptionsModel := zonesSettingsService.NewGetHttp2Options()
				getHttp2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHttp2OptionsModel).ToNot(BeNil())
				Expect(getHttp2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetImageLoadOptimizationOptions successfully`, func() {
				// Construct an instance of the GetImageLoadOptimizationOptions model
				getImageLoadOptimizationOptionsModel := zonesSettingsService.NewGetImageLoadOptimizationOptions()
				getImageLoadOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getImageLoadOptimizationOptionsModel).ToNot(BeNil())
				Expect(getImageLoadOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetImageSizeOptimizationOptions successfully`, func() {
				// Construct an instance of the GetImageSizeOptimizationOptions model
				getImageSizeOptimizationOptionsModel := zonesSettingsService.NewGetImageSizeOptimizationOptions()
				getImageSizeOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getImageSizeOptimizationOptionsModel).ToNot(BeNil())
				Expect(getImageSizeOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIpGeolocationOptions successfully`, func() {
				// Construct an instance of the GetIpGeolocationOptions model
				getIpGeolocationOptionsModel := zonesSettingsService.NewGetIpGeolocationOptions()
				getIpGeolocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIpGeolocationOptionsModel).ToNot(BeNil())
				Expect(getIpGeolocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetIpv6Options successfully`, func() {
				// Construct an instance of the GetIpv6Options model
				getIpv6OptionsModel := zonesSettingsService.NewGetIpv6Options()
				getIpv6OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIpv6OptionsModel).ToNot(BeNil())
				Expect(getIpv6OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMaxUploadOptions successfully`, func() {
				// Construct an instance of the GetMaxUploadOptions model
				getMaxUploadOptionsModel := zonesSettingsService.NewGetMaxUploadOptions()
				getMaxUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMaxUploadOptionsModel).ToNot(BeNil())
				Expect(getMaxUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMinTlsVersionOptions successfully`, func() {
				// Construct an instance of the GetMinTlsVersionOptions model
				getMinTlsVersionOptionsModel := zonesSettingsService.NewGetMinTlsVersionOptions()
				getMinTlsVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMinTlsVersionOptionsModel).ToNot(BeNil())
				Expect(getMinTlsVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMinifyOptions successfully`, func() {
				// Construct an instance of the GetMinifyOptions model
				getMinifyOptionsModel := zonesSettingsService.NewGetMinifyOptions()
				getMinifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMinifyOptionsModel).ToNot(BeNil())
				Expect(getMinifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMobileRedirectOptions successfully`, func() {
				// Construct an instance of the GetMobileRedirectOptions model
				getMobileRedirectOptionsModel := zonesSettingsService.NewGetMobileRedirectOptions()
				getMobileRedirectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMobileRedirectOptionsModel).ToNot(BeNil())
				Expect(getMobileRedirectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOpportunisticEncryptionOptions successfully`, func() {
				// Construct an instance of the GetOpportunisticEncryptionOptions model
				getOpportunisticEncryptionOptionsModel := zonesSettingsService.NewGetOpportunisticEncryptionOptions()
				getOpportunisticEncryptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOpportunisticEncryptionOptionsModel).ToNot(BeNil())
				Expect(getOpportunisticEncryptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPrefetchPreloadOptions successfully`, func() {
				// Construct an instance of the GetPrefetchPreloadOptions model
				getPrefetchPreloadOptionsModel := zonesSettingsService.NewGetPrefetchPreloadOptions()
				getPrefetchPreloadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPrefetchPreloadOptionsModel).ToNot(BeNil())
				Expect(getPrefetchPreloadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPseudoIpv4Options successfully`, func() {
				// Construct an instance of the GetPseudoIpv4Options model
				getPseudoIpv4OptionsModel := zonesSettingsService.NewGetPseudoIpv4Options()
				getPseudoIpv4OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPseudoIpv4OptionsModel).ToNot(BeNil())
				Expect(getPseudoIpv4OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResponseBufferingOptions successfully`, func() {
				// Construct an instance of the GetResponseBufferingOptions model
				getResponseBufferingOptionsModel := zonesSettingsService.NewGetResponseBufferingOptions()
				getResponseBufferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResponseBufferingOptionsModel).ToNot(BeNil())
				Expect(getResponseBufferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetScriptLoadOptimizationOptions successfully`, func() {
				// Construct an instance of the GetScriptLoadOptimizationOptions model
				getScriptLoadOptimizationOptionsModel := zonesSettingsService.NewGetScriptLoadOptimizationOptions()
				getScriptLoadOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getScriptLoadOptimizationOptionsModel).ToNot(BeNil())
				Expect(getScriptLoadOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSecurityHeaderOptions successfully`, func() {
				// Construct an instance of the GetSecurityHeaderOptions model
				getSecurityHeaderOptionsModel := zonesSettingsService.NewGetSecurityHeaderOptions()
				getSecurityHeaderOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSecurityHeaderOptionsModel).ToNot(BeNil())
				Expect(getSecurityHeaderOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServerSideExcludeOptions successfully`, func() {
				// Construct an instance of the GetServerSideExcludeOptions model
				getServerSideExcludeOptionsModel := zonesSettingsService.NewGetServerSideExcludeOptions()
				getServerSideExcludeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServerSideExcludeOptionsModel).ToNot(BeNil())
				Expect(getServerSideExcludeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTlsClientAuthOptions successfully`, func() {
				// Construct an instance of the GetTlsClientAuthOptions model
				getTlsClientAuthOptionsModel := zonesSettingsService.NewGetTlsClientAuthOptions()
				getTlsClientAuthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTlsClientAuthOptionsModel).ToNot(BeNil())
				Expect(getTlsClientAuthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTrueClientIpOptions successfully`, func() {
				// Construct an instance of the GetTrueClientIpOptions model
				getTrueClientIpOptionsModel := zonesSettingsService.NewGetTrueClientIpOptions()
				getTrueClientIpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTrueClientIpOptionsModel).ToNot(BeNil())
				Expect(getTrueClientIpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWebApplicationFirewallOptions successfully`, func() {
				// Construct an instance of the GetWebApplicationFirewallOptions model
				getWebApplicationFirewallOptionsModel := zonesSettingsService.NewGetWebApplicationFirewallOptions()
				getWebApplicationFirewallOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWebApplicationFirewallOptionsModel).ToNot(BeNil())
				Expect(getWebApplicationFirewallOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWebSocketsOptions successfully`, func() {
				// Construct an instance of the GetWebSocketsOptions model
				getWebSocketsOptionsModel := zonesSettingsService.NewGetWebSocketsOptions()
				getWebSocketsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWebSocketsOptionsModel).ToNot(BeNil())
				Expect(getWebSocketsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneCnameFlatteningOptions successfully`, func() {
				// Construct an instance of the GetZoneCnameFlatteningOptions model
				getZoneCnameFlatteningOptionsModel := zonesSettingsService.NewGetZoneCnameFlatteningOptions()
				getZoneCnameFlatteningOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneCnameFlatteningOptionsModel).ToNot(BeNil())
				Expect(getZoneCnameFlatteningOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetZoneDnssecOptions successfully`, func() {
				// Construct an instance of the GetZoneDnssecOptions model
				getZoneDnssecOptionsModel := zonesSettingsService.NewGetZoneDnssecOptions()
				getZoneDnssecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getZoneDnssecOptionsModel).ToNot(BeNil())
				Expect(getZoneDnssecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMinifySettingValue successfully`, func() {
				css := "false"
				html := "false"
				js := "false"
				model, err := zonesSettingsService.NewMinifySettingValue(css, html, js)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMobileRedirecSettingValue successfully`, func() {
				status := "true"
				mobileSubdomain := "m"
				stripURI := false
				model, err := zonesSettingsService.NewMobileRedirecSettingValue(status, mobileSubdomain, stripURI)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSecurityHeaderSettingValue successfully`, func() {
				var strictTransportSecurity *zonessettingsv1.SecurityHeaderSettingValueStrictTransportSecurity = nil
				_, err := zonesSettingsService.NewSecurityHeaderSettingValue(strictTransportSecurity)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewSecurityHeaderSettingValueStrictTransportSecurity successfully`, func() {
				enabled := true
				maxAge := int64(86400)
				includeSubdomains := true
				nosniff := true
				model, err := zonesSettingsService.NewSecurityHeaderSettingValueStrictTransportSecurity(enabled, maxAge, includeSubdomains, nosniff)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAlwaysUseHttpsOptions successfully`, func() {
				// Construct an instance of the UpdateAlwaysUseHttpsOptions model
				updateAlwaysUseHttpsOptionsModel := zonesSettingsService.NewUpdateAlwaysUseHttpsOptions()
				updateAlwaysUseHttpsOptionsModel.SetValue("true")
				updateAlwaysUseHttpsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAlwaysUseHttpsOptionsModel).ToNot(BeNil())
				Expect(updateAlwaysUseHttpsOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateAlwaysUseHttpsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAutomaticHttpsRewritesOptions successfully`, func() {
				// Construct an instance of the UpdateAutomaticHttpsRewritesOptions model
				updateAutomaticHttpsRewritesOptionsModel := zonesSettingsService.NewUpdateAutomaticHttpsRewritesOptions()
				updateAutomaticHttpsRewritesOptionsModel.SetValue("false")
				updateAutomaticHttpsRewritesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAutomaticHttpsRewritesOptionsModel).ToNot(BeNil())
				Expect(updateAutomaticHttpsRewritesOptionsModel.Value).To(Equal(core.StringPtr("false")))
				Expect(updateAutomaticHttpsRewritesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBrowserCheckOptions successfully`, func() {
				// Construct an instance of the UpdateBrowserCheckOptions model
				updateBrowserCheckOptionsModel := zonesSettingsService.NewUpdateBrowserCheckOptions()
				updateBrowserCheckOptionsModel.SetValue("true")
				updateBrowserCheckOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBrowserCheckOptionsModel).ToNot(BeNil())
				Expect(updateBrowserCheckOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateBrowserCheckOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateChallengeTtlOptions successfully`, func() {
				// Construct an instance of the UpdateChallengeTtlOptions model
				updateChallengeTtlOptionsModel := zonesSettingsService.NewUpdateChallengeTtlOptions()
				updateChallengeTtlOptionsModel.SetValue(int64(1800))
				updateChallengeTtlOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateChallengeTtlOptionsModel).ToNot(BeNil())
				Expect(updateChallengeTtlOptionsModel.Value).To(Equal(core.Int64Ptr(int64(1800))))
				Expect(updateChallengeTtlOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCiphersOptions successfully`, func() {
				// Construct an instance of the UpdateCiphersOptions model
				updateCiphersOptionsModel := zonesSettingsService.NewUpdateCiphersOptions()
				updateCiphersOptionsModel.SetValue([]string{"ECDHE-ECDSA-AES128-GCM-SHA256"})
				updateCiphersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCiphersOptionsModel).ToNot(BeNil())
				Expect(updateCiphersOptionsModel.Value).To(Equal([]string{"ECDHE-ECDSA-AES128-GCM-SHA256"}))
				Expect(updateCiphersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEnableErrorPagesOnOptions successfully`, func() {
				// Construct an instance of the UpdateEnableErrorPagesOnOptions model
				updateEnableErrorPagesOnOptionsModel := zonesSettingsService.NewUpdateEnableErrorPagesOnOptions()
				updateEnableErrorPagesOnOptionsModel.SetValue("true")
				updateEnableErrorPagesOnOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEnableErrorPagesOnOptionsModel).ToNot(BeNil())
				Expect(updateEnableErrorPagesOnOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateEnableErrorPagesOnOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateHotlinkProtectionOptions successfully`, func() {
				// Construct an instance of the UpdateHotlinkProtectionOptions model
				updateHotlinkProtectionOptionsModel := zonesSettingsService.NewUpdateHotlinkProtectionOptions()
				updateHotlinkProtectionOptionsModel.SetValue("true")
				updateHotlinkProtectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateHotlinkProtectionOptionsModel).ToNot(BeNil())
				Expect(updateHotlinkProtectionOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateHotlinkProtectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateHttp2Options successfully`, func() {
				// Construct an instance of the UpdateHttp2Options model
				updateHttp2OptionsModel := zonesSettingsService.NewUpdateHttp2Options()
				updateHttp2OptionsModel.SetValue("true")
				updateHttp2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateHttp2OptionsModel).ToNot(BeNil())
				Expect(updateHttp2OptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateHttp2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateImageLoadOptimizationOptions successfully`, func() {
				// Construct an instance of the UpdateImageLoadOptimizationOptions model
				updateImageLoadOptimizationOptionsModel := zonesSettingsService.NewUpdateImageLoadOptimizationOptions()
				updateImageLoadOptimizationOptionsModel.SetValue("true")
				updateImageLoadOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateImageLoadOptimizationOptionsModel).ToNot(BeNil())
				Expect(updateImageLoadOptimizationOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateImageLoadOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateImageSizeOptimizationOptions successfully`, func() {
				// Construct an instance of the UpdateImageSizeOptimizationOptions model
				updateImageSizeOptimizationOptionsModel := zonesSettingsService.NewUpdateImageSizeOptimizationOptions()
				updateImageSizeOptimizationOptionsModel.SetValue("lossless")
				updateImageSizeOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateImageSizeOptimizationOptionsModel).ToNot(BeNil())
				Expect(updateImageSizeOptimizationOptionsModel.Value).To(Equal(core.StringPtr("lossless")))
				Expect(updateImageSizeOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateIpGeolocationOptions successfully`, func() {
				// Construct an instance of the UpdateIpGeolocationOptions model
				updateIpGeolocationOptionsModel := zonesSettingsService.NewUpdateIpGeolocationOptions()
				updateIpGeolocationOptionsModel.SetValue("true")
				updateIpGeolocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateIpGeolocationOptionsModel).ToNot(BeNil())
				Expect(updateIpGeolocationOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateIpGeolocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateIpv6Options successfully`, func() {
				// Construct an instance of the UpdateIpv6Options model
				updateIpv6OptionsModel := zonesSettingsService.NewUpdateIpv6Options()
				updateIpv6OptionsModel.SetValue("true")
				updateIpv6OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateIpv6OptionsModel).ToNot(BeNil())
				Expect(updateIpv6OptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateIpv6OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMaxUploadOptions successfully`, func() {
				// Construct an instance of the UpdateMaxUploadOptions model
				updateMaxUploadOptionsModel := zonesSettingsService.NewUpdateMaxUploadOptions()
				updateMaxUploadOptionsModel.SetValue(int64(300))
				updateMaxUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMaxUploadOptionsModel).ToNot(BeNil())
				Expect(updateMaxUploadOptionsModel.Value).To(Equal(core.Int64Ptr(int64(300))))
				Expect(updateMaxUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMinTlsVersionOptions successfully`, func() {
				// Construct an instance of the UpdateMinTlsVersionOptions model
				updateMinTlsVersionOptionsModel := zonesSettingsService.NewUpdateMinTlsVersionOptions()
				updateMinTlsVersionOptionsModel.SetValue("1.2")
				updateMinTlsVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMinTlsVersionOptionsModel).ToNot(BeNil())
				Expect(updateMinTlsVersionOptionsModel.Value).To(Equal(core.StringPtr("1.2")))
				Expect(updateMinTlsVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMinifyOptions successfully`, func() {
				// Construct an instance of the MinifySettingValue model
				minifySettingValueModel := new(zonessettingsv1.MinifySettingValue)
				Expect(minifySettingValueModel).ToNot(BeNil())
				minifySettingValueModel.Css = core.StringPtr("false")
				minifySettingValueModel.HTML = core.StringPtr("false")
				minifySettingValueModel.Js = core.StringPtr("false")
				Expect(minifySettingValueModel.Css).To(Equal(core.StringPtr("false")))
				Expect(minifySettingValueModel.HTML).To(Equal(core.StringPtr("false")))
				Expect(minifySettingValueModel.Js).To(Equal(core.StringPtr("false")))

				// Construct an instance of the UpdateMinifyOptions model
				updateMinifyOptionsModel := zonesSettingsService.NewUpdateMinifyOptions()
				updateMinifyOptionsModel.SetValue(minifySettingValueModel)
				updateMinifyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMinifyOptionsModel).ToNot(BeNil())
				Expect(updateMinifyOptionsModel.Value).To(Equal(minifySettingValueModel))
				Expect(updateMinifyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateMobileRedirectOptions successfully`, func() {
				// Construct an instance of the MobileRedirecSettingValue model
				mobileRedirecSettingValueModel := new(zonessettingsv1.MobileRedirecSettingValue)
				Expect(mobileRedirecSettingValueModel).ToNot(BeNil())
				mobileRedirecSettingValueModel.Status = core.StringPtr("true")
				mobileRedirecSettingValueModel.MobileSubdomain = core.StringPtr("m")
				mobileRedirecSettingValueModel.StripURI = core.BoolPtr(false)
				Expect(mobileRedirecSettingValueModel.Status).To(Equal(core.StringPtr("true")))
				Expect(mobileRedirecSettingValueModel.MobileSubdomain).To(Equal(core.StringPtr("m")))
				Expect(mobileRedirecSettingValueModel.StripURI).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the UpdateMobileRedirectOptions model
				updateMobileRedirectOptionsModel := zonesSettingsService.NewUpdateMobileRedirectOptions()
				updateMobileRedirectOptionsModel.SetValue(mobileRedirecSettingValueModel)
				updateMobileRedirectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateMobileRedirectOptionsModel).ToNot(BeNil())
				Expect(updateMobileRedirectOptionsModel.Value).To(Equal(mobileRedirecSettingValueModel))
				Expect(updateMobileRedirectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOpportunisticEncryptionOptions successfully`, func() {
				// Construct an instance of the UpdateOpportunisticEncryptionOptions model
				updateOpportunisticEncryptionOptionsModel := zonesSettingsService.NewUpdateOpportunisticEncryptionOptions()
				updateOpportunisticEncryptionOptionsModel.SetValue("false")
				updateOpportunisticEncryptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOpportunisticEncryptionOptionsModel).ToNot(BeNil())
				Expect(updateOpportunisticEncryptionOptionsModel.Value).To(Equal(core.StringPtr("false")))
				Expect(updateOpportunisticEncryptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePrefetchPreloadOptions successfully`, func() {
				// Construct an instance of the UpdatePrefetchPreloadOptions model
				updatePrefetchPreloadOptionsModel := zonesSettingsService.NewUpdatePrefetchPreloadOptions()
				updatePrefetchPreloadOptionsModel.SetValue("true")
				updatePrefetchPreloadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePrefetchPreloadOptionsModel).ToNot(BeNil())
				Expect(updatePrefetchPreloadOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updatePrefetchPreloadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePseudoIpv4Options successfully`, func() {
				// Construct an instance of the UpdatePseudoIpv4Options model
				updatePseudoIpv4OptionsModel := zonesSettingsService.NewUpdatePseudoIpv4Options()
				updatePseudoIpv4OptionsModel.SetValue("add_header")
				updatePseudoIpv4OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePseudoIpv4OptionsModel).ToNot(BeNil())
				Expect(updatePseudoIpv4OptionsModel.Value).To(Equal(core.StringPtr("add_header")))
				Expect(updatePseudoIpv4OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResponseBufferingOptions successfully`, func() {
				// Construct an instance of the UpdateResponseBufferingOptions model
				updateResponseBufferingOptionsModel := zonesSettingsService.NewUpdateResponseBufferingOptions()
				updateResponseBufferingOptionsModel.SetValue("true")
				updateResponseBufferingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResponseBufferingOptionsModel).ToNot(BeNil())
				Expect(updateResponseBufferingOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateResponseBufferingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateScriptLoadOptimizationOptions successfully`, func() {
				// Construct an instance of the UpdateScriptLoadOptimizationOptions model
				updateScriptLoadOptimizationOptionsModel := zonesSettingsService.NewUpdateScriptLoadOptimizationOptions()
				updateScriptLoadOptimizationOptionsModel.SetValue("true")
				updateScriptLoadOptimizationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateScriptLoadOptimizationOptionsModel).ToNot(BeNil())
				Expect(updateScriptLoadOptimizationOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateScriptLoadOptimizationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSecurityHeaderOptions successfully`, func() {
				// Construct an instance of the SecurityHeaderSettingValueStrictTransportSecurity model
				securityHeaderSettingValueStrictTransportSecurityModel := new(zonessettingsv1.SecurityHeaderSettingValueStrictTransportSecurity)
				Expect(securityHeaderSettingValueStrictTransportSecurityModel).ToNot(BeNil())
				securityHeaderSettingValueStrictTransportSecurityModel.Enabled = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.MaxAge = core.Int64Ptr(int64(86400))
				securityHeaderSettingValueStrictTransportSecurityModel.IncludeSubdomains = core.BoolPtr(true)
				securityHeaderSettingValueStrictTransportSecurityModel.Nosniff = core.BoolPtr(true)
				Expect(securityHeaderSettingValueStrictTransportSecurityModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(securityHeaderSettingValueStrictTransportSecurityModel.MaxAge).To(Equal(core.Int64Ptr(int64(86400))))
				Expect(securityHeaderSettingValueStrictTransportSecurityModel.IncludeSubdomains).To(Equal(core.BoolPtr(true)))
				Expect(securityHeaderSettingValueStrictTransportSecurityModel.Nosniff).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SecurityHeaderSettingValue model
				securityHeaderSettingValueModel := new(zonessettingsv1.SecurityHeaderSettingValue)
				Expect(securityHeaderSettingValueModel).ToNot(BeNil())
				securityHeaderSettingValueModel.StrictTransportSecurity = securityHeaderSettingValueStrictTransportSecurityModel
				Expect(securityHeaderSettingValueModel.StrictTransportSecurity).To(Equal(securityHeaderSettingValueStrictTransportSecurityModel))

				// Construct an instance of the UpdateSecurityHeaderOptions model
				updateSecurityHeaderOptionsModel := zonesSettingsService.NewUpdateSecurityHeaderOptions()
				updateSecurityHeaderOptionsModel.SetValue(securityHeaderSettingValueModel)
				updateSecurityHeaderOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSecurityHeaderOptionsModel).ToNot(BeNil())
				Expect(updateSecurityHeaderOptionsModel.Value).To(Equal(securityHeaderSettingValueModel))
				Expect(updateSecurityHeaderOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServerSideExcludeOptions successfully`, func() {
				// Construct an instance of the UpdateServerSideExcludeOptions model
				updateServerSideExcludeOptionsModel := zonesSettingsService.NewUpdateServerSideExcludeOptions()
				updateServerSideExcludeOptionsModel.SetValue("true")
				updateServerSideExcludeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServerSideExcludeOptionsModel).ToNot(BeNil())
				Expect(updateServerSideExcludeOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateServerSideExcludeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTlsClientAuthOptions successfully`, func() {
				// Construct an instance of the UpdateTlsClientAuthOptions model
				updateTlsClientAuthOptionsModel := zonesSettingsService.NewUpdateTlsClientAuthOptions()
				updateTlsClientAuthOptionsModel.SetValue("true")
				updateTlsClientAuthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTlsClientAuthOptionsModel).ToNot(BeNil())
				Expect(updateTlsClientAuthOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateTlsClientAuthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTrueClientIpOptions successfully`, func() {
				// Construct an instance of the UpdateTrueClientIpOptions model
				updateTrueClientIpOptionsModel := zonesSettingsService.NewUpdateTrueClientIpOptions()
				updateTrueClientIpOptionsModel.SetValue("true")
				updateTrueClientIpOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTrueClientIpOptionsModel).ToNot(BeNil())
				Expect(updateTrueClientIpOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateTrueClientIpOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWebApplicationFirewallOptions successfully`, func() {
				// Construct an instance of the UpdateWebApplicationFirewallOptions model
				updateWebApplicationFirewallOptionsModel := zonesSettingsService.NewUpdateWebApplicationFirewallOptions()
				updateWebApplicationFirewallOptionsModel.SetValue("true")
				updateWebApplicationFirewallOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWebApplicationFirewallOptionsModel).ToNot(BeNil())
				Expect(updateWebApplicationFirewallOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateWebApplicationFirewallOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWebSocketsOptions successfully`, func() {
				// Construct an instance of the UpdateWebSocketsOptions model
				updateWebSocketsOptionsModel := zonesSettingsService.NewUpdateWebSocketsOptions()
				updateWebSocketsOptionsModel.SetValue("true")
				updateWebSocketsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWebSocketsOptionsModel).ToNot(BeNil())
				Expect(updateWebSocketsOptionsModel.Value).To(Equal(core.StringPtr("true")))
				Expect(updateWebSocketsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneCnameFlatteningOptions successfully`, func() {
				// Construct an instance of the UpdateZoneCnameFlatteningOptions model
				updateZoneCnameFlatteningOptionsModel := zonesSettingsService.NewUpdateZoneCnameFlatteningOptions()
				updateZoneCnameFlatteningOptionsModel.SetValue("flatten_all")
				updateZoneCnameFlatteningOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneCnameFlatteningOptionsModel).ToNot(BeNil())
				Expect(updateZoneCnameFlatteningOptionsModel.Value).To(Equal(core.StringPtr("flatten_all")))
				Expect(updateZoneCnameFlatteningOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateZoneDnssecOptions successfully`, func() {
				// Construct an instance of the UpdateZoneDnssecOptions model
				updateZoneDnssecOptionsModel := zonesSettingsService.NewUpdateZoneDnssecOptions()
				updateZoneDnssecOptionsModel.SetStatus("active")
				updateZoneDnssecOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateZoneDnssecOptionsModel).ToNot(BeNil())
				Expect(updateZoneDnssecOptionsModel.Status).To(Equal(core.StringPtr("active")))
				Expect(updateZoneDnssecOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
