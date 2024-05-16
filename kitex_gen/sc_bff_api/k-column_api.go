// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_bff_api

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = sc_misc_base.KitexUnusedProtection
)

func (p *ColumnOmnibusRequest) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.STRING {
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
			if fieldTypeId == thrift.STRING {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnOmnibusRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnOmnibusRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.AuthorName = v

	}
	return offset, nil
}

func (p *ColumnOmnibusRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ColumnName = v

	}
	return offset, nil
}

func (p *ColumnOmnibusRequest) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ColumnPrice = v

	}
	return offset, nil
}

func (p *ColumnOmnibusRequest) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBase()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Base = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnOmnibusRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnOmnibusRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnOmnibusRequest")
	if p != nil {
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnOmnibusRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnOmnibusRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "authorName", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.AuthorName)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnName", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.ColumnName)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusRequest) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "columnPrice", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.ColumnPrice)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnOmnibusRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("authorName", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.AuthorName)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnOmnibusRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnName", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.ColumnName)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnOmnibusRequest) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("columnPrice", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.ColumnPrice)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnOmnibusRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnOmnibusResponse) FastRead(buf []byte) (int, error) {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnOmnibusResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnOmnibusResponse) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBaseResp()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.BaseResp = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnOmnibusResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnOmnibusResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnOmnibusResponse")
	if p != nil {
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnOmnibusResponse")
	if p != nil {
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnOmnibusResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnOmnibusResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnPageRequest) FastRead(buf []byte) (int, error) {
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
		case 4:
			if fieldTypeId == thrift.STRING {
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
			if fieldTypeId == thrift.I32 {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnPageRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnPageRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PageNum = v

	}
	return offset, nil
}

func (p *ColumnPageRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.PageSize = v

	}
	return offset, nil
}

func (p *ColumnPageRequest) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Name = &v

	}
	return offset, nil
}

func (p *ColumnPageRequest) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.AuthorId = &v

	}
	return offset, nil
}

func (p *ColumnPageRequest) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Status = &v

	}
	return offset, nil
}

func (p *ColumnPageRequest) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBase()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Base = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnPageRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnPageRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnPageRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnPageRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnPageRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnPageRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pageNum", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PageNum)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnPageRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pageSize", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.PageSize)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnPageRequest) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetName() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "name", thrift.STRING, 4)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Name)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnPageRequest) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetAuthorId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "authorId", thrift.I64, 5)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.AuthorId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnPageRequest) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetStatus() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "status", thrift.I32, 6)
		offset += bthrift.Binary.WriteI32(buf[offset:], *p.Status)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnPageRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnPageRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pageNum", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.PageNum)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnPageRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pageSize", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.PageSize)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnPageRequest) field4Length() int {
	l := 0
	if p.IsSetName() {
		l += bthrift.Binary.FieldBeginLength("name", thrift.STRING, 4)
		l += bthrift.Binary.StringLengthNocopy(*p.Name)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnPageRequest) field5Length() int {
	l := 0
	if p.IsSetAuthorId() {
		l += bthrift.Binary.FieldBeginLength("authorId", thrift.I64, 5)
		l += bthrift.Binary.I64Length(*p.AuthorId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnPageRequest) field6Length() int {
	l := 0
	if p.IsSetStatus() {
		l += bthrift.Binary.FieldBeginLength("status", thrift.I32, 6)
		l += bthrift.Binary.I32Length(*p.Status)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnPageRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnPageResponse) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.LIST {
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
			if fieldTypeId == thrift.STRUCT {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnPageResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnPageResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Data = make([]*RickColumnDTO, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewRickColumnDTO()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.Data = append(p.Data, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *ColumnPageResponse) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewPagination()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Pagination = tmp
	return offset, nil
}

func (p *ColumnPageResponse) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBaseResp()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.BaseResp = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnPageResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnPageResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnPageResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnPageResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnPageResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnPageResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "data", thrift.LIST, 1)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.Data {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnPageResponse) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "pagination", thrift.STRUCT, 2)
	offset += p.Pagination.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnPageResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnPageResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("data", thrift.LIST, 1)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.Data))
	for _, v := range p.Data {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnPageResponse) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("pagination", thrift.STRUCT, 2)
	l += p.Pagination.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnPageResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDetailRequest) FastRead(buf []byte) (int, error) {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnDetailRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnDetailRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Id = v

	}
	return offset, nil
}

func (p *ColumnDetailRequest) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBase()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Base = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnDetailRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnDetailRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnDetailRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnDetailRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnDetailRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnDetailRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDetailRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ColumnDetailRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDetailRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ColumnDetailResponse) FastRead(buf []byte) (int, error) {
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
			if fieldTypeId == thrift.STRUCT {
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
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ColumnDetailResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ColumnDetailResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewColumnDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Data = tmp
	return offset, nil
}

func (p *ColumnDetailResponse) FastReadField255(buf []byte) (int, error) {
	offset := 0

	tmp := sc_misc_base.NewBaseResp()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.BaseResp = tmp
	return offset, nil
}

// for compatibility
func (p *ColumnDetailResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *ColumnDetailResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ColumnDetailResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ColumnDetailResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ColumnDetailResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ColumnDetailResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "data", thrift.STRUCT, 1)
	offset += p.Data.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDetailResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ColumnDetailResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("data", thrift.STRUCT, 1)
	l += p.Data.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ColumnDetailResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
