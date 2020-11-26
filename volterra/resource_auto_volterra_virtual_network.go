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
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_virtual_network "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_network"
)

// resourceVolterraVirtualNetwork is implementation of Volterra's VirtualNetwork resources
func resourceVolterraVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraVirtualNetworkCreate,
		Read:   resourceVolterraVirtualNetworkRead,
		Update: resourceVolterraVirtualNetworkUpdate,
		Delete: resourceVolterraVirtualNetworkDelete,

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

			"global_network": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"legacy_type": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"site_local_inside_network": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"site_local_network": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"static_routes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attrs": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ip_prefixes": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"default_gateway": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"kind": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"ip_address": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraVirtualNetworkCreate creates VirtualNetwork resource
func resourceVolterraVirtualNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_virtual_network.CreateSpecType{}
	createReq := &ves_io_schema_virtual_network.CreateRequest{
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

	networkChoiceTypeFound := false

	if v, ok := d.GetOk("global_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.CreateSpecType_GlobalNetwork{}
			networkChoiceInt.GlobalNetwork = &ves_io_schema.Empty{}
			createSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("legacy_type"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true
		networkChoiceInt := &ves_io_schema_virtual_network.CreateSpecType_LegacyType{}

		createSpec.NetworkChoice = networkChoiceInt

		networkChoiceInt.LegacyType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

	}

	if v, ok := d.GetOk("site_local_inside_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.CreateSpecType_SiteLocalInsideNetwork{}
			networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
			createSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("site_local_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.CreateSpecType_SiteLocalNetwork{}
			networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
			createSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("static_routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		staticRoutes := make([]*ves_io_schema_virtual_network.StaticRouteViewType, len(sl))
		createSpec.StaticRoutes = staticRoutes
		for i, set := range sl {
			staticRoutes[i] = &ves_io_schema_virtual_network.StaticRouteViewType{}

			staticRoutesMapStrToI := set.(map[string]interface{})

			if v, ok := staticRoutesMapStrToI["attrs"]; ok && !isIntfNil(v) {

				attrsList := []ves_io_schema.RouteAttrType{}
				for _, j := range v.([]interface{}) {
					attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
				}
				staticRoutes[i].Attrs = attrsList

			}

			if w, ok := staticRoutesMapStrToI["ip_prefixes"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				staticRoutes[i].IpPrefixes = ls
			}

			nextHopChoiceTypeFound := false

			if v, ok := staticRoutesMapStrToI["default_gateway"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true

				if v.(bool) {
					nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_DefaultGateway{}
					nextHopChoiceInt.DefaultGateway = &ves_io_schema.Empty{}
					staticRoutes[i].NextHopChoice = nextHopChoiceInt
				}

			}

			if v, ok := staticRoutesMapStrToI["interface"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true
				nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_Interface{}
				nextHopChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
				staticRoutes[i].NextHopChoice = nextHopChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Tenant = v.(string)
					}

				}

			}

			if v, ok := staticRoutesMapStrToI["ip_address"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true
				nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_IpAddress{}

				staticRoutes[i].NextHopChoice = nextHopChoiceInt

				nextHopChoiceInt.IpAddress = v.(string)

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra VirtualNetwork object with struct: %+v", createReq)

	createVirtualNetworkResp, err := client.CreateObject(context.Background(), ves_io_schema_virtual_network.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating VirtualNetwork: %s", err)
	}
	d.SetId(createVirtualNetworkResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraVirtualNetworkRead(d, meta)
}

func resourceVolterraVirtualNetworkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_virtual_network.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualNetwork %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualNetwork %q: %s", d.Id(), err)
	}
	return setVirtualNetworkFields(client, d, resp)
}

func setVirtualNetworkFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraVirtualNetworkUpdate updates VirtualNetwork resource
func resourceVolterraVirtualNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_virtual_network.ReplaceSpecType{}
	updateReq := &ves_io_schema_virtual_network.ReplaceRequest{
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

	networkChoiceTypeFound := false

	if v, ok := d.GetOk("global_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.ReplaceSpecType_GlobalNetwork{}
			networkChoiceInt.GlobalNetwork = &ves_io_schema.Empty{}
			updateSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("legacy_type"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true
		networkChoiceInt := &ves_io_schema_virtual_network.ReplaceSpecType_LegacyType{}

		updateSpec.NetworkChoice = networkChoiceInt

		networkChoiceInt.LegacyType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

	}

	if v, ok := d.GetOk("site_local_inside_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.ReplaceSpecType_SiteLocalInsideNetwork{}
			networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
			updateSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("site_local_network"); ok && !networkChoiceTypeFound {

		networkChoiceTypeFound = true

		if v.(bool) {
			networkChoiceInt := &ves_io_schema_virtual_network.ReplaceSpecType_SiteLocalNetwork{}
			networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
			updateSpec.NetworkChoice = networkChoiceInt
		}

	}

	if v, ok := d.GetOk("static_routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		staticRoutes := make([]*ves_io_schema_virtual_network.StaticRouteViewType, len(sl))
		updateSpec.StaticRoutes = staticRoutes
		for i, set := range sl {
			staticRoutes[i] = &ves_io_schema_virtual_network.StaticRouteViewType{}

			staticRoutesMapStrToI := set.(map[string]interface{})

			if v, ok := staticRoutesMapStrToI["attrs"]; ok && !isIntfNil(v) {

				attrsList := []ves_io_schema.RouteAttrType{}
				for _, j := range v.([]interface{}) {
					attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
				}
				staticRoutes[i].Attrs = attrsList

			}

			if w, ok := staticRoutesMapStrToI["ip_prefixes"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				staticRoutes[i].IpPrefixes = ls
			}

			nextHopChoiceTypeFound := false

			if v, ok := staticRoutesMapStrToI["default_gateway"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true

				if v.(bool) {
					nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_DefaultGateway{}
					nextHopChoiceInt.DefaultGateway = &ves_io_schema.Empty{}
					staticRoutes[i].NextHopChoice = nextHopChoiceInt
				}

			}

			if v, ok := staticRoutesMapStrToI["interface"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true
				nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_Interface{}
				nextHopChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
				staticRoutes[i].NextHopChoice = nextHopChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						nextHopChoiceInt.Interface.Tenant = v.(string)
					}

				}

			}

			if v, ok := staticRoutesMapStrToI["ip_address"]; ok && !isIntfNil(v) && !nextHopChoiceTypeFound {

				nextHopChoiceTypeFound = true
				nextHopChoiceInt := &ves_io_schema_virtual_network.StaticRouteViewType_IpAddress{}

				staticRoutes[i].NextHopChoice = nextHopChoiceInt

				nextHopChoiceInt.IpAddress = v.(string)

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra VirtualNetwork obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_virtual_network.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating VirtualNetwork: %s", err)
	}

	return resourceVolterraVirtualNetworkRead(d, meta)
}

func resourceVolterraVirtualNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_virtual_network.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualNetwork %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualNetwork before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra VirtualNetwork obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_virtual_network.ObjectType, namespace, name)
}
