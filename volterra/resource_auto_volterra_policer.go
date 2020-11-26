//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_policer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policer"
)

// resourceVolterraPolicer is implementation of Volterra's Policer resources
func resourceVolterraPolicer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraPolicerCreate,
		Read:   resourceVolterraPolicerRead,
		Update: resourceVolterraPolicerUpdate,
		Delete: resourceVolterraPolicerDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"burst_size": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"committed_information_rate": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"policer_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"policer_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraPolicerCreate creates Policer resource
func resourceVolterraPolicerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_policer.CreateSpecType{}
	createReq := &ves_io_schema_policer.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("burst_size"); ok && !isIntfNil(v) {

		createSpec.BurstSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("committed_information_rate"); ok && !isIntfNil(v) {

		createSpec.CommittedInformationRate =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("policer_mode"); ok && !isIntfNil(v) {

		createSpec.PolicerMode = ves_io_schema_policer.PolicerMode(ves_io_schema_policer.PolicerMode_value[v.(string)])

	}

	if v, ok := d.GetOk("policer_type"); ok && !isIntfNil(v) {

		createSpec.PolicerType = ves_io_schema_policer.PolicerType(ves_io_schema_policer.PolicerType_value[v.(string)])

	}

	log.Printf("[DEBUG] Creating Volterra Policer object with struct: %+v", createReq)

	createPolicerResp, err := client.CreateObject(context.Background(), ves_io_schema_policer.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Policer: %s", err)
	}
	d.SetId(createPolicerResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraPolicerRead(d, meta)
}

func resourceVolterraPolicerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_policer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Policer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Policer %q: %s", d.Id(), err)
	}
	return setPolicerFields(client, d, resp)
}

func setPolicerFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraPolicerUpdate updates Policer resource
func resourceVolterraPolicerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_policer.ReplaceSpecType{}
	updateReq := &ves_io_schema_policer.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("burst_size"); ok && !isIntfNil(v) {

		updateSpec.BurstSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("committed_information_rate"); ok && !isIntfNil(v) {

		updateSpec.CommittedInformationRate =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("policer_mode"); ok && !isIntfNil(v) {

		updateSpec.PolicerMode = ves_io_schema_policer.PolicerMode(ves_io_schema_policer.PolicerMode_value[v.(string)])

	}

	if v, ok := d.GetOk("policer_type"); ok && !isIntfNil(v) {

		updateSpec.PolicerType = ves_io_schema_policer.PolicerType(ves_io_schema_policer.PolicerType_value[v.(string)])

	}

	log.Printf("[DEBUG] Updating Volterra Policer obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_policer.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Policer: %s", err)
	}

	return resourceVolterraPolicerRead(d, meta)
}

func resourceVolterraPolicerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_policer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Policer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Policer before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Policer obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_policer.ObjectType, namespace, name)
}
