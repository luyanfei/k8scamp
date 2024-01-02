## istio运行结果
将 httpserver 服务以 Istio Ingress Gateway 的形式发布，运行结果如下：
```
luyanfei@hw1:~/k8scamp/httpserver$ curl -H'Host: simple.cncamp.io' 10.106.146.75/ -v
*   Trying 10.106.146.75:80...
*   * TCP_NODELAY set
*   * Connected to 10.106.146.75 (10.106.146.75) port 80 (#0)
*   > GET / HTTP/1.1
*   > Host: simple.cncamp.io
*   > User-Agent: curl/7.68.0
*   > Accept: */*
*   > 
*   * Mark bundle as not supporting multiuse
*   < HTTP/1.1 200 OK
*   < x-version: 1.1
*   < x-request-accept: */*
*   < x-request-user-agent: curl/7.68.0
*   < x-request-x-b3-sampled: 0
*   < x-request-x-b3-spanid: 0d5d824a07117bd5
*   < x-request-x-b3-traceid: ac3a42ce912bc7fc0d5d824a07117bd5
*   < x-request-x-envoy-attempt-count: 1
*   < x-request-x-envoy-decorator-operation: simple.simple.svc.cluster.local:80/*
*   < x-request-x-envoy-internal: true
*   < x-request-x-envoy-peer-metadata: ChQKDkFQUF9DT05UQUlORVJTEgIaAAoaCgpDTFVTVEVSX0lEEgw
*   aCkt1YmVybmV0ZXMKGQoNSVNUSU9fVkVSU0lPThIIGgYxLjEyLjEKvQMKBkxBQkVMUxKyAyqvAwodCgNhcHASF
*   hoUaXN0aW8taW5ncmVzc2dhdGV3YXkKEwoFY2hhcnQSChoIZ2F0ZXdheXMKFAoIaGVyaXRhZ2USCBoGVGlsbGV
*   yCjYKKWluc3RhbGwub3BlcmF0b3IuaXN0aW8uaW8vb3duaW5nLXJlc291cmNlEgkaB3Vua25vd24KGQoFaXN0a
*   W8SEBoOaW5ncmVzc2dhdGV3YXkKGQoMaXN0aW8uaW8vcmV2EgkaB2RlZmF1bHQKMAobb3BlcmF0b3IuaXN0aW8
*   uaW8vY29tcG9uZW50EhEaD0luZ3Jlc3NHYXRld2F5cwofChFwb2QtdGVtcGxhdGUtaGFzaBIKGgg4YzQ4ZDg3N
*   QoSCgdyZWxlYXNlEgcaBWlzdGlvCjkKH3NlcnZpY2UuaXN0aW8uaW8vY2Fub25pY2FsLW5hbWUSFhoUaXN0aW8
*   taW5ncmVzc2dhdGV3YXkKLwojc2VydmljZS5pc3Rpby5pby9jYW5vbmljYWwtcmV2aXNpb24SCBoGbGF0ZXN0C
*   iIKF3NpZGVjYXIuaXN0aW8uaW8vaW5qZWN0EgcaBWZhbHNlChoKB01FU0hfSUQSDxoNY2x1c3Rlci5sb2NhbAo
*   tCgROQU1FEiUaI2lzdGlvLWluZ3Jlc3NnYXRld2F5LThjNDhkODc1LXRtd2JuChsKCU5BTUVTUEFDRRIOGgxpc
*   3Rpby1zeXN0ZW0KXQoFT1dORVISVBpSa3ViZXJuZXRlczovL2FwaXMvYXBwcy92MS9uYW1lc3BhY2VzL2lzdGl
*   vLXN5c3RlbS9kZXBsb3ltZW50cy9pc3Rpby1pbmdyZXNzZ2F0ZXdheQoXChFQTEFURk9STV9NRVRBREFUQRICK
*   gAKJwoNV09SS0xPQURfTkFNRRIWGhRpc3Rpby1pbmdyZXNzZ2F0ZXdheQ==
*   < x-request-x-envoy-peer-metadata-id: router~172.16.214.145~istio-ingressgateway-8c48d
*   875-tmwbn.istio-system~istio-system.svc.cluster.local
*   < x-request-x-forwarded-for: 192.168.0.159
*   < x-request-x-forwarded-proto: http
*   < x-request-x-request-id: 3ca304ac-1d56-478a-8e82-fa357e5afafa
*   < date: Sun, 26 Dec 2021 13:06:37 GMT
*   < content-length: 0
*   < x-envoy-upstream-service-time: 322
*   < server: istio-envoy
*   < 
*   * Connection #0 to host 10.106.146.75 left intact
*
```

## Dockerfile最佳实践
1. 构建较稳定的层的命令先写，容易变化的层的命令后写。
2. COPY好过ADD。

httpserver只要拷贝一个文件，考虑的东西不需要太多。

## 镜像仓库
    Docker官方镜像仓库太难访问了，这次的作业用了阿里云的镜像仓库，具体地址为：
```
registry.cn-shanghai.aliyuncs.com/luyanfei
```

## 本地启动httpserver
```
docker run -d -p8080:8080 registry.cn-shanghai.aliyuncs.com/luyanfei/httpserver:v1.1
```
用curl命令访问：
```
❯ curl -v http://localhost:8080/healthz
*   Trying 127.0.0.1:8080...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /healthz HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 17 Oct 2021 07:16:27 GMT
< Content-Length: 0
<
```

## 查看IP配置
```
❯ docker inspect 160f92bd373b | grep -i pid           
            "Pid": 16840,
            "PidMode": "",
            "PidsLimit": null,
❯ sudo nsenter -t 16840 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
16: eth0@if17: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

## httpserver
### ingress
ingress配置成功后，使用下面的命令测试https:
```
curl -v -H 'Host: httpserver.cncamp.com' https://192.168.0.159:30021/ -k
```
其中IP地址为k8s服务器结点的地址。
