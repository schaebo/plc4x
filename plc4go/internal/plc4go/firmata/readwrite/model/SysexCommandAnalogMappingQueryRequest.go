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
type SysexCommandAnalogMappingQueryRequest struct {
	*SysexCommand
}

// The corresponding interface
type ISysexCommandAnalogMappingQueryRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandAnalogMappingQueryRequest) CommandType() uint8 {
	return 0x69
}

func (m *SysexCommandAnalogMappingQueryRequest) Response() bool {
	return bool(false)
}

func (m *SysexCommandAnalogMappingQueryRequest) InitializeParent(parent *SysexCommand) {}

func NewSysexCommandAnalogMappingQueryRequest() *SysexCommand {
	child := &SysexCommandAnalogMappingQueryRequest{
		SysexCommand: NewSysexCommand(),
	}
	child.Child = child
	return child.SysexCommand
}

func CastSysexCommandAnalogMappingQueryRequest(structType interface{}) *SysexCommandAnalogMappingQueryRequest {
	castFunc := func(typ interface{}) *SysexCommandAnalogMappingQueryRequest {
		if casted, ok := typ.(SysexCommandAnalogMappingQueryRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandAnalogMappingQueryRequest); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandAnalogMappingQueryRequest(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandAnalogMappingQueryRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandAnalogMappingQueryRequest) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryRequest"
}

func (m *SysexCommandAnalogMappingQueryRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandAnalogMappingQueryRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *SysexCommandAnalogMappingQueryRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandAnalogMappingQueryRequestParse(readBuffer utils.ReadBuffer, response bool) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandAnalogMappingQueryRequest{
		SysexCommand: &SysexCommand{},
	}
	_child.SysexCommand.Child = _child
	return _child.SysexCommand, nil
}

func (m *SysexCommandAnalogMappingQueryRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandAnalogMappingQueryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
