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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BVLC_BACNETTYPE uint8 = 0x81

// The data-structure of this message
type BVLC struct {
	BvlcPayloadLength uint16
	Child             IBVLCChild
}

// The corresponding interface
type IBVLC interface {
	// BvlcFunction returns BvlcFunction
	BvlcFunction() uint8
	// GetBvlcPayloadLength returns BvlcPayloadLength
	GetBvlcPayloadLength() uint16
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBVLCParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBVLCChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BVLC, bvlcPayloadLength uint16)
	GetTypeName() string
	IBVLC
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BVLC) GetBvlcPayloadLength() uint16 {
	// TODO: calculation should happen here instead accessing the stored field
	return m.BvlcPayloadLength
}

func NewBVLC(bvlcPayloadLength uint16) *BVLC {
	return &BVLC{BvlcPayloadLength: bvlcPayloadLength}
}

func CastBVLC(structType interface{}) *BVLC {
	castFunc := func(typ interface{}) *BVLC {
		if casted, ok := typ.(BVLC); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLC); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLC) GetTypeName() string {
	return "BVLC"
}

func (m *BVLC) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLC) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BVLC) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetType)
	lengthInBits += 8
	// Discriminator Field (bvlcFunction)
	lengthInBits += 8

	// Implicit Field (bvlcLength)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BVLC) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLC"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (bacnetType)
	bacnetType, _bacnetTypeErr := readBuffer.ReadUint8("bacnetType", 8)
	if _bacnetTypeErr != nil {
		return nil, errors.Wrap(_bacnetTypeErr, "Error parsing 'bacnetType' field")
	}
	if bacnetType != BVLC_BACNETTYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BVLC_BACNETTYPE) + " but got " + fmt.Sprintf("%d", bacnetType))
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction, _bvlcFunctionErr := readBuffer.ReadUint8("bvlcFunction", 8)
	if _bvlcFunctionErr != nil {
		return nil, errors.Wrap(_bvlcFunctionErr, "Error parsing 'bvlcFunction' field")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	bvlcLength, _bvlcLengthErr := readBuffer.ReadUint16("bvlcLength", 16)
	_ = bvlcLength
	if _bvlcLengthErr != nil {
		return nil, errors.Wrap(_bvlcLengthErr, "Error parsing 'bvlcLength' field")
	}

	// Virtual field
	_bvlcPayloadLength := uint16(bvlcLength) - uint16(uint16(4))
	bvlcPayloadLength := uint16(_bvlcPayloadLength)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BVLC
	var typeSwitchError error
	switch {
	case bvlcFunction == 0x00: // BVLCResult
		_parent, typeSwitchError = BVLCResultParse(readBuffer)
	case bvlcFunction == 0x01: // BVLCWriteBroadcastDistributionTable
		_parent, typeSwitchError = BVLCWriteBroadcastDistributionTableParse(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x02: // BVLCReadBroadcastDistributionTable
		_parent, typeSwitchError = BVLCReadBroadcastDistributionTableParse(readBuffer)
	case bvlcFunction == 0x03: // BVLCReadBroadcastDistributionTableAck
		_parent, typeSwitchError = BVLCReadBroadcastDistributionTableAckParse(readBuffer)
	case bvlcFunction == 0x04: // BVLCForwardedNPDU
		_parent, typeSwitchError = BVLCForwardedNPDUParse(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x05: // BVLCRegisterForeignDevice
		_parent, typeSwitchError = BVLCRegisterForeignDeviceParse(readBuffer)
	case bvlcFunction == 0x06: // BVLCReadForeignDeviceTable
		_parent, typeSwitchError = BVLCReadForeignDeviceTableParse(readBuffer)
	case bvlcFunction == 0x07: // BVLCReadForeignDeviceTableAck
		_parent, typeSwitchError = BVLCReadForeignDeviceTableAckParse(readBuffer)
	case bvlcFunction == 0x08: // BVLCDeleteForeignDeviceTableEntry
		_parent, typeSwitchError = BVLCDeleteForeignDeviceTableEntryParse(readBuffer)
	case bvlcFunction == 0x09: // BVLCDistributeBroadcastToNetwork
		_parent, typeSwitchError = BVLCDistributeBroadcastToNetworkParse(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0A: // BVLCOriginalUnicastNPDU
		_parent, typeSwitchError = BVLCOriginalUnicastNPDUParse(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0B: // BVLCOriginalBroadcastNPDU
		_parent, typeSwitchError = BVLCOriginalBroadcastNPDUParse(readBuffer, bvlcPayloadLength)
	case bvlcFunction == 0x0C: // BVLCSecureBVLL
		_parent, typeSwitchError = BVLCSecureBVLLParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BVLC"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, bvlcPayloadLength)
	return _parent, nil
}

func (m *BVLC) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BVLC) SerializeParent(writeBuffer utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BVLC"); pushErr != nil {
		return pushErr
	}

	// Const Field (bacnetType)
	_bacnetTypeErr := writeBuffer.WriteUint8("bacnetType", 8, 0x81)
	if _bacnetTypeErr != nil {
		return errors.Wrap(_bacnetTypeErr, "Error serializing 'bacnetType' field")
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction := uint8(child.BvlcFunction())
	_bvlcFunctionErr := writeBuffer.WriteUint8("bvlcFunction", 8, (bvlcFunction))

	if _bvlcFunctionErr != nil {
		return errors.Wrap(_bvlcFunctionErr, "Error serializing 'bvlcFunction' field")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	bvlcLength := uint16(uint16(m.LengthInBytes()))
	_bvlcLengthErr := writeBuffer.WriteUint16("bvlcLength", 16, (bvlcLength))
	if _bvlcLengthErr != nil {
		return errors.Wrap(_bvlcLengthErr, "Error serializing 'bvlcLength' field")
	}
	// Virtual field
	if _bvlcPayloadLengthErr := writeBuffer.WriteVirtual("bvlcPayloadLength", m.BvlcPayloadLength); _bvlcPayloadLengthErr != nil {
		return errors.Wrap(_bvlcPayloadLengthErr, "Error serializing 'bvlcPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BVLC"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BVLC) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
