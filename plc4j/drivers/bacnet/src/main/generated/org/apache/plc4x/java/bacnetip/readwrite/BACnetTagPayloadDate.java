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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetTagPayloadDate implements Message {

  // Properties.
  protected final short yearMinus1900;
  protected final short month;
  protected final short dayOfMonth;
  protected final short dayOfWeek;

  public BACnetTagPayloadDate(short yearMinus1900, short month, short dayOfMonth, short dayOfWeek) {
    super();
    this.yearMinus1900 = yearMinus1900;
    this.month = month;
    this.dayOfMonth = dayOfMonth;
    this.dayOfWeek = dayOfWeek;
  }

  public short getYearMinus1900() {
    return yearMinus1900;
  }

  public short getMonth() {
    return month;
  }

  public short getDayOfMonth() {
    return dayOfMonth;
  }

  public short getDayOfWeek() {
    return dayOfWeek;
  }

  public short getWildcard() {
    return (short) (0xFF);
  }

  public boolean getYearIsWildcard() {
    return (boolean) ((getYearMinus1900()) == (getWildcard()));
  }

  public int getYear() {
    return (int) ((getYearMinus1900()) + (1900));
  }

  public boolean getMonthIsWildcard() {
    return (boolean) ((getMonth()) == (getWildcard()));
  }

  public boolean getOddMonthWildcard() {
    return (boolean) ((getMonth()) == (13));
  }

  public boolean getEvenMonthWildcard() {
    return (boolean) ((getMonth()) == (14));
  }

  public boolean getDayOfMonthIsWildcard() {
    return (boolean) ((getDayOfMonth()) == (getWildcard()));
  }

  public boolean getLastDayOfMonthWildcard() {
    return (boolean) ((getDayOfMonth()) == (32));
  }

  public boolean getOddDayOfMonthWildcard() {
    return (boolean) ((getDayOfMonth()) == (33));
  }

  public boolean getEvenDayOfMonthWildcard() {
    return (boolean) ((getDayOfMonth()) == (34));
  }

  public boolean getDayOfWeekIsWildcard() {
    return (boolean) ((getDayOfWeek()) == (getWildcard()));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetTagPayloadDate");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short wildcard = getWildcard();
    writeBuffer.writeVirtual("wildcard", wildcard);

    // Simple Field (yearMinus1900)
    writeSimpleField("yearMinus1900", yearMinus1900, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean yearIsWildcard = getYearIsWildcard();
    writeBuffer.writeVirtual("yearIsWildcard", yearIsWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    int year = getYear();
    writeBuffer.writeVirtual("year", year);

    // Simple Field (month)
    writeSimpleField("month", month, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean monthIsWildcard = getMonthIsWildcard();
    writeBuffer.writeVirtual("monthIsWildcard", monthIsWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean oddMonthWildcard = getOddMonthWildcard();
    writeBuffer.writeVirtual("oddMonthWildcard", oddMonthWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean evenMonthWildcard = getEvenMonthWildcard();
    writeBuffer.writeVirtual("evenMonthWildcard", evenMonthWildcard);

    // Simple Field (dayOfMonth)
    writeSimpleField("dayOfMonth", dayOfMonth, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean dayOfMonthIsWildcard = getDayOfMonthIsWildcard();
    writeBuffer.writeVirtual("dayOfMonthIsWildcard", dayOfMonthIsWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean lastDayOfMonthWildcard = getLastDayOfMonthWildcard();
    writeBuffer.writeVirtual("lastDayOfMonthWildcard", lastDayOfMonthWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean oddDayOfMonthWildcard = getOddDayOfMonthWildcard();
    writeBuffer.writeVirtual("oddDayOfMonthWildcard", oddDayOfMonthWildcard);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean evenDayOfMonthWildcard = getEvenDayOfMonthWildcard();
    writeBuffer.writeVirtual("evenDayOfMonthWildcard", evenDayOfMonthWildcard);

    // Simple Field (dayOfWeek)
    writeSimpleField("dayOfWeek", dayOfWeek, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean dayOfWeekIsWildcard = getDayOfWeekIsWildcard();
    writeBuffer.writeVirtual("dayOfWeekIsWildcard", dayOfWeekIsWildcard);

    writeBuffer.popContext("BACnetTagPayloadDate");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetTagPayloadDate _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Simple field (yearMinus1900)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (month)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (dayOfMonth)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Simple field (dayOfWeek)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetTagPayloadDate staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetTagPayloadDate staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetTagPayloadDate");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    short wildcard = readVirtualField("wildcard", short.class, 0xFF);

    short yearMinus1900 = readSimpleField("yearMinus1900", readUnsignedShort(readBuffer, 8));
    boolean yearIsWildcard =
        readVirtualField("yearIsWildcard", boolean.class, (yearMinus1900) == (wildcard));
    int year = readVirtualField("year", int.class, (yearMinus1900) + (1900));

    short month = readSimpleField("month", readUnsignedShort(readBuffer, 8));
    boolean monthIsWildcard =
        readVirtualField("monthIsWildcard", boolean.class, (month) == (wildcard));
    boolean oddMonthWildcard = readVirtualField("oddMonthWildcard", boolean.class, (month) == (13));
    boolean evenMonthWildcard =
        readVirtualField("evenMonthWildcard", boolean.class, (month) == (14));

    short dayOfMonth = readSimpleField("dayOfMonth", readUnsignedShort(readBuffer, 8));
    boolean dayOfMonthIsWildcard =
        readVirtualField("dayOfMonthIsWildcard", boolean.class, (dayOfMonth) == (wildcard));
    boolean lastDayOfMonthWildcard =
        readVirtualField("lastDayOfMonthWildcard", boolean.class, (dayOfMonth) == (32));
    boolean oddDayOfMonthWildcard =
        readVirtualField("oddDayOfMonthWildcard", boolean.class, (dayOfMonth) == (33));
    boolean evenDayOfMonthWildcard =
        readVirtualField("evenDayOfMonthWildcard", boolean.class, (dayOfMonth) == (34));

    short dayOfWeek = readSimpleField("dayOfWeek", readUnsignedShort(readBuffer, 8));
    boolean dayOfWeekIsWildcard =
        readVirtualField("dayOfWeekIsWildcard", boolean.class, (dayOfWeek) == (wildcard));

    readBuffer.closeContext("BACnetTagPayloadDate");
    // Create the instance
    BACnetTagPayloadDate _bACnetTagPayloadDate;
    _bACnetTagPayloadDate = new BACnetTagPayloadDate(yearMinus1900, month, dayOfMonth, dayOfWeek);
    return _bACnetTagPayloadDate;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTagPayloadDate)) {
      return false;
    }
    BACnetTagPayloadDate that = (BACnetTagPayloadDate) o;
    return (getYearMinus1900() == that.getYearMinus1900())
        && (getMonth() == that.getMonth())
        && (getDayOfMonth() == that.getDayOfMonth())
        && (getDayOfWeek() == that.getDayOfWeek())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getYearMinus1900(), getMonth(), getDayOfMonth(), getDayOfWeek());
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
