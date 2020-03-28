// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogTarget Log Target
//
// Per-instance logging of events and traffic.
// swagger:model log_target
type LogTarget struct {

	// address
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// facility
	// Enum: [kern user mail daemon auth syslog lpr news uucp cron auth2 ftp ntp audit alert cron2 local0 local1 local2 local3 local4 local5 local6 local7]
	Facility string `json:"facility,omitempty"`

	// format
	// Enum: [rfc3164 rfc5424 short raw]
	Format string `json:"format,omitempty"`

	// global
	Global bool `json:"global,omitempty"`

	// index
	// Required: true
	Index *int64 `json:"index"`

	// length
	Length int64 `json:"length,omitempty"`

	// level
	// Enum: [emerg alert crit err warning notice info debug]
	Level string `json:"level,omitempty"`

	// minlevel
	// Enum: [emerg alert crit err warning notice info debug]
	Minlevel string `json:"minlevel,omitempty"`

	// nolog
	Nolog bool `json:"nolog,omitempty"`
}

// Validate validates this log target
func (m *LogTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinlevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogTarget) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("address", "body", string(m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var logTargetTypeFacilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kern","user","mail","daemon","auth","syslog","lpr","news","uucp","cron","auth2","ftp","ntp","audit","alert","cron2","local0","local1","local2","local3","local4","local5","local6","local7"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logTargetTypeFacilityPropEnum = append(logTargetTypeFacilityPropEnum, v)
	}
}

const (

	// LogTargetFacilityKern captures enum value "kern"
	LogTargetFacilityKern string = "kern"

	// LogTargetFacilityUser captures enum value "user"
	LogTargetFacilityUser string = "user"

	// LogTargetFacilityMail captures enum value "mail"
	LogTargetFacilityMail string = "mail"

	// LogTargetFacilityDaemon captures enum value "daemon"
	LogTargetFacilityDaemon string = "daemon"

	// LogTargetFacilityAuth captures enum value "auth"
	LogTargetFacilityAuth string = "auth"

	// LogTargetFacilitySyslog captures enum value "syslog"
	LogTargetFacilitySyslog string = "syslog"

	// LogTargetFacilityLpr captures enum value "lpr"
	LogTargetFacilityLpr string = "lpr"

	// LogTargetFacilityNews captures enum value "news"
	LogTargetFacilityNews string = "news"

	// LogTargetFacilityUucp captures enum value "uucp"
	LogTargetFacilityUucp string = "uucp"

	// LogTargetFacilityCron captures enum value "cron"
	LogTargetFacilityCron string = "cron"

	// LogTargetFacilityAuth2 captures enum value "auth2"
	LogTargetFacilityAuth2 string = "auth2"

	// LogTargetFacilityFtp captures enum value "ftp"
	LogTargetFacilityFtp string = "ftp"

	// LogTargetFacilityNtp captures enum value "ntp"
	LogTargetFacilityNtp string = "ntp"

	// LogTargetFacilityAudit captures enum value "audit"
	LogTargetFacilityAudit string = "audit"

	// LogTargetFacilityAlert captures enum value "alert"
	LogTargetFacilityAlert string = "alert"

	// LogTargetFacilityCron2 captures enum value "cron2"
	LogTargetFacilityCron2 string = "cron2"

	// LogTargetFacilityLocal0 captures enum value "local0"
	LogTargetFacilityLocal0 string = "local0"

	// LogTargetFacilityLocal1 captures enum value "local1"
	LogTargetFacilityLocal1 string = "local1"

	// LogTargetFacilityLocal2 captures enum value "local2"
	LogTargetFacilityLocal2 string = "local2"

	// LogTargetFacilityLocal3 captures enum value "local3"
	LogTargetFacilityLocal3 string = "local3"

	// LogTargetFacilityLocal4 captures enum value "local4"
	LogTargetFacilityLocal4 string = "local4"

	// LogTargetFacilityLocal5 captures enum value "local5"
	LogTargetFacilityLocal5 string = "local5"

	// LogTargetFacilityLocal6 captures enum value "local6"
	LogTargetFacilityLocal6 string = "local6"

	// LogTargetFacilityLocal7 captures enum value "local7"
	LogTargetFacilityLocal7 string = "local7"
)

// prop value enum
func (m *LogTarget) validateFacilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logTargetTypeFacilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogTarget) validateFacility(formats strfmt.Registry) error {

	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	// value enum
	if err := m.validateFacilityEnum("facility", "body", m.Facility); err != nil {
		return err
	}

	return nil
}

var logTargetTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rfc3164","rfc5424","short","raw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logTargetTypeFormatPropEnum = append(logTargetTypeFormatPropEnum, v)
	}
}

const (

	// LogTargetFormatRfc3164 captures enum value "rfc3164"
	LogTargetFormatRfc3164 string = "rfc3164"

	// LogTargetFormatRfc5424 captures enum value "rfc5424"
	LogTargetFormatRfc5424 string = "rfc5424"

	// LogTargetFormatShort captures enum value "short"
	LogTargetFormatShort string = "short"

	// LogTargetFormatRaw captures enum value "raw"
	LogTargetFormatRaw string = "raw"
)

// prop value enum
func (m *LogTarget) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logTargetTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogTarget) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *LogTarget) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

var logTargetTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logTargetTypeLevelPropEnum = append(logTargetTypeLevelPropEnum, v)
	}
}

const (

	// LogTargetLevelEmerg captures enum value "emerg"
	LogTargetLevelEmerg string = "emerg"

	// LogTargetLevelAlert captures enum value "alert"
	LogTargetLevelAlert string = "alert"

	// LogTargetLevelCrit captures enum value "crit"
	LogTargetLevelCrit string = "crit"

	// LogTargetLevelErr captures enum value "err"
	LogTargetLevelErr string = "err"

	// LogTargetLevelWarning captures enum value "warning"
	LogTargetLevelWarning string = "warning"

	// LogTargetLevelNotice captures enum value "notice"
	LogTargetLevelNotice string = "notice"

	// LogTargetLevelInfo captures enum value "info"
	LogTargetLevelInfo string = "info"

	// LogTargetLevelDebug captures enum value "debug"
	LogTargetLevelDebug string = "debug"
)

// prop value enum
func (m *LogTarget) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logTargetTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogTarget) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

var logTargetTypeMinlevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logTargetTypeMinlevelPropEnum = append(logTargetTypeMinlevelPropEnum, v)
	}
}

const (

	// LogTargetMinlevelEmerg captures enum value "emerg"
	LogTargetMinlevelEmerg string = "emerg"

	// LogTargetMinlevelAlert captures enum value "alert"
	LogTargetMinlevelAlert string = "alert"

	// LogTargetMinlevelCrit captures enum value "crit"
	LogTargetMinlevelCrit string = "crit"

	// LogTargetMinlevelErr captures enum value "err"
	LogTargetMinlevelErr string = "err"

	// LogTargetMinlevelWarning captures enum value "warning"
	LogTargetMinlevelWarning string = "warning"

	// LogTargetMinlevelNotice captures enum value "notice"
	LogTargetMinlevelNotice string = "notice"

	// LogTargetMinlevelInfo captures enum value "info"
	LogTargetMinlevelInfo string = "info"

	// LogTargetMinlevelDebug captures enum value "debug"
	LogTargetMinlevelDebug string = "debug"
)

// prop value enum
func (m *LogTarget) validateMinlevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logTargetTypeMinlevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogTarget) validateMinlevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Minlevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateMinlevelEnum("minlevel", "body", m.Minlevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogTarget) UnmarshalBinary(b []byte) error {
	var res LogTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
