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
	ves_io_schema_fast_acl "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl"
	ves_io_schema_fast_acl_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl_rule"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraFastAcl is implementation of Volterra's FastAcl resources
func resourceVolterraFastAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraFastAclCreate,
		Read:   resourceVolterraFastAclRead,
		Update: resourceVolterraFastAclUpdate,
		Delete: resourceVolterraFastAclDelete,

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

			"protocol_policer": {

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

			"legacy_acl": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"destination_type": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_services": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"destination_ip_address": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"ipv6": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"ports": {

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

												"protocol": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"interface_services": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"selected_vip_address": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"address": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"ipv6": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"addr": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},

									"shared_vip_services": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"vip_services": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"network_type": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"public": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"site_local": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"site_local_inside": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"source_rules": {

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

			"re_acl": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fast_acl_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"name": {
										Type:     schema.TypeString,
										Optional: true,
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
							},
						},

						"all_public_vips": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_tenant_vip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"selected_tenant_vip": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_tenant_vip": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"public_ip_refs": {

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
					},
				},
			},

			"site_acl": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fast_acl_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"name": {
										Type:     schema.TypeString,
										Optional: true,
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
							},
						},

						"inside_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"all_services": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"destination_ip_address": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ipv4": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"addr": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"ipv6": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"addr": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"ports": {

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

									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"interface_services": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"shared_vip_services": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"vip_services": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraFastAclCreate creates FastAcl resource
func resourceVolterraFastAclCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_fast_acl.CreateSpecType{}
	createReq := &ves_io_schema_fast_acl.CreateRequest{
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

	if v, ok := d.GetOk("protocol_policer"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		protocolPolicer := &ves_io_schema_views.ObjectRefType{}
		createSpec.ProtocolPolicer = protocolPolicer
		for _, set := range sl {

			protocolPolicerMapStrToI := set.(map[string]interface{})

			if w, ok := protocolPolicerMapStrToI["name"]; ok && !isIntfNil(w) {
				protocolPolicer.Name = w.(string)
			}

			if w, ok := protocolPolicerMapStrToI["namespace"]; ok && !isIntfNil(w) {
				protocolPolicer.Namespace = w.(string)
			}

			if w, ok := protocolPolicerMapStrToI["tenant"]; ok && !isIntfNil(w) {
				protocolPolicer.Tenant = w.(string)
			}

		}

	}

	siteChoiceTypeFound := false

	if v, ok := d.GetOk("legacy_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.CreateSpecType_LegacyAcl{}
		siteChoiceInt.LegacyAcl = &ves_io_schema_fast_acl.LegacyACLType{}
		createSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["destination_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				destinationType := &ves_io_schema_fast_acl.DestinationType{}
				siteChoiceInt.LegacyAcl.DestinationType = destinationType
				for _, set := range sl {

					destinationTypeMapStrToI := set.(map[string]interface{})

					destinationTypeChoiceTypeFound := false

					if v, ok := destinationTypeMapStrToI["all_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_AllServices{}
							destinationTypeChoiceInt.AllServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["destination_ip_address"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true
						destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_DestinationIpAddress{}
						destinationTypeChoiceInt.DestinationIpAddress = &ves_io_schema_fast_acl.DestinationIPAddressType{}
						destinationType.DestinationTypeChoice = destinationTypeChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["address"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								address := make([]*ves_io_schema.IpAddressType, len(sl))
								destinationTypeChoiceInt.DestinationIpAddress.Address = address
								for i, set := range sl {
									address[i] = &ves_io_schema.IpAddressType{}

									addressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

							if v, ok := cs["ports"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								ports := make([]*ves_io_schema.PortValueType, len(sl))
								destinationTypeChoiceInt.DestinationIpAddress.Ports = ports
								for i, set := range sl {
									ports[i] = &ves_io_schema.PortValueType{}

									portsMapStrToI := set.(map[string]interface{})

									portValueTypeChoiceTypeFound := false

									if v, ok := portsMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true

										if v.(bool) {
											portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
											portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
											ports[i].PortValueTypeChoice = portValueTypeChoiceInt
										}

									}

									if v, ok := portsMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true

										if v.(bool) {
											portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
											portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
											ports[i].PortValueTypeChoice = portValueTypeChoiceInt
										}

									}

									if v, ok := portsMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true
										portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

										ports[i].PortValueTypeChoice = portValueTypeChoiceInt

										portValueTypeChoiceInt.UserDefined =
											uint32(v.(int))

									}

								}

							}

							if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

								destinationTypeChoiceInt.DestinationIpAddress.Protocol = v.(string)
							}

						}

					}

					if v, ok := destinationTypeMapStrToI["interface_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_InterfaceServices{}
							destinationTypeChoiceInt.InterfaceServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["selected_vip_address"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true
						destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_SelectedVipAddress{}
						destinationTypeChoiceInt.SelectedVipAddress = &ves_io_schema_fast_acl.SelectedVIPAddressType{}
						destinationType.DestinationTypeChoice = destinationTypeChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["address"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								address := make([]*ves_io_schema.IpAddressType, len(sl))
								destinationTypeChoiceInt.SelectedVipAddress.Address = address
								for i, set := range sl {
									address[i] = &ves_io_schema.IpAddressType{}

									addressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

						}

					}

					if v, ok := destinationTypeMapStrToI["shared_vip_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_SharedVipServices{}
							destinationTypeChoiceInt.SharedVipServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["vip_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_VipServices{}
							destinationTypeChoiceInt.VipServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

				}

			}

			if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				networkType := &ves_io_schema.VirtualNetworkSelectorType{}
				siteChoiceInt.LegacyAcl.NetworkType = networkType
				for _, set := range sl {

					networkTypeMapStrToI := set.(map[string]interface{})

					vnTypeChoiceTypeFound := false

					if v, ok := networkTypeMapStrToI["public"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_Public{}
							vnTypeChoiceInt.Public = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

					if v, ok := networkTypeMapStrToI["site_local"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_SiteLocal{}
							vnTypeChoiceInt.SiteLocal = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

					if v, ok := networkTypeMapStrToI["site_local_inside"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_SiteLocalInside{}
							vnTypeChoiceInt.SiteLocalInside = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

				}

			}

			if v, ok := cs["source_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				sourceRulesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				siteChoiceInt.LegacyAcl.SourceRules = sourceRulesInt
				for i, ps := range sl {

					srMapToStrVal := ps.(map[string]interface{})
					sourceRulesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := srMapToStrVal["name"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Name = v.(string)
					}

					if v, ok := srMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Namespace = v.(string)
					}

					if v, ok := srMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("re_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.CreateSpecType_ReAcl{}
		siteChoiceInt.ReAcl = &ves_io_schema_fast_acl.ReACLType{}
		createSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acl_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclRules := make([]*ves_io_schema_fast_acl.FastACLRuleType, len(sl))
				siteChoiceInt.ReAcl.FastAclRules = fastAclRules
				for i, set := range sl {
					fastAclRules[i] = &ves_io_schema_fast_acl.FastACLRuleType{}

					fastAclRulesMapStrToI := set.(map[string]interface{})

					if v, ok := fastAclRulesMapStrToI["action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
						fastAclRules[i].Action = action
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

					if w, ok := fastAclRulesMapStrToI["name"]; ok && !isIntfNil(w) {
						fastAclRules[i].Name = w.(string)
					}

					if v, ok := fastAclRulesMapStrToI["port"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						port := make([]*ves_io_schema.PortValueType, len(sl))
						fastAclRules[i].Port = port
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

					if v, ok := fastAclRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_IpPrefixSet{}
						sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
						fastAclRules[i].Source = sourceInt

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

					if v, ok := fastAclRulesMapStrToI["prefix"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_Prefix{}
						sourceInt.Prefix = &ves_io_schema.PrefixListType{}
						fastAclRules[i].Source = sourceInt

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

				}

			}

			vipChoiceTypeFound := false

			if v, ok := cs["all_public_vips"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_AllPublicVips{}
					vipChoiceInt.AllPublicVips = &ves_io_schema.Empty{}
					siteChoiceInt.ReAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["default_tenant_vip"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_DefaultTenantVip{}
					vipChoiceInt.DefaultTenantVip = &ves_io_schema.Empty{}
					siteChoiceInt.ReAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["selected_tenant_vip"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true
				vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_SelectedTenantVip{}
				vipChoiceInt.SelectedTenantVip = &ves_io_schema_fast_acl.SelectedTenantVIPsType{}
				siteChoiceInt.ReAcl.VipChoice = vipChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["default_tenant_vip"]; ok && !isIntfNil(v) {

						vipChoiceInt.SelectedTenantVip.DefaultTenantVip = v.(bool)
					}

					if v, ok := cs["public_ip_refs"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						publicIpRefsInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						vipChoiceInt.SelectedTenantVip.PublicIpRefs = publicIpRefsInt
						for i, ps := range sl {

							pirMapToStrVal := ps.(map[string]interface{})
							publicIpRefsInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := pirMapToStrVal["name"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Name = v.(string)
							}

							if v, ok := pirMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Namespace = v.(string)
							}

							if v, ok := pirMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("site_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.CreateSpecType_SiteAcl{}
		siteChoiceInt.SiteAcl = &ves_io_schema_fast_acl.SiteACLType{}
		createSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acl_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclRules := make([]*ves_io_schema_fast_acl.FastACLRuleType, len(sl))
				siteChoiceInt.SiteAcl.FastAclRules = fastAclRules
				for i, set := range sl {
					fastAclRules[i] = &ves_io_schema_fast_acl.FastACLRuleType{}

					fastAclRulesMapStrToI := set.(map[string]interface{})

					if v, ok := fastAclRulesMapStrToI["action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
						fastAclRules[i].Action = action
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

					if w, ok := fastAclRulesMapStrToI["name"]; ok && !isIntfNil(w) {
						fastAclRules[i].Name = w.(string)
					}

					if v, ok := fastAclRulesMapStrToI["port"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						port := make([]*ves_io_schema.PortValueType, len(sl))
						fastAclRules[i].Port = port
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

					if v, ok := fastAclRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_IpPrefixSet{}
						sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
						fastAclRules[i].Source = sourceInt

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

					if v, ok := fastAclRulesMapStrToI["prefix"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_Prefix{}
						sourceInt.Prefix = &ves_io_schema.PrefixListType{}
						fastAclRules[i].Source = sourceInt

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

				}

			}

			networkChoiceTypeFound := false

			if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

				networkChoiceTypeFound = true

				if v.(bool) {
					networkChoiceInt := &ves_io_schema_fast_acl.SiteACLType_InsideNetwork{}
					networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.NetworkChoice = networkChoiceInt
				}

			}

			if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

				networkChoiceTypeFound = true

				if v.(bool) {
					networkChoiceInt := &ves_io_schema_fast_acl.SiteACLType_OutsideNetwork{}
					networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.NetworkChoice = networkChoiceInt
				}

			}

			vipChoiceTypeFound := false

			if v, ok := cs["all_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_AllServices{}
					vipChoiceInt.AllServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["destination_ip_address"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true
				vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_DestinationIpAddress{}
				vipChoiceInt.DestinationIpAddress = &ves_io_schema_fast_acl.DestinationIPAddressType{}
				siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["address"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						address := make([]*ves_io_schema.IpAddressType, len(sl))
						vipChoiceInt.DestinationIpAddress.Address = address
						for i, set := range sl {
							address[i] = &ves_io_schema.IpAddressType{}

							addressMapStrToI := set.(map[string]interface{})

							verTypeFound := false

							if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpAddressType_Ipv4{}
								verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
								address[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["addr"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Addr = v.(string)
									}

								}

							}

							if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpAddressType_Ipv6{}
								verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
								address[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["addr"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Addr = v.(string)
									}

								}

							}

						}

					}

					if v, ok := cs["ports"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						ports := make([]*ves_io_schema.PortValueType, len(sl))
						vipChoiceInt.DestinationIpAddress.Ports = ports
						for i, set := range sl {
							ports[i] = &ves_io_schema.PortValueType{}

							portsMapStrToI := set.(map[string]interface{})

							portValueTypeChoiceTypeFound := false

							if v, ok := portsMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true

								if v.(bool) {
									portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
									portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
									ports[i].PortValueTypeChoice = portValueTypeChoiceInt
								}

							}

							if v, ok := portsMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true

								if v.(bool) {
									portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
									portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
									ports[i].PortValueTypeChoice = portValueTypeChoiceInt
								}

							}

							if v, ok := portsMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true
								portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

								ports[i].PortValueTypeChoice = portValueTypeChoiceInt

								portValueTypeChoiceInt.UserDefined =
									uint32(v.(int))

							}

						}

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						vipChoiceInt.DestinationIpAddress.Protocol = v.(string)
					}

				}

			}

			if v, ok := cs["interface_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_InterfaceServices{}
					vipChoiceInt.InterfaceServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["shared_vip_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_SharedVipServices{}
					vipChoiceInt.SharedVipServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["vip_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_VipServices{}
					vipChoiceInt.VipServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra FastAcl object with struct: %+v", createReq)

	createFastAclResp, err := client.CreateObject(context.Background(), ves_io_schema_fast_acl.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating FastAcl: %s", err)
	}
	d.SetId(createFastAclResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraFastAclRead(d, meta)
}

func resourceVolterraFastAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_fast_acl.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAcl %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAcl %q: %s", d.Id(), err)
	}
	return setFastAclFields(client, d, resp)
}

func setFastAclFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraFastAclUpdate updates FastAcl resource
func resourceVolterraFastAclUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_fast_acl.ReplaceSpecType{}
	updateReq := &ves_io_schema_fast_acl.ReplaceRequest{
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

	if v, ok := d.GetOk("protocol_policer"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		protocolPolicer := &ves_io_schema_views.ObjectRefType{}
		updateSpec.ProtocolPolicer = protocolPolicer
		for _, set := range sl {

			protocolPolicerMapStrToI := set.(map[string]interface{})

			if w, ok := protocolPolicerMapStrToI["name"]; ok && !isIntfNil(w) {
				protocolPolicer.Name = w.(string)
			}

			if w, ok := protocolPolicerMapStrToI["namespace"]; ok && !isIntfNil(w) {
				protocolPolicer.Namespace = w.(string)
			}

			if w, ok := protocolPolicerMapStrToI["tenant"]; ok && !isIntfNil(w) {
				protocolPolicer.Tenant = w.(string)
			}

		}

	}

	siteChoiceTypeFound := false

	if v, ok := d.GetOk("legacy_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.ReplaceSpecType_LegacyAcl{}
		siteChoiceInt.LegacyAcl = &ves_io_schema_fast_acl.LegacyACLType{}
		updateSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["destination_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				destinationType := &ves_io_schema_fast_acl.DestinationType{}
				siteChoiceInt.LegacyAcl.DestinationType = destinationType
				for _, set := range sl {

					destinationTypeMapStrToI := set.(map[string]interface{})

					destinationTypeChoiceTypeFound := false

					if v, ok := destinationTypeMapStrToI["all_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_AllServices{}
							destinationTypeChoiceInt.AllServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["destination_ip_address"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true
						destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_DestinationIpAddress{}
						destinationTypeChoiceInt.DestinationIpAddress = &ves_io_schema_fast_acl.DestinationIPAddressType{}
						destinationType.DestinationTypeChoice = destinationTypeChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["address"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								address := make([]*ves_io_schema.IpAddressType, len(sl))
								destinationTypeChoiceInt.DestinationIpAddress.Address = address
								for i, set := range sl {
									address[i] = &ves_io_schema.IpAddressType{}

									addressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

							if v, ok := cs["ports"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								ports := make([]*ves_io_schema.PortValueType, len(sl))
								destinationTypeChoiceInt.DestinationIpAddress.Ports = ports
								for i, set := range sl {
									ports[i] = &ves_io_schema.PortValueType{}

									portsMapStrToI := set.(map[string]interface{})

									portValueTypeChoiceTypeFound := false

									if v, ok := portsMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true

										if v.(bool) {
											portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
											portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
											ports[i].PortValueTypeChoice = portValueTypeChoiceInt
										}

									}

									if v, ok := portsMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true

										if v.(bool) {
											portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
											portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
											ports[i].PortValueTypeChoice = portValueTypeChoiceInt
										}

									}

									if v, ok := portsMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

										portValueTypeChoiceTypeFound = true
										portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

										ports[i].PortValueTypeChoice = portValueTypeChoiceInt

										portValueTypeChoiceInt.UserDefined =
											uint32(v.(int))

									}

								}

							}

							if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

								destinationTypeChoiceInt.DestinationIpAddress.Protocol = v.(string)
							}

						}

					}

					if v, ok := destinationTypeMapStrToI["interface_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_InterfaceServices{}
							destinationTypeChoiceInt.InterfaceServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["selected_vip_address"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true
						destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_SelectedVipAddress{}
						destinationTypeChoiceInt.SelectedVipAddress = &ves_io_schema_fast_acl.SelectedVIPAddressType{}
						destinationType.DestinationTypeChoice = destinationTypeChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["address"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								address := make([]*ves_io_schema.IpAddressType, len(sl))
								destinationTypeChoiceInt.SelectedVipAddress.Address = address
								for i, set := range sl {
									address[i] = &ves_io_schema.IpAddressType{}

									addressMapStrToI := set.(map[string]interface{})

									verTypeFound := false

									if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv4{}
										verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv4.Addr = v.(string)
											}

										}

									}

									if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

										verTypeFound = true
										verInt := &ves_io_schema.IpAddressType_Ipv6{}
										verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
										address[i].Ver = verInt

										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["addr"]; ok && !isIntfNil(v) {

												verInt.Ipv6.Addr = v.(string)
											}

										}

									}

								}

							}

						}

					}

					if v, ok := destinationTypeMapStrToI["shared_vip_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_SharedVipServices{}
							destinationTypeChoiceInt.SharedVipServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

					if v, ok := destinationTypeMapStrToI["vip_services"]; ok && !isIntfNil(v) && !destinationTypeChoiceTypeFound {

						destinationTypeChoiceTypeFound = true

						if v.(bool) {
							destinationTypeChoiceInt := &ves_io_schema_fast_acl.DestinationType_VipServices{}
							destinationTypeChoiceInt.VipServices = &ves_io_schema.Empty{}
							destinationType.DestinationTypeChoice = destinationTypeChoiceInt
						}

					}

				}

			}

			if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				networkType := &ves_io_schema.VirtualNetworkSelectorType{}
				siteChoiceInt.LegacyAcl.NetworkType = networkType
				for _, set := range sl {

					networkTypeMapStrToI := set.(map[string]interface{})

					vnTypeChoiceTypeFound := false

					if v, ok := networkTypeMapStrToI["public"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_Public{}
							vnTypeChoiceInt.Public = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

					if v, ok := networkTypeMapStrToI["site_local"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_SiteLocal{}
							vnTypeChoiceInt.SiteLocal = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

					if v, ok := networkTypeMapStrToI["site_local_inside"]; ok && !isIntfNil(v) && !vnTypeChoiceTypeFound {

						vnTypeChoiceTypeFound = true

						if v.(bool) {
							vnTypeChoiceInt := &ves_io_schema.VirtualNetworkSelectorType_SiteLocalInside{}
							vnTypeChoiceInt.SiteLocalInside = &ves_io_schema.Empty{}
							networkType.VnTypeChoice = vnTypeChoiceInt
						}

					}

				}

			}

			if v, ok := cs["source_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				sourceRulesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				siteChoiceInt.LegacyAcl.SourceRules = sourceRulesInt
				for i, ps := range sl {

					srMapToStrVal := ps.(map[string]interface{})
					sourceRulesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := srMapToStrVal["name"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Name = v.(string)
					}

					if v, ok := srMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Namespace = v.(string)
					}

					if v, ok := srMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						sourceRulesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("re_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.ReplaceSpecType_ReAcl{}
		siteChoiceInt.ReAcl = &ves_io_schema_fast_acl.ReACLType{}
		updateSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acl_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclRules := make([]*ves_io_schema_fast_acl.FastACLRuleType, len(sl))
				siteChoiceInt.ReAcl.FastAclRules = fastAclRules
				for i, set := range sl {
					fastAclRules[i] = &ves_io_schema_fast_acl.FastACLRuleType{}

					fastAclRulesMapStrToI := set.(map[string]interface{})

					if v, ok := fastAclRulesMapStrToI["action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
						fastAclRules[i].Action = action
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

					if w, ok := fastAclRulesMapStrToI["name"]; ok && !isIntfNil(w) {
						fastAclRules[i].Name = w.(string)
					}

					if v, ok := fastAclRulesMapStrToI["port"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						port := make([]*ves_io_schema.PortValueType, len(sl))
						fastAclRules[i].Port = port
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

					if v, ok := fastAclRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_IpPrefixSet{}
						sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
						fastAclRules[i].Source = sourceInt

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

					if v, ok := fastAclRulesMapStrToI["prefix"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_Prefix{}
						sourceInt.Prefix = &ves_io_schema.PrefixListType{}
						fastAclRules[i].Source = sourceInt

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

				}

			}

			vipChoiceTypeFound := false

			if v, ok := cs["all_public_vips"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_AllPublicVips{}
					vipChoiceInt.AllPublicVips = &ves_io_schema.Empty{}
					siteChoiceInt.ReAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["default_tenant_vip"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_DefaultTenantVip{}
					vipChoiceInt.DefaultTenantVip = &ves_io_schema.Empty{}
					siteChoiceInt.ReAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["selected_tenant_vip"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true
				vipChoiceInt := &ves_io_schema_fast_acl.ReACLType_SelectedTenantVip{}
				vipChoiceInt.SelectedTenantVip = &ves_io_schema_fast_acl.SelectedTenantVIPsType{}
				siteChoiceInt.ReAcl.VipChoice = vipChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["default_tenant_vip"]; ok && !isIntfNil(v) {

						vipChoiceInt.SelectedTenantVip.DefaultTenantVip = v.(bool)
					}

					if v, ok := cs["public_ip_refs"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						publicIpRefsInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						vipChoiceInt.SelectedTenantVip.PublicIpRefs = publicIpRefsInt
						for i, ps := range sl {

							pirMapToStrVal := ps.(map[string]interface{})
							publicIpRefsInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := pirMapToStrVal["name"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Name = v.(string)
							}

							if v, ok := pirMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Namespace = v.(string)
							}

							if v, ok := pirMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								publicIpRefsInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("site_acl"); ok && !siteChoiceTypeFound {

		siteChoiceTypeFound = true
		siteChoiceInt := &ves_io_schema_fast_acl.ReplaceSpecType_SiteAcl{}
		siteChoiceInt.SiteAcl = &ves_io_schema_fast_acl.SiteACLType{}
		updateSpec.SiteChoice = siteChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acl_rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclRules := make([]*ves_io_schema_fast_acl.FastACLRuleType, len(sl))
				siteChoiceInt.SiteAcl.FastAclRules = fastAclRules
				for i, set := range sl {
					fastAclRules[i] = &ves_io_schema_fast_acl.FastACLRuleType{}

					fastAclRulesMapStrToI := set.(map[string]interface{})

					if v, ok := fastAclRulesMapStrToI["action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						action := &ves_io_schema_fast_acl_rule.FastAclRuleAction{}
						fastAclRules[i].Action = action
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

					if w, ok := fastAclRulesMapStrToI["name"]; ok && !isIntfNil(w) {
						fastAclRules[i].Name = w.(string)
					}

					if v, ok := fastAclRulesMapStrToI["port"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						port := make([]*ves_io_schema.PortValueType, len(sl))
						fastAclRules[i].Port = port
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

					if v, ok := fastAclRulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_IpPrefixSet{}
						sourceInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
						fastAclRules[i].Source = sourceInt

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

					if v, ok := fastAclRulesMapStrToI["prefix"]; ok && !isIntfNil(v) && !sourceTypeFound {

						sourceTypeFound = true
						sourceInt := &ves_io_schema_fast_acl.FastACLRuleType_Prefix{}
						sourceInt.Prefix = &ves_io_schema.PrefixListType{}
						fastAclRules[i].Source = sourceInt

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

				}

			}

			networkChoiceTypeFound := false

			if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

				networkChoiceTypeFound = true

				if v.(bool) {
					networkChoiceInt := &ves_io_schema_fast_acl.SiteACLType_InsideNetwork{}
					networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.NetworkChoice = networkChoiceInt
				}

			}

			if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

				networkChoiceTypeFound = true

				if v.(bool) {
					networkChoiceInt := &ves_io_schema_fast_acl.SiteACLType_OutsideNetwork{}
					networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.NetworkChoice = networkChoiceInt
				}

			}

			vipChoiceTypeFound := false

			if v, ok := cs["all_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_AllServices{}
					vipChoiceInt.AllServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["destination_ip_address"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true
				vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_DestinationIpAddress{}
				vipChoiceInt.DestinationIpAddress = &ves_io_schema_fast_acl.DestinationIPAddressType{}
				siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["address"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						address := make([]*ves_io_schema.IpAddressType, len(sl))
						vipChoiceInt.DestinationIpAddress.Address = address
						for i, set := range sl {
							address[i] = &ves_io_schema.IpAddressType{}

							addressMapStrToI := set.(map[string]interface{})

							verTypeFound := false

							if v, ok := addressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpAddressType_Ipv4{}
								verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
								address[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["addr"]; ok && !isIntfNil(v) {

										verInt.Ipv4.Addr = v.(string)
									}

								}

							}

							if v, ok := addressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

								verTypeFound = true
								verInt := &ves_io_schema.IpAddressType_Ipv6{}
								verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
								address[i].Ver = verInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["addr"]; ok && !isIntfNil(v) {

										verInt.Ipv6.Addr = v.(string)
									}

								}

							}

						}

					}

					if v, ok := cs["ports"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						ports := make([]*ves_io_schema.PortValueType, len(sl))
						vipChoiceInt.DestinationIpAddress.Ports = ports
						for i, set := range sl {
							ports[i] = &ves_io_schema.PortValueType{}

							portsMapStrToI := set.(map[string]interface{})

							portValueTypeChoiceTypeFound := false

							if v, ok := portsMapStrToI["all"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true

								if v.(bool) {
									portValueTypeChoiceInt := &ves_io_schema.PortValueType_All{}
									portValueTypeChoiceInt.All = &ves_io_schema.Empty{}
									ports[i].PortValueTypeChoice = portValueTypeChoiceInt
								}

							}

							if v, ok := portsMapStrToI["dns"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true

								if v.(bool) {
									portValueTypeChoiceInt := &ves_io_schema.PortValueType_Dns{}
									portValueTypeChoiceInt.Dns = &ves_io_schema.Empty{}
									ports[i].PortValueTypeChoice = portValueTypeChoiceInt
								}

							}

							if v, ok := portsMapStrToI["user_defined"]; ok && !isIntfNil(v) && !portValueTypeChoiceTypeFound {

								portValueTypeChoiceTypeFound = true
								portValueTypeChoiceInt := &ves_io_schema.PortValueType_UserDefined{}

								ports[i].PortValueTypeChoice = portValueTypeChoiceInt

								portValueTypeChoiceInt.UserDefined =
									uint32(v.(int))

							}

						}

					}

					if v, ok := cs["protocol"]; ok && !isIntfNil(v) {

						vipChoiceInt.DestinationIpAddress.Protocol = v.(string)
					}

				}

			}

			if v, ok := cs["interface_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_InterfaceServices{}
					vipChoiceInt.InterfaceServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["shared_vip_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_SharedVipServices{}
					vipChoiceInt.SharedVipServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

			if v, ok := cs["vip_services"]; ok && !isIntfNil(v) && !vipChoiceTypeFound {

				vipChoiceTypeFound = true

				if v.(bool) {
					vipChoiceInt := &ves_io_schema_fast_acl.SiteACLType_VipServices{}
					vipChoiceInt.VipServices = &ves_io_schema.Empty{}
					siteChoiceInt.SiteAcl.VipChoice = vipChoiceInt
				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra FastAcl obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_fast_acl.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating FastAcl: %s", err)
	}

	return resourceVolterraFastAclRead(d, meta)
}

func resourceVolterraFastAclDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_fast_acl.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAcl %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAcl before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra FastAcl obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_fast_acl.ObjectType, namespace, name)
}
