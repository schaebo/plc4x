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
package org.apache.plc4x.java.cbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class SALDataHvacActuator extends SALData implements Message {

  // Accessors for discriminator values.
  public ApplicationId getApplicationId() {
    return ApplicationId.HVAC_ACTUATOR;
  }

  // Properties.
  protected final LightingData hvacActuatorData;

  public SALDataHvacActuator(SALData salData, LightingData hvacActuatorData) {
    super(salData);
    this.hvacActuatorData = hvacActuatorData;
  }

  public LightingData getHvacActuatorData() {
    return hvacActuatorData;
  }

  @Override
  protected void serializeSALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("SALDataHvacActuator");

    // Simple Field (hvacActuatorData)
    writeSimpleField(
        "hvacActuatorData", hvacActuatorData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SALDataHvacActuator");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SALDataHvacActuator _value = this;

    // Simple field (hvacActuatorData)
    lengthInBits += hvacActuatorData.getLengthInBits();

    return lengthInBits;
  }

  public static SALDataHvacActuatorBuilder staticParseBuilder(
      ReadBuffer readBuffer, ApplicationId applicationId) throws ParseException {
    readBuffer.pullContext("SALDataHvacActuator");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    LightingData hvacActuatorData =
        readSimpleField(
            "hvacActuatorData",
            new DataReaderComplexDefault<>(() -> LightingData.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SALDataHvacActuator");
    // Create the instance
    return new SALDataHvacActuatorBuilder(hvacActuatorData);
  }

  public static class SALDataHvacActuatorBuilder implements SALData.SALDataBuilder {
    private final LightingData hvacActuatorData;

    public SALDataHvacActuatorBuilder(LightingData hvacActuatorData) {

      this.hvacActuatorData = hvacActuatorData;
    }

    public SALDataHvacActuator build(SALData salData) {
      SALDataHvacActuator sALDataHvacActuator = new SALDataHvacActuator(salData, hvacActuatorData);
      return sALDataHvacActuator;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SALDataHvacActuator)) {
      return false;
    }
    SALDataHvacActuator that = (SALDataHvacActuator) o;
    return (getHvacActuatorData() == that.getHvacActuatorData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getHvacActuatorData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}