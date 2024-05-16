include "../base/base.thrift"
include "../base/pagination.thrift"
include "model/payment_model.thrift"
namespace go sc_bff_api


// 请求体
struct PaymentPaidCallbackRequest {
    1: string orderNo //订单编号

    255: optional base.Base base //标准请求内容
}

// 响应体
struct PaymentPaidCallbackResponse {
    1: payment_model.PaymentDTO payment //数据信息

    255: base.BaseResp baseResp //标准返回内容
}
