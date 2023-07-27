/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// DeadbandType is an enum
type DeadbandType uint32

type IDeadbandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DeadbandType_deadbandTypeNone     DeadbandType = 0
	DeadbandType_deadbandTypeAbsolute DeadbandType = 1
	DeadbandType_deadbandTypePercent  DeadbandType = 2
)

var DeadbandTypeValues []DeadbandType

func init() {
	_ = errors.New
	DeadbandTypeValues = []DeadbandType{
		DeadbandType_deadbandTypeNone,
		DeadbandType_deadbandTypeAbsolute,
		DeadbandType_deadbandTypePercent,
	}
}

func DeadbandTypeByValue(value uint32) (enum DeadbandType, ok bool) {
	switch value {
	case 0:
		return DeadbandType_deadbandTypeNone, true
	case 1:
		return DeadbandType_deadbandTypeAbsolute, true
	case 2:
		return DeadbandType_deadbandTypePercent, true
	}
	return 0, false
}

func DeadbandTypeByName(value string) (enum DeadbandType, ok bool) {
	switch value {
	case "deadbandTypeNone":
		return DeadbandType_deadbandTypeNone, true
	case "deadbandTypeAbsolute":
		return DeadbandType_deadbandTypeAbsolute, true
	case "deadbandTypePercent":
		return DeadbandType_deadbandTypePercent, true
	}
	return 0, false
}

func DeadbandTypeKnows(value uint32) bool {
	for _, typeValue := range DeadbandTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDeadbandType(structType any) DeadbandType {
	castFunc := func(typ any) DeadbandType {
		if sDeadbandType, ok := typ.(DeadbandType); ok {
			return sDeadbandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeadbandType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m DeadbandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeadbandTypeParse(ctx context.Context, theBytes []byte) (DeadbandType, error) {
	return DeadbandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeadbandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeadbandType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("DeadbandType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DeadbandType")
	}
	if enum, ok := DeadbandTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return DeadbandType(val), nil
	} else {
		return enum, nil
	}
}

func (e DeadbandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DeadbandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("DeadbandType", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DeadbandType) PLC4XEnumName() string {
	switch e {
	case DeadbandType_deadbandTypeNone:
		return "deadbandTypeNone"
	case DeadbandType_deadbandTypeAbsolute:
		return "deadbandTypeAbsolute"
	case DeadbandType_deadbandTypePercent:
		return "deadbandTypePercent"
	}
	return ""
}

func (e DeadbandType) String() string {
	return e.PLC4XEnumName()
}