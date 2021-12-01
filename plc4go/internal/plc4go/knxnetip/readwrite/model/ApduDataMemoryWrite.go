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
type ApduDataMemoryWrite struct {
	*ApduData
}

// The corresponding interface
type IApduDataMemoryWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryWrite) ApciType() uint8 {
	return 0xA
}

func (m *ApduDataMemoryWrite) InitializeParent(parent *ApduData) {
}

func NewApduDataMemoryWrite() *ApduData {
	child := &ApduDataMemoryWrite{
		ApduData: NewApduData(),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataMemoryWrite(structType interface{}) *ApduDataMemoryWrite {
	castFunc := func(typ interface{}) *ApduDataMemoryWrite {
		if casted, ok := typ.(ApduDataMemoryWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataMemoryWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataMemoryWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataMemoryWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataMemoryWrite) GetTypeName() string {
	return "ApduDataMemoryWrite"
}

func (m *ApduDataMemoryWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataMemoryWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataMemoryWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataMemoryWriteParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataMemoryWrite"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataMemoryWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataMemoryWrite{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataMemoryWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataMemoryWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
