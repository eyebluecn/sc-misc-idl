// Code generated by Kitex v0.9.1. DO NOT EDIT.

package sc_misc_api

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

func (p *ReaderLoginRequest) FastRead(buf []byte) (int, error) {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderLoginRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderLoginRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Username = v

	}
	return offset, nil
}

func (p *ReaderLoginRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Password = v

	}
	return offset, nil
}

func (p *ReaderLoginRequest) FastReadField255(buf []byte) (int, error) {
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
func (p *ReaderLoginRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ReaderLoginRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ReaderLoginRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ReaderLoginRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ReaderLoginRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "username", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Username)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginRequest) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "password", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Password)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ReaderLoginRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("username", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.Username)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderLoginRequest) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("password", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.Password)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderLoginRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ReaderLoginResponse) FastRead(buf []byte) (int, error) {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderLoginResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderLoginResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewReaderDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Data = tmp
	return offset, nil
}

func (p *ReaderLoginResponse) FastReadField255(buf []byte) (int, error) {
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
func (p *ReaderLoginResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *ReaderLoginResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ReaderLoginResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ReaderLoginResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ReaderLoginResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "data", thrift.STRUCT, 1)
	offset += p.Data.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderLoginResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("data", thrift.STRUCT, 1)
	l += p.Data.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderLoginResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderQueryByIdRequest) FastRead(buf []byte) (int, error) {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderQueryByIdRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderQueryByIdRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.ReaderId = v

	}
	return offset, nil
}

func (p *ReaderQueryByIdRequest) FastReadField255(buf []byte) (int, error) {
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
func (p *ReaderQueryByIdRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ReaderQueryByIdRequest) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ReaderQueryByIdRequest")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ReaderQueryByIdRequest) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ReaderQueryByIdRequest")
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ReaderQueryByIdRequest) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "readerId", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.ReaderId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderQueryByIdRequest) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetBase() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "base", thrift.STRUCT, 255)
		offset += p.Base.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *ReaderQueryByIdRequest) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("readerId", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.ReaderId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderQueryByIdRequest) field255Length() int {
	l := 0
	if p.IsSetBase() {
		l += bthrift.Binary.FieldBeginLength("base", thrift.STRUCT, 255)
		l += p.Base.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *ReaderQueryByIdResponse) FastRead(buf []byte) (int, error) {
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
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderQueryByIdResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderQueryByIdResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewReaderDTO()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Data = tmp
	return offset, nil
}

func (p *ReaderQueryByIdResponse) FastReadField255(buf []byte) (int, error) {
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
func (p *ReaderQueryByIdResponse) FastWrite(buf []byte) int {
	return 0
}

func (p *ReaderQueryByIdResponse) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "ReaderQueryByIdResponse")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField255(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *ReaderQueryByIdResponse) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("ReaderQueryByIdResponse")
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *ReaderQueryByIdResponse) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "data", thrift.STRUCT, 1)
	offset += p.Data.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderQueryByIdResponse) fastWriteField255(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "baseResp", thrift.STRUCT, 255)
	offset += p.BaseResp.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *ReaderQueryByIdResponse) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("data", thrift.STRUCT, 1)
	l += p.Data.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *ReaderQueryByIdResponse) field255Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("baseResp", thrift.STRUCT, 255)
	l += p.BaseResp.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
