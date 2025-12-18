// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_firewall

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
var _ resource.ResourceWithConfigure = (*NetworkFirewallResource)(nil)
var _ resource.ResourceWithModifyPlan = (*NetworkFirewallResource)(nil)
var _ resource.ResourceWithImportState = (*NetworkFirewallResource)(nil)

const pollInterval = 15 * time.Second

func NewResource() resource.Resource {
	return &NetworkFirewallResource{}
}

// NetworkFirewallResource defines the resource implementation.
type NetworkFirewallResource struct {
	client *gocloudcix.Client
}

func (r *NetworkFirewallResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_firewall"
}

func (r *NetworkFirewallResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NetworkFirewallResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *NetworkFirewallModel

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

	tflog.Info(ctx, "Creating network firewall")

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := NetworkFirewallContentEnvelope{*data}
	_, err = r.client.Network.Firewalls.New(
		ctx,
		gocloudcix.NetworkFirewallNewParams{},
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

	tflog.Info(ctx, "Network firewall created, waiting for running state", map[string]interface{}{
		"firewall_id": data.ID.ValueInt64(),
	})

	// Wait for state to become "running"
	err = r.waitForRunningState(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for firewall to reach running state", err.Error())
		return
	}

	tflog.Info(ctx, "Network firewall is now running")

	// Refresh data after state is running
	res = new(http.Response)
	env = NetworkFirewallContentEnvelope{*data}
	_, err = r.client.Network.Firewalls.Get(
		ctx,
		data.ID.ValueInt64(),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to read firewall after creation", err.Error())
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

// waitForRunningState polls the firewall until its state is "running"
func (r *NetworkFirewallResource) waitForRunningState(ctx context.Context, firewallID int64) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for firewall to reach running state: %w", ctx.Err())
		case <-ticker.C:
			elapsed := time.Since(startTime)

			res := new(http.Response)
			_, err := r.client.Network.Firewalls.Get(
				ctx,
				firewallID,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)
			if err != nil {
				return fmt.Errorf("failed to check firewall state: %w", err)
			}

			bytes, _ := io.ReadAll(res.Body)
			var env NetworkFirewallContentEnvelope
			err = apijson.Unmarshal(bytes, &env)
			if err != nil {
				return fmt.Errorf("failed to parse firewall state: %w", err)
			}

			state := env.Content.State.ValueString()

			// Log current state with elapsed time
			tflog.Debug(ctx, "Checking firewall state", map[string]interface{}{
				"firewall_id": firewallID,
				"state":       state,
				"elapsed":     elapsed.String(),
			})

			if state == "running" {
				return nil
			}

			// Check for error states
			if state == "error" || state == "failed" {
				return fmt.Errorf("firewall entered error state: %s", state)
			}
		}
	}
}

// waitForDeletion polls until the firewall is deleted or reaches a terminal state
func (r *NetworkFirewallResource) waitForDeletion(ctx context.Context, firewallID int64) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	startTime := time.Now()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for firewall to be deleted: %w", ctx.Err())
		case <-ticker.C:
			elapsed := time.Since(startTime)

			res := new(http.Response)
			_, err := r.client.Network.Firewalls.Get(
				ctx,
				firewallID,
				option.WithResponseBodyInto(&res),
				option.WithMiddleware(logging.Middleware(ctx)),
			)

			if err != nil {
				return fmt.Errorf("failed to check firewall state: %w", err)
			}

			bytes, _ := io.ReadAll(res.Body)
			var env NetworkFirewallContentEnvelope
			err = apijson.Unmarshal(bytes, &env)
			
			if err != nil {
				return fmt.Errorf("failed to parse firewall state: %w", err)
			}

			state := env.Content.State.ValueString()
			
			// Log current state with elapsed time
			tflog.Debug(ctx, "Checking firewall state", map[string]interface{}{
				"firewall_id": firewallID,
				"state":       state,
				"elapsed":     elapsed.String(),
			})

			if state == "deleted" {
				return nil
			}

			// Check for error states
			if state == "unresourced" {
				return fmt.Errorf("firewall entered error state: %s", state)
			}
		}
	}
}

func (r *NetworkFirewallResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *NetworkFirewallModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *NetworkFirewallModel

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
	env := NetworkFirewallContentEnvelope{*data}
	_, err = r.client.Network.Firewalls.Update(
		ctx,
		data.ID.ValueInt64(),
		gocloudcix.NetworkFirewallUpdateParams{},
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

func (r *NetworkFirewallResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *NetworkFirewallModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := NetworkFirewallContentEnvelope{*data}
	_, err := r.client.Network.Firewalls.Get(
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

func (r *NetworkFirewallResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *NetworkFirewallModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get timeout from configuration
	deleteTimeout, diags := data.Timeouts.Delete(ctx, 30*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()

	tflog.Info(ctx, "Deleting network firewall", map[string]interface{}{
		"router_id": data.ID.ValueInt64(),
	})

	err := r.client.Network.Firewalls.Delete(
		ctx,
		data.ID.ValueInt64(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	// Wait for resource to be deleted
	err = r.waitForDeletion(ctx, data.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError("failed waiting for firewall to be deleted", err.Error())
		return
	}

	tflog.Info(ctx, "Network firewall deleted successfully")
}

func (r *NetworkFirewallResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(NetworkFirewallModel)

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
	env := NetworkFirewallContentEnvelope{*data}
	_, err := r.client.Network.Firewalls.Get(
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

func (r *NetworkFirewallResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
