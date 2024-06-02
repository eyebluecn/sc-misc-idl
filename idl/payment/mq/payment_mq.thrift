include "../enum/payment_enums.thrift"
include "../model/payment_model.thrift"
namespace go sc_misc_api



//支付单mq事件
enum PaymentMqEvent {
    PAYMENT_CREATED = 0, //已创建
    PAYMENT_PAID = 1, //已支付
    PAYMENT_CLOSED = 2, //已关闭
}


//支付单mq消息内容
struct PaymentMqPayload {
    1: payment_model.PaymentDTO paymentDTO //支付单
    2: PaymentMqEvent event //mq事件
    3: i64 occurTime //消息发生时间
}
