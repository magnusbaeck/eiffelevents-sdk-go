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
func (e *IssueDefinedV3) MarshalJSON() ([]byte, error) {
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
		links = make([]IDV3Link, 0)
	}
	s := struct {
		Data  *IDV3Data  `json:"data"`
		Links []IDV3Link `json:"links"`
		Meta  *IDV3Meta  `json:"meta"`
	}{
		Data:  &e.Data,
		Links: links,
		Meta:  &e.Meta,
	}
	return json.Marshal(s)
}

// String returns the JSON encoding of the event.
func (e *IssueDefinedV3) String() string {
	b, err := e.MarshalJSON()
	if err != nil {
		// This should never happen, and if it does happen it's not clear that
		// there's a reasonable way to recover. Returning an error message
		// instead of the JSON string won't end well.
		panic(err)
	}
	return string(b)
}

type IssueDefinedV3 struct {
	// Mandatory fields

	// Optional fields
	Data  IDV3Data   `json:"data,omitempty"`
	Links []IDV3Link `json:"links,omitempty"`
	Meta  IDV3Meta   `json:"meta,omitempty"`
}

type IDV3Data struct {
	// Mandatory fields
	ID      string       `json:"id"`
	Tracker string       `json:"tracker"`
	Type    IDV3DataType `json:"type"`
	URI     string       `json:"uri"`

	// Optional fields
	CustomData []IDV3DataCustomDatum `json:"customData,omitempty"`
	Title      string                `json:"title,omitempty"`
}

type IDV3DataCustomDatum struct {
	// Mandatory fields
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// Optional fields

}

type IDV3DataType string

const (
	IDV3DataType_Bug         IDV3DataType = "BUG"
	IDV3DataType_Improvement IDV3DataType = "IMPROVEMENT"
	IDV3DataType_Feature     IDV3DataType = "FEATURE"
	IDV3DataType_WorkItem    IDV3DataType = "WORK_ITEM"
	IDV3DataType_Requirement IDV3DataType = "REQUIREMENT"
	IDV3DataType_Other       IDV3DataType = "OTHER"
)

type IDV3Link struct {
	// Mandatory fields
	Target string `json:"target"`
	Type   string `json:"type"`

	// Optional fields
	DomainID string `json:"domainId,omitempty"`
}

type IDV3Meta struct {
	// Mandatory fields
	ID      string `json:"id"`
	Time    int64  `json:"time"`
	Type    string `json:"type"`
	Version string `json:"version"`

	// Optional fields
	Security IDV3MetaSecurity `json:"security,omitempty"`
	Source   IDV3MetaSource   `json:"source,omitempty"`
	Tags     []string         `json:"tags,omitempty"`
}

type IDV3MetaSecurity struct {
	// Mandatory fields
	AuthorIdentity string `json:"authorIdentity"`

	// Optional fields
	IntegrityProtection IDV3MetaSecurityIntegrityProtection  `json:"integrityProtection,omitempty"`
	SequenceProtection  []IDV3MetaSecuritySequenceProtection `json:"sequenceProtection,omitempty"`
}

type IDV3MetaSecurityIntegrityProtection struct {
	// Mandatory fields
	Alg       IDV3MetaSecurityIntegrityProtectionAlg `json:"alg"`
	Signature string                                 `json:"signature"`

	// Optional fields
	PublicKey string `json:"publicKey,omitempty"`
}

type IDV3MetaSecurityIntegrityProtectionAlg string

const (
	IDV3MetaSecurityIntegrityProtectionAlg_HS256 IDV3MetaSecurityIntegrityProtectionAlg = "HS256"
	IDV3MetaSecurityIntegrityProtectionAlg_HS384 IDV3MetaSecurityIntegrityProtectionAlg = "HS384"
	IDV3MetaSecurityIntegrityProtectionAlg_HS512 IDV3MetaSecurityIntegrityProtectionAlg = "HS512"
	IDV3MetaSecurityIntegrityProtectionAlg_RS256 IDV3MetaSecurityIntegrityProtectionAlg = "RS256"
	IDV3MetaSecurityIntegrityProtectionAlg_RS384 IDV3MetaSecurityIntegrityProtectionAlg = "RS384"
	IDV3MetaSecurityIntegrityProtectionAlg_RS512 IDV3MetaSecurityIntegrityProtectionAlg = "RS512"
	IDV3MetaSecurityIntegrityProtectionAlg_ES256 IDV3MetaSecurityIntegrityProtectionAlg = "ES256"
	IDV3MetaSecurityIntegrityProtectionAlg_ES384 IDV3MetaSecurityIntegrityProtectionAlg = "ES384"
	IDV3MetaSecurityIntegrityProtectionAlg_ES512 IDV3MetaSecurityIntegrityProtectionAlg = "ES512"
	IDV3MetaSecurityIntegrityProtectionAlg_PS256 IDV3MetaSecurityIntegrityProtectionAlg = "PS256"
	IDV3MetaSecurityIntegrityProtectionAlg_PS384 IDV3MetaSecurityIntegrityProtectionAlg = "PS384"
	IDV3MetaSecurityIntegrityProtectionAlg_PS512 IDV3MetaSecurityIntegrityProtectionAlg = "PS512"
)

type IDV3MetaSecuritySequenceProtection struct {
	// Mandatory fields
	Position     int64  `json:"position"`
	SequenceName string `json:"sequenceName"`

	// Optional fields

}

type IDV3MetaSource struct {
	// Mandatory fields

	// Optional fields
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}
