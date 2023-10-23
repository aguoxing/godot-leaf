Leaf server
===========
A game server based on [Leaf framework](https://github.com/name5566/leaf).

Licensing
---------

Leaf server is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/name5566/leafserver/blob/master/LICENSE) for the full license text.



### 注意
1. 把[Leaf](https://github.com/name5566/leaf)源码拉下来放到指定位置
2. 修改leafsrv下的go.mod，替换leaf包位置
```
replace github.com/name5566/leaf v0.0.0-20221021105039-af71eb082cda => ../xxx/leaf
```

```
protoc --go_out=. --go-grpc_out=. *.proto
```