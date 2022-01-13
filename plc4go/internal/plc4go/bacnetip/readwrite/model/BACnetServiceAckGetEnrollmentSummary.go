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
type BACnetServiceAckGetEnrollmentSummary struct {
	*BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckGetEnrollmentSummary interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckGetEnrollmentSummary) ServiceChoice() uint8 {
	return 0x04
}

func (m *BACnetServiceAckGetEnrollmentSummary) InitializeParent(parent *BACnetServiceAck) {}

func NewBACnetServiceAckGetEnrollmentSummary() *BACnetServiceAck {
	child := &BACnetServiceAckGetEnrollmentSummary{
		BACnetServiceAck: NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckGetEnrollmentSummary(structType interface{}) *BACnetServiceAckGetEnrollmentSummary {
	castFunc := func(typ interface{}) *BACnetServiceAckGetEnrollmentSummary {
		if casted, ok := typ.(BACnetServiceAckGetEnrollmentSummary); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckGetEnrollmentSummary); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckGetEnrollmentSummary(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckGetEnrollmentSummary(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckGetEnrollmentSummary) GetTypeName() string {
	return "BACnetServiceAckGetEnrollmentSummary"
}

func (m *BACnetServiceAckGetEnrollmentSummary) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckGetEnrollmentSummary) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckGetEnrollmentSummary) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckGetEnrollmentSummaryParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEnrollmentSummary"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEnrollmentSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckGetEnrollmentSummary{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckGetEnrollmentSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEnrollmentSummary"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEnrollmentSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
