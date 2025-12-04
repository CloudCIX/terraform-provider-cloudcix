// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_backup_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/compute_backup"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestComputeBackupDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_backup.ComputeBackupDataSourceModel)(nil)
	schema := compute_backup.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
