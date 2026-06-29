// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/option"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/apijson"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/importpath"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ComputeGPUResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ComputeGPUResource)(nil)
var _ resource.ResourceWithImportState = (*ComputeGPUResource)(nil)

func NewResource() resource.Resource {
	return &ComputeGPUResource{}
}

// ComputeGPUResource defines the resource implementation.
type ComputeGPUResource struct {
	client *gocloudcix.Client
}

func (r *ComputeGPUResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compute_gpu"
}

func (r *ComputeGPUResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*gocloudcix.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *gocloudcix.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ComputeGPUResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ComputeGPUModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := computeGPUAttachParams(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := ComputeGPUContentEnvelope{*data}
	gpu, err := r.client.Compute.GPUs.Attach(
		ctx,
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if gpu == nil {
		resp.Diagnostics.AddError("failed to make http request", "received empty response")
		return
	}
	bytes, err := json.Marshal(gocloudcix.ComputeGPUResponse{Content: *gpu})
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http response", err.Error())
		return
	}
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func computeGPUAttachParams(ctx context.Context, data *ComputeGPUModel) (gocloudcix.ComputeGPUAttachParams, diag.Diagnostics) {
	specs, diags := data.Specs.AsStructSliceT(ctx)
	if diags.HasError() {
		return gocloudcix.ComputeGPUAttachParams{}, diags
	}

	params := gocloudcix.ComputeGPUAttachParams{
		InstanceID: data.InstanceID.ValueInt64(),
		ProjectID:  data.ProjectID.ValueInt64(),
		Specs:      make([]gocloudcix.ComputeGPUAttachParamsSpec, len(specs)),
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		params.Name = gocloudcix.String(data.Name.ValueString())
	}
	for i, spec := range specs {
		if !spec.SKUName.IsNull() && !spec.SKUName.IsUnknown() {
			params.Specs[i].SKUName = gocloudcix.String(spec.SKUName.ValueString())
		}
	}
	return params, diags
}

func (r *ComputeGPUResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ComputeGPUModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ComputeGPUModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ComputeGPUContentEnvelope{*data}
	_, err = r.client.Compute.GPUs.Update(
		ctx,
		data.ID.ValueInt64(),
		gocloudcix.ComputeGPUUpdateParams{},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeGPUResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ComputeGPUModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ComputeGPUContentEnvelope{*data}
	_, err := r.client.Compute.GPUs.Get(
		ctx,
		data.ID.ValueInt64(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeGPUResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ComputeGPUModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Compute.GPUs.Delete(
		ctx,
		data.ID.ValueInt64(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeGPUResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ComputeGPUModel)

	path := int64(0)
	diags := importpath.ParseImportID(
		req.ID,
		"<id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.Int64Value(path)

	res := new(http.Response)
	env := ComputeGPUContentEnvelope{*data}
	_, err := r.client.Compute.GPUs.Get(
		ctx,
		path,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ComputeGPUResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
