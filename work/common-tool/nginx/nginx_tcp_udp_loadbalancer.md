## tcp udp load balancer
https://docs.nginx.com/nginx/admin-guide/load-balancer/tcp-udp-load-balancer/

nginx plus r5 支持 tcp
nginx plus r9 支持 udp

stream block 定义tcp、udp反向代理
tce是默认的，udp需要指定
```conf
stream {
    upstream stream_backend {
        least_conn;  # 最少连接 loadbalance
        server backend1.example.com:12345 weight=5;
        server backend2.example.com:12345 max_fails=2 fail_timeout=30s;
        server backend3.example.com:12345 max_conns=3;
    }

    upstream dns_servers {
        least_conn;  # 最少连接 loadbalance
        server 192.168.136.130:53;
        server 192.168.136.131:53;
        server 192.168.136.132:53;
    }

    server {
        listen        12345;   # 默认是tcp连接
        proxy_pass    stream_backend;
        proxy_timeout 3s;
        proxy_connect_timeout 1s;
    }

    server {
        listen     53 udp;    # udp连接
        proxy_pass dns_servers;
    }

    server {
        listen     12346;   # 默认是tcp连接、Round Robin(轮询)负载均衡
        proxy_pass backend4.example.com:12346;
    }
}
```