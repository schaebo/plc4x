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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type APDUSimpleAck struct {
	*APDU
	OriginalInvokeId uint8
	ServiceChoice    uint8
}

// The corresponding interface
type IAPDUSimpleAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUSimpleAck) ApduType() uint8 {
	return 0x2
}

func (m *APDUSimpleAck) InitializeParent(parent *APDU) {
}

func NewAPDUSimpleAck(originalInvokeId uint8, serviceChoice uint8) *APDU {
	child := &APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		APDU:             NewAPDU(),
	}
	child.Child = child
	return child.APDU
}

func CastAPDUSimpleAck(structType interface{}) *APDUSimpleAck {
	castFunc := func(typ interface{}) *APDUSimpleAck {
		if casted, ok := typ.(APDUSimpleAck); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUSimpleAck); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUSimpleAck(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUSimpleAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUSimpleAck) GetTypeName() string {
	return "APDUSimpleAck"
}

func (m *APDUSimpleAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUSimpleAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *APDUSimpleAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUSimpleAckParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUSimpleAck"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (serviceChoice)
	_serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}
	serviceChoice := _serviceChoice

	if closeErr := readBuffer.CloseContext("APDUSimpleAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice:    serviceChoice,
		APDU:             &APDU{},
	}
	_child.APDU.Child = _child
	return _child.APDU, nil
}

func (m *APDUSimpleAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSimpleAck"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (serviceChoice)
		serviceChoice := uint8(m.ServiceChoice)
		_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))
		if _serviceChoiceErr != nil {
			return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
		}

		if popErr := writeBuffer.PopContext("APDUSimpleAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUSimpleAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
