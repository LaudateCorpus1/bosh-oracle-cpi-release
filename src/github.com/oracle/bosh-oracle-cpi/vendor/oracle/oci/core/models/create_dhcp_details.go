package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateDhcpDetails create dhcp details
// swagger:model CreateDhcpDetails
type CreateDhcpDetails struct {

	// The OCID of the compartment to contain the set of DHCP options.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	Options []DhcpOption `json:"options"`

	// The OCID of the VCN the set of DHCP options belongs to.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CreateDhcpDetails) UnmarshalJSON(raw []byte) error {
	var data struct {
		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		Options json.RawMessage `json:"options"`

		VcnID *string `json:"vcnId"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var options []DhcpOption

	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	untypedObj := make(map[string]interface{})
	if err := dec.Decode(&untypedObj); err != nil {
		return err
	}
	if untypedOptions, ok := untypedObj["options"]; ok {
		if slcOptions, ok := untypedOptions.([]interface{}); ok {
			for _, slcEl := range slcOptions {
				slcJSON, _ := json.Marshal(slcEl)
				slcObj, err := UnmarshalDhcpOption(bytes.NewBuffer(slcJSON), runtime.JSONConsumer())
				if err != nil && err != io.EOF {
					return err
				}
				options = append(options, slcObj)
			}
		}
	}

	var result CreateDhcpDetails
	result.CompartmentID = data.CompartmentID
	result.DisplayName = data.DisplayName
	result.Options = options
	result.VcnID = data.VcnID
	*m = result
	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CreateDhcpDetails) MarshalJSON() ([]byte, error) {
	var b1, b2 []byte
	var err error
	b1, err = json.Marshal(struct {
		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		VcnID *string `json:"vcnId"`
	}{
		CompartmentID: m.CompartmentID,
		DisplayName:   m.DisplayName,
		VcnID:         m.VcnID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Options []DhcpOption `json:"options"`
	}{
		Options: m.Options,
	})
	if err != nil {
		return nil, err
	}
	return swag.ConcatJSON(b1, b2), nil
}

// Validate validates this create dhcp details
func (m *CreateDhcpDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDhcpDetails) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateDhcpDetails) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateDhcpDetails) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("options", "body", m.Options); err != nil {
		return err
	}

	if err := validate.UniqueItems("options", "body", m.Options); err != nil {
		return err
	}

	for i := 0; i < len(m.Options); i++ {

		if err := m.Options[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateDhcpDetails) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}