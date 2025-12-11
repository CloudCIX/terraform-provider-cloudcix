// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_instance

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/option"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/apijson"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/importpath"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ComputeInstanceResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ComputeInstanceResource)(nil)
var _ resource.ResourceWithImportState = (*ComputeInstanceResource)(nil)

func NewResource() resource.Resource {
	return &ComputeInstanceResource{}
}

// ComputeInstanceResource defines the resource implementation.
type ComputeInstanceResource struct {
	client *gocloudcix.Client
}

func (r *ComputeInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compute_instance"
}

func (r *ComputeInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ComputeInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ComputeInstanceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get timeout from configuration
	createTimeout, diags := data.Timeouts.Create(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	tflog.Info(ctx, "Creating compute instance")

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ComputeInstanceContentEnvelope{*data}
	_, err = r.client.Compute.Instances.New(
		ctx,
		gocloudcix.ComputeInstanceNewParams{},
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

	tflog.Info(ctx, "Compute instance created, waiting for running state", map[string]interface{}{
		"instance_id": data.ID.ValueInt64(),
	})

	// Wait for state to become "running"
	err = r.waitForRunningState(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for instance to reach running state", err.Error())
		return
	}

	tflog.Info(ctx, "Compute instance is now running")

	// Refresh data after state is running
	res = new(http.Response)
	env = ComputeInstanceContentEnvelope{*data}
	_, err = r.client.Compute.Instances.Get(
		ctx,
		data.ID.ValueInt64(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to read instance after creation", err.Error())
		return
	}
	bytes, _ = io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Content

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// waitForRunningState polls the instance until its state is "running"
func (r *ComputeInstanceResource) waitForRunningState(ctx context.Context, instanceID int64) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for instance to reach running state: %w", ctx.Err())
		case <-ticker.C:
			elapsed := time.Since(startTime)

			res := new(http.Response)
			_, err := r.client.Compute.Instances.Get(
				ctx,
				instanceID,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				return fmt.Errorf("failed to check instance state: %w", err)
			}

			bytes, _ := io.ReadAll(res.Body)
			var env ComputeInstanceContentEnvelope
			err = apijson.Unmarshal(bytes, &env)
			if err != nil {
				return fmt.Errorf("failed to parse instance state: %w", err)
			}

			state := env.Content.State.ValueString()

			// Log current state with elapsed time
			tflog.Debug(ctx, "Checking instance state", map[string]interface{}{
				"instance_id": instanceID,
				"state":       state,
				"elapsed":     elapsed.String(),
			})

			if state == "running" {
				return nil
			}

			// Check for error states
			if state == "error" || state == "failed" {
				return fmt.Errorf("instance entered error state: %s", state)
			}
		}
	}
}

// waitForDeletion polls until the instance is deleted (404) or reaches a terminal state
func (r *ComputeInstanceResource) waitForDeletion(ctx context.Context, instanceID int64) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for instance to be deleted: %w", ctx.Err())
		case <-ticker.C:
			elapsed := time.Since(startTime)

			res := new(http.Response)
			_, err := r.client.Compute.Instances.Get(
				ctx,
				instanceID,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)

			// Resource deleted successfully (404)
			if res != nil && res.StatusCode == 404 {
				return nil
			}

			if err != nil {
				return fmt.Errorf("failed to check instance deletion status: %w", err)
			}

			bytes, _ := io.ReadAll(res.Body)
			var env ComputeInstanceContentEnvelope
			err = apijson.Unmarshal(bytes, &env)
			if err != nil {
				return fmt.Errorf("failed to parse instance state: %w", err)
			}

			state := env.Content.State.ValueString()

			// Log current state with elapsed time
			tflog.Debug(ctx, "Checking instance deletion state", map[string]interface{}{
				"instance_id": instanceID,
				"state":       state,
				"elapsed":     elapsed.String(),
			})

			// Check for error states
			if state == "error" || state == "failed" {
				return fmt.Errorf("instance entered error state during deletion: %s", state)
			}
		}
	}
}

func (r *ComputeInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ComputeInstanceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ComputeInstanceModel

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
	env := ComputeInstanceContentEnvelope{*data}
	_, err = r.client.Compute.Instances.Update(
		ctx,
		data.ID.ValueInt64(),
		gocloudcix.ComputeInstanceUpdateParams{},
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

func (r *ComputeInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ComputeInstanceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := ComputeInstanceContentEnvelope{*data}
	_, err := r.client.Compute.Instances.Get(
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

func (r *ComputeInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ComputeInstanceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Compute.Instances.Delete(
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

func (r *ComputeInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ComputeInstanceModel)

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
	env := ComputeInstanceContentEnvelope{*data}
	_, err := r.client.Compute.Instances.Get(
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

func (r *ComputeInstanceResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
