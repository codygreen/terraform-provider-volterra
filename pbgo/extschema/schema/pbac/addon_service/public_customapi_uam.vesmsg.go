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

func (m *GetActivationStatusReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActivationStatusReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActivationStatusReq) DeepCopy() *GetActivationStatusReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActivationStatusReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActivationStatusReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActivationStatusReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActivationStatusReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetActivationStatusReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActivationStatusReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActivationStatusReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActivationStatusReq got type %s", t)
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
var DefaultGetActivationStatusReqValidator = func() *ValidateGetActivationStatusReq {
	v := &ValidateGetActivationStatusReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetActivationStatusReqValidator() db.Validator {
	return DefaultGetActivationStatusReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetActivationStatusResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActivationStatusResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActivationStatusResp) DeepCopy() *GetActivationStatusResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActivationStatusResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActivationStatusResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActivationStatusResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActivationStatusRespValidator().Validate(ctx, m, opts...)
}

type ValidateGetActivationStatusResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActivationStatusResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActivationStatusResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActivationStatusResp got type %s", t)
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
var DefaultGetActivationStatusRespValidator = func() *ValidateGetActivationStatusResp {
	v := &ValidateGetActivationStatusResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetActivationStatusRespValidator() db.Validator {
	return DefaultGetActivationStatusRespValidator
}
