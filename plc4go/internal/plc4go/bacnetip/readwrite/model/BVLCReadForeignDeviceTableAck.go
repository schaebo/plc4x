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
type BVLCReadForeignDeviceTableAck struct {
	*BVLC
}

// The corresponding interface
type IBVLCReadForeignDeviceTableAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCReadForeignDeviceTableAck) BvlcFunction() uint8 {
	return 0x07
}

func (m *BVLCReadForeignDeviceTableAck) InitializeParent(parent *BVLC, bvlcPayloadLength uint16) {}

func NewBVLCReadForeignDeviceTableAck(bvlcPayloadLength uint16) *BVLC {
	child := &BVLCReadForeignDeviceTableAck{
		BVLC: NewBVLC(bvlcPayloadLength),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCReadForeignDeviceTableAck(structType interface{}) *BVLCReadForeignDeviceTableAck {
	castFunc := func(typ interface{}) *BVLCReadForeignDeviceTableAck {
		if casted, ok := typ.(BVLCReadForeignDeviceTableAck); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCReadForeignDeviceTableAck); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCReadForeignDeviceTableAck(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCReadForeignDeviceTableAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCReadForeignDeviceTableAck) GetTypeName() string {
	return "BVLCReadForeignDeviceTableAck"
}

func (m *BVLCReadForeignDeviceTableAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCReadForeignDeviceTableAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BVLCReadForeignDeviceTableAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCReadForeignDeviceTableAckParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCReadForeignDeviceTableAck"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BVLCReadForeignDeviceTableAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCReadForeignDeviceTableAck{
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCReadForeignDeviceTableAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadForeignDeviceTableAck"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCReadForeignDeviceTableAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCReadForeignDeviceTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
