namespace go sc_misc_base

//操作者：读者
struct ReaderOperator {
    1: i64 readerId //读者id
    4: string username //昵称
}
//操作者：小编
struct EditorOperator {
    1: i64 editorId //小编id
    4: string username //昵称
}