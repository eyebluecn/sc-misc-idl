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
    1: string topic //主题
    2: string tags //mq事件
    3: string keys //事件标识
    4: PaymentMqEvent event //mq事件
    5: i64 occurTime //消息发生时间
    6: payment_model.PaymentDTO paymentDTO //支付单
}
