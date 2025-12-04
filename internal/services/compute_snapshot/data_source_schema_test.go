// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_snapshot_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/compute_snapshot"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestComputeSnapshotDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_snapshot.ComputeSnapshotDataSourceModel)(nil)
	schema := compute_snapshot.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
