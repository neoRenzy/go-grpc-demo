/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package java_exception

// StreamCorruptedException represents an exception of the same name in java
type StreamCorruptedException struct {
	SerialVersionUID     int64
	DetailMessage        string
	StackTrace           []StackTraceElement
	SuppressedExceptions []Throwabler
	Cause                Throwabler
}

// NewStreamCorruptedException is the constructor
func NewStreamCorruptedException(detailMessage string) *StreamCorruptedException {
	return &StreamCorruptedException{DetailMessage: detailMessage, StackTrace: []StackTraceElement{}}
}

// Error output error message
func (e StreamCorruptedException) Error() string {
	return e.DetailMessage
}

// JavaClassName  java fully qualified path
func (StreamCorruptedException) JavaClassName() string {
	return "java.io.StreamCorruptedException"
}

// equals to getStackTrace in java
func (e StreamCorruptedException) GetStackTrace() []StackTraceElement {
	return e.StackTrace
}