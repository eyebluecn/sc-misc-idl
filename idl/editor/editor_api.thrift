include "../base/base.thrift"
include "../base/pagination.thrift"
include "enums/enums.thrift"
include "model/model.thrift"
namespace go sc_misc_api


//编辑登录 请求体
struct EditorLoginRequest {
	1: string username //用户名
	2: string password //密码

	255: optional base.Base base //标准请求内容
}

//编辑登录 响应体
struct EditorLoginResponse {
	1: model.EditorDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容

}
