//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package addon_service

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CanSubscribeReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CanSubscribeReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CanSubscribeReq) DeepCopy() *CanSubscribeReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CanSubscribeReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CanSubscribeReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CanSubscribeReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CanSubscribeReqValidator().Validate(ctx, m, opts...)
}

type ValidateCanSubscribeReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCanSubscribeReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CanSubscribeReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CanSubscribeReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["addon_service"]; exists {

		vOpts := append(opts, db.WithValidateField("addon_service"))
		if err := fv(ctx, m.GetAddonService(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCanSubscribeReqValidator = func() *ValidateCanSubscribeReq {
	v := &ValidateCanSubscribeReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CanSubscribeReqValidator() db.Validator {
	return DefaultCanSubscribeReqValidator
}

// augmented methods on protoc/std generated struct

func (m *CanSubscribeResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CanSubscribeResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CanSubscribeResp) DeepCopy() *CanSubscribeResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CanSubscribeResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CanSubscribeResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CanSubscribeResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CanSubscribeRespValidator().Validate(ctx, m, opts...)
}

type ValidateCanSubscribeResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCanSubscribeResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CanSubscribeResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CanSubscribeResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["result"]; exists {

		vOpts := append(opts, db.WithValidateField("result"))
		if err := fv(ctx, m.GetResult(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCanSubscribeRespValidator = func() *ValidateCanSubscribeResp {
	v := &ValidateCanSubscribeResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CanSubscribeRespValidator() db.Validator {
	return DefaultCanSubscribeRespValidator
}

// augmented methods on protoc/std generated struct

func (m *SetSubscriptionReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetSubscriptionReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetSubscriptionReq) DeepCopy() *SetSubscriptionReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetSubscriptionReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetSubscriptionReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetSubscriptionReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetSubscriptionReqValidator().Validate(ctx, m, opts...)
}

type ValidateSetSubscriptionReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetSubscriptionReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetSubscriptionReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetSubscriptionReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["action"]; exists {

		vOpts := append(opts, db.WithValidateField("action"))
		if err := fv(ctx, m.GetAction(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["addon_service"]; exists {

		vOpts := append(opts, db.WithValidateField("addon_service"))
		if err := fv(ctx, m.GetAddonService(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetSubscriptionReqValidator = func() *ValidateSetSubscriptionReq {
	v := &ValidateSetSubscriptionReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetSubscriptionReqValidator() db.Validator {
	return DefaultSetSubscriptionReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SetSubscriptionResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetSubscriptionResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetSubscriptionResp) DeepCopy() *SetSubscriptionResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetSubscriptionResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetSubscriptionResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetSubscriptionResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetSubscriptionRespValidator().Validate(ctx, m, opts...)
}

type ValidateSetSubscriptionResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetSubscriptionResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetSubscriptionResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetSubscriptionResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetSubscriptionRespValidator = func() *ValidateSetSubscriptionResp {
	v := &ValidateSetSubscriptionResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetSubscriptionRespValidator() db.Validator {
	return DefaultSetSubscriptionRespValidator
}
