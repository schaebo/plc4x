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
type ModbusPDUReadExceptionStatusRequest struct {
	*ModbusPDU
}

// The corresponding interface
type IModbusPDUReadExceptionStatusRequest interface {
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
func (m *ModbusPDUReadExceptionStatusRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) FunctionFlag() uint8 {
	return 0x07
}

func (m *ModbusPDUReadExceptionStatusRequest) GetFunctionFlag() uint8 {
	return 0x07
}

func (m *ModbusPDUReadExceptionStatusRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) GetResponse() bool {
	return bool(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewModbusPDUReadExceptionStatusRequest() *ModbusPDU {
	child := &ModbusPDUReadExceptionStatusRequest{
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadExceptionStatusRequest(structType interface{}) *ModbusPDUReadExceptionStatusRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadExceptionStatusRequest {
		if casted, ok := typ.(ModbusPDUReadExceptionStatusRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadExceptionStatusRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadExceptionStatusRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadExceptionStatusRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadExceptionStatusRequest) GetTypeName() string {
	return "ModbusPDUReadExceptionStatusRequest"
}

func (m *ModbusPDUReadExceptionStatusRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadExceptionStatusRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUReadExceptionStatusRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadExceptionStatusRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadExceptionStatusRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadExceptionStatusRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadExceptionStatusRequest{
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadExceptionStatusRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadExceptionStatusRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadExceptionStatusRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadExceptionStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
