// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeImageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Management of Operating System Images\n\nThis module provides API endpoints for browsing available operating system images that can be\nused when creating virtual instances in the CloudCIX Compute platform. Images represent\npre-configured OS templates including various Linux distributions and Windows versions.\n\nAvailable operations:\n- List and filter available OS images by region, name, or variant\n- Retrieve detailed information about a specific image including its SKU and OS variant",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"filename": schema.StringAttribute{
				Description: "The name of the file containing the Image.",
				Computed:    true,
			},
			"os_variant": schema.StringAttribute{
				Description: "Is a unique word to define each Image.",
				Computed:    true,
			},
			"sku_name": schema.StringAttribute{
				Description: "The name of the Image.",
				Computed:    true,
			},
		},
	}
}

func (d *ComputeImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeImageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
