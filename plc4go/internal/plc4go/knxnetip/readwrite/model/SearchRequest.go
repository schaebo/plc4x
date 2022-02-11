/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type SearchRequest struct {
	*KnxNetIpMessage
	HpaiIDiscoveryEndpoint *HPAIDiscoveryEndpoint
}

// The corresponding interface
type ISearchRequest interface {
	// GetHpaiIDiscoveryEndpoint returns HpaiIDiscoveryEndpoint
	GetHpaiIDiscoveryEndpoint() *HPAIDiscoveryEndpoint
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SearchRequest) MsgType() uint16 {
	return 0x0201
}

func (m *SearchRequest) GetMsgType() uint16 {
	return 0x0201
}

func (m *SearchRequest) InitializeParent(parent *KnxNetIpMessage) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *SearchRequest) GetHpaiIDiscoveryEndpoint() *HPAIDiscoveryEndpoint {
	return m.HpaiIDiscoveryEndpoint
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewSearchRequest(hpaiIDiscoveryEndpoint *HPAIDiscoveryEndpoint) *KnxNetIpMessage {
	child := &SearchRequest{
		HpaiIDiscoveryEndpoint: hpaiIDiscoveryEndpoint,
		KnxNetIpMessage:        NewKnxNetIpMessage(),
	}
	child.Child = child
	return child.KnxNetIpMessage
}

func CastSearchRequest(structType interface{}) *SearchRequest {
	castFunc := func(typ interface{}) *SearchRequest {
		if casted, ok := typ.(SearchRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*SearchRequest); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastSearchRequest(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastSearchRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SearchRequest) GetTypeName() string {
	return "SearchRequest"
}

func (m *SearchRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SearchRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (hpaiIDiscoveryEndpoint)
	lengthInBits += m.HpaiIDiscoveryEndpoint.LengthInBits()

	return lengthInBits
}

func (m *SearchRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SearchRequestParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("SearchRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (hpaiIDiscoveryEndpoint)
	if pullErr := readBuffer.PullContext("hpaiIDiscoveryEndpoint"); pullErr != nil {
		return nil, pullErr
	}
	_hpaiIDiscoveryEndpoint, _hpaiIDiscoveryEndpointErr := HPAIDiscoveryEndpointParse(readBuffer)
	if _hpaiIDiscoveryEndpointErr != nil {
		return nil, errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error parsing 'hpaiIDiscoveryEndpoint' field")
	}
	hpaiIDiscoveryEndpoint := CastHPAIDiscoveryEndpoint(_hpaiIDiscoveryEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiIDiscoveryEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SearchRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SearchRequest{
		HpaiIDiscoveryEndpoint: CastHPAIDiscoveryEndpoint(hpaiIDiscoveryEndpoint),
		KnxNetIpMessage:        &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child.KnxNetIpMessage, nil
}

func (m *SearchRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hpaiIDiscoveryEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiIDiscoveryEndpoint"); pushErr != nil {
			return pushErr
		}
		_hpaiIDiscoveryEndpointErr := m.HpaiIDiscoveryEndpoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("hpaiIDiscoveryEndpoint"); popErr != nil {
			return popErr
		}
		if _hpaiIDiscoveryEndpointErr != nil {
			return errors.Wrap(_hpaiIDiscoveryEndpointErr, "Error serializing 'hpaiIDiscoveryEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("SearchRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SearchRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
