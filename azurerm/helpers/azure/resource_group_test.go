package azure_test

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestValidateResourceGroupName(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		{
			Value:    "",
			ErrCount: 1,
		},
		{
			Value:    "hello",
			ErrCount: 0,
		},
		{
			Value:    "Hello",
			ErrCount: 0,
		},
		{
			Value:    "hello-world",
			ErrCount: 0,
		},
		{
			Value:    "Hello_World",
			ErrCount: 0,
		},
		{
			Value:    "HelloWithNumbers12345",
			ErrCount: 0,
		},
		{
			Value:    "(Did)You(Know)That(Brackets)Are(Allowed)In(Resource)Group(Names)",
			ErrCount: 0,
		},
		{
			Value:    "EndingWithAPeriod.",
			ErrCount: 1,
		},
		{
			Value:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo",
			ErrCount: 1,
		},
		{
			Value:    acceptance.RandString(90),
			ErrCount: 0,
		},
		{
			Value:    acceptance.RandString(91),
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		_, errors := azure.ValidateResourceGroupName(tc.Value, "azurerm_resource_group")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected "+
				"validateResourceGroupName to trigger '%d' errors for '%s' - got '%d'", tc.ErrCount, tc.Value, len(errors))
		}
	}
}
