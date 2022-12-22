// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package fetcher

import (
	"context"
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-framework/apis/run/util/sets"
	"github.com/vmware-tanzu/tanzu-framework/apis/run/util/version"
	"github.com/vmware-tanzu/tanzu-framework/tkr/controller/tkr-source/registry"
)

func TestCompatibility(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteConfig, _ := GinkgoConfiguration()
	suiteConfig.FailFast = true
	RunSpecs(t, "TKR Compatibility: unit tests", suiteConfig)
}

var _ = Describe("Fetcher", func() {
	var (
		f             *Fetcher
		compatibility version.Compatibility
		r             registry.Registry
	)

	JustBeforeEach(func() {
		f = &Fetcher{
			Log:           logr.Discard(),
			Config:        Config{},
			Registry:      r,
			Compatibility: compatibility,
		}
	})

	Describe("imageTagsToPull()", func() {
		Context("For versions that are known to be compatible", func() {
			var (
				compatibleVersions sets.StringSet
			)

			BeforeEach(func() {
				compatibleVersions = sets.Strings("v1.24.6+vmware.1-tkg.1-rc.1", "v1.23.13+vmware.1-tkg.1-rc.1", "v1.22.11+vmware.2-tkg.2-rc.1")
				compatibility = &fakeCompat{
					result: compatibleVersions,
				}
			})
			When("corresponding TKR package bundle image tags exist", func() {
				BeforeEach(func() {
					r = &fakeRegistry{
						imageTags: []string{
							"v1.22.11_vmware.2-tkg.2",
							"v1.22.11_vmware.2-tkg.2-fc.1",
							"v1.22.11_vmware.2-tkg.2-fc.2",
							"v1.22.11_vmware.2-tkg.2-rc.1",
							"v1.22.11_vmware.2-tkg.2-tf-v0.26.0",
							"v1.23.12_vmware.1-tkg.2-zshippable",
							"v1.23.13_vmware.1-tkg.1-fc.1",
							"v1.23.13_vmware.1-tkg.1-fc.2",
							"v1.23.13_vmware.1-tkg.1-rc.1",
							"v1.23.13_vmware.1-tkg.1-zshippable",
							"v1.24.6_vmware.1-tkg.1-fc.1",
							"v1.24.6_vmware.1-tkg.1-fc.2",
							"v1.24.6_vmware.1-tkg.1-rc.1",
							"v1.24.6_vmware.1-tkg.1-zshippable",
						},
					}
				})
				It("should include the tags in the returned set", func() {
					imageTagsToPull, err := f.imageTagsToPull(context.Background())
					Expect(err).ToNot(HaveOccurred())
					Expect(imageTagsToPull).To(HaveLen(len(compatibleVersions)))
				})
			})
		})
	})
})

type fakeCompat struct {
	result sets.StringSet
	err    error
}

func (f fakeCompat) CompatibleVersions(ctx context.Context) (sets.StringSet, error) {
	return f.result, f.err
}

type fakeRegistry struct {
	imageTags []string
	err       error
}

func (f fakeRegistry) ListImageTags(imageName string) ([]string, error) {
	return f.imageTags, f.err
}

func (fakeRegistry) GetFile(imageWithTag string, filename string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (fakeRegistry) GetFiles(imageWithTag string) (map[string][]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (fakeRegistry) DownloadBundle(imageName, outputDir string) error {
	// TODO implement me
	panic("implement me")
}
