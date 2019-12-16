### Ingress 服务访问

需要在本地/etc/hosts中定义Ingress中定义的域名：

你的k8s集群节点IP goappk8s.local

获取ingress-nginx服务的ip

```
$kubectl get svc -n ingress-nginx
NAME            TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx   NodePort   10.109.210.224   <none>        80:30321/TCP,443:31864/TCP   3h35m
```

然后，在浏览器中访问连接http://goappk8s.local:30321/ping，即可在页面上发现已经打印出PONG