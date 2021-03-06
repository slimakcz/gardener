// Copyright 2018 The Gardener Authors.
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

package openstackbotanist

import (
	"github.com/gardener/gardener/pkg/operation/common"
	"github.com/gardener/gardener/pkg/operation/terraformer"
)

// DeployInfrastructure kicks off a Terraform job which deploys the infrastructure.
func (b *OpenStackBotanist) DeployInfrastructure() error {
	var (
		routerID     = "${openstack_networking_router_v2.router.id}"
		createRouter = true
	)

	// check if we should use an existing Router or create a new one
	if b.Shoot.Info.Spec.Cloud.OpenStack.Networks.Router != nil {
		routerID = b.Shoot.Info.Spec.Cloud.OpenStack.Networks.Router.ID
		createRouter = false
	}

	return terraformer.
		New(b.Operation, common.TerraformerPurposeInfra).
		SetVariablesEnvironment(b.generateTerraformInfraVariablesEnvironment()).
		DefineConfig("openstack-infra", b.generateTerraformInfraConfig(createRouter, routerID)).
		Apply()
}

// DestroyInfrastructure kicks off a Terraform job which destroys the infrastructure.
func (b *OpenStackBotanist) DestroyInfrastructure() error {
	return terraformer.
		New(b.Operation, common.TerraformerPurposeInfra).
		SetVariablesEnvironment(b.generateTerraformInfraVariablesEnvironment()).
		Destroy()
}

// generateTerraformInfraVariablesEnvironment generates the environment containing the credentials which
// are required to validate/apply/destroy the Terraform configuration. These environment must contain
// Terraform variables which are prefixed with TF_VAR_.
func (b *OpenStackBotanist) generateTerraformInfraVariablesEnvironment() []map[string]interface{} {
	return common.GenerateTerraformVariablesEnvironment(b.Shoot.Secret, map[string]string{
		"USER_NAME": UserName,
		"PASSWORD":  Password,
	})
}

// generateTerraformInfraConfig creates the Terraform variables and the Terraform config (for the infrastructure)
// and returns them (these values will be stored as a ConfigMap and a Secret in the Garden cluster.
func (b *OpenStackBotanist) generateTerraformInfraConfig(createRouter bool, routerID string) map[string]interface{} {
	return map[string]interface{}{
		"openstack": map[string]interface{}{
			"authURL":              b.Shoot.CloudProfile.Spec.OpenStack.KeyStoneURL,
			"domainName":           string(b.Shoot.Secret.Data[DomainName]),
			"tenantName":           string(b.Shoot.Secret.Data[TenantName]),
			"region":               b.Shoot.Info.Spec.Cloud.Region,
			"floatingPoolName":     b.Shoot.Info.Spec.Cloud.OpenStack.FloatingPoolName,
			"loadBalancerProvider": b.Shoot.Info.Spec.Cloud.OpenStack.LoadBalancerProvider,
		},
		"create": map[string]interface{}{
			"router": createRouter,
		},
		"sshPublicKey": string(b.Secrets["ssh-keypair"].Data["id_rsa.pub"]),
		"router": map[string]interface{}{
			"id": routerID,
		},
		"clusterName": b.Shoot.SeedNamespace,
		"networks": map[string]interface{}{
			"worker": b.Shoot.Info.Spec.Cloud.OpenStack.Networks.Workers[0],
		},
	}
}

// DeployBackupInfrastructure kicks off a Terraform job which creates the infrastructure resources for backup.
func (b *OpenStackBotanist) DeployBackupInfrastructure() error {
	return nil
}

// DestroyBackupInfrastructure kicks off a Terraform job which destroys the infrastructure for backup.
func (b *OpenStackBotanist) DestroyBackupInfrastructure() error {
	return nil
}
