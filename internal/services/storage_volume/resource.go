// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume

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
var _ resource.ResourceWithConfigure = (*StorageVolumeResource)(nil)
var _ resource.ResourceWithModifyPlan = (*StorageVolumeResource)(nil)
var _ resource.ResourceWithImportState = (*StorageVolumeResource)(nil)

func NewResource() resource.Resource {
	return &StorageVolumeResource{}
}

// StorageVolumeResource defines the resource implementation.
type StorageVolumeResource struct {
	client *gocloudcix.Client
}

func (r *StorageVolumeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_volume"
}

func (r *StorageVolumeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *StorageVolumeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *StorageVolumeModel

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

	tflog.Info(ctx, "Creating storage volume")

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := StorageVolumeContentEnvelope{*data}
	_, err = r.client.Storage.Volumes.New(
		ctx,
		gocloudcix.StorageVolumeNewParams{},
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

	tflog.Info(ctx, "Storage volume created, waiting for running state", map[string]interface{}{
		"volume_id": data.ID.ValueInt64(),
	})

	// Wait for state to become "running"
	err = r.waitForRunningState(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for volume to reach running state", err.Error())
		return
	}

	tflog.Info(ctx, "Storage volume is now running")

	// Refresh data after state is running
	res = new(http.Response)
	env = StorageVolumeContentEnvelope{*data}
	_, err = r.client.Storage.Volumes.Get(
		ctx,
		data.ID.ValueInt64(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to read volume after creation", err.Error())
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

// waitForRunningState polls the volume until its state is "running"
func (r *StorageVolumeResource) waitForRunningState(ctx context.Context, volumeID int64) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for volume to reach running state: %w", ctx.Err())
		case <-ticker.C:
			elapsed := time.Since(startTime)

			res := new(http.Response)
			_, err := r.client.Storage.Volumes.Get(
				ctx,
				volumeID,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				return fmt.Errorf("failed to check volume state: %w", err)
			}

			bytes, _ := io.ReadAll(res.Body)
			var env StorageVolumeContentEnvelope
			err = apijson.Unmarshal(bytes, &env)
			if err != nil {
				return fmt.Errorf("failed to parse volume state: %w", err)
			}

			state := env.Content.State.ValueString()

			// Log current state with elapsed time
			tflog.Debug(ctx, "Checking volume state", map[string]interface{}{
				"volume_id": volumeID,
				"state":     state,
				"elapsed":   elapsed.String(),
			})

			if state == "running" {
				return nil
			}

			// Check for error states
			if state == "error" || state == "failed" {
				return fmt.Errorf("volume entered error state: %s", state)
			}
		}
	}
}

func (r *StorageVolumeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *StorageVolumeModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *StorageVolumeModel

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
	env := StorageVolumeContentEnvelope{*data}
	_, err = r.client.Storage.Volumes.Update(
		ctx,
		data.ID.ValueInt64(),
		gocloudcix.StorageVolumeUpdateParams{},
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

func (r *StorageVolumeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *StorageVolumeModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := StorageVolumeContentEnvelope{*data}
	_, err := r.client.Storage.Volumes.Get(
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

func (r *StorageVolumeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *StorageVolumeModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Storage.Volumes.Delete(
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

func (r *StorageVolumeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(StorageVolumeModel)

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
	env := StorageVolumeContentEnvelope{*data}
	_, err := r.client.Storage.Volumes.Get(
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

func (r *StorageVolumeResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
