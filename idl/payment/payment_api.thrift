include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/payment_model.thrift"

namespace go sc_misc_api


//详情 请求体
struct PaymentQueryByIdRequest {
	1: i64 paymentId //Id

    255: optional base.Base base //标准请求内容
}

//详情 响应体
struct PaymentQueryByIdResponse {
	1: payment_model.PaymentDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}



// 请求体
struct PaymentPrepareRequest {
	1: i64 paymentId //Id

    255: optional base.Base base //标准请求内容
}

// 响应体
struct PaymentPrepareResponse {
	1: PaymentPrepareData data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}

struct PaymentPrepareData {
    1: payment_model.PaymentDTO paymentDTO //数据信息
    2: string thirdTransactionNo // 支付的一些token及信息
    3: string nonceStr // 支付时候的验证信息等
}


// 请求体
struct PaymentCreateRequest {
	1: string orderNo //订单编号
	2: string method //支付方式 ALIPAY/ WEIXIN 支付宝/ 微信
	3: i64 amount //金额(单位：分)

    255: optional base.Base base //标准请求内容
}

// 响应体
struct PaymentCreateResponse {
	1: PaymentPrepareData data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}

// 数据体
struct PaymentCreateData {
    1: payment_model.PaymentDTO paymentDTO //数据信息
    2: string thirdTransactionNo // 支付的一些token及信息
    3: string nonceStr // 支付时候的验证信息等
}


// 请求体
struct PaymentPaidCallbackRequest {
    1: string orderNo //订单编号

    255: optional base.Base base //标准请求内容
}

// 响应体
struct PaymentPaidCallbackResponse {
    1: payment_model.PaymentDTO data //数据信息

    255: base.BaseResp baseResp //标准返回内容
}


