include "../base/base.thrift"
include "../base/pagination.thrift"
include "enums/enums.thrift"
include "model/model.thrift"
namespace go sc_misc_api

//读者登录 请求体
struct ReaderLoginRequest {
	1: string username (go.tag="json:\"username\" query:\"username\"") //用户名
	2: string password (go.tag="json:\"password\" query:\"password\"") //密码

	255: optional base.Base base //标准请求内容
}

//读者登录 响应体
struct ReaderLoginResponse {
	1: model.ReaderDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容

}
