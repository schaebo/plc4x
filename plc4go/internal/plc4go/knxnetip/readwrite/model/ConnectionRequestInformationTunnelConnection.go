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
type ConnectionRequestInformationTunnelConnection struct {
	*ConnectionRequestInformation
	KnxLayer KnxLayer
}

// The corresponding interface
type IConnectionRequestInformationTunnelConnection interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ConnectionRequestInformationTunnelConnection) ConnectionType() uint8 {
	return 0x04
}

func (m *ConnectionRequestInformationTunnelConnection) InitializeParent(parent *ConnectionRequestInformation) {
}

func NewConnectionRequestInformationTunnelConnection(knxLayer KnxLayer) *ConnectionRequestInformation {
	child := &ConnectionRequestInformationTunnelConnection{
		KnxLayer:                     knxLayer,
		ConnectionRequestInformation: NewConnectionRequestInformation(),
	}
	child.Child = child
	return child.ConnectionRequestInformation
}

func CastConnectionRequestInformationTunnelConnection(structType interface{}) *ConnectionRequestInformationTunnelConnection {
	castFunc := func(typ interface{}) *ConnectionRequestInformationTunnelConnection {
		if casted, ok := typ.(ConnectionRequestInformationTunnelConnection); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionRequestInformationTunnelConnection); ok {
			return casted
		}
		if casted, ok := typ.(ConnectionRequestInformation); ok {
			return CastConnectionRequestInformationTunnelConnection(casted.Child)
		}
		if casted, ok := typ.(*ConnectionRequestInformation); ok {
			return CastConnectionRequestInformationTunnelConnection(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionRequestInformationTunnelConnection) GetTypeName() string {
	return "ConnectionRequestInformationTunnelConnection"
}

func (m *ConnectionRequestInformationTunnelConnection) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ConnectionRequestInformationTunnelConnection) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (knxLayer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionRequestInformationTunnelConnection) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ConnectionRequestInformationTunnelConnectionParse(readBuffer utils.ReadBuffer) (*ConnectionRequestInformation, error) {
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationTunnelConnection"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (knxLayer)
	if pullErr := readBuffer.PullContext("knxLayer"); pullErr != nil {
		return nil, pullErr
	}
	_knxLayer, _knxLayerErr := KnxLayerParse(readBuffer)
	if _knxLayerErr != nil {
		return nil, errors.Wrap(_knxLayerErr, "Error parsing 'knxLayer' field")
	}
	knxLayer := _knxLayer
	if closeErr := readBuffer.CloseContext("knxLayer"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationTunnelConnection"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ConnectionRequestInformationTunnelConnection{
		KnxLayer:                     knxLayer,
		ConnectionRequestInformation: &ConnectionRequestInformation{},
	}
	_child.ConnectionRequestInformation.Child = _child
	return _child.ConnectionRequestInformation, nil
}

func (m *ConnectionRequestInformationTunnelConnection) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationTunnelConnection"); pushErr != nil {
			return pushErr
		}

		// Simple Field (knxLayer)
		if pushErr := writeBuffer.PushContext("knxLayer"); pushErr != nil {
			return pushErr
		}
		_knxLayerErr := m.KnxLayer.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("knxLayer"); popErr != nil {
			return popErr
		}
		if _knxLayerErr != nil {
			return errors.Wrap(_knxLayerErr, "Error serializing 'knxLayer' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationTunnelConnection"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionRequestInformationTunnelConnection) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
