package projects

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
)

func ProjectResource() *schema.Resource {
	return &schema.Resource{
		Create:        Create,
		Read:          Read,
		Delete:        Delete,
		Update:        Update,
		Exists:        Exists,
		SchemaVersion: 1,
		Schema:        map[string]*schema.Schema{},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	fmt.Println(client)
	return nil
}

func Read(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	fmt.Println(client)
	return nil
}

func Delete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	fmt.Println(client)
	return nil
}

func Update(d *schema.ResourceData, meta interface{}) error {
	client := meta.(pt.ClientCaller)
	fmt.Println(client)
	return nil
}

func Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(pt.ClientCaller)
	fmt.Println(client)
	return false, nil
}
