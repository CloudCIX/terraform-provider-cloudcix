// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_ssh_key

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSSHKeyContentDataSourceEnvelope struct {
	Content ComputeSSHKeyDataSourceModel `json:"content,computed"`
}

type ComputeSSHKeyDataSourceModel struct {
	ID         types.Int64  `tfsdk:"id" path:"id,required"`
	Created    types.String `tfsdk:"created" json:"created,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
	PrivateKey types.String `tfsdk:"private_key" json:"private_key,computed"`
	PublicKey  types.String `tfsdk:"public_key" json:"public_key,computed"`
}
