package json

import (
	"fmt"
	"io"
	"net/http"
)

var (
	beginObjectBytes = []byte("{")
	endObjectBytes   = []byte("}")
	beginArrayBytes  = []byte("[")
	endArrayBytes    = []byte("]")
	separatorBytes   = []byte(",")
)

type StreamWriter interface {
	io.Writer
	http.Flusher
}

type Stream interface {
	BeginObject() error
	EndObject() error
	BeginArray() error
	EndArray() error
	WriteKey(keyName string) error
	WriteString(value string) error
	WriteInt(value int) error
	WriteInt8(value int8) error
	WriteInt16(value int16) error
	WriteInt32(value int32) error
	WriteInt64(value int64) error
	WriteUint(value uint) error
	WriteUint8(value uint8) error
	WriteUint16(value uint16) error
	WriteUint32(value uint32) error
	WriteUint64(value uint64) error
	WriteFloat32(value float32) error
	WriteFloat64(value float64) error
	WriteObject(obj Writable) error
	Encode(obj Writable) error
}

type stream struct {
	writer         StreamWriter
	shouldSeparate bool
}

type Writable interface {
	Encode(stream Stream) error
}

func NewEncoder(w StreamWriter) Stream {
	return &stream{writer: w}
}

func (s *stream) Encode(obj Writable) error {
	err := s.WriteObject(obj)
	s.shouldSeparate = false

	return err
}

func (s *stream) BeginObject() error {
	_, err := s.writeSeparated(beginObjectBytes)
	s.shouldSeparate = false
	return err
}

func (s *stream) EndObject() error {
	_, err := s.writer.Write(endObjectBytes)
	s.shouldSeparate = true
	return err
}

func (s *stream) BeginArray() error {
	_, err := s.writeSeparated(beginArrayBytes)
	s.shouldSeparate = false
	return err
}

func (s *stream) EndArray() error {
	_, err := s.writer.Write(endArrayBytes)
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteKey(keyName string) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%q:", keyName)))
	s.shouldSeparate = false
	return err
}

func (s *stream) WriteString(value string) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%q", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteInt(value int) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteInt8(value int8) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteInt16(value int16) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteInt32(value int32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteInt64(value int64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteUint(value uint) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteUint8(value uint8) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteUint16(value uint16) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteUint32(value uint32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteUint64(value uint64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%d", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteFloat32(value float32) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%g", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteFloat64(value float64) error {
	_, err := s.writeSeparated([]byte(fmt.Sprintf("%g", value)))
	s.shouldSeparate = true
	return err
}

func (s *stream) WriteObject(obj Writable) error {
	err := obj.Encode(s)
	s.shouldSeparate = true

	return err
}

func (s *stream) writeSeparated(data []byte) (bytes int, err error) {
	if s.shouldSeparate {
		bytes, err = s.writer.Write(separatorBytes)
	}

	prevBytes := bytes
	bytes, err = s.writer.Write(data)
	bytes += prevBytes

	return
}
