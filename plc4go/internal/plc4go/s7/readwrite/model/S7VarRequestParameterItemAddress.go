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
type S7VarRequestParameterItemAddress struct {
	*S7VarRequestParameterItem
	Address *S7Address
}

// The corresponding interface
type IS7VarRequestParameterItemAddress interface {
	// GetAddress returns Address
	GetAddress() *S7Address
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
func (m *S7VarRequestParameterItemAddress) ItemType() uint8 {
	return 0x12
}

func (m *S7VarRequestParameterItemAddress) GetItemType() uint8 {
	return 0x12
}

func (m *S7VarRequestParameterItemAddress) InitializeParent(parent *S7VarRequestParameterItem) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *S7VarRequestParameterItemAddress) GetAddress() *S7Address {
	return m.Address
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewS7VarRequestParameterItemAddress(address *S7Address) *S7VarRequestParameterItem {
	child := &S7VarRequestParameterItemAddress{
		Address:                   address,
		S7VarRequestParameterItem: NewS7VarRequestParameterItem(),
	}
	child.Child = child
	return child.S7VarRequestParameterItem
}

func CastS7VarRequestParameterItemAddress(structType interface{}) *S7VarRequestParameterItemAddress {
	castFunc := func(typ interface{}) *S7VarRequestParameterItemAddress {
		if casted, ok := typ.(S7VarRequestParameterItemAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*S7VarRequestParameterItemAddress); ok {
			return casted
		}
		if casted, ok := typ.(S7VarRequestParameterItem); ok {
			return CastS7VarRequestParameterItemAddress(casted.Child)
		}
		if casted, ok := typ.(*S7VarRequestParameterItem); ok {
			return CastS7VarRequestParameterItemAddress(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7VarRequestParameterItemAddress) GetTypeName() string {
	return "S7VarRequestParameterItemAddress"
}

func (m *S7VarRequestParameterItemAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7VarRequestParameterItemAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += m.Address.LengthInBits()

	return lengthInBits
}

func (m *S7VarRequestParameterItemAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7VarRequestParameterItemAddressParse(readBuffer utils.ReadBuffer) (*S7VarRequestParameterItem, error) {
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItemAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := readBuffer.ReadUint8("itemLength", 8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field")
	}

	// Simple Field (address)
	if pullErr := readBuffer.PullContext("address"); pullErr != nil {
		return nil, pullErr
	}
	_address, _addressErr := S7AddressParse(readBuffer)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := CastS7Address(_address)
	if closeErr := readBuffer.CloseContext("address"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItemAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7VarRequestParameterItemAddress{
		Address:                   CastS7Address(address),
		S7VarRequestParameterItem: &S7VarRequestParameterItem{},
	}
	_child.S7VarRequestParameterItem.Child = _child
	return _child.S7VarRequestParameterItem, nil
}

func (m *S7VarRequestParameterItemAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7VarRequestParameterItemAddress"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(m.Address.LengthInBytes())
		_itemLengthErr := writeBuffer.WriteUint8("itemLength", 8, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (address)
		if pushErr := writeBuffer.PushContext("address"); pushErr != nil {
			return pushErr
		}
		_addressErr := m.Address.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("address"); popErr != nil {
			return popErr
		}
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("S7VarRequestParameterItemAddress"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7VarRequestParameterItemAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
