//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package api_credential

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetUsersDRefInfo()

}

func (m *GlobalSpecType) GetUsersDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetUsers()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.users[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "user.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "users",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetUsersDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetUsersDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "user.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: user")
	}
	for _, ref := range m.GetUsers() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) UsersValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for users")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated users")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items users")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) CertificateSerialNumValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for certificate_serial_num")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["active"]; exists {

		vOpts := append(opts, db.WithValidateField("active"))
		if err := fv(ctx, m.GetActive(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["certificate_serial_num"]; exists {

		vOpts := append(opts, db.WithValidateField("certificate_serial_num"))
		if err := fv(ctx, m.GetCertificateSerialNum(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["created_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("created_timestamp"))
		if err := fv(ctx, m.GetCreatedTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["digest"]; exists {

		vOpts := append(opts, db.WithValidateField("digest"))
		if err := fv(ctx, m.GetDigest(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_timestamp"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_timestamp"))
		if err := fv(ctx, m.GetExpirationTimestamp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["users"]; exists {
		vOpts := append(opts, db.WithValidateField("users"))
		if err := fv(ctx, m.GetUsers(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_name"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_name"))
		if err := fv(ctx, m.GetVirtualK8SName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_k8s_namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_k8s_namespace"))
		if err := fv(ctx, m.GetVirtualK8SNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhUsers := v.UsersValidationRuleHandler
	rulesUsers := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhUsers(rulesUsers)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.users: %s", err)
		panic(errMsg)
	}
	v.FldValidators["users"] = vFn

	vrhCertificateSerialNum := v.CertificateSerialNumValidationRuleHandler
	rulesCertificateSerialNum := map[string]string{
		"ves.io.schema.rules.string.max_len": "32",
	}
	vFn, err = vrhCertificateSerialNum(rulesCertificateSerialNum)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.certificate_serial_num: %s", err)
		panic(errMsg)
	}
	v.FldValidators["certificate_serial_num"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}
