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
type S7PayloadUserDataItemCpuFunctionReadSzlRequest struct {
	*S7PayloadUserDataItem
	SzlId    *SzlId
	SzlIndex uint16
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionReadSzlRequest interface {
	// GetSzlId returns SzlId
	GetSzlId() *SzlId
	// GetSzlIndex returns SzlIndex
	GetSzlIndex() uint16
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
func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) CpuFunctionType() uint8 {
	return 0x04
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) CpuSubfunction() uint8 {
	return 0x01
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetCpuSubfunction() uint8 {
	return 0x01
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) DataLength() uint16 {
	return 0
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetDataLength() uint16 {
	return 0
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetSzlId() *SzlId {
	return m.SzlId
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetSzlIndex() uint16 {
	return m.SzlIndex
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewS7PayloadUserDataItemCpuFunctionReadSzlRequest(szlId *SzlId, szlIndex uint16, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadUserDataItemCpuFunctionReadSzlRequest{
		SzlId:                 szlId,
		SzlIndex:              szlIndex,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Child = child
	return child.S7PayloadUserDataItem
}

func CastS7PayloadUserDataItemCpuFunctionReadSzlRequest(structType interface{}) *S7PayloadUserDataItemCpuFunctionReadSzlRequest {
	castFunc := func(typ interface{}) *S7PayloadUserDataItemCpuFunctionReadSzlRequest {
		if casted, ok := typ.(S7PayloadUserDataItemCpuFunctionReadSzlRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserDataItemCpuFunctionReadSzlRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionReadSzlRequest(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionReadSzlRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlRequest"
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (szlId)
	lengthInBits += m.SzlId.LengthInBits()

	// Simple field (szlIndex)
	lengthInBits += 16

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlRequestParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (szlId)
	if pullErr := readBuffer.PullContext("szlId"); pullErr != nil {
		return nil, pullErr
	}
	_szlId, _szlIdErr := SzlIdParse(readBuffer)
	if _szlIdErr != nil {
		return nil, errors.Wrap(_szlIdErr, "Error parsing 'szlId' field")
	}
	szlId := CastSzlId(_szlId)
	if closeErr := readBuffer.CloseContext("szlId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (szlIndex)
	_szlIndex, _szlIndexErr := readBuffer.ReadUint16("szlIndex", 16)
	if _szlIndexErr != nil {
		return nil, errors.Wrap(_szlIndexErr, "Error parsing 'szlIndex' field")
	}
	szlIndex := _szlIndex

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionReadSzlRequest{
		SzlId:                 CastSzlId(szlId),
		SzlIndex:              szlIndex,
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child.S7PayloadUserDataItem, nil
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (szlId)
		if pushErr := writeBuffer.PushContext("szlId"); pushErr != nil {
			return pushErr
		}
		_szlIdErr := m.SzlId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("szlId"); popErr != nil {
			return popErr
		}
		if _szlIdErr != nil {
			return errors.Wrap(_szlIdErr, "Error serializing 'szlId' field")
		}

		// Simple Field (szlIndex)
		szlIndex := uint16(m.SzlIndex)
		_szlIndexErr := writeBuffer.WriteUint16("szlIndex", 16, (szlIndex))
		if _szlIndexErr != nil {
			return errors.Wrap(_szlIndexErr, "Error serializing 'szlIndex' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
