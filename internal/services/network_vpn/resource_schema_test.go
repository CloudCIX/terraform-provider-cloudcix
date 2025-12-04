// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_vpn"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkVpnModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_vpn.NetworkVpnModel)(nil)
	schema := network_vpn.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
