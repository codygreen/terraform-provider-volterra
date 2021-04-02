//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package advertise_policy

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

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetTlsParameters().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateSpecType.tls_parameters")
	}

	return nil
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetPublicIpDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetWhereDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *CreateSpecType) GetPublicIpDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetPublicIp() {
		if ref == nil {
			return nil, fmt.Errorf("CreateSpecType.public_ip[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "public_ip.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "public_ip",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetPublicIpDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CreateSpecType) GetPublicIpDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "public_ip.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: public_ip")
	}
	for _, ref := range m.GetPublicIp() {
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

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetWhereDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GetWhere() == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GetWhere().GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "where." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) AddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for address")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) ProtocolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for protocol")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address"]; exists {

		vOpts := append(opts, db.WithValidateField("address"))
		if err := fv(ctx, m.GetAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["protocol"]; exists {

		vOpts := append(opts, db.WithValidateField("protocol"))
		if err := fv(ctx, m.GetProtocol(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("public_ip"))
		for idx, item := range m.GetPublicIp() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["skip_xff_append"]; exists {

		vOpts := append(opts, db.WithValidateField("skip_xff_append"))
		if err := fv(ctx, m.GetSkipXffAppend(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tls_parameters"]; exists {

		vOpts := append(opts, db.WithValidateField("tls_parameters"))
		if err := fv(ctx, m.GetTlsParameters(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["where"]; exists {

		vOpts := append(opts, db.WithValidateField("where"))
		if err := fv(ctx, m.GetWhere(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAddress := v.AddressValidationRuleHandler
	rulesAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhAddress(rulesAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address"] = vFn

	vrhProtocol := v.ProtocolValidationRuleHandler
	rulesProtocol := map[string]string{
		"ves.io.schema.rules.string.in": "[\"\",\"TCP\",\"UDP\"]",
	}
	vFn, err = vrhProtocol(rulesProtocol)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.protocol: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	v.FldValidators["where"] = ves_io_schema.NetworkSiteRefSelectorValidator().Validate

	v.FldValidators["tls_parameters"] = ves_io_schema.DownstreamTlsParamsTypeValidator().Validate

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GetSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetTlsParameters().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GetSpecType.tls_parameters")
	}

	return nil
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetPublicIpDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetWhereDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GetSpecType) GetPublicIpDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetPublicIp() {
		if ref == nil {
			return nil, fmt.Errorf("GetSpecType.public_ip[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "public_ip.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "public_ip",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetPublicIpDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetPublicIpDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "public_ip.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: public_ip")
	}
	for _, ref := range m.GetPublicIp() {
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

// GetDRefInfo for the field's type
func (m *GetSpecType) GetWhereDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GetWhere() == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GetWhere().GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "where." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) AddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for address")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) ProtocolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for protocol")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address"]; exists {

		vOpts := append(opts, db.WithValidateField("address"))
		if err := fv(ctx, m.GetAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["protocol"]; exists {

		vOpts := append(opts, db.WithValidateField("protocol"))
		if err := fv(ctx, m.GetProtocol(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("public_ip"))
		for idx, item := range m.GetPublicIp() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["skip_xff_append"]; exists {

		vOpts := append(opts, db.WithValidateField("skip_xff_append"))
		if err := fv(ctx, m.GetSkipXffAppend(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tls_parameters"]; exists {

		vOpts := append(opts, db.WithValidateField("tls_parameters"))
		if err := fv(ctx, m.GetTlsParameters(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["where"]; exists {

		vOpts := append(opts, db.WithValidateField("where"))
		if err := fv(ctx, m.GetWhere(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAddress := v.AddressValidationRuleHandler
	rulesAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhAddress(rulesAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address"] = vFn

	vrhProtocol := v.ProtocolValidationRuleHandler
	rulesProtocol := map[string]string{
		"ves.io.schema.rules.string.in": "[\"\",\"TCP\",\"UDP\"]",
	}
	vFn, err = vrhProtocol(rulesProtocol)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.protocol: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	v.FldValidators["where"] = ves_io_schema.NetworkSiteRefSelectorValidator().Validate

	v.FldValidators["tls_parameters"] = ves_io_schema.DownstreamTlsParamsTypeValidator().Validate

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetTlsParameters().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.tls_parameters")
	}

	return nil
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
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetPublicIpDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetWhereDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *GlobalSpecType) GetPublicIpDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetPublicIp() {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.public_ip[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "public_ip.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "public_ip",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetPublicIpDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetPublicIpDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "public_ip.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: public_ip")
	}
	for _, ref := range m.GetPublicIp() {
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

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetWhereDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GetWhere() == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GetWhere().GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "where." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) AddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for address")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ProtocolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for protocol")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
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

	if fv, exists := v.FldValidators["address"]; exists {

		vOpts := append(opts, db.WithValidateField("address"))
		if err := fv(ctx, m.GetAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["protocol"]; exists {

		vOpts := append(opts, db.WithValidateField("protocol"))
		if err := fv(ctx, m.GetProtocol(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("public_ip"))
		for idx, item := range m.GetPublicIp() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["skip_xff_append"]; exists {

		vOpts := append(opts, db.WithValidateField("skip_xff_append"))
		if err := fv(ctx, m.GetSkipXffAppend(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tls_parameters"]; exists {

		vOpts := append(opts, db.WithValidateField("tls_parameters"))
		if err := fv(ctx, m.GetTlsParameters(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["where"]; exists {

		vOpts := append(opts, db.WithValidateField("where"))
		if err := fv(ctx, m.GetWhere(), vOpts...); err != nil {
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

	vrhAddress := v.AddressValidationRuleHandler
	rulesAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhAddress(rulesAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address"] = vFn

	vrhProtocol := v.ProtocolValidationRuleHandler
	rulesProtocol := map[string]string{
		"ves.io.schema.rules.string.in": "[\"\",\"TCP\",\"UDP\"]",
	}
	vFn, err = vrhProtocol(rulesProtocol)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.protocol: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	v.FldValidators["where"] = ves_io_schema.NetworkSiteRefSelectorValidator().Validate

	v.FldValidators["tls_parameters"] = ves_io_schema.DownstreamTlsParamsTypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *ReplaceSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetTlsParameters().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting ReplaceSpecType.tls_parameters")
	}

	return nil
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetPublicIpDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetWhereDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

func (m *ReplaceSpecType) GetPublicIpDRefInfo() ([]db.DRefInfo, error) {
	drInfos := []db.DRefInfo{}
	for i, ref := range m.GetPublicIp() {
		if ref == nil {
			return nil, fmt.Errorf("ReplaceSpecType.public_ip[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "public_ip.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "public_ip",
			Ref:        ref,
		})
	}

	return drInfos, nil
}

// GetPublicIpDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ReplaceSpecType) GetPublicIpDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "public_ip.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: public_ip")
	}
	for _, ref := range m.GetPublicIp() {
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

// GetDRefInfo for the field's type
func (m *ReplaceSpecType) GetWhereDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GetWhere() == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GetWhere().GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "where." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) AddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for address")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) ProtocolValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for protocol")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) PortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for port")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address"]; exists {

		vOpts := append(opts, db.WithValidateField("address"))
		if err := fv(ctx, m.GetAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["port"]; exists {

		vOpts := append(opts, db.WithValidateField("port"))
		if err := fv(ctx, m.GetPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["protocol"]; exists {

		vOpts := append(opts, db.WithValidateField("protocol"))
		if err := fv(ctx, m.GetProtocol(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("public_ip"))
		for idx, item := range m.GetPublicIp() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["skip_xff_append"]; exists {

		vOpts := append(opts, db.WithValidateField("skip_xff_append"))
		if err := fv(ctx, m.GetSkipXffAppend(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tls_parameters"]; exists {

		vOpts := append(opts, db.WithValidateField("tls_parameters"))
		if err := fv(ctx, m.GetTlsParameters(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["where"]; exists {

		vOpts := append(opts, db.WithValidateField("where"))
		if err := fv(ctx, m.GetWhere(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAddress := v.AddressValidationRuleHandler
	rulesAddress := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhAddress(rulesAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address"] = vFn

	vrhProtocol := v.ProtocolValidationRuleHandler
	rulesProtocol := map[string]string{
		"ves.io.schema.rules.string.in": "[\"\",\"TCP\",\"UDP\"]",
	}
	vFn, err = vrhProtocol(rulesProtocol)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.protocol: %s", err)
		panic(errMsg)
	}
	v.FldValidators["protocol"] = vFn

	vrhPort := v.PortValidationRuleHandler
	rulesPort := map[string]string{
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhPort(rulesPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["port"] = vFn

	v.FldValidators["where"] = ves_io_schema.NetworkSiteRefSelectorValidator().Validate

	v.FldValidators["tls_parameters"] = ves_io_schema.DownstreamTlsParamsTypeValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Address = f.GetAddress()
	m.Port = f.GetPort()
	m.Protocol = f.GetProtocol()
	m.PublicIp = f.GetPublicIp()
	m.SkipXffAppend = f.GetSkipXffAppend()
	m.TlsParameters = f.GetTlsParameters()
	m.Where = f.GetWhere()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Address = m1.Address
	f.Port = m1.Port
	f.Protocol = m1.Protocol
	f.PublicIp = m1.PublicIp
	f.SkipXffAppend = m1.SkipXffAppend
	f.TlsParameters = m1.TlsParameters
	f.Where = m1.Where
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Address = f.GetAddress()
	m.Port = f.GetPort()
	m.Protocol = f.GetProtocol()
	m.PublicIp = f.GetPublicIp()
	m.SkipXffAppend = f.GetSkipXffAppend()
	m.TlsParameters = f.GetTlsParameters()
	m.Where = f.GetWhere()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Address = m1.Address
	f.Port = m1.Port
	f.Protocol = m1.Protocol
	f.PublicIp = m1.PublicIp
	f.SkipXffAppend = m1.SkipXffAppend
	f.TlsParameters = m1.TlsParameters
	f.Where = m1.Where
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Address = f.GetAddress()
	m.Port = f.GetPort()
	m.Protocol = f.GetProtocol()
	m.PublicIp = f.GetPublicIp()
	m.SkipXffAppend = f.GetSkipXffAppend()
	m.TlsParameters = f.GetTlsParameters()
	m.Where = f.GetWhere()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Address = m1.Address
	f.Port = m1.Port
	f.Protocol = m1.Protocol
	f.PublicIp = m1.PublicIp
	f.SkipXffAppend = m1.SkipXffAppend
	f.TlsParameters = m1.TlsParameters
	f.Where = m1.Where
}
