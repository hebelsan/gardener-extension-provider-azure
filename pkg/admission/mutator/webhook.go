// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package mutator

import (
	extensionspredicate "github.com/gardener/gardener/extensions/pkg/predicate"
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/gardener/gardener-extension-provider-azure/pkg/azure"
)

const (
	// Name is a name for a mutation webhook.
	Name = "mutator"
)

var logger = log.Log.WithName("azure-mutator-webhook")

// New creates a new webhook that mutates Shoot resources.
func New(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	logger.Info("Setting up webhook", "name", Name)

	return extensionswebhook.New(mgr, extensionswebhook.Args{
		Provider:   azure.Type,
		Name:       Name,
		Path:       "/webhooks/mutate",
		Predicates: []predicate.Predicate{extensionspredicate.GardenCoreProviderType(azure.Type)},
		Mutators: map[extensionswebhook.Mutator][]extensionswebhook.Type{
			NewShootMutator(mgr): {{Obj: &gardencorev1beta1.Shoot{}}},
		},
		Target: extensionswebhook.TargetSeed,
		ObjectSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{"provider.extensions.gardener.cloud/azure": "true"},
		},
	})
}
