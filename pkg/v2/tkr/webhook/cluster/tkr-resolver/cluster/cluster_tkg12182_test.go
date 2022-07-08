// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

//go:build tkg12182
// +build tkg12182

package cluster

import (
	"os"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/yaml"

	runv1 "github.com/vmware-tanzu/tanzu-framework/apis/run/v1alpha3"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v2/tkr/resolver"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v2/tkr/resolver/data"
)

var _ = Describe("TKG-12182", func() {
	var (
		cw           *Webhook
		clusterClass *clusterv1.ClusterClass
		cluster      *clusterv1.Cluster
		tkrs         data.TKRs
		osImages     data.OSImages
	)

	BeforeEach(func() {
		tkrResolver := resolver.New()

		tkrs, osImages = loadTKG12182Objects()
		for _, tkr := range tkrs {
			tkrResolver.Add(tkr)
		}
		for _, osImage := range osImages {
			tkrResolver.Add(osImage)
		}

		cw = &Webhook{
			TKRResolver: tkrResolver,
			Log:         logr.Discard(),
		}

		cluster = loadTKG12182Cluster()
		clusterClass = loadTKG12182ClusterClass()
	})

	FIt("the webhook should resolve the cluster", func() {
		err := cw.ResolveAndSetMetadata(cluster, clusterClass)
		Expect(err).To(BeNil())
		Expect(cluster.Labels[runv1.LabelTKR]).To(Equal("v1.22.8---vmware.1-tkg.2-zshippable"))
	})
})

func loadTKG12182Objects() (data.TKRs, data.OSImages) {
	tkrList := loadTKG12182TKRList()
	tkrs := data.TKRs{}
	for i := range tkrList.Items {
		tkr := &tkrList.Items[i]
		tkrs[tkr.Name] = tkr
	}

	osImageList := loadTKG12182OSImageList()
	osImages := data.OSImages{}
	for i := range osImageList.Items {
		osImage := &osImageList.Items[i]
		osImages[osImage.Name] = osImage
	}

	return tkrs, osImages
}

func loadTKG12182TKRList() *runv1.TanzuKubernetesReleaseList {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12182-cannot-resolve/tkr.yaml")
	if err != nil {
		panic(err)
	}
	tkrList := &runv1.TanzuKubernetesReleaseList{}
	if err := yaml.Unmarshal(bytes, tkrList); err != nil {
		panic(err)
	}
	return tkrList
}

func loadTKG12182OSImageList() *runv1.OSImageList {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12182-cannot-resolve/osimages.yaml")
	if err != nil {
		panic(err)
	}
	osImageList := &runv1.OSImageList{}
	if err := yaml.Unmarshal(bytes, osImageList); err != nil {
		panic(err)
	}
	return osImageList
}

func loadTKG12182Cluster() *clusterv1.Cluster {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12182-cannot-resolve/cluster.yaml")
	if err != nil {
		panic(err)
	}
	cluster := &clusterv1.Cluster{}
	if err := yaml.Unmarshal(bytes, cluster); err != nil {
		panic(err)
	}
	return cluster
}

func loadTKG12182ClusterClass() *clusterv1.ClusterClass {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12182-cannot-resolve/clusterclass.yaml")
	if err != nil {
		panic(err)
	}
	cluster := &clusterv1.ClusterClass{}
	if err := yaml.Unmarshal(bytes, cluster); err != nil {
		panic(err)
	}
	return cluster
}
