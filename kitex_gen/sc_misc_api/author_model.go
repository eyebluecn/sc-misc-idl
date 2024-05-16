// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_misc_api

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type AuthorDTO struct {
	Id         int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	CreateTime int64  `thrift:"createTime,2" frugal:"2,default,i64" json:"createTime"`
	UpdateTime int64  `thrift:"updateTime,3" frugal:"3,default,i64" json:"updateTime"`
	Username   string `thrift:"username,4" frugal:"4,default,string" json:"username"`
	Realname   string `thrift:"realname,5" frugal:"5,default,string" json:"realname"`
}

func NewAuthorDTO() *AuthorDTO {
	return &AuthorDTO{}
}

func (p *AuthorDTO) InitDefault() {
	*p = AuthorDTO{}
}

func (p *AuthorDTO) GetId() (v int64) {
	return p.Id
}

func (p *AuthorDTO) GetCreateTime() (v int64) {
	return p.CreateTime
}

func (p *AuthorDTO) GetUpdateTime() (v int64) {
	return p.UpdateTime
}

func (p *AuthorDTO) GetUsername() (v string) {
	return p.Username
}

func (p *AuthorDTO) GetRealname() (v string) {
	return p.Realname
}
func (p *AuthorDTO) SetId(val int64) {
	p.Id = val
}
func (p *AuthorDTO) SetCreateTime(val int64) {
	p.CreateTime = val
}
func (p *AuthorDTO) SetUpdateTime(val int64) {
	p.UpdateTime = val
}
func (p *AuthorDTO) SetUsername(val string) {
	p.Username = val
}
func (p *AuthorDTO) SetRealname(val string) {
	p.Realname = val
}

var fieldIDToName_AuthorDTO = map[int16]string{
	1: "id",
	2: "createTime",
	3: "updateTime",
	4: "username",
	5: "realname",
}

func (p *AuthorDTO) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AuthorDTO[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AuthorDTO) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Id = _field
	return nil
}
func (p *AuthorDTO) ReadField2(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreateTime = _field
	return nil
}
func (p *AuthorDTO) ReadField3(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UpdateTime = _field
	return nil
}
func (p *AuthorDTO) ReadField4(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Username = _field
	return nil
}
func (p *AuthorDTO) ReadField5(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Realname = _field
	return nil
}

func (p *AuthorDTO) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AuthorDTO"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AuthorDTO) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AuthorDTO) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("createTime", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *AuthorDTO) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("updateTime", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UpdateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *AuthorDTO) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("username", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Username); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *AuthorDTO) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("realname", thrift.STRING, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Realname); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *AuthorDTO) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuthorDTO(%+v)", *p)

}

func (p *AuthorDTO) DeepEqual(ano *AuthorDTO) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.CreateTime) {
		return false
	}
	if !p.Field3DeepEqual(ano.UpdateTime) {
		return false
	}
	if !p.Field4DeepEqual(ano.Username) {
		return false
	}
	if !p.Field5DeepEqual(ano.Realname) {
		return false
	}
	return true
}

func (p *AuthorDTO) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *AuthorDTO) Field2DeepEqual(src int64) bool {

	if p.CreateTime != src {
		return false
	}
	return true
}
func (p *AuthorDTO) Field3DeepEqual(src int64) bool {

	if p.UpdateTime != src {
		return false
	}
	return true
}
func (p *AuthorDTO) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Username, src) != 0 {
		return false
	}
	return true
}
func (p *AuthorDTO) Field5DeepEqual(src string) bool {

	if strings.Compare(p.Realname, src) != 0 {
		return false
	}
	return true
}