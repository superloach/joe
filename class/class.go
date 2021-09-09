package class

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// Magic is the magic number used by all JVM ClassFiles.
const Magic uint32 = 0xCAFEBABE

// ErrMagic occurs when the Magic field of a ClassFile doesn't match Magic.
var ErrMagic = errors.New("magic number mismatch")

// ClassFile represents a JVM class file. Taken from Section 4.1 "The ClassFile Structure".
//
// XyzCount fields are omitted from the actual structure, because Go slices contain len. These
// fields are still encoded and decoded per the spec.
type ClassFile struct {
	Magic uint32

	MinorVersion uint16
	MajorVersion uint16

	// ConstantPoolCount uint16
	ConstantPool []CPInfo

	AccessFlags uint16

	ThisClass  uint16
	SuperClass uint16

	// InterfacesCount uint16
	Interfaces []uint16

	// FieldsCount uint16
	Fields []FieldInfo

	// MethodsCount uint16
	Methods []MethodInfo

	// AttributesCount uint16
	Attributes []AttributeInfo
}

func (cf *ClassFile) Check() error {
	if cf.Magic != Magic {
		return ErrMagic
	}

	return nil
}

func (cf *ClassFile) MarshalBinary() ([]byte, error) {
	err := cf.Check()
	if err != nil {
		return nil, fmt.Errorf("check: %w", err)
	}

	buf := &bytes.Buffer{}

	err = binary.Write(buf, binary.BigEndian, cf.Magic)
	if err != nil {
		return nil, fmt.Errorf("write magic: %w", err)
	}

	return buf.Bytes(), nil
}
