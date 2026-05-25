// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_ssh_key

import (
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeSSHKeyContentEnvelope struct {
	Content ComputeSSHKeyModel `json:"content"`
}

type ComputeSSHKeyModel struct {
	ID         types.Int64  `tfsdk:"id" json:"id,computed"`
	Created    types.String `tfsdk:"created" json:"created,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
	PrivateKey types.String `tfsdk:"private_key" json:"private_key,computed"`
	PublicKey  types.String `tfsdk:"public_key" json:"public_key,computed"`
}

func (m ComputeSSHKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeSSHKeyModel) MarshalJSONForUpdate(state ComputeSSHKeyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
