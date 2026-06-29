package compute_gpu

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestComputeGPUAttachParams(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	params, diags := computeGPUAttachParams(ctx, &ComputeGPUModel{
		ProjectID:  types.Int64Value(123),
		InstanceID: types.Int64Value(456),
		Name:       types.StringValue("gpu01"),
		Specs: customfield.NewObjectListMust(ctx, []ComputeGPUSpecsModel{{
			SKUName: types.StringValue("A100_GPU"),
		}}),
	})
	if diags.HasError() {
		t.Fatalf("computeGPUAttachParams returned diagnostics: %v", diags)
	}

	if params.ProjectID != 123 || params.InstanceID != 456 {
		t.Fatalf("unexpected IDs: project_id=%d instance_id=%d", params.ProjectID, params.InstanceID)
	}
	if !params.Name.Valid() || params.Name.Value != "gpu01" {
		t.Fatalf("unexpected name: %#v", params.Name)
	}
	if len(params.Specs) != 1 || !params.Specs[0].SKUName.Valid() || params.Specs[0].SKUName.Value != "A100_GPU" {
		t.Fatalf("unexpected specs: %#v", params.Specs)
	}
}
