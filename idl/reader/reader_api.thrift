include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/reader_model.thrift"
namespace go sc_misc_api

//读者登录 请求体
struct ReaderLoginRequest {
	1: string username //用户名
	2: string password //密码

	255: optional base.Base base //标准请求内容
}

//读者登录 响应体
struct ReaderLoginResponse {
	1: reader_model.ReaderDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容

}

//读者查询 请求体
struct ReaderQueryByIdRequest {
	1: i64 readerId //id

	255: optional base.Base base //标准请求内容
}

//读者查询 响应体
struct ReaderQueryByIdResponse {
	1: reader_model.ReaderDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容

}

//读者查询 请求体
struct ReaderQueryByIdsRequest {
	1: list<i64> readerIds //id

	255: optional base.Base base //标准请求内容
}

//读者查询 响应体
struct ReaderQueryByIdsResponse {
	1: list<reader_model.ReaderDTO> data //数据信息

    255: base.BaseResp baseResp //标准返回内容

}
