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
type BACnetReadAccessSpecification struct {
	ObjectIdentifier         *BACnetContextTagObjectIdentifier
	OpeningTag               *BACnetOpeningTag
	ListOfPropertyReferences []*BACnetPropertyReference
	ClosingTag               *BACnetClosingTag
}

// The corresponding interface
type IBACnetReadAccessSpecification interface {
	// GetObjectIdentifier returns ObjectIdentifier
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfPropertyReferences returns ListOfPropertyReferences
	GetListOfPropertyReferences() []*BACnetPropertyReference
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
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
func (m *BACnetReadAccessSpecification) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetReadAccessSpecification) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetReadAccessSpecification) GetListOfPropertyReferences() []*BACnetPropertyReference {
	return m.ListOfPropertyReferences
}

func (m *BACnetReadAccessSpecification) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewBACnetReadAccessSpecification(objectIdentifier *BACnetContextTagObjectIdentifier, openingTag *BACnetOpeningTag, listOfPropertyReferences []*BACnetPropertyReference, closingTag *BACnetClosingTag) *BACnetReadAccessSpecification {
	return &BACnetReadAccessSpecification{ObjectIdentifier: objectIdentifier, OpeningTag: openingTag, ListOfPropertyReferences: listOfPropertyReferences, ClosingTag: closingTag}
}

func CastBACnetReadAccessSpecification(structType interface{}) *BACnetReadAccessSpecification {
	castFunc := func(typ interface{}) *BACnetReadAccessSpecification {
		if casted, ok := typ.(BACnetReadAccessSpecification); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetReadAccessSpecification); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetReadAccessSpecification) GetTypeName() string {
	return "BACnetReadAccessSpecification"
}

func (m *BACnetReadAccessSpecification) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetReadAccessSpecification) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.LengthInBits()

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.LengthInBits()

	// Array field
	if len(m.ListOfPropertyReferences) > 0 {
		for _, element := range m.ListOfPropertyReferences {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetReadAccessSpecification) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetReadAccessSpecificationParse(readBuffer utils.ReadBuffer) (*BACnetReadAccessSpecification, error) {
	if pullErr := readBuffer.PullContext("BACnetReadAccessSpecification"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType_OPENING_TAG)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (listOfPropertyReferences)
	if pullErr := readBuffer.PullContext("listOfPropertyReferences", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	listOfPropertyReferences := make([]*BACnetPropertyReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 1)) {
			_item, _err := BACnetPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfPropertyReferences' field")
			}
			listOfPropertyReferences = append(listOfPropertyReferences, CastBACnetPropertyReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfPropertyReferences", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType_CLOSING_TAG)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessSpecification"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetReadAccessSpecification(objectIdentifier, openingTag, listOfPropertyReferences, closingTag), nil
}

func (m *BACnetReadAccessSpecification) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetReadAccessSpecification"); pushErr != nil {
		return pushErr
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return pushErr
	}
	_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return popErr
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfPropertyReferences)
	if m.ListOfPropertyReferences != nil {
		if pushErr := writeBuffer.PushContext("listOfPropertyReferences", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.ListOfPropertyReferences {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfPropertyReferences' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfPropertyReferences", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessSpecification"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetReadAccessSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
