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

// CreateSubscriptionRequest is the corresponding interface of CreateSubscriptionRequest
type CreateSubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetRequestedPublishingInterval returns RequestedPublishingInterval (property field)
	GetRequestedPublishingInterval() float64
	// GetRequestedLifetimeCount returns RequestedLifetimeCount (property field)
	GetRequestedLifetimeCount() uint32
	// GetRequestedMaxKeepAliveCount returns RequestedMaxKeepAliveCount (property field)
	GetRequestedMaxKeepAliveCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetPriority returns Priority (property field)
	GetPriority() uint8
}

// CreateSubscriptionRequestExactly can be used when we want exactly this type and not a type which fulfills CreateSubscriptionRequest.
// This is useful for switch cases.
type CreateSubscriptionRequestExactly interface {
	CreateSubscriptionRequest
	isCreateSubscriptionRequest() bool
}

// _CreateSubscriptionRequest is the data-structure of this message
type _CreateSubscriptionRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader               ExtensionObjectDefinition
	RequestedPublishingInterval float64
	RequestedLifetimeCount      uint32
	RequestedMaxKeepAliveCount  uint32
	MaxNotificationsPerPublish  uint32
	PublishingEnabled           bool
	Priority                    uint8
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSubscriptionRequest) GetIdentifier() string {
	return "787"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSubscriptionRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CreateSubscriptionRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSubscriptionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateSubscriptionRequest) GetRequestedPublishingInterval() float64 {
	return m.RequestedPublishingInterval
}

func (m *_CreateSubscriptionRequest) GetRequestedLifetimeCount() uint32 {
	return m.RequestedLifetimeCount
}

func (m *_CreateSubscriptionRequest) GetRequestedMaxKeepAliveCount() uint32 {
	return m.RequestedMaxKeepAliveCount
}

