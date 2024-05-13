include "../enums/enums.thrift"
namespace go sc_misc_api


//订阅模型
struct SubscriptionDTO {
    1: i64 id, //id
    2: i64 createTime, //创建时间
    3: i64 updateTime, //编辑时间
    4: i64 readerId, //读者id
    5: i64 columnId, //专栏id
    6: i64 orderId, //订单id
    7: enums.SubscriptionStatus status, //状态
}
