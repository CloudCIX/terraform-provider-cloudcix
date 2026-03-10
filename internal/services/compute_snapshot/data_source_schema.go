// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot

import (
	"context"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeSnapshotDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Management of Instance Snapshots\n\nThis module provides API endpoints for managing snapshots of virtual machine and container instances\nin the CloudCIX Compute platform. Snapshots are point-in-time backups of running instances that can\nbe used for data recovery from a known state. Two snapshot types are\nsupported: LXD snapshots (for Linux containers and VMs) and Hyper-V snapshots (for Windows VMs).\n\nAvailable operations:\n- List and filter snapshots across your projects by type, instance, or state\n- Create new snapshots from running LXD or Hyper-V instances\n- Retrieve detailed information about individual snapshots\n- Delete snapshots by updating their state\n\nEach snapshot includes its associated instance, project, creation timestamp, and current state.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute Snapshots record was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name given to this Compute Snapshots instance",
				Computed:    true,
			},
			"project_id": schema.Int64Attribute{
				Description: "The id of the Project that this Compute Snapshots belongs to",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The current state of the Compute Snapshots",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the Compute Snapshots",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the Compute Snapshots record was last updated.",
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "URL that can be used to run methods in the API associated with the Compute Snapshots instance.",
				Computed:    true,
			},
			"instance": schema.SingleNestedAttribute{
				Description: "The Compute Instance the Compute Snapshot record is of.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeSnapshotInstanceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "The ID of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "The user-friendly name of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
					"state": schema.StringAttribute{
						Description: "The current state of the Compute Instance the Compute Snapshot is of.",
						Computed:    true,
					},
				},
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "The metadata details of the The metadata details of the \"hyperv\" Compute Snapshot. Returned if the type\nis \"hyperv\".",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ComputeSnapshotMetadataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Description: `Indicates if the "hyperv" Compute Snapshot is currently active.`,
						Computed:    true,
					},
					"remove_subtree": schema.BoolAttribute{
						Description: `Indicates if the "hyperv" Compute Snapshot should remove the subtree when deleted.`,
						Computed:    true,
					},
				},
			},
			"specs": schema.ListNestedAttribute{
				Description: "An array of the specs for the Compute Snapshots",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ComputeSnapshotSpecsDataSourceModel](ctx),
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

func (d *ComputeSnapshotDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeSnapshotDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
