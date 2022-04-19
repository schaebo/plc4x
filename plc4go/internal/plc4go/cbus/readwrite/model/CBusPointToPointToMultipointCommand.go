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

// CBusPointToPointToMultipointCommand is the data-structure of this message
type CBusPointToPointToMultipointCommand struct {
	BridgeAddress     *BridgeAddress
	NetworkRoute      *NetworkRoute
	PeekedApplication byte

	// Arguments.
	Srchk bool
	Child ICBusPointToPointToMultipointCommandChild
}

// ICBusPointToPointToMultipointCommand is the corresponding interface of CBusPointToPointToMultipointCommand
type ICBusPointToPointToMultipointCommand interface {
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() *BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() *NetworkRoute
	// GetPeekedApplication returns PeekedApplication (property field)
	GetPeekedApplication() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICBusPointToPointToMultipointCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICBusPointToPointToMultipointCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICBusPointToPointToMultipointCommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CBusPointToPointToMultipointCommand, bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte)
	GetParent() *CBusPointToPointToMultipointCommand

	GetTypeName() string
	ICBusPointToPointToMultipointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CBusPointToPointToMultipointCommand) GetBridgeAddress() *BridgeAddress {
	return m.BridgeAddress
}

func (m *CBusPointToPointToMultipointCommand) GetNetworkRoute() *NetworkRoute {
	return m.NetworkRoute
}

func (m *CBusPointToPointToMultipointCommand) GetPeekedApplication() byte {
	return m.PeekedApplication
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointToMultipointCommand factory function for CBusPointToPointToMultipointCommand
func NewCBusPointToPointToMultipointCommand(bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte, srchk bool) *CBusPointToPointToMultipointCommand {
	return &CBusPointToPointToMultipointCommand{BridgeAddress: bridgeAddress, NetworkRoute: networkRoute, PeekedApplication: peekedApplication, Srchk: srchk}
}

func CastCBusPointToPointToMultipointCommand(structType interface{}) *CBusPointToPointToMultipointCommand {
	if casted, ok := structType.(CBusPointToPointToMultipointCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultipointCommand); ok {
		return casted
	}
	if casted, ok := structType.(ICBusPointToPointToMultipointCommandChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *CBusPointToPointToMultipointCommand) GetTypeName() string {
	return "CBusPointToPointToMultipointCommand"
}

func (m *CBusPointToPointToMultipointCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusPointToPointToMultipointCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CBusPointToPointToMultipointCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits()

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits()

	return lengthInBits
}

func (m *CBusPointToPointToMultipointCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToPointToMultipointCommandParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointToMultipointCommand, error) {
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultipointCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (bridgeAddress)
	if pullErr := readBuffer.PullContext("bridgeAddress"); pullErr != nil {
		return nil, pullErr
	}
	_bridgeAddress, _bridgeAddressErr := BridgeAddressParse(readBuffer)
	if _bridgeAddressErr != nil {
		return nil, errors.Wrap(_bridgeAddressErr, "Error parsing 'bridgeAddress' field")
	}
	bridgeAddress := CastBridgeAddress(_bridgeAddress)
	if closeErr := readBuffer.CloseContext("bridgeAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (networkRoute)
	if pullErr := readBuffer.PullContext("networkRoute"); pullErr != nil {
		return nil, pullErr
	}
	_networkRoute, _networkRouteErr := NetworkRouteParse(readBuffer)
	if _networkRouteErr != nil {
		return nil, errors.Wrap(_networkRouteErr, "Error parsing 'networkRoute' field")
	}
	networkRoute := CastNetworkRoute(_networkRoute)
	if closeErr := readBuffer.CloseContext("networkRoute"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedApplication)
	currentPos = readBuffer.GetPos()
	peekedApplication, _err := readBuffer.ReadByte("peekedApplication")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekedApplication' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CBusPointToPointToMultipointCommandChild interface {
		InitializeParent(*CBusPointToPointToMultipointCommand, *BridgeAddress, *NetworkRoute, byte)
		GetParent() *CBusPointToPointToMultipointCommand
	}
	var _child CBusPointToPointToMultipointCommandChild
	var typeSwitchError error
	switch {
	case peekedApplication == 0xFF: // CBusCommandPointToPointToMultiPointStatus
		_child, typeSwitchError = CBusCommandPointToPointToMultiPointStatusParse(readBuffer, srchk)
	case true: // CBusCommandPointToPointToMultiPointNormal
		_child, typeSwitchError = CBusCommandPointToPointToMultiPointNormalParse(readBuffer, srchk)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultipointCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), bridgeAddress, networkRoute, peekedApplication)
	return _child.GetParent(), nil
}

func (m *CBusPointToPointToMultipointCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CBusPointToPointToMultipointCommand) SerializeParent(writeBuffer utils.WriteBuffer, child ICBusPointToPointToMultipointCommand, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("CBusPointToPointToMultipointCommand"); pushErr != nil {
		return pushErr
	}

	// Simple Field (bridgeAddress)
	if pushErr := writeBuffer.PushContext("bridgeAddress"); pushErr != nil {
		return pushErr
	}
	_bridgeAddressErr := m.BridgeAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("bridgeAddress"); popErr != nil {
		return popErr
	}
	if _bridgeAddressErr != nil {
		return errors.Wrap(_bridgeAddressErr, "Error serializing 'bridgeAddress' field")
	}

	// Simple Field (networkRoute)
	if pushErr := writeBuffer.PushContext("networkRoute"); pushErr != nil {
		return pushErr
	}
	_networkRouteErr := m.NetworkRoute.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("networkRoute"); popErr != nil {
		return popErr
	}
	if _networkRouteErr != nil {
		return errors.Wrap(_networkRouteErr, "Error serializing 'networkRoute' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointToMultipointCommand"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CBusPointToPointToMultipointCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}