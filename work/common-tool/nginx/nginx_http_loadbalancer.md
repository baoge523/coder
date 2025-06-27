## load balance

https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/

### http load balancing
upstream backend 定义一个服务组 叫做 backend； 并通过proxy_pass 引用
```conf
http {
    upstream backend {
        server backend1.example.com;
        server backend2.example.com;
        server 192.0.0.1 backup;
    }

    server {
        location / {
            proxy_pass http://backend;
        }
    }
}
```
#### Choosing a Load-Balancing Method
1、Round Robin 轮询 (默认)
可以加权重，不变成加权轮询
```conf
upstream backend {
   # no load balancing method is specified for Round Robin
   server backend1.example.com;
   server backend2.example.com;
}
```
2、Least Connections 最少连接
也可以配置权重
```conf
upstream backend {
    least_conn;
    server backend1.example.com;
    server backend2.example.com;
}
```
3、IP Hash 
```conf
upstream backend {
    ip_hash;
    server backend1.example.com;
    server backend2.example.com;
}
```
4、Random
random two least_time=last_byte; 表示随机选择两个，然后选择响应时间最短的那个
```conf
upstream backend {
    random two least_time=last_byte;
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com;
    server backend4.example.com;
}
```
Server Weights
显示设置weight=5，其他的weight=1
```conf
upstream backend {
    server backend1.example.com weight=5;
    server backend2.example.com;
    server 192.0.0.1 backup;
}
```

