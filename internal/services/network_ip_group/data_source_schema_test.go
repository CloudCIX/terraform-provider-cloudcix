// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_ip_group"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkIPGroupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_ip_group.NetworkIPGroupDataSourceModel)(nil)
	schema := network_ip_group.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
