// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/compute_instance"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestComputeInstanceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*compute_instance.ComputeInstanceModel)(nil)
	schema := compute_instance.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)

	errs.Ignore(".@ComputeInstanceModel.timeouts.@ObjectValue.create")
	errs.Ignore(".@ComputeInstanceModel.timeouts.@ObjectValue.update")
	errs.Ignore(".@ComputeInstanceModel.timeouts.@ObjectValue.read")
	errs.Ignore(".@ComputeInstanceModel.timeouts.@ObjectValue.delete")

	errs.Report(t)
}
