/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dbwatch

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("parseOrganizationData", func() {
	It("parses displayName and emailDomains from JSONB", func() {
		data := []byte(`{"displayName": "ACME Corp", "emailDomains": ["acme.com", "acme.org"]}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(Equal("ACME Corp"))
		Expect(emailDomains).To(Equal([]string{"acme.com", "acme.org"}))
	})

	It("returns empty values for empty JSON object", func() {
		data := []byte(`{}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(BeEmpty())
		Expect(emailDomains).To(BeNil())
	})

	It("returns empty values for null data", func() {
		displayName, emailDomains := parseOrganizationData(nil)
		Expect(displayName).To(BeEmpty())
		Expect(emailDomains).To(BeNil())
	})

	It("returns empty values for empty data", func() {
		displayName, emailDomains := parseOrganizationData([]byte{})
		Expect(displayName).To(BeEmpty())
		Expect(emailDomains).To(BeNil())
	})

	It("returns empty values for invalid JSON", func() {
		displayName, emailDomains := parseOrganizationData([]byte(`not json`))
		Expect(displayName).To(BeEmpty())
		Expect(emailDomains).To(BeNil())
	})

	It("handles missing emailDomains field", func() {
		data := []byte(`{"displayName": "ACME Corp"}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(Equal("ACME Corp"))
		Expect(emailDomains).To(BeNil())
	})

	It("handles missing displayName field", func() {
		data := []byte(`{"emailDomains": ["acme.com"]}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(BeEmpty())
		Expect(emailDomains).To(Equal([]string{"acme.com"}))
	})

	It("handles extra fields gracefully", func() {
		data := []byte(`{"displayName": "ACME Corp", "emailDomains": ["acme.com"], "someOtherField": 42}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(Equal("ACME Corp"))
		Expect(emailDomains).To(Equal([]string{"acme.com"}))
	})

	It("handles empty emailDomains array", func() {
		data := []byte(`{"displayName": "ACME Corp", "emailDomains": []}`)
		displayName, emailDomains := parseOrganizationData(data)
		Expect(displayName).To(Equal("ACME Corp"))
		Expect(emailDomains).To(Equal([]string{}))
	})
})
