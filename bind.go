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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Bind Bind
//
// HAProxy frontend bind configuration
//
// swagger:model bind
type Bind struct {

	// accept netscaler cip
	AcceptNetscalerCip int64 `json:"accept_netscaler_cip,omitempty"`

	// accept proxy
	AcceptProxy bool `json:"accept_proxy,omitempty"`

	// address
	// Pattern: ^[^\s]+$
	Address string `json:"address,omitempty"`

	// allow 0rtt
	Allow0rtt bool `json:"allow_0rtt,omitempty"`

	// alpn
	// Pattern: ^[^\s]+$
	Alpn string `json:"alpn,omitempty"`

	// backlog
	Backlog string `json:"backlog,omitempty"`

	// ca ignore err
	CaIgnoreErr string `json:"ca_ignore_err,omitempty"`

	// ca sign file
	CaSignFile string `json:"ca_sign_file,omitempty"`

	// ca sign pass
	CaSignPass string `json:"ca_sign_pass,omitempty"`

	// ca verify file
	CaVerifyFile string `json:"ca_verify_file,omitempty"`

	// ciphers
	Ciphers string `json:"ciphers,omitempty"`

	// ciphersuites
	Ciphersuites string `json:"ciphersuites,omitempty"`

	// crl file
	CrlFile string `json:"crl_file,omitempty"`

	// crt ignore err
	CrtIgnoreErr string `json:"crt_ignore_err,omitempty"`

	// crt list
	CrtList string `json:"crt_list,omitempty"`

	// curves
	Curves string `json:"curves,omitempty"`

	// defer accept
	DeferAccept bool `json:"defer_accept,omitempty"`

	// ecdhe
	Ecdhe string `json:"ecdhe,omitempty"`

	// expose fd listeners
	ExposeFdListeners bool `json:"expose_fd_listeners,omitempty"`

	// force sslv3
	ForceSslv3 bool `json:"force_sslv3,omitempty"`

	// force tlsv10
	ForceTlsv10 bool `json:"force_tlsv10,omitempty"`

	// force tlsv11
	ForceTlsv11 bool `json:"force_tlsv11,omitempty"`

	// force tlsv12
	ForceTlsv12 bool `json:"force_tlsv12,omitempty"`

	// force tlsv13
	ForceTlsv13 bool `json:"force_tlsv13,omitempty"`

	// generate certificates
	GenerateCertificates bool `json:"generate_certificates,omitempty"`

	// gid
	Gid int64 `json:"gid,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// interface
	Interface string `json:"interface,omitempty"`

	// level
	// Enum: [user operator admin]
	Level string `json:"level,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// mss
	Mss string `json:"mss,omitempty"`

	// name
	// Required: true
	// Pattern: ^[^\s]+$
	Name string `json:"name"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// nice
	Nice int64 `json:"nice,omitempty"`

	// no ca names
	NoCaNames bool `json:"no_ca_names,omitempty"`

	// no sslv3
	NoSslv3 bool `json:"no_sslv3,omitempty"`

	// no tls tickets
	NoTLSTickets bool `json:"no_tls_tickets,omitempty"`

	// no tlsv10
	NoTlsv10 bool `json:"no_tlsv10,omitempty"`

	// no tlsv11
	NoTlsv11 bool `json:"no_tlsv11,omitempty"`

	// no tlsv12
	NoTlsv12 bool `json:"no_tlsv12,omitempty"`

	// no tlsv13
	NoTlsv13 bool `json:"no_tlsv13,omitempty"`

	// npn
	Npn string `json:"npn,omitempty"`

	// port
	// Maximum: 65535
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`

	// prefer client ciphers
	PreferClientCiphers bool `json:"prefer_client_ciphers,omitempty"`

	// process
	// Pattern: ^[^\s]+$
	Process string `json:"process,omitempty"`

	// proto
	Proto string `json:"proto,omitempty"`

	// severity output
	// Enum: [none number string]
	SeverityOutput string `json:"severity_output,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// ssl cafile
	// Pattern: ^[^\s]+$
	SslCafile string `json:"ssl_cafile,omitempty"`

	// ssl certificate
	// Pattern: ^[^\s]+$
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// ssl max ver
	// Enum: [SSLv3 TLSv1.0 TLSv1.1 TLSv1.2 TLSv1.3]
	SslMaxVer string `json:"ssl_max_ver,omitempty"`

	// ssl min ver
	// Enum: [SSLv3 TLSv1.0 TLSv1.1 TLSv1.2 TLSv1.3]
	SslMinVer string `json:"ssl_min_ver,omitempty"`

	// strict sni
	StrictSni bool `json:"strict_sni,omitempty"`

	// tcp user timeout
	TCPUserTimeout *int64 `json:"tcp_user_timeout,omitempty"`

	// tfo
	Tfo bool `json:"tfo,omitempty"`

	// tls ticket keys
	TLSTicketKeys string `json:"tls_ticket_keys,omitempty"`

	// transparent
	Transparent bool `json:"transparent,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// v4v6
	V4v6 bool `json:"v4v6,omitempty"`

	// v6only
	V6only bool `json:"v6only,omitempty"`

	// verify
	// Enum: [none optional required]
	Verify string `json:"verify,omitempty"`
}

