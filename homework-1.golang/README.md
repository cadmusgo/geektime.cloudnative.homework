# Homework-1.golang

## 作業要求

1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 
4. 当访问 localhost/healthz 时，应返回200



## survy

1. 要求一
   - [ ] 如何讀取 request header
   - [ ] 如何寫入 response header
2. 要求二
   
   - [ ] 如何讀取環境變量
3. 要求三
   - [ ] 如何讀取 client ip
   - [ ] 如何 log 訪問日誌
4. 要求四
   
   - [ ] 如何回應 http 200
   
     

## 擴展需求
1. 如何透過 goroutine 發動 reqeust
2. 如何透過 goroutine 調用 shell command