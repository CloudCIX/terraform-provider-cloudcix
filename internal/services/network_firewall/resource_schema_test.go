// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/network_firewall"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestNetworkFirewallModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*network_firewall.NetworkFirewallModel)(nil)
	schema := network_firewall.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)

	errs.Ignore(".@NetworkFirewallModel.timeouts.@ObjectValue.create")
	errs.Ignore(".@NetworkFirewallModel.timeouts.@ObjectValue.update")
	errs.Ignore(".@NetworkFirewallModel.timeouts.@ObjectValue.read")
	errs.Ignore(".@NetworkFirewallModel.timeouts.@ObjectValue.delete")

	errs.Report(t)
}
