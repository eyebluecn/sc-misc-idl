include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/column_quote_model.thrift"
namespace go sc_misc_api

// 请求体
struct ColumnQuoteQueryByColumnIdRequest {
	1: i64 columnId //专栏Id

    255: optional base.Base base //标准请求内容
}

// 响应体
struct ColumnQuoteQueryByColumnIdResponse {
	1: column_quote_model.ColumnQuoteDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}
