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
type ModbusPDUReadDeviceIdentificationResponse struct {
	*ModbusPDU
}

// The corresponding interface
type IModbusPDUReadDeviceIdentificationResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadDeviceIdentificationResponse) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) FunctionFlag() uint8 {
	return 0x2B
}

func (m *ModbusPDUReadDeviceIdentificationResponse) Response() bool {
	return bool(true)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) InitializeParent(parent *ModbusPDU) {}

func NewModbusPDUReadDeviceIdentificationResponse() *ModbusPDU {
	child := &ModbusPDUReadDeviceIdentificationResponse{
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadDeviceIdentificationResponse(structType interface{}) *ModbusPDUReadDeviceIdentificationResponse {
	castFunc := func(typ interface{}) *ModbusPDUReadDeviceIdentificationResponse {
		if casted, ok := typ.(ModbusPDUReadDeviceIdentificationResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadDeviceIdentificationResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadDeviceIdentificationResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadDeviceIdentificationResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationResponse"
}

func (m *ModbusPDUReadDeviceIdentificationResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUReadDeviceIdentificationResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadDeviceIdentificationResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadDeviceIdentificationResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDeviceIdentificationResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadDeviceIdentificationResponse{
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadDeviceIdentificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDeviceIdentificationResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDeviceIdentificationResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadDeviceIdentificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
