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

// BridgeCount is the data-structure of this message
type BridgeCount struct {
	Count uint8
}

// IBridgeCount is the corresponding interface of BridgeCount
type IBridgeCount interface {
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BridgeCount) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBridgeCount factory function for BridgeCount
func NewBridgeCount(count uint8) *BridgeCount {
	return &BridgeCount{Count: count}
}

func CastBridgeCount(structType interface{}) *BridgeCount {
	if casted, ok := structType.(BridgeCount); ok {
		return &casted
	}
	if casted, ok := structType.(*BridgeCount); ok {
		return casted
	}
	return nil
}

func (m *BridgeCount) GetTypeName() string {
	return "BridgeCount"
}

func (m *BridgeCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BridgeCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *BridgeCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BridgeCountParse(readBuffer utils.ReadBuffer) (*BridgeCount, error) {
	if pullErr := readBuffer.PullContext("BridgeCount"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("BridgeCount"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBridgeCount(count), nil
}

func (m *BridgeCount) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BridgeCount"); pushErr != nil {
		return pushErr
	}

	// Simple Field (count)
	count := uint8(m.Count)
	_countErr := writeBuffer.WriteUint8("count", 8, (count))
	if _countErr != nil {
		return errors.Wrap(_countErr, "Error serializing 'count' field")
	}

	if popErr := writeBuffer.PopContext("BridgeCount"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BridgeCount) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}