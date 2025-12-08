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

	ignore_list := []string{
		".@NetworkFirewallModel.timeouts.@ObjectValue.read",
		".@NetworkFirewallModel.timeouts.@ObjectValue.create",
		".@NetworkFirewallModel.timeouts.@ObjectValue.update",
		".@NetworkFirewallModel.timeouts.@ObjectValue.delete",
	}

	for _, item := range ignore_list {
		errs.IgnoreAll(t, item)
	}

	errs.Report(t)
}
