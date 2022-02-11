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
type AlarmMessageAckType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []*AlarmMessageObjectAckType
}

// The corresponding interface
type IAlarmMessageAckType interface {
	// GetFunctionId returns FunctionId
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects
	GetMessageObjects() []*AlarmMessageObjectAckType
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *AlarmMessageAckType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *AlarmMessageAckType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *AlarmMessageAckType) GetMessageObjects() []*AlarmMessageObjectAckType {
	return m.MessageObjects
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewAlarmMessageAckType(functionId uint8, numberOfObjects uint8, messageObjects []*AlarmMessageObjectAckType) *AlarmMessageAckType {
	return &AlarmMessageAckType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

func CastAlarmMessageAckType(structType interface{}) *AlarmMessageAckType {
	castFunc := func(typ interface{}) *AlarmMessageAckType {
		if casted, ok := typ.(AlarmMessageAckType); ok {
			return &casted
		}
		if casted, ok := typ.(*AlarmMessageAckType); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AlarmMessageAckType) GetTypeName() string {
	return "AlarmMessageAckType"
}

func (m *AlarmMessageAckType) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AlarmMessageAckType) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AlarmMessageAckType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AlarmMessageAckTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageAckType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageAckType"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}
	numberOfObjects := _numberOfObjects

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	messageObjects := make([]*AlarmMessageObjectAckType, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := AlarmMessageObjectAckTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
			}
			messageObjects[curItem] = CastAlarmMessageObjectAckType(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageAckType(functionId, numberOfObjects, messageObjects), nil
}

func (m *AlarmMessageAckType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageAckType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (functionId)
	functionId := uint8(m.FunctionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.NumberOfObjects)
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if m.MessageObjects != nil {
		if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.MessageObjects {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageAckType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
