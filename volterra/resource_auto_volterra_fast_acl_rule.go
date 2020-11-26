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
	ves_io_schema_fast_acl_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl_rule"
)

// resourceVolterraFastAclRule is implementation of Volterra's FastAclRule resources
func resourceVolterraFastAclRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraFastAclRuleCreate,
		Read:   resourceVolterraFastAclRuleRead,
		Update: resourceVolterraFastAclRuleUpdate,
		Delete: resourceVolterraFastAclRuleDelete,

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

			"action": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"policer_action": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

										Type:     schema.TypeList,
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
								},
							},
						},

						"protocol_policer_action": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

										Type:     schema.TypeList,
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
								},
							},
						},

						"simple_action": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"port": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"dns": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"user_defined": {

							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"ip_prefix_set": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ref": {

							Type:     schema.TypeList,
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
					},
				},
			},

			"prefix": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"prefix": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraFastAclRuleCreate creates FastAclRule resource
func resourceVolterraFastAclRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_fast_acl_rule.CreateSpecType{}
	createReq := &ves_io_schema_fast_acl_rule.CreateRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
		createSpec.Action = action
		for _, set := range sl {

			actionMapStrToI := set.(map[string]interface{})

			actionTypeFound := false

			if v, ok := actionMapStrToI["policer_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_PolicerAction{}
				actionInt.PolicerAction = &ves_io_schema.PolicerRefType{}
				action.Action = actionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionInt.PolicerAction.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "policer"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := actionMapStrToI["protocol_policer_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_ProtocolPolicerAction{}
				actionInt.ProtocolPolicerAction = &ves_io_schema.ProtocolPolicerRefType{}
				action.Action = actionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionInt.ProtocolPolicerAction.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "protocol_policer"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := actionMapStrToI["simple_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_SimpleAction{}

				action.Action = actionInt

				actionInt.SimpleAction = ves_io_schema_fast_acl_rule.FastAclRuleSimpleAction(ves_io_schema_fast_acl_rule.FastAclRuleSimpleAction_value[v.(string)])

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		port := make([]*ves_io_schema.PortValueType, len(sl))
		createSpec.Port = port
		for i, set := range sl {
			port[i] = &ves_io_schema.PortValueType{}

			portMapStrToI := set.(map[string]interface{})

			portValueTypeChoiceTypeFound := false

			if v, ok := portMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true

				if v.(bool) {
					portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
					portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
					port[i].PortValueTypeChoice = portValueTypeChoiceInt
				}

			}

			if v, ok := portMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true

				if v.(bool) {
					portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
					portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
					port[i].PortValueTypeChoice = portValueTypeChoiceInt
				}

			}

			if v, ok := portMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true
				portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

				port[i].PortValueTypeChoice = portValueTypeChoiceInt

				portValueTypeChoiceInt.UserDefined =
					uint32(v.(int))

			}

		}

	}

	sourceTypeFound := false

	if v, ok := d.GetOk("ip_prefix_set"); ok && !sourceTypeFound {

		sourceTypeFound = true
		sourceInt := &ves_io_schema_fast_acl_rule.CreateSpecType_IpPrefixSet{}
		sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
		createSpec.Source = sourceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["ref"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				sourceInt.IpPrefixSet.Ref = refInt
				for i, ps := range sl {

					rMapToStrVal := ps.(map[string]interface{})
					refInt[i] = &ves_io_schema.ObjectRefType{}

					refInt[i].Kind = "ip_prefix_set"

					if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
						refInt[i].Name = v.(string)
					}

					if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						refInt[i].Namespace = v.(string)
					}

					if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						refInt[i].Tenant = v.(string)
					}

					if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
						refInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("prefix"); ok && !sourceTypeFound {

		sourceTypeFound = true
		sourceInt := &ves_io_schema_fast_acl_rule.CreateSpecType_Prefix{}
		sourceInt.Prefix = &ves_io_schema.PrefixListType{}
		createSpec.Source = sourceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				sourceInt.Prefix.Prefix = ls

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra FastAclRule object with struct: %+v", createReq)

	createFastAclRuleResp, err := client.CreateObject(context.Background(), ves_io_schema_fast_acl_rule.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating FastAclRule: %s", err)
	}
	d.SetId(createFastAclRuleResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraFastAclRuleRead(d, meta)
}

func resourceVolterraFastAclRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_fast_acl_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAclRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAclRule %q: %s", d.Id(), err)
	}
	return setFastAclRuleFields(client, d, resp)
}

func setFastAclRuleFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraFastAclRuleUpdate updates FastAclRule resource
func resourceVolterraFastAclRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_fast_acl_rule.ReplaceSpecType{}
	updateReq := &ves_io_schema_fast_acl_rule.ReplaceRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
		updateSpec.Action = action
		for _, set := range sl {

			actionMapStrToI := set.(map[string]interface{})

			actionTypeFound := false

			if v, ok := actionMapStrToI["policer_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_PolicerAction{}
				actionInt.PolicerAction = &ves_io_schema.PolicerRefType{}
				action.Action = actionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionInt.PolicerAction.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "policer"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := actionMapStrToI["protocol_policer_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_ProtocolPolicerAction{}
				actionInt.ProtocolPolicerAction = &ves_io_schema.ProtocolPolicerRefType{}
				action.Action = actionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionInt.ProtocolPolicerAction.Ref = refInt
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refInt[i] = &ves_io_schema.ObjectRefType{}

							refInt[i].Kind = "protocol_policer"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refInt[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refInt[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refInt[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := actionMapStrToI["simple_action"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true
				actionInt := &ves_io_schema_fast_acl_rule.FastAclRuleAction_SimpleAction{}

				action.Action = actionInt

				actionInt.SimpleAction = ves_io_schema_fast_acl_rule.FastAclRuleSimpleAction(ves_io_schema_fast_acl_rule.FastAclRuleSimpleAction_value[v.(string)])

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		port := make([]*ves_io_schema.PortValueType, len(sl))
		updateSpec.Port = port
		for i, set := range sl {
			port[i] = &ves_io_schema.PortValueType{}

			portMapStrToI := set.(map[string]interface{})

			portValueTypeChoiceTypeFound := false

			if v, ok := portMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true

				if v.(bool) {
					portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
					portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
					port[i].PortValueTypeChoice = portValueTypeChoiceInt
				}

			}

			if v, ok := portMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true

				if v.(bool) {
					portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
					portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
					port[i].PortValueTypeChoice = portValueTypeChoiceInt
				}

			}

			if v, ok := portMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

				portValueTypeChoiceTypeFound = true
				portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

				port[i].PortValueTypeChoice = portValueTypeChoiceInt

				portValueTypeChoiceInt.UserDefined =
					uint32(v.(int))

			}

		}

	}

	sourceTypeFound := false

	if v, ok := d.GetOk("ip_prefix_set"); ok && !sourceTypeFound {

		sourceTypeFound = true
		sourceInt := &ves_io_schema_fast_acl_rule.ReplaceSpecType_IpPrefixSet{}
		sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
		updateSpec.Source = sourceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["ref"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				sourceInt.IpPrefixSet.Ref = refInt
				for i, ps := range sl {

					rMapToStrVal := ps.(map[string]interface{})
					refInt[i] = &ves_io_schema.ObjectRefType{}

					refInt[i].Kind = "ip_prefix_set"

					if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
						refInt[i].Name = v.(string)
					}

					if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						refInt[i].Namespace = v.(string)
					}

					if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						refInt[i].Tenant = v.(string)
					}

					if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
						refInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("prefix"); ok && !sourceTypeFound {

		sourceTypeFound = true
		sourceInt := &ves_io_schema_fast_acl_rule.ReplaceSpecType_Prefix{}
		sourceInt.Prefix = &ves_io_schema.PrefixListType{}
		updateSpec.Source = sourceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				sourceInt.Prefix.Prefix = ls

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra FastAclRule obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_fast_acl_rule.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating FastAclRule: %s", err)
	}

	return resourceVolterraFastAclRuleRead(d, meta)
}

func resourceVolterraFastAclRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_fast_acl_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAclRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAclRule before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra FastAclRule obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_fast_acl_rule.ObjectType, namespace, name)
}
