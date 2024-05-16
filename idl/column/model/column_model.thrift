include "../../author/model/author_model.thrift"
include "../../column_quote/model/column_quote_model.thrift"
include "../../subscription/model/subscription_model.thrift"
namespace go sc_bff_api


//专栏模型
struct ColumnDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间
    4: string name //名称
    5: i64 authorId //作者id
    6: i32 status //状态
}


//富专栏模型
struct RickColumnDTO {
    1: ColumnDTO column //专栏本身
    2: author_model.AuthorDTO author //关联的作者
    3: column_quote_model.ColumnQuoteDTO columnQuote //关联的报价
    4: subscription_model.SubscriptionDTO subscription //关联的订阅情况
}