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

// Code generated by "plc4xgenerator -type=DeviceInfo"; DO NOT EDIT.

package bacnetip

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DeviceInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DeviceInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("DeviceInfo"); err != nil {
		return err
	}

	if d.DeviceIdentifier != nil {
		if serializableField, ok := d.DeviceIdentifier.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("deviceIdentifier"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("deviceIdentifier"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.DeviceIdentifier)
			if err := writeBuffer.WriteString("deviceIdentifier", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	{
		_value := fmt.Sprintf("%v", d.Address)

		if err := writeBuffer.WriteString("address", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	if d.MaximumApduLengthAccepted != nil {
		if err := writeBuffer.WriteString("maximumApduLengthAccepted", uint32(len(d.MaximumApduLengthAccepted.String())*8), "UTF-8", d.MaximumApduLengthAccepted.String()); err != nil {
			return err
		}
	}
	if d.SegmentationSupported != nil {
		if err := writeBuffer.WriteString("segmentationSupported", uint32(len(d.SegmentationSupported.String())*8), "UTF-8", d.SegmentationSupported.String()); err != nil {
			return err
		}
	}
	if d.MaxSegmentsAccepted != nil {
		if err := writeBuffer.WriteString("maxSegmentsAccepted", uint32(len(d.MaxSegmentsAccepted.String())*8), "UTF-8", d.MaxSegmentsAccepted.String()); err != nil {
			return err
		}
	}
	if d.VendorId != nil {
		if err := writeBuffer.WriteString("vendorId", uint32(len(d.VendorId.String())*8), "UTF-8", d.VendorId.String()); err != nil {
			return err
		}
	}
	{
		_value := fmt.Sprintf("%v", d.MaximumNpduLength)

		if err := writeBuffer.WriteString("maximumNpduLength", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteInt64("_refCount", 64, int64(d._refCount)); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d._cacheKey)

		if err := writeBuffer.WriteString("_cacheKey", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("DeviceInfo"); err != nil {
		return err
	}
	return nil
}

func (d *DeviceInfo) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
