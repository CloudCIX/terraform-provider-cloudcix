// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"context"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*NetworkIPGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Management of Network IP Groups\n\nIP address groups organise sets of CIDR networks for use in firewall rules and access control.\nTwo types are available:\n\n- Geo Groups (type=\"geo\"): Maintained by CloudCIX and accessible to all members\n  * Used for geo-filtering based on country IP ranges (e.g., 'Ireland', 'USA', 'China')\n  * Essential for creating geo firewalls that block/allow traffic from specific countries\n  * To list country groups: GET /ip_address_groups?search[member_id]=0\n  * Referenced in geo firewall rules by numeric ID: \"ip_address_group_id\": 123\n\n- Project Groups (type=\"project\"): Created and managed by individual members for their own use\n  * Used for project firewalls with fine-grained access control\n  * Examples: office networks, VPN endpoints, admin workstations\n  * Referenced in project firewall rules using @groupname syntax: \"source\": \"@office_networks\"\n\nUsage in Firewall Rules:\n- Project Firewall: \"source\": \"@office_networks\" (uses project type groups only)\n- Geo Firewall: \"group_name\": \"@ie_v4\" (uses geo type groups only)\n\nExamples:\n- Block traffic from Ireland: Create geo firewall rule with group_name of Ireland group\n- Allow access from office: Create project firewall rule with source \"@office_networks\"\n- Compliance geo-blocking: Use global country groups referenced by ID in geo firewall rules",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Network IP Group.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Network IP Group",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The absolute URL of the Network IP Group record that can be used to perform `Read`, `Update` and `Delete`",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The IP Version of the CIDRs in the group.",
				Computed:    true,
			},
			"cidrs": schema.ListAttribute{
				Description: "An array of CIDR addresses in the Network IP Group.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *NetworkIPGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NetworkIPGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
