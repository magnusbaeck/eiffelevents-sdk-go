// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
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

// THIS FILE IS AUTOMATICALLY GENERATED AND MUST NOT BE EDITED BY HAND.

package eiffelevents

import (
	"github.com/clarketm/json"
)

// MarshalJSON returns the JSON encoding of the event.
func (e *ArtifactPublishedV3) MarshalJSON() ([]byte, error) {
	// The standard encoding/json package doesn't honor omitempty for
	// non-pointer structs (it doesn't recurse into values, only examines
	// the immediate value). This is a not terribly elegant way of making
	// sure that this struct is marshaled by github.com/clarketm/json
	// without the infinite loop we'd get by just passing the struct to
	// github.com/clarketm/json.Marshal.
	//
	// Make sure the links slice is non-null so that non-initialized slices
	// get serialized as "[]" instead of "null".
	links := e.Links
	if links == nil {
		links = make([]ArtPV3Link, 0)
	}
	s := struct {
		Data  *ArtPV3Data  `json:"data"`
		Links []ArtPV3Link `json:"links"`
		Meta  *ArtPV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *ArtifactPublishedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type ArtifactPublishedV3 struct {
	// Mandatory fields
	Data  ArtPV3Data   `json:"data"`
	Links []ArtPV3Link `json:"links"`
	Meta  ArtPV3Meta   `json:"meta"`

	// Optional fields

}

type ArtPV3Data struct {
	// Mandatory fields
	Locations []ArtPV3DataLocation `json:"locations"`

	// Optional fields
	CustomData []ArtPV3DataCustomDatum `json:"customData,omitempty"`
}

type ArtPV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type ArtPV3DataLocation struct {
	// Mandatory fields
	Type ArtPV3DataLocationType `json:"type"`
	URI  string                 `json:"uri"`

	// Optional fields
	Name string `json:"name,omitempty"`
}

type ArtPV3DataLocationType string

const (
	ArtPV3DataLocationType_Artifactory ArtPV3DataLocationType = "ARTIFACTORY"
	ArtPV3DataLocationType_Nexus       ArtPV3DataLocationType = "NEXUS"
	ArtPV3DataLocationType_Plain       ArtPV3DataLocationType = "PLAIN"
	ArtPV3DataLocationType_Other       ArtPV3DataLocationType = "OTHER"
)

type ArtPV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type ArtPV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security ArtPV3MetaSecurity `json:"security,omitempty"`
	Source   ArtPV3MetaSource   `json:"source,omitempty"`
	Tags     []string           `json:"tags,omitempty"`
}

type ArtPV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection ArtPV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []ArtPV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type ArtPV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       ArtPV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                   `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type ArtPV3MetaSecurityIntegrityProtectionAlg string

const (
	ArtPV3MetaSecurityIntegrityProtectionAlg_HS256 ArtPV3MetaSecurityIntegrityProtectionAlg = "HS256"
	ArtPV3MetaSecurityIntegrityProtectionAlg_HS384 ArtPV3MetaSecurityIntegrityProtectionAlg = "HS384"
	ArtPV3MetaSecurityIntegrityProtectionAlg_HS512 ArtPV3MetaSecurityIntegrityProtectionAlg = "HS512"
	ArtPV3MetaSecurityIntegrityProtectionAlg_RS256 ArtPV3MetaSecurityIntegrityProtectionAlg = "RS256"
	ArtPV3MetaSecurityIntegrityProtectionAlg_RS384 ArtPV3MetaSecurityIntegrityProtectionAlg = "RS384"
	ArtPV3MetaSecurityIntegrityProtectionAlg_RS512 ArtPV3MetaSecurityIntegrityProtectionAlg = "RS512"
	ArtPV3MetaSecurityIntegrityProtectionAlg_ES256 ArtPV3MetaSecurityIntegrityProtectionAlg = "ES256"
	ArtPV3MetaSecurityIntegrityProtectionAlg_ES384 ArtPV3MetaSecurityIntegrityProtectionAlg = "ES384"
	ArtPV3MetaSecurityIntegrityProtectionAlg_ES512 ArtPV3MetaSecurityIntegrityProtectionAlg = "ES512"
	ArtPV3MetaSecurityIntegrityProtectionAlg_PS256 ArtPV3MetaSecurityIntegrityProtectionAlg = "PS256"
	ArtPV3MetaSecurityIntegrityProtectionAlg_PS384 ArtPV3MetaSecurityIntegrityProtectionAlg = "PS384"
	ArtPV3MetaSecurityIntegrityProtectionAlg_PS512 ArtPV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type ArtPV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type ArtPV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
