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

	ignore_list := []string{
		".@ComputeInstanceModel.timeouts.@ObjectValue.read",
		".@ComputeInstanceModel.timeouts.@ObjectValue.create",
		".@ComputeInstanceModel.timeouts.@ObjectValue.update",
		".@ComputeInstanceModel.timeouts.@ObjectValue.delete",
	}

	for _, item := range ignore_list {
		errs.IgnoreAll(t, item)
	}

	errs.Report(t)
}
