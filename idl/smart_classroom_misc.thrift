include "./reader/reader_api.thrift"
include "./editor/editor_api.thrift"
namespace go sc_misc_api

//定义misc应用中的所有RPC服务
service MiscService {

    //读者登录
    reader_api.ReaderLoginResponse ReaderLogin(1: reader_api.ReaderLoginRequest request)

    //编辑登录
    editor_api.EditorLoginResponse EditorLogin(1: editor_api.EditorLoginRequest request)

}
