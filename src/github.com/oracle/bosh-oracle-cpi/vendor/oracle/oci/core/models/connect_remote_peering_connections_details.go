// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectRemotePeeringConnectionsDetails Information about the other remote peering connection (RPC).
// swagger:model ConnectRemotePeeringConnectionsDetails
type ConnectRemotePeeringConnectionsDetails struct {

	// The OCID of the RPC you want to peer with.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	PeerID *string `json:"peerId"`

	// The name of the region that contains the RPC you want to peer with.
	//
	// Example: `us-ashburn-1`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	PeerRegionName *string `json:"peerRegionName"`
}

// Validate validates this connect remote peering connections details
func (m *ConnectRemotePeeringConnectionsDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePeerRegionName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectRemotePeeringConnectionsDetails) validatePeerID(formats strfmt.Registry) error {

	if err := validate.Required("peerId", "body", m.PeerID); err != nil {
		return err
	}

	if err := validate.MinLength("peerId", "body", string(*m.PeerID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("peerId", "body", string(*m.PeerID), 255); err != nil {
		return err
	}

	return nil
}

func (m *ConnectRemotePeeringConnectionsDetails) validatePeerRegionName(formats strfmt.Registry) error {

	if err := validate.Required("peerRegionName", "body", m.PeerRegionName); err != nil {
		return err
	}

	if err := validate.MinLength("peerRegionName", "body", string(*m.PeerRegionName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("peerRegionName", "body", string(*m.PeerRegionName), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectRemotePeeringConnectionsDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectRemotePeeringConnectionsDetails) UnmarshalBinary(b []byte) error {
	var res ConnectRemotePeeringConnectionsDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}