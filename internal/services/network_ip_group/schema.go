// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_ip_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*NetworkIPGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Management of Network IP Groups\n\nIP address groups organise sets of CIDR networks for use in firewall rules and access control.\nTwo types are available:\n\n- Geo Groups (type=\"geo\"): Maintained by CloudCIX and accessible to all members\n  * Used for geo-filtering based on country IP ranges (e.g., 'Ireland', 'USA', 'China')\n  * Essential for creating geo firewalls that block/allow traffic from specific countries\n  * To list country groups: GET /ip_address_groups?search[member_id]=0\n  * Referenced in geo firewall rules by numeric ID: \"ip_address_group_id\": 123\n\n- Project Groups (type=\"project\"): Created and managed by individual members for their own use\n  * Used for project firewalls with fine-grained access control\n  * Examples: office networks, VPN endpoints, admin workstations\n  * Referenced in project firewall rules using @groupname syntax: \"source\": \"@office_networks\"\n\nUsage in Firewall Rules:\n- Project Firewall: \"source\": \"@office_networks\" (uses project type groups only)\n- Geo Firewall: \"group_name\": \"@ie_v4\" (uses geo type groups only)\n\nExamples:\n- Block traffic from Ireland: Create geo firewall rule with group_name of Ireland group\n- Allow access from office: Create project firewall rule with source \"@office_networks\"\n- Compliance geo-blocking: Use global country groups referenced by ID in geo firewall rules",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:   "The ID of the Network IP Goup record",
				Computed:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description: "The name to be given to the new IP Address Group. Used to identify the group when creating\nfirewall rules or geo-filters. Must start with a letter and contain only letters, numbers,\nunderscores, and hyphens.",
				Required:    true,
			},
			"cidrs": schema.ListAttribute{
				Description: "An array of CIDR addresses in the IP Address Group. Can include individual IP addresses\n(e.g., \"91.103.3.36\") or network ranges (e.g., \"90.103.2.0/24\"). All addresses must match\nthe specified IP version. Use these groups in firewall rules to allow/block traffic from\nmultiple locations with a single rule.",
				Required:    true,
				ElementType: types.StringType,
			},
			"version": schema.Int64Attribute{
				Description: "The IP version of the IP Address Group Objects in the IP Address Group. Accepted versions are 4 and 6.\nIf not sent, it will default to 4.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Network IP Group was created.",
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
		},
	}
}

func (r *NetworkIPGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *NetworkIPGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
