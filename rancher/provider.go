package rancher

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider The provider interface
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"rancher_namespace": resourceNamespace(),
		},
	}
}
