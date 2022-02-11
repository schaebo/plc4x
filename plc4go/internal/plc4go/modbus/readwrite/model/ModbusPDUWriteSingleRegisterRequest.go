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
type ModbusPDUWriteSingleRegisterRequest struct {
	*ModbusPDU
	Address uint16
	Value   uint16
}

// The corresponding interface
type IModbusPDUWriteSingleRegisterRequest interface {
	// GetAddress returns Address
	GetAddress() uint16
	// GetValue returns Value
	GetValue() uint16
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
func (m *ModbusPDUWriteSingleRegisterRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleRegisterRequest) FunctionFlag() uint8 {
	return 0x06
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetFunctionFlag() uint8 {
	return 0x06
}

func (m *ModbusPDUWriteSingleRegisterRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetResponse() bool {
	return bool(false)
}

func (m *ModbusPDUWriteSingleRegisterRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteSingleRegisterRequest) GetAddress() uint16 {
	return m.Address
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetValue() uint16 {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewModbusPDUWriteSingleRegisterRequest(address uint16, value uint16) *ModbusPDU {
	child := &ModbusPDUWriteSingleRegisterRequest{
		Address:   address,
		Value:     value,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUWriteSingleRegisterRequest(structType interface{}) *ModbusPDUWriteSingleRegisterRequest {
	castFunc := func(typ interface{}) *ModbusPDUWriteSingleRegisterRequest {
		if casted, ok := typ.(ModbusPDUWriteSingleRegisterRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteSingleRegisterRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteSingleRegisterRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteSingleRegisterRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteSingleRegisterRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleRegisterRequest"
}

func (m *ModbusPDUWriteSingleRegisterRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteSingleRegisterRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUWriteSingleRegisterRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteSingleRegisterRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleRegisterRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleRegisterRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteSingleRegisterRequest{
		Address:   address,
		Value:     value,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUWriteSingleRegisterRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleRegisterRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Simple Field (value)
		value := uint16(m.Value)
		_valueErr := writeBuffer.WriteUint16("value", 16, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleRegisterRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteSingleRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
