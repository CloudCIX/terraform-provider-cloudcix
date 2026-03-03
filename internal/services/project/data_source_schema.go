// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ProjectDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Management of Cloud Projects\n\nThis module provides API endpoints for managing cloud projects in the CloudCIX Compute platform.\nProjects are logical containers that organise and group your cloud infrastructure resources such as\nvirtual machines, routers, firewalls, and storage. Each project belongs to a specific region and\nhas its own isolated network environment.\n\nAvailable operations:\n- List and filter projects across your organization\n- Create new projects in available cloud regions\n- Retrieve detailed project information including region and manager\n- Update project details such as name and notes\n\nEach project includes its associated address, region, manager, and creation metadata.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"address_id": schema.Int64Attribute{
				Description: "The ID of the Project address.",
				Computed:    true,
			},
			"closed": schema.BoolAttribute{
				Description: "A flag stating whether or not the Project is classified as closed. A Project is classified as closed when\nall the infrastructure in it is in a Closed (99) state.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date that the Project entry was created",
				Computed:    true,
			},
			"manager_id": schema.Int64Attribute{
				Description: "The ID of the User that manages the Project",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Project.",
				Computed:    true,
			},
			"note": schema.StringAttribute{
				Description: "The note attached to the Project.",
				Computed:    true,
			},
			"region_id": schema.Int64Attribute{
				Description: "The region ID that the Project is in.",
				Computed:    true,
			},
			"reseller_id": schema.Int64Attribute{
				Description: "The Address ID that will send the bill for the Project to the customer.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "The date that the Project entry was last updated",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The absolute URL of the Project that can be used to perform `Read`, `Update` and `Delete`",
				Computed:    true,
			},
		},
	}
}

func (d *ProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ProjectDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
