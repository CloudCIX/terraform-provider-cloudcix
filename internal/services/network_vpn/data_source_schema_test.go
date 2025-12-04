// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_vpn_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_vpn"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkVpnDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_vpn.NetworkVpnDataSourceModel)(nil)
	schema := network_vpn.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
