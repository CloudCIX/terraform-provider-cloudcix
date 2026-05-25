// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_ssh_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ComputeSSHKeyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Management of SSH Key records.\n\nThis module provides API endpoints for managing SSH Key pairs used when provisioning\ncompute instances. SSH Keys are stored in the Membership service and proxied through here.\n\nAvailable operations:\n- List SSH Keys belonging to the requesting User's Address\n- Create a new SSH Key (optionally auto-generate an Ed25519 key pair)\n- Read a single SSH Key record\n- Delete an SSH Key record",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "Timestamp, in ISO format, of when the SSH Key record was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The user-friendly name for the SSH Key.",
				Computed:    true,
			},
			"private_key": schema.StringAttribute{
				Description: "The PEM-encoded Ed25519 private key. Only present in the create (POST) response\nwhen no public_key was supplied and the key pair was auto-generated. Not stored\nby the API — save it immediately.",
				Computed:    true,
			},
			"public_key": schema.StringAttribute{
				Description: "The SSH public key string.",
				Computed:    true,
			},
		},
	}
}

func (d *ComputeSSHKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ComputeSSHKeyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
