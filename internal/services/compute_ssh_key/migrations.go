// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_ssh_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ComputeSSHKeyResource)(nil)

func (r *ComputeSSHKeyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
