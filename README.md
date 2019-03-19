> Author:Cmpeax
>
> Build Date:2018-3-19

## 这是一个和下位机进行通讯的socket框架.

## HOW TO BUILD?
```
git clone https://github.com/cmpeax/go-socket-framework.git
vim ~/.bash_profile //设置环境变量
//追加 gopath 为 当前pwd路径 + '/go-socket-framework'
source ~/.bash_profile
cd go-socket-framework/server/src/cmpeax.tech
govendor sync
```
### 2018-3-19
> * 对Socket进行简单工厂模式封装。 
> * 该框架进行了map映射字符串匹配回调函数，以return键值对的方式进行中间件的添加。
