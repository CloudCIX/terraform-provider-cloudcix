// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

import (
	"context"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkFirewallDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Management of Network Firewall Rules\n\nThis module provides API endpoints for managing network firewall rules in the CloudCIX Compute platform.\nEach project can have ONE project firewall and ONE geo firewall. These firewalls control traffic\nflow to and from your cloud resources by defining rules that allow or block network traffic.\n\nFirewall Types:\n1. Project Firewall (type: \"project\") - Fine-grained rules controlling traffic based on specific source/destination\n   IP addresses, ports, and protocols. Can reference your member's IP Address Groups using '@groupname' syntax\n   in source/destination fields (e.g., \"source\": \"@office_networks\").\n\n2. Geo Firewall (type: \"geo\") - Country-based IP filtering using global IP Address Groups (member_id = 0) that\n   contain IP ranges for specific countries/regions. References groups by numeric ID using the\n   \"ip_address_group_id\" field (e.g., \"ip_address_group_id\": 123).\n\nIP Address Group Usage:\n- Project Firewalls: Use \"source\": \"@groupname\" or \"destination\": \"@groupname\" (member groups only)\n- Geo Firewalls: Use \"ip_address_group_id\": 123 (global groups with member_id=0 only)\n- Project firewalls cannot use global groups, and geo firewalls cannot use member groups\n\nCRITICAL: When updating firewall rules, you MUST include ALL rules you want to keep. The update operation\nreplaces the entire rule list - any rules not included in the update will be permanently deleted.\n\nAvailable operations:\n- List and filter firewall rules across your projects by type\n- Create a project's single firewall or geo firewall with complete rule definitions\n- Retrieve detailed information about individual firewall configurations\n- Update firewall rules (replaces ALL existing rules) or delete firewalls by changing their state\n\nRule Direction:\nEach firewall rule specifies traffic direction using the 'inbound' flag:\n- inbound: true = Inbound rule (traffic coming INTO your project/network)\n- inbound: false = Outbound rule (traffic going OUT FROM your project/network)\nRules default to outbound (inbound: false) if not specified.\n\nEach firewall includes its associated project, router, rule definitions, and current state.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network Firewall record was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name given to this Network Firewall instance",
				Computed:    true,
			},
			"project_id": schema.Int64Attribute{
				Description: "The id of the Project that this Network Firewall belongs to",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The current state of the Network Firewall",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Network Firewall",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network Firewall record was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Network Firewall instance.",
				Computed:    true,
			},
			"rules": schema.ListNestedAttribute{
				Description: "List of rules for this Network Firewall.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkFirewallRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow": schema.BoolAttribute{
							Description: "True to allow traffic, False to deny.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: `Optional description of the rule. Returned if the type is "project".`,
							Computed:    true,
						},
						"destination": schema.StringAttribute{
							Description: `Destination address or subnet. Use * for any. Returned if the type is "project".`,
							Computed:    true,
						},
						"group_name": schema.StringAttribute{
							Description: `The name of the Geo IP Address Group. Returned if the type is "geo".`,
							Computed:    true,
						},
						"inbound": schema.BoolAttribute{
							Description: "True if the rule applies to inbound traffic.",
							Computed:    true,
						},
						"order": schema.Int64Attribute{
							Description: `Order of rule evaluation (lower runs first). Returned if the type is "project".`,
							Computed:    true,
						},
						"port": schema.StringAttribute{
							Description: "Port or port range (e.g. 80, 443, 1000-2000). Not required for ICMP or ANY.\nReturned if the type is \"project\".",
							Computed:    true,
						},
						"protocol": schema.StringAttribute{
							Description: `Network protocol (any, icmp, tcp, udp). Returned if the type is "project".`,
							Computed:    true,
						},
						"source": schema.StringAttribute{
							Description: `Source address or subnet. Use * for any. Returned if the type is "project".`,
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: `IP version (4 or 6). Returned if the type is "project".`,
							Computed:    true,
						},
						"zone": schema.StringAttribute{
							Description: "The zone in the firewall that the rule is applied to.",
							Computed:    true,
						},
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Network Firewall",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[NetworkFirewallSpecsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"quantity": schema.Int64Attribute{
							Description: "How many units of a billable entity that a Resource utilises",
							Computed:    true,
						},
						"sku_name": schema.StringAttribute{
							Description: "An identifier for a billable entity that a Resource utilises",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkFirewallDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkFirewallDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
