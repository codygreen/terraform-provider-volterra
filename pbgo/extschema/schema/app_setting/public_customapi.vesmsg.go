//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package app_setting

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

func (m *SuspiciousUser) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuspiciousUser) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuspiciousUser) DeepCopy() *SuspiciousUser {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuspiciousUser{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuspiciousUser) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuspiciousUser) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuspiciousUserValidator().Validate(ctx, m, opts...)
}

type ValidateSuspiciousUser struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuspiciousUser) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuspiciousUser)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuspiciousUser got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["logs"]; exists {

		vOpts := append(opts, db.WithValidateField("logs"))
		for idx, item := range m.GetLogs() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["suspicion_score"]; exists {

		vOpts := append(opts, db.WithValidateField("suspicion_score"))
		if err := fv(ctx, m.GetSuspicionScore(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user_id"]; exists {

		vOpts := append(opts, db.WithValidateField("user_id"))
		if err := fv(ctx, m.GetUserId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuspiciousUserValidator = func() *ValidateSuspiciousUser {
	v := &ValidateSuspiciousUser{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuspiciousUserValidator() db.Validator {
	return DefaultSuspiciousUserValidator
}

// augmented methods on protoc/std generated struct

func (m *SuspiciousUserStatusReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuspiciousUserStatusReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuspiciousUserStatusReq) DeepCopy() *SuspiciousUserStatusReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuspiciousUserStatusReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuspiciousUserStatusReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuspiciousUserStatusReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuspiciousUserStatusReqValidator().Validate(ctx, m, opts...)
}

type ValidateSuspiciousUserStatusReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuspiciousUserStatusReq) TopnValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for topn")
	}

	return validatorFn, nil
}

func (v *ValidateSuspiciousUserStatusReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuspiciousUserStatusReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuspiciousUserStatusReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query"]; exists {

		vOpts := append(opts, db.WithValidateField("query"))
		if err := fv(ctx, m.GetQuery(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["topn"]; exists {

		vOpts := append(opts, db.WithValidateField("topn"))
		if err := fv(ctx, m.GetTopn(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuspiciousUserStatusReqValidator = func() *ValidateSuspiciousUserStatusReq {
	v := &ValidateSuspiciousUserStatusReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhTopn := v.TopnValidationRuleHandler
	rulesTopn := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "100",
	}
	vFn, err = vrhTopn(rulesTopn)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SuspiciousUserStatusReq.topn: %s", err)
		panic(errMsg)
	}
	v.FldValidators["topn"] = vFn

	return v
}()

func SuspiciousUserStatusReqValidator() db.Validator {
	return DefaultSuspiciousUserStatusReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SuspiciousUserStatusRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SuspiciousUserStatusRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SuspiciousUserStatusRsp) DeepCopy() *SuspiciousUserStatusRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SuspiciousUserStatusRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SuspiciousUserStatusRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SuspiciousUserStatusRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SuspiciousUserStatusRspValidator().Validate(ctx, m, opts...)
}

type ValidateSuspiciousUserStatusRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSuspiciousUserStatusRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SuspiciousUserStatusRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SuspiciousUserStatusRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["suspicious_users"]; exists {

		vOpts := append(opts, db.WithValidateField("suspicious_users"))
		for idx, item := range m.GetSuspiciousUsers() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSuspiciousUserStatusRspValidator = func() *ValidateSuspiciousUserStatusRsp {
	v := &ValidateSuspiciousUserStatusRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SuspiciousUserStatusRspValidator() db.Validator {
	return DefaultSuspiciousUserStatusRspValidator
}
