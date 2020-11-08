//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetTagApplicationObjectIdentifier struct {
    Parent *BACnetTag
    IBACnetTagApplicationObjectIdentifier
}

// The corresponding interface
type IBACnetTagApplicationObjectIdentifier interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationObjectIdentifier) ContextSpecificTag() uint8 {
    return 0
}


func (m *BACnetTagApplicationObjectIdentifier) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
    m.Parent.TypeOrTagNumber = typeOrTagNumber
    m.Parent.LengthValueType = lengthValueType
    m.Parent.ExtTagNumber = extTagNumber
    m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationObjectIdentifier(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
    child := &BACnetTagApplicationObjectIdentifier{
        Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetTagApplicationObjectIdentifier(structType interface{}) BACnetTagApplicationObjectIdentifier {
    castFunc := func(typ interface{}) BACnetTagApplicationObjectIdentifier {
        if casted, ok := typ.(BACnetTagApplicationObjectIdentifier); ok {
            return casted
        }
        if casted, ok := typ.(*BACnetTagApplicationObjectIdentifier); ok {
            return *casted
        }
        if casted, ok := typ.(BACnetTag); ok {
            return CastBACnetTagApplicationObjectIdentifier(casted.Child)
        }
        if casted, ok := typ.(*BACnetTag); ok {
            return CastBACnetTagApplicationObjectIdentifier(casted.Child)
        }
        return BACnetTagApplicationObjectIdentifier{}
    }
    return castFunc(structType)
}

func (m *BACnetTagApplicationObjectIdentifier) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BACnetTagApplicationObjectIdentifier) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetTagApplicationObjectIdentifierParse(io *utils.ReadBuffer) (*BACnetTag, error) {

    // Create a partially initialized instance
    _child := &BACnetTagApplicationObjectIdentifier{
        Parent: &BACnetTag{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetTagApplicationObjectIdentifier) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetTagApplicationObjectIdentifier) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *BACnetTagApplicationObjectIdentifier) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}
