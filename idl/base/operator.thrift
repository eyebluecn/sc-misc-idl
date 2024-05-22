namespace go sc_misc_base

enum OperatorType {
    READER = 0 //读者
    EDITOR = 1 //小编
}

//操作者：读者
struct Operator {
    1: i64 operatorId //操作者id
    2: OperatorType type //操作者类型
    4: string username //昵称
}
