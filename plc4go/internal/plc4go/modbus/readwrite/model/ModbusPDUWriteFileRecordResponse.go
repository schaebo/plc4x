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
type ModbusPDUWriteFileRecordResponse struct {
	*ModbusPDU
	Items []*ModbusPDUWriteFileRecordResponseItem
}

// The corresponding interface
type IModbusPDUWriteFileRecordResponse interface {
	// GetItems returns Items
	GetItems() []*ModbusPDUWriteFileRecordResponseItem
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
func (m *ModbusPDUWriteFileRecordResponse) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteFileRecordResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUWriteFileRecordResponse) FunctionFlag() uint8 {
	return 0x15
}

func (m *ModbusPDUWriteFileRecordResponse) GetFunctionFlag() uint8 {
	return 0x15
}

func (m *ModbusPDUWriteFileRecordResponse) Response() bool {
	return bool(true)
}

func (m *ModbusPDUWriteFileRecordResponse) GetResponse() bool {
	return bool(true)
}

func (m *ModbusPDUWriteFileRecordResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteFileRecordResponse) GetItems() []*ModbusPDUWriteFileRecordResponseItem {
	return m.Items
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewModbusPDUWriteFileRecordResponse(items []*ModbusPDUWriteFileRecordResponseItem) *ModbusPDU {
	child := &ModbusPDUWriteFileRecordResponse{
		Items:     items,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUWriteFileRecordResponse(structType interface{}) *ModbusPDUWriteFileRecordResponse {
	castFunc := func(typ interface{}) *ModbusPDUWriteFileRecordResponse {
		if casted, ok := typ.(ModbusPDUWriteFileRecordResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteFileRecordResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteFileRecordResponse) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponse"
}

func (m *ModbusPDUWriteFileRecordResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteFileRecordResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteFileRecordResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	items := make([]*ModbusPDUWriteFileRecordResponseItem, 0)
	{
		_itemsLength := byteCount
		_itemsEndPos := readBuffer.GetPos() + uint16(_itemsLength)
		for readBuffer.GetPos() < _itemsEndPos {
			_item, _err := ModbusPDUWriteFileRecordResponseItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items = append(items, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteFileRecordResponse{
		Items:     items,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUWriteFileRecordResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	itemsArraySizeInBytes := func(items []*ModbusPDUWriteFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.LengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.Items)))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUWriteFileRecordResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