func (m *_CreateSubscriptionRequest) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_CreateSubscriptionRequest) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_CreateSubscriptionRequest) GetPriority() uint8 {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSubscriptionRequest factory function for _CreateSubscriptionRequest
func NewCreateSubscriptionRequest(requestHeader ExtensionObjectDefinition, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, priority uint8) *_CreateSubscriptionRequest {
	_result := &_CreateSubscriptionRequest{
		RequestHeader:               requestHeader,
		RequestedPublishingInterval: requestedPublishingInterval,
		RequestedLifetimeCount:      requestedLifetimeCount,
		RequestedMaxKeepAliveCount:  requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:  maxNotificationsPerPublish,
		PublishingEnabled:           publishingEnabled,
		Priority:                    priority,
		_ExtensionObjectDefinition:  NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSubscriptionRequest(structType any) CreateSubscriptionRequest {
	if casted, ok := structType.(CreateSubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSubscriptionRequest) GetTypeName() string {
	return "CreateSubscriptionRequest"
}

func (m *_CreateSubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (requestedPublishingInterval)
	lengthInBits += 64

	// Simple field (requestedLifetimeCount)
	lengthInBits += 32

	// Simple field (requestedMaxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CreateSubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateSubscriptionRequestParse(ctx context.Context, theBytes []byte, identifier string) (CreateSubscriptionRequest, error) {
	return CreateSubscriptionRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CreateSubscriptionRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CreateSubscriptionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateSubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of CreateSubscriptionRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (requestedPublishingInterval)
	_requestedPublishingInterval, _requestedPublishingIntervalErr := readBuffer.ReadFloat64("requestedPublishingInterval", 64)
	if _requestedPublishingIntervalErr != nil {
		return nil, errors.Wrap(_requestedPublishingIntervalErr, "Error parsing 'requestedPublishingInterval' field of CreateSubscriptionRequest")
	}
	requestedPublishingInterval := _requestedPublishingInterval

	// Simple Field (requestedLifetimeCount)
	_requestedLifetimeCount, _requestedLifetimeCountErr := readBuffer.ReadUint32("requestedLifetimeCount", 32)
	if _requestedLifetimeCountErr != nil {
		return nil, errors.Wrap(_requestedLifetimeCountErr, "Error parsing 'requestedLifetimeCount' field of CreateSubscriptionRequest")
	}
	requestedLifetimeCount := _requestedLifetimeCount

	// Simple Field (requestedMaxKeepAliveCount)
	_requestedMaxKeepAliveCount, _requestedMaxKeepAliveCountErr := readBuffer.ReadUint32("requestedMaxKeepAliveCount", 32)
	if _requestedMaxKeepAliveCountErr != nil {
		return nil, errors.Wrap(_requestedMaxKeepAliveCountErr, "Error parsing 'requestedMaxKeepAliveCount' field of CreateSubscriptionRequest")
	}
	requestedMaxKeepAliveCount := _requestedMaxKeepAliveCount

	// Simple Field (maxNotificationsPerPublish)
	_maxNotificationsPerPublish, _maxNotificationsPerPublishErr := readBuffer.ReadUint32("maxNotificationsPerPublish", 32)
	if _maxNotificationsPerPublishErr != nil {
		return nil, errors.Wrap(_maxNotificationsPerPublishErr, "Error parsing 'maxNotificationsPerPublish' field of CreateSubscriptionRequest")
	}
	maxNotificationsPerPublish := _maxNotificationsPerPublish

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CreateSubscriptionRequest")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (publishingEnabled)
	_publishingEnabled, _publishingEnabledErr := readBuffer.ReadBit("publishingEnabled")
	if _publishingEnabledErr != nil {
		return nil, errors.Wrap(_publishingEnabledErr, "Error parsing 'publishingEnabled' field of CreateSubscriptionRequest")
	}
	publishingEnabled := _publishingEnabled

	// Simple Field (priority)
	_priority, _priorityErr := readBuffer.ReadUint8("priority", 8)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of CreateSubscriptionRequest")
	}
	priority := _priority

	if closeErr := readBuffer.CloseContext("CreateSubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSubscriptionRequest")
	}

	// Create a partially initialized instance
	_child := &_CreateSubscriptionRequest{
		_ExtensionObjectDefinition:  &_ExtensionObjectDefinition{},
		RequestHeader:               requestHeader,
		RequestedPublishingInterval: requestedPublishingInterval,
		RequestedLifetimeCount:      requestedLifetimeCount,
		RequestedMaxKeepAliveCount:  requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:  maxNotificationsPerPublish,
		PublishingEnabled:           publishingEnabled,
		Priority:                    priority,
		reservedField0:              reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CreateSubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSubscriptionRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (requestedPublishingInterval)
		requestedPublishingInterval := float64(m.GetRequestedPublishingInterval())
		_requestedPublishingIntervalErr := writeBuffer.WriteFloat64("requestedPublishingInterval", 64, (requestedPublishingInterval))
		if _requestedPublishingIntervalErr != nil {
			return errors.Wrap(_requestedPublishingIntervalErr, "Error serializing 'requestedPublishingInterval' field")
		}

		// Simple Field (requestedLifetimeCount)
		requestedLifetimeCount := uint32(m.GetRequestedLifetimeCount())
		_requestedLifetimeCountErr := writeBuffer.WriteUint32("requestedLifetimeCount", 32, uint32((requestedLifetimeCount)))
		if _requestedLifetimeCountErr != nil {
			return errors.Wrap(_requestedLifetimeCountErr, "Error serializing 'requestedLifetimeCount' field")
		}

		// Simple Field (requestedMaxKeepAliveCount)
		requestedMaxKeepAliveCount := uint32(m.GetRequestedMaxKeepAliveCount())
		_requestedMaxKeepAliveCountErr := writeBuffer.WriteUint32("requestedMaxKeepAliveCount", 32, uint32((requestedMaxKeepAliveCount)))
		if _requestedMaxKeepAliveCountErr != nil {
			return errors.Wrap(_requestedMaxKeepAliveCountErr, "Error serializing 'requestedMaxKeepAliveCount' field")
		}

		// Simple Field (maxNotificationsPerPublish)
		maxNotificationsPerPublish := uint32(m.GetMaxNotificationsPerPublish())
		_maxNotificationsPerPublishErr := writeBuffer.WriteUint32("maxNotificationsPerPublish", 32, uint32((maxNotificationsPerPublish)))
		if _maxNotificationsPerPublishErr != nil {
			return errors.Wrap(_maxNotificationsPerPublishErr, "Error serializing 'maxNotificationsPerPublish' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (publishingEnabled)
		publishingEnabled := bool(m.GetPublishingEnabled())
		_publishingEnabledErr := writeBuffer.WriteBit("publishingEnabled", (publishingEnabled))
		if _publishingEnabledErr != nil {
			return errors.Wrap(_publishingEnabledErr, "Error serializing 'publishingEnabled' field")
		}

		// Simple Field (priority)
		priority := uint8(m.GetPriority())
		_priorityErr := writeBuffer.WriteUint8("priority", 8, uint8((priority)))
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("CreateSubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSubscriptionRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSubscriptionRequest) isCreateSubscriptionRequest() bool {
	return true
}

func (m *_CreateSubscriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
