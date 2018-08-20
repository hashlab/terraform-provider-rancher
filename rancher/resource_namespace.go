package rancher

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNamespaceCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespaceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespaceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespaceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceNamespaceCreate,
		Read:   resourceNamespaceRead,
		Update: resourceNamespaceUpdate,
		Delete: resourceNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
