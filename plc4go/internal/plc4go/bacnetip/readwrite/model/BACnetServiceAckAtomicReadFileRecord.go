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
type BACnetServiceAckAtomicReadFileRecord struct {
	*BACnetServiceAckAtomicReadFileStreamOrRecord
	FileStartRecord     *BACnetApplicationTagSignedInteger
	ReturnedRecordCount *BACnetApplicationTagUnsignedInteger
	FileRecordData      []*BACnetApplicationTagOctetString
}

// The corresponding interface
type IBACnetServiceAckAtomicReadFileRecord interface {
	// GetFileStartRecord returns FileStartRecord
	GetFileStartRecord() *BACnetApplicationTagSignedInteger
	// GetReturnedRecordCount returns ReturnedRecordCount
	GetReturnedRecordCount() *BACnetApplicationTagUnsignedInteger
	// GetFileRecordData returns FileRecordData
	GetFileRecordData() []*BACnetApplicationTagOctetString
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicReadFileRecord) PeekedTagNumber() uint8 {
	return 0x1
}

func (m *BACnetServiceAckAtomicReadFileRecord) GetPeekedTagNumber() uint8 {
	return 0x1
}

func (m *BACnetServiceAckAtomicReadFileRecord) InitializeParent(parent *BACnetServiceAckAtomicReadFileStreamOrRecord, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, peekedTagNumber uint8) {
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.PeekedTagHeader = peekedTagHeader
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.OpeningTag = openingTag
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.ClosingTag = closingTag
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.PeekedTagNumber = peekedTagNumber
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicReadFileRecord) GetFileStartRecord() *BACnetApplicationTagSignedInteger {
	return m.FileStartRecord
}

func (m *BACnetServiceAckAtomicReadFileRecord) GetReturnedRecordCount() *BACnetApplicationTagUnsignedInteger {
	return m.ReturnedRecordCount
}

func (m *BACnetServiceAckAtomicReadFileRecord) GetFileRecordData() []*BACnetApplicationTagOctetString {
	return m.FileRecordData
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewBACnetServiceAckAtomicReadFileRecord(fileStartRecord *BACnetApplicationTagSignedInteger, returnedRecordCount *BACnetApplicationTagUnsignedInteger, fileRecordData []*BACnetApplicationTagOctetString, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetServiceAckAtomicReadFileStreamOrRecord {
	child := &BACnetServiceAckAtomicReadFileRecord{
		FileStartRecord:     fileStartRecord,
		ReturnedRecordCount: returnedRecordCount,
		FileRecordData:      fileRecordData,
		BACnetServiceAckAtomicReadFileStreamOrRecord: NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag, peekedTagNumber),
	}
	child.Child = child
	return child.BACnetServiceAckAtomicReadFileStreamOrRecord
}

func CastBACnetServiceAckAtomicReadFileRecord(structType interface{}) *BACnetServiceAckAtomicReadFileRecord {
	castFunc := func(typ interface{}) *BACnetServiceAckAtomicReadFileRecord {
		if casted, ok := typ.(BACnetServiceAckAtomicReadFileRecord); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckAtomicReadFileRecord); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
			return CastBACnetServiceAckAtomicReadFileRecord(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
			return CastBACnetServiceAckAtomicReadFileRecord(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckAtomicReadFileRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileRecord"
}

func (m *BACnetServiceAckAtomicReadFileRecord) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicReadFileRecord) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (fileStartRecord)
	lengthInBits += m.FileStartRecord.LengthInBits()

	// Simple field (returnedRecordCount)
	lengthInBits += m.ReturnedRecordCount.LengthInBits()

	// Array field
	if len(m.FileRecordData) > 0 {
		for i, element := range m.FileRecordData {
			last := i == len(m.FileRecordData)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *BACnetServiceAckAtomicReadFileRecord) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckAtomicReadFileRecordParse(readBuffer utils.ReadBuffer) (*BACnetServiceAckAtomicReadFileStreamOrRecord, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileRecord"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (fileStartRecord)
	if pullErr := readBuffer.PullContext("fileStartRecord"); pullErr != nil {
		return nil, pullErr
	}
	_fileStartRecord, _fileStartRecordErr := BACnetApplicationTagParse(readBuffer)
	if _fileStartRecordErr != nil {
		return nil, errors.Wrap(_fileStartRecordErr, "Error parsing 'fileStartRecord' field")
	}
	fileStartRecord := CastBACnetApplicationTagSignedInteger(_fileStartRecord)
	if closeErr := readBuffer.CloseContext("fileStartRecord"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (returnedRecordCount)
	if pullErr := readBuffer.PullContext("returnedRecordCount"); pullErr != nil {
		return nil, pullErr
	}
	_returnedRecordCount, _returnedRecordCountErr := BACnetApplicationTagParse(readBuffer)
	if _returnedRecordCountErr != nil {
		return nil, errors.Wrap(_returnedRecordCountErr, "Error parsing 'returnedRecordCount' field")
	}
	returnedRecordCount := CastBACnetApplicationTagUnsignedInteger(_returnedRecordCount)
	if closeErr := readBuffer.CloseContext("returnedRecordCount"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (fileRecordData)
	if pullErr := readBuffer.PullContext("fileRecordData", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	fileRecordData := make([]*BACnetApplicationTagOctetString, returnedRecordCount.Payload.ActualValue)
	{
		for curItem := uint16(0); curItem < uint16(returnedRecordCount.Payload.ActualValue); curItem++ {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'fileRecordData' field")
			}
			fileRecordData[curItem] = CastBACnetApplicationTagOctetString(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("fileRecordData", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileRecord"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckAtomicReadFileRecord{
		FileStartRecord:     CastBACnetApplicationTagSignedInteger(fileStartRecord),
		ReturnedRecordCount: CastBACnetApplicationTagUnsignedInteger(returnedRecordCount),
		FileRecordData:      fileRecordData,
		BACnetServiceAckAtomicReadFileStreamOrRecord: &BACnetServiceAckAtomicReadFileStreamOrRecord{},
	}
	_child.BACnetServiceAckAtomicReadFileStreamOrRecord.Child = _child
	return _child.BACnetServiceAckAtomicReadFileStreamOrRecord, nil
}

func (m *BACnetServiceAckAtomicReadFileRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileRecord"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fileStartRecord)
		if pushErr := writeBuffer.PushContext("fileStartRecord"); pushErr != nil {
			return pushErr
		}
		_fileStartRecordErr := m.FileStartRecord.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileStartRecord"); popErr != nil {
			return popErr
		}
		if _fileStartRecordErr != nil {
			return errors.Wrap(_fileStartRecordErr, "Error serializing 'fileStartRecord' field")
		}

		// Simple Field (returnedRecordCount)
		if pushErr := writeBuffer.PushContext("returnedRecordCount"); pushErr != nil {
			return pushErr
		}
		_returnedRecordCountErr := m.ReturnedRecordCount.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("returnedRecordCount"); popErr != nil {
			return popErr
		}
		if _returnedRecordCountErr != nil {
			return errors.Wrap(_returnedRecordCountErr, "Error serializing 'returnedRecordCount' field")
		}

		// Array Field (fileRecordData)
		if m.FileRecordData != nil {
			if pushErr := writeBuffer.PushContext("fileRecordData", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.FileRecordData {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'fileRecordData' field")
				}
			}
			if popErr := writeBuffer.PopContext("fileRecordData", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileRecord"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckAtomicReadFileRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
