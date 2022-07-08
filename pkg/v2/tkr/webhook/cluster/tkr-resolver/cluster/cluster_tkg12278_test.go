// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

//go:build tkg12278
// +build tkg12278

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

var _ = Describe("TKG-12278", func() {
	var (
		cw           *Webhook
		clusterClass *clusterv1.ClusterClass
		cluster      *clusterv1.Cluster
		tkrs         data.TKRs
		osImages     data.OSImages
	)

	BeforeEach(func() {
		tkrResolver := resolver.New()

		tkrs, osImages = loadObjects()
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

		cluster = loadCluster()
		clusterClass = loadClusterClass()
	})

	FIt("the webhook should resolve the cluster", func() {
		err := cw.ResolveAndSetMetadata(cluster, clusterClass)
		Expect(err).To(BeNil())
		Expect(cluster.Labels[runv1.LabelTKR]).To(Equal("v1.20.12---vmware.1-tkg.1.10b2767"))
		Expect(cluster.Spec.Topology.Version).To(Equal("v1.20.12+vmware.1"))
	})
})

func loadObjects() (data.TKRs, data.OSImages) {
	tkrList := loadTKRList()
	tkrs := data.TKRs{}
	for i := range tkrList.Items {
		tkr := &tkrList.Items[i]
		tkrs[tkr.Name] = tkr
	}

	osImageList := loadOSImageList()
	osImages := data.OSImages{}
	for i := range osImageList.Items {
		osImage := &osImageList.Items[i]
		osImages[osImage.Name] = osImage
	}

	return tkrs, osImages
}

func loadTKRList() *runv1.TanzuKubernetesReleaseList {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12278-should-not-re-resolve/tkr.yaml")
	if err != nil {
		panic(err)
	}
	tkrList := &runv1.TanzuKubernetesReleaseList{}
	if err := yaml.Unmarshal(bytes, tkrList); err != nil {
		panic(err)
	}
	return tkrList
}

func loadOSImageList() *runv1.OSImageList {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12278-should-not-re-resolve/osimage.yaml")
	if err != nil {
		panic(err)
	}
	osImageList := &runv1.OSImageList{}
	if err := yaml.Unmarshal(bytes, osImageList); err != nil {
		panic(err)
	}
	return osImageList
}

func loadCluster() *clusterv1.Cluster {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12278-should-not-re-resolve/cluster.yaml")
	if err != nil {
		panic(err)
	}
	cluster := &clusterv1.Cluster{}
	if err := yaml.Unmarshal(bytes, cluster); err != nil {
		panic(err)
	}
	return cluster
}

func loadClusterClass() *clusterv1.ClusterClass {
	bytes, err := os.ReadFile("/Users/ivan/src/vmware-tanzu/tanzu-framework/tmp/TKG-12278-should-not-re-resolve/clusterclass.yaml")
	if err != nil {
		panic(err)
	}
	clusterClass := &clusterv1.ClusterClass{}
	if err := yaml.Unmarshal(bytes, clusterClass); err != nil {
		panic(err)
	}
	return clusterClass
}
