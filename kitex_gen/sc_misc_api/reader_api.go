// Code generated by thriftgo (0.3.12). DO NOT EDIT.

package sc_misc_api

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_base"
	"strings"
)

type ReaderLoginRequest struct {
	Username string             `thrift:"username,1" frugal:"1,default,string" json:"username" query:"username"`
	Password string             `thrift:"password,2" frugal:"2,default,string" json:"password" query:"password"`
	Base     *sc_misc_base.Base `thrift:"base,255,optional" frugal:"255,optional,sc_misc_base.Base" json:"base,omitempty"`
}

func NewReaderLoginRequest() *ReaderLoginRequest {
	return &ReaderLoginRequest{}
}

func (p *ReaderLoginRequest) InitDefault() {
	*p = ReaderLoginRequest{}
}

func (p *ReaderLoginRequest) GetUsername() (v string) {
	return p.Username
}

func (p *ReaderLoginRequest) GetPassword() (v string) {
	return p.Password
}

var ReaderLoginRequest_Base_DEFAULT *sc_misc_base.Base

func (p *ReaderLoginRequest) GetBase() (v *sc_misc_base.Base) {
	if !p.IsSetBase() {
		return ReaderLoginRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *ReaderLoginRequest) SetUsername(val string) {
	p.Username = val
}
func (p *ReaderLoginRequest) SetPassword(val string) {
	p.Password = val
}
func (p *ReaderLoginRequest) SetBase(val *sc_misc_base.Base) {
	p.Base = val
}

var fieldIDToName_ReaderLoginRequest = map[int16]string{
	1:   "username",
	2:   "password",
	255: "base",
}

func (p *ReaderLoginRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *ReaderLoginRequest) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderLoginRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderLoginRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Username = _field
	return nil
}
func (p *ReaderLoginRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Password = _field
	return nil
}
func (p *ReaderLoginRequest) ReadField255(iprot thrift.TProtocol) error {
	_field := sc_misc_base.NewBase()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}

func (p *ReaderLoginRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ReaderLoginRequest"); err != nil {
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
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
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

func (p *ReaderLoginRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ReaderLoginRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("password", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Password); err != nil {
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

func (p *ReaderLoginRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *ReaderLoginRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReaderLoginRequest(%+v)", *p)

}

func (p *ReaderLoginRequest) DeepEqual(ano *ReaderLoginRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Username) {
		return false
	}
	if !p.Field2DeepEqual(ano.Password) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *ReaderLoginRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Username, src) != 0 {
		return false
	}
	return true
}
func (p *ReaderLoginRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Password, src) != 0 {
		return false
	}
	return true
}
func (p *ReaderLoginRequest) Field255DeepEqual(src *sc_misc_base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type ReaderLoginResponse struct {
	Data     *ReaderDTO             `thrift:"data,1" frugal:"1,default,ReaderDTO" json:"data"`
	BaseResp *sc_misc_base.BaseResp `thrift:"baseResp,255" frugal:"255,default,sc_misc_base.BaseResp" json:"baseResp"`
}

func NewReaderLoginResponse() *ReaderLoginResponse {
	return &ReaderLoginResponse{}
}

func (p *ReaderLoginResponse) InitDefault() {
	*p = ReaderLoginResponse{}
}

var ReaderLoginResponse_Data_DEFAULT *ReaderDTO

func (p *ReaderLoginResponse) GetData() (v *ReaderDTO) {
	if !p.IsSetData() {
		return ReaderLoginResponse_Data_DEFAULT
	}
	return p.Data
}

var ReaderLoginResponse_BaseResp_DEFAULT *sc_misc_base.BaseResp

func (p *ReaderLoginResponse) GetBaseResp() (v *sc_misc_base.BaseResp) {
	if !p.IsSetBaseResp() {
		return ReaderLoginResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}
func (p *ReaderLoginResponse) SetData(val *ReaderDTO) {
	p.Data = val
}
func (p *ReaderLoginResponse) SetBaseResp(val *sc_misc_base.BaseResp) {
	p.BaseResp = val
}

var fieldIDToName_ReaderLoginResponse = map[int16]string{
	1:   "data",
	255: "baseResp",
}

func (p *ReaderLoginResponse) IsSetData() bool {
	return p.Data != nil
}

func (p *ReaderLoginResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *ReaderLoginResponse) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ReaderLoginResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ReaderLoginResponse) ReadField1(iprot thrift.TProtocol) error {
	_field := NewReaderDTO()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Data = _field
	return nil
}
func (p *ReaderLoginResponse) ReadField255(iprot thrift.TProtocol) error {
	_field := sc_misc_base.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResp = _field
	return nil
}

func (p *ReaderLoginResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ReaderLoginResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
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

func (p *ReaderLoginResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Data.Write(oprot); err != nil {
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

func (p *ReaderLoginResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("baseResp", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *ReaderLoginResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReaderLoginResponse(%+v)", *p)

}

func (p *ReaderLoginResponse) DeepEqual(ano *ReaderLoginResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Data) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResp) {
		return false
	}
	return true
}

func (p *ReaderLoginResponse) Field1DeepEqual(src *ReaderDTO) bool {

	if !p.Data.DeepEqual(src) {
		return false
	}
	return true
}
func (p *ReaderLoginResponse) Field255DeepEqual(src *sc_misc_base.BaseResp) bool {

	if !p.BaseResp.DeepEqual(src) {
		return false
	}
	return true
}