// Validate validates this bind
func (m *Bind) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlpn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverityOutput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCafile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslMaxVer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslMinVer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bind) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.Pattern("address", "body", string(m.Address), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateAlpn(formats strfmt.Registry) error {

	if swag.IsZero(m.Alpn) { // not required
		return nil
	}

	if err := validate.Pattern("alpn", "body", string(m.Alpn), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","operator","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindTypeLevelPropEnum = append(bindTypeLevelPropEnum, v)
	}
}

const (

	// BindLevelUser captures enum value "user"
	BindLevelUser string = "user"

	// BindLevelOperator captures enum value "operator"
	BindLevelOperator string = "operator"

	// BindLevelAdmin captures enum value "admin"
	BindLevelAdmin string = "admin"
)

// prop value enum
func (m *Bind) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bindTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Bind) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if err := validate.Pattern("process", "body", string(m.Process), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindTypeSeverityOutputPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","number","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindTypeSeverityOutputPropEnum = append(bindTypeSeverityOutputPropEnum, v)
	}
}

const (

	// BindSeverityOutputNone captures enum value "none"
	BindSeverityOutputNone string = "none"

	// BindSeverityOutputNumber captures enum value "number"
	BindSeverityOutputNumber string = "number"

	// BindSeverityOutputString captures enum value "string"
	BindSeverityOutputString string = "string"
)

// prop value enum
func (m *Bind) validateSeverityOutputEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bindTypeSeverityOutputPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Bind) validateSeverityOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.SeverityOutput) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityOutputEnum("severity_output", "body", m.SeverityOutput); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateSslCafile(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCafile) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_cafile", "body", string(m.SslCafile), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Bind) validateSslCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCertificate) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_certificate", "body", string(m.SslCertificate), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindTypeSslMaxVerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSLv3","TLSv1.0","TLSv1.1","TLSv1.2","TLSv1.3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindTypeSslMaxVerPropEnum = append(bindTypeSslMaxVerPropEnum, v)
	}
}

const (

	// BindSslMaxVerSSLv3 captures enum value "SSLv3"
	BindSslMaxVerSSLv3 string = "SSLv3"

	// BindSslMaxVerTLSv10 captures enum value "TLSv1.0"
	BindSslMaxVerTLSv10 string = "TLSv1.0"

	// BindSslMaxVerTLSv11 captures enum value "TLSv1.1"
	BindSslMaxVerTLSv11 string = "TLSv1.1"

	// BindSslMaxVerTLSv12 captures enum value "TLSv1.2"
	BindSslMaxVerTLSv12 string = "TLSv1.2"

	// BindSslMaxVerTLSv13 captures enum value "TLSv1.3"
	BindSslMaxVerTLSv13 string = "TLSv1.3"
)

// prop value enum
func (m *Bind) validateSslMaxVerEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bindTypeSslMaxVerPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Bind) validateSslMaxVer(formats strfmt.Registry) error {

	if swag.IsZero(m.SslMaxVer) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslMaxVerEnum("ssl_max_ver", "body", m.SslMaxVer); err != nil {
		return err
	}

	return nil
}

var bindTypeSslMinVerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSLv3","TLSv1.0","TLSv1.1","TLSv1.2","TLSv1.3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindTypeSslMinVerPropEnum = append(bindTypeSslMinVerPropEnum, v)
	}
}

const (

	// BindSslMinVerSSLv3 captures enum value "SSLv3"
	BindSslMinVerSSLv3 string = "SSLv3"

	// BindSslMinVerTLSv10 captures enum value "TLSv1.0"
	BindSslMinVerTLSv10 string = "TLSv1.0"

	// BindSslMinVerTLSv11 captures enum value "TLSv1.1"
	BindSslMinVerTLSv11 string = "TLSv1.1"

	// BindSslMinVerTLSv12 captures enum value "TLSv1.2"
	BindSslMinVerTLSv12 string = "TLSv1.2"

	// BindSslMinVerTLSv13 captures enum value "TLSv1.3"
	BindSslMinVerTLSv13 string = "TLSv1.3"
)

// prop value enum
func (m *Bind) validateSslMinVerEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bindTypeSslMinVerPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Bind) validateSslMinVer(formats strfmt.Registry) error {

	if swag.IsZero(m.SslMinVer) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslMinVerEnum("ssl_min_ver", "body", m.SslMinVer); err != nil {
		return err
	}

	return nil
}

var bindTypeVerifyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","optional","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindTypeVerifyPropEnum = append(bindTypeVerifyPropEnum, v)
	}
}

const (

	// BindVerifyNone captures enum value "none"
	BindVerifyNone string = "none"

	// BindVerifyOptional captures enum value "optional"
	BindVerifyOptional string = "optional"

	// BindVerifyRequired captures enum value "required"
	BindVerifyRequired string = "required"
)

// prop value enum
func (m *Bind) validateVerifyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bindTypeVerifyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Bind) validateVerify(formats strfmt.Registry) error {

	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerifyEnum("verify", "body", m.Verify); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Bind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bind) UnmarshalBinary(b []byte) error {
	var res Bind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
