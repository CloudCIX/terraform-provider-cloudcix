// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_router_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_router"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkRouterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_router.NetworkRouterModel)(nil)
	schema := network_router.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)

	errs.Ignore(".@NetworkRouterModel.timeouts.@ObjectValue.create")
	errs.Ignore(".@NetworkRouterModel.timeouts.@ObjectValue.update")
	errs.Ignore(".@NetworkRouterModel.timeouts.@ObjectValue.read")
	errs.Ignore(".@NetworkRouterModel.timeouts.@ObjectValue.delete")
	
	errs.Report(t)
}
