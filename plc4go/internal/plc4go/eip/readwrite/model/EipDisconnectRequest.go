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
type EipDisconnectRequest struct {
	*EipPacket
}

// The corresponding interface
type IEipDisconnectRequest interface {
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
func (m *EipDisconnectRequest) Command() uint16 {
	return 0x0066
}

func (m *EipDisconnectRequest) GetCommand() uint16 {
	return 0x0066
}

func (m *EipDisconnectRequest) InitializeParent(parent *EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.EipPacket.SessionHandle = sessionHandle
	m.EipPacket.Status = status
	m.EipPacket.SenderContext = senderContext
	m.EipPacket.Options = options
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewEipDisconnectRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *EipPacket {
	child := &EipDisconnectRequest{
		EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	child.Child = child
	return child.EipPacket
}

func CastEipDisconnectRequest(structType interface{}) *EipDisconnectRequest {
	castFunc := func(typ interface{}) *EipDisconnectRequest {
		if casted, ok := typ.(EipDisconnectRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*EipDisconnectRequest); ok {
			return casted
		}
		if casted, ok := typ.(EipPacket); ok {
			return CastEipDisconnectRequest(casted.Child)
		}
		if casted, ok := typ.(*EipPacket); ok {
			return CastEipDisconnectRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *EipDisconnectRequest) GetTypeName() string {
	return "EipDisconnectRequest"
}

func (m *EipDisconnectRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *EipDisconnectRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *EipDisconnectRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func EipDisconnectRequestParse(readBuffer utils.ReadBuffer) (*EipPacket, error) {
	if pullErr := readBuffer.PullContext("EipDisconnectRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("EipDisconnectRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &EipDisconnectRequest{
		EipPacket: &EipPacket{},
	}
	_child.EipPacket.Child = _child
	return _child.EipPacket, nil
}

func (m *EipDisconnectRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipDisconnectRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("EipDisconnectRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *EipDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
