// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package project_test

import (
	"context"
	"testing"

	"github.com/CloudCIX/terraform-provider-cloudcix/internal/services/project"
	"github.com/CloudCIX/terraform-provider-cloudcix/internal/test_helpers"
)

func TestProjectModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*project.ProjectModel)(nil)
	schema := project.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
