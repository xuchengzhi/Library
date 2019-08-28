namespace go testone.rpc
namespace java testone.rpc
namespace py testone.rpc

// 测试服务
service RpcService {
 
    // 发起远程调用
    list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
    string testOne(1:string msg,2:string ip)

}
