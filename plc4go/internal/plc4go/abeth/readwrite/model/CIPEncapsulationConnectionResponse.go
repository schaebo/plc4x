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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CIPEncapsulationConnectionResponse struct {
	*CIPEncapsulationPacket
}

// The corresponding interface
type ICIPEncapsulationConnectionResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CIPEncapsulationConnectionResponse) CommandType() uint16 {
	return 0x0201
}

func (m *CIPEncapsulationConnectionResponse) InitializeParent(parent *CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func NewCIPEncapsulationConnectionResponse(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *CIPEncapsulationPacket {
	child := &CIPEncapsulationConnectionResponse{
		CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	child.Child = child
	return child.CIPEncapsulationPacket
}

func CastCIPEncapsulationConnectionResponse(structType interface{}) *CIPEncapsulationConnectionResponse {
	castFunc := func(typ interface{}) *CIPEncapsulationConnectionResponse {
		if casted, ok := typ.(CIPEncapsulationConnectionResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*CIPEncapsulationConnectionResponse); ok {
			return casted
		}
		if casted, ok := typ.(CIPEncapsulationPacket); ok {
			return CastCIPEncapsulationConnectionResponse(casted.Child)
		}
		if casted, ok := typ.(*CIPEncapsulationPacket); ok {
			return CastCIPEncapsulationConnectionResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CIPEncapsulationConnectionResponse) GetTypeName() string {
	return "CIPEncapsulationConnectionResponse"
}

func (m *CIPEncapsulationConnectionResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CIPEncapsulationConnectionResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *CIPEncapsulationConnectionResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CIPEncapsulationConnectionResponseParse(readBuffer utils.ReadBuffer) (*CIPEncapsulationPacket, error) {
	if pullErr := readBuffer.PullContext("CIPEncapsulationConnectionResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationConnectionResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CIPEncapsulationConnectionResponse{
		CIPEncapsulationPacket: &CIPEncapsulationPacket{},
	}
	_child.CIPEncapsulationPacket.Child = _child
	return _child.CIPEncapsulationPacket, nil
}

func (m *CIPEncapsulationConnectionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationConnectionResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationConnectionResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CIPEncapsulationConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
