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

// LocaleId is the corresponding interface of LocaleId
type LocaleId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// LocaleIdExactly can be used when we want exactly this type and not a type which fulfills LocaleId.
// This is useful for switch cases.
type LocaleIdExactly interface {
	LocaleId
	isLocaleId() bool
}

// _LocaleId is the data-structure of this message
type _LocaleId struct {
}

// NewLocaleId factory function for _LocaleId
func NewLocaleId() *_LocaleId {
	return &_LocaleId{}
}

// Deprecated: use the interface for direct cast
func CastLocaleId(structType any) LocaleId {
	if casted, ok := structType.(LocaleId); ok {
		return casted
	}
	if casted, ok := structType.(*LocaleId); ok {
		return *casted
	}
	return nil
}

func (m *_LocaleId) GetTypeName() string {
	return "LocaleId"
}

func (m *_LocaleId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_LocaleId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LocaleIdParse(ctx context.Context, theBytes []byte) (LocaleId, error) {
	return LocaleIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LocaleIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LocaleId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("LocaleId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LocaleId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LocaleId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LocaleId")
	}

	// Create the instance
	return &_LocaleId{}, nil
}

func (m *_LocaleId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LocaleId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LocaleId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LocaleId")
	}

	if popErr := writeBuffer.PopContext("LocaleId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LocaleId")
	}
	return nil
}

func (m *_LocaleId) isLocaleId() bool {
	return true
}

func (m *_LocaleId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}