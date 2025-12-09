// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage_volume_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/storage_volume"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestStorageVolumeModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*storage_volume.StorageVolumeModel)(nil)
	schema := storage_volume.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)

	ignore_list := []string{
		".@StorageVolumeModel.timeouts.@ObjectValue.read",
		".@StorageVolumeModel.timeouts.@ObjectValue.create",
		".@StorageVolumeModel.timeouts.@ObjectValue.update",
		".@StorageVolumeModel.timeouts.@ObjectValue.delete",
	}

	for _, item := range ignore_list {
		errs.IgnoreAll(t, item)
	}

	errs.Report(t)
}
