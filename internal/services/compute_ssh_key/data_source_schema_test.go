// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_ssh_key_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/compute_ssh_key"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestComputeSSHKeyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_ssh_key.ComputeSSHKeyDataSourceModel)(nil)
	schema := compute_ssh_key.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
