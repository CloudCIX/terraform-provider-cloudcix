// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compute_gpu

import (
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/apijson"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComputeGPUContentEnvelope struct {
	Content ComputeGPUModel `json:"content"`
}

type ComputeGPUModel struct {
	ID         types.Int64                                        `tfsdk:"id" json:"id,computed"`
	ProjectID  types.Int64                                        `tfsdk:"project_id" json:"project_id,required"`
	InstanceID types.Int64                                        `tfsdk:"instance_id" json:"instance_id,required,no_refresh"`
	Name       types.String                                       `tfsdk:"name" json:"name,optional"`
	State      types.String                                       `tfsdk:"state" json:"state,optional"`
	Specs      customfield.NestedObjectList[ComputeGPUSpecsModel] `tfsdk:"specs" json:"specs,required"`
	Created    types.String                                       `tfsdk:"created" json:"created,computed"`
	Updated    types.String                                       `tfsdk:"updated" json:"updated,computed"`
	Uri        types.String                                       `tfsdk:"uri" json:"uri,computed"`
	Instance   customfield.NestedObject[ComputeGPUInstanceModel]  `tfsdk:"instance" json:"instance,computed"`
}

func (m ComputeGPUModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ComputeGPUModel) MarshalJSONForUpdate(state ComputeGPUModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(
		ComputeGPUUpdateModel{Name: m.Name, State: m.State},
		ComputeGPUUpdateModel{Name: state.Name, State: state.State},
	)
}

type ComputeGPUUpdateModel struct {
	Name  types.String `tfsdk:"name" json:"name,optional"`
	State types.String `tfsdk:"state" json:"state,optional"`
}

type ComputeGPUInstanceModel struct {
	ID    types.Int64  `tfsdk:"id" json:"id,computed"`
	Name  types.String `tfsdk:"name" json:"name,computed"`
	State types.String `tfsdk:"state" json:"state,computed"`
}

type ComputeGPUSpecsModel struct {
	Quantity types.Int64  `tfsdk:"quantity" json:"quantity,computed"`
	SKUName  types.String `tfsdk:"sku_name" json:"sku_name,optional"`
}
