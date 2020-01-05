/*
 * Rumble API
 *
 * Rumble Network Discovery API
 *
 * API version: 1.0.0
 * Contact: support@rumble.run
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// AgentSiteId struct for AgentSiteId
type AgentSiteId struct {
	SiteId string `json:"site_id"`
}

// GetSiteId returns the SiteId field value
func (o *AgentSiteId) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// SetSiteId sets field value
func (o *AgentSiteId) SetSiteId(v string) {
	o.SiteId = v
}

type NullableAgentSiteId struct {
	Value        AgentSiteId
	ExplicitNull bool
}

func (v NullableAgentSiteId) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAgentSiteId) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
