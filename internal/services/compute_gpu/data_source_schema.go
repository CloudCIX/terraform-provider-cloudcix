// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"context"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeGPUDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Management of GPU Resources\n\nThis module provides API endpoints for managing GPU (Graphics Processing Unit) resources\nwithin the CloudCIX Compute platform. GPUs are physical hardware accelerators that can be\nattached to LXD instances to provide enhanced computational capabilities for workloads such\nas machine learning, AI training, scientific simulations, and graphics rendering.\n\nAvailable operations:\n- List and filter GPU resources across your projects\n- Attach GPUs to running LXD instances by creating new GPU resources\n- Retrieve individual GPU configuration and status details\n- Detach GPUs from instances by updating the state to delete\n\nEach GPU resource includes its associated LXD instance, capacity specifications (SKUs),\ncurrent state, and project information.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute GPU record was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name given to this Compute GPU",
				Computed:    true,
			},
			"project_id": schema.Int64Attribute{
				Description: "The id of the Project that this Compute GPU belongs to",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The current state of the Compute GPU",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute GPU record was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Compute GPU instance.",
				Computed:    true,
			},
			"instance": schema.SingleNestedAttribute{
				Description: `The "lxd" Compute Instance the Compute GPU is attached to.`,
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeGPUInstanceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: `The ID of the "lxd" Compute Instance the Compute GPU is attached to.`,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: `The user-friendly name of the "lxd" Compute Instance the Compute GPU is attached to.`,
						Computed:    true,
					},
					"state": schema.StringAttribute{
						Description: `The current state of the "lxd" Compute Instance the Compute GPU is attached to.`,
						Computed:    true,
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Compute GPU",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ComputeGPUSpecsDataSourceModel](ctx),
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

func (d *ComputeGPUDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeGPUDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
