include "./reader/reader_api.thrift"
include "./editor/editor_api.thrift"
include "./column/column_api.thrift"
include "./payment/payment_api.thrift"
namespace go sc_misc_api

//定义misc应用中的所有RPC服务
service MiscService {


    //**************** 专栏START ****************//
    //演示接口 创建专栏等相关所有实体
    column_api.ColumnOmnibusResponse ColumnOmnibus(1: column_api.ColumnOmnibusRequest request)

    //查看专栏列表
    column_api.RichColumnPageResponse RichColumnPage(1: column_api.RichColumnPageRequest request)

    //查看专栏详情
    column_api.ColumnQueryByIdResponse ColumnQueryById(1: column_api.ColumnQueryByIdRequest request)


    //**************** 专栏END ****************//


    //**************** 小编START ****************//

    //小编登录
    editor_api.EditorLoginResponse EditorLogin(1: editor_api.EditorLoginRequest request)



    //**************** 小编END ****************//



    //**************** 读者START ****************//

    //读者登录
    reader_api.ReaderLoginResponse ReaderLogin(1: reader_api.ReaderLoginRequest request)

    //查询读者
    reader_api.ReaderQueryByIdResponse ReaderQueryById(1: reader_api.ReaderQueryByIdRequest request)


    //**************** 读者END ****************//


    //**************** 支付单START ****************//

    //根据id查询支付单
    payment_api.PaymentQueryByIdResponse PaymentQueryById(1: payment_api.PaymentQueryByIdRequest request)

    //获取一个支付单对应的支付准备物料等信息
    payment_api.PaymentPrepareResponse PaymentPrepare(1: payment_api.PaymentPrepareRequest request)

    //创建一个支付单同时返回支付准备物料等信息
    payment_api.PaymentCreateResponse PaymentCreate(1: payment_api.PaymentCreateRequest request)

    //第三方支付平台，支付成功后的回调接口。
    payment_api.PaymentPaidCallbackResponse PaymentPaidCallback(1: payment_api.PaymentPaidCallbackRequest request)


    //**************** 支付单END ****************//

}
