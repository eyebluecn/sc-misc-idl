include "../base/base.thrift"
include "../base/pagination.thrift"
include "../base/operator.thrift"
include "model/column_model.thrift"
include "enum/column_enums.thrift"
namespace go sc_misc_api

//请求体
struct ColumnOmnibusRequest {
	1: string authorName //作者名
	2: string columnName //专栏名
	3: i64 columnPrice //专栏价格

    255: optional base.Base base //标准请求内容
}

//响应体
struct ColumnOmnibusResponse {

    255: base.BaseResp baseResp //标准返回内容


}

//专栏列表 请求体
struct ColumnPageRequest {
	1: i64 pageNum //当前页码 1基
	2: i64 pageSize //每页大小
	4: optional string name //名称
	5: optional i64 authorId //作者id
	6: optional column_enums.ColumnStatus status //状态
	7: operator.ReaderOperator readerOperator //读者操作者

    255: optional base.Base base //标准请求内容
}

//专栏列表 响应体
struct ColumnPageResponse {
	1: list<column_model.RichColumnDTO> data //数据信息
    2: pagination.Pagination pagination //分页指示器

    255: base.BaseResp baseResp //标准返回内容
}

//专栏详情 请求体
struct ColumnDetailRequest {
	1: i64 id //专栏Id

    255: optional base.Base base //标准请求内容
}

//专栏详情 响应体
struct ColumnDetailResponse {
	1: column_model.ColumnDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}
