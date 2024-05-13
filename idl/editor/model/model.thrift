include "../enums/enums.thrift"
namespace go sc_misc_api


//编辑模型
struct EditorDTO {
    1: i64 id //id
    2: i64 createTime //创建时间
    3: i64 updateTime //编辑时间
    4: string username //昵称
}
