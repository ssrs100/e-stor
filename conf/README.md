# 最简单的生成步骤

## 生成私钥

```
openssl genrsa -out ssl.key 2048
```
## 证书请求
```
openssl req -new -key ssl.key -out ssl.csr
```

## 生成证书

```
openssl x509 -req -in ssl.csr -signkey ssl.key -out ssl.crt
```
