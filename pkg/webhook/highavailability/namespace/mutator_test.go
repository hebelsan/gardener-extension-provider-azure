// Copyright (c) 2024 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespace

import (
	"context"
	"testing"

	"github.com/gardener/gardener/pkg/apis/resources/v1alpha1"
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/gardener/gardener-extension-provider-azure/pkg/azure"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Topology Webhook Suite")
}

var _ = Describe("Topology", func() {
	var (
		ctrl *gomock.Controller

		ctx     context.Context
		ns      *corev1.Namespace
		mutator *handler
		region  = "westeurope"
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mutator = New(&admission.Decoder{}, logr.Discard(), AddOptions{
			SeedRegion:   region,
			SeedProvider: azure.Type,
		})

		ns = &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "foo",
				Annotations: map[string]string{
					v1alpha1.HighAvailabilityConfigZones: "2,westeurope-1,3",
				},
			},
		}

		var err error
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("#Webhook", func() {
		Describe("#Mutate", func() {
			It("it should correctly mutate required", func() {
				err := mutator.Mutate(ctx, ns, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(ns.Annotations[v1alpha1.HighAvailabilityConfigZones]).To(Equal("1,2,3"))
			})

			It("it should not mutate if the zone annotation is missing", func() {
				ns.Annotations[v1alpha1.HighAvailabilityConfigZones] = ""
				nsOld := ns.DeepCopy()
				err := mutator.Mutate(ctx, ns, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(nsOld).To(Equal(ns))
			})

		})
	})
})
