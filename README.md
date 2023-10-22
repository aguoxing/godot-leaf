服务端运行
1. 把leaf源码拉下来放到指定位置
2. 修改leafsrv下的go.mod，替换leaf包位置
```
replace github.com/name5566/leaf v0.0.0-20221021105039-af71eb082cda => ../xxx/leaf
```

3. 直接运行leafsrv/out/下的server.exe