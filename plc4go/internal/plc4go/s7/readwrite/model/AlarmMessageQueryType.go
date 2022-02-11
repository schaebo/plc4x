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
const AlarmMessageQueryType_DATALENGTH uint16 = 0xFFFF

// The data-structure of this message
type AlarmMessageQueryType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	ReturnCode      DataTransportErrorCode
	TransportSize   DataTransportSize
	MessageObjects  []*AlarmMessageObjectQueryType
}

// The corresponding interface
type IAlarmMessageQueryType interface {
	// GetFunctionId returns FunctionId
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects
	GetNumberOfObjects() uint8
	// GetReturnCode returns ReturnCode
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize
	GetTransportSize() DataTransportSize
	// GetMessageObjects returns MessageObjects
	GetMessageObjects() []*AlarmMessageObjectQueryType
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
func (m *AlarmMessageQueryType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *AlarmMessageQueryType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *AlarmMessageQueryType) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *AlarmMessageQueryType) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

func (m *AlarmMessageQueryType) GetMessageObjects() []*AlarmMessageObjectQueryType {
	return m.MessageObjects
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewAlarmMessageQueryType(functionId uint8, numberOfObjects uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize, messageObjects []*AlarmMessageObjectQueryType) *AlarmMessageQueryType {
	return &AlarmMessageQueryType{FunctionId: functionId, NumberOfObjects: numberOfObjects, ReturnCode: returnCode, TransportSize: transportSize, MessageObjects: messageObjects}
}

func CastAlarmMessageQueryType(structType interface{}) *AlarmMessageQueryType {
	castFunc := func(typ interface{}) *AlarmMessageQueryType {
		if casted, ok := typ.(AlarmMessageQueryType); ok {
			return &casted
		}
		if casted, ok := typ.(*AlarmMessageQueryType); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AlarmMessageQueryType) GetTypeName() string {
	return "AlarmMessageQueryType"
}

func (m *AlarmMessageQueryType) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AlarmMessageQueryType) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Const Field (DataLength)
	lengthInBits += 16

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AlarmMessageQueryType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AlarmMessageQueryTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageQueryType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageQueryType"); pullErr != nil {
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

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, pullErr
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, pullErr
	}
	_transportSize, _transportSizeErr := DataTransportSizeParse(readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (DataLength)
	DataLength, _DataLengthErr := readBuffer.ReadUint16("DataLength", 16)
	if _DataLengthErr != nil {
		return nil, errors.Wrap(_DataLengthErr, "Error parsing 'DataLength' field")
	}
	if DataLength != AlarmMessageQueryType_DATALENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageQueryType_DATALENGTH) + " but got " + fmt.Sprintf("%d", DataLength))
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	messageObjects := make([]*AlarmMessageObjectQueryType, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := AlarmMessageObjectQueryTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
			}
			messageObjects[curItem] = CastAlarmMessageObjectQueryType(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageQueryType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageQueryType(functionId, numberOfObjects, returnCode, transportSize, messageObjects), nil
}

func (m *AlarmMessageQueryType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageQueryType"); pushErr != nil {
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

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return pushErr
	}
	_returnCodeErr := m.ReturnCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return popErr
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return pushErr
	}
	_transportSizeErr := m.TransportSize.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return popErr
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Const Field (DataLength)
	_DataLengthErr := writeBuffer.WriteUint16("DataLength", 16, 0xFFFF)
	if _DataLengthErr != nil {
		return errors.Wrap(_DataLengthErr, "Error serializing 'DataLength' field")
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

	if popErr := writeBuffer.PopContext("AlarmMessageQueryType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
