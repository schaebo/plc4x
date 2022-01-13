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
type BACnetServiceAckReadRange struct {
	*BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckReadRange interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckReadRange) ServiceChoice() uint8 {
	return 0x1A
}

func (m *BACnetServiceAckReadRange) InitializeParent(parent *BACnetServiceAck) {}

func NewBACnetServiceAckReadRange() *BACnetServiceAck {
	child := &BACnetServiceAckReadRange{
		BACnetServiceAck: NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckReadRange(structType interface{}) *BACnetServiceAckReadRange {
	castFunc := func(typ interface{}) *BACnetServiceAckReadRange {
		if casted, ok := typ.(BACnetServiceAckReadRange); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckReadRange); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckReadRange(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckReadRange(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckReadRange) GetTypeName() string {
	return "BACnetServiceAckReadRange"
}

func (m *BACnetServiceAckReadRange) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadRange) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckReadRange) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckReadRangeParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadRange"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadRange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadRange{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckReadRange) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadRange"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadRange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
