// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_misc_api

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *ColumnQuoteDTO) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnQuoteDTO[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnQuoteDTO) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CreateTime = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UpdateTime = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ColumnId = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Editor = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Price = v

	}
	return offset, nil
}

func (p *ColumnQuoteDTO) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Status = v

	}
	return offset, nil
}

// for compatibility
func (p *ColumnQuoteDTO) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnQuoteDTO) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnQuoteDTO")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnQuoteDTO")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnQuoteDTO) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "createTime", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.CreateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "updateTime", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.UpdateTime)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnId", thrift.I64, 4)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.ColumnId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "editor", thrift.I64, 5)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Editor)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "price", thrift.I64, 6)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Price)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status", thrift.I32, 7)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Status)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnQuoteDTO) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("createTime", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.CreateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("updateTime", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.UpdateTime)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnId", thrift.I64, 4)
	l += bthrift.Binary.I64Length(p.ColumnId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("editor", thrift.I64, 5)
	l += bthrift.Binary.I64Length(p.Editor)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("price", thrift.I64, 6)
	l += bthrift.Binary.I64Length(p.Price)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnQuoteDTO) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("status", thrift.I32, 7)
	l += bthrift.Binary.I32Length(p.Status)

	l += bthrift.Binary.FieldEndLength()
	return l
}