// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_router"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkRouterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_router.NetworkRouterDataSourceModel)(nil)
	schema := network_router.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
