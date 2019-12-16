### Secret and ConfigMap

#### 创建Secret和ConfigMap对象
```
$ kubectl create secret generic token --from-literal=TOKEN=abcd123456000
```

```
$ kubectl create configmap language --from-literal=LANGUAGE=English
```

#### 更新Secret和ConfigMap

首先更新Secret或者ConfigMap：
```
$ kubectl create configmap language --from-literal=LANGUAGE=Chinese -o yaml --dry-run | kubectl replace -f -
```

```
$ kubectl create secret generic token --from-literal=TOKEN=bbbbb123456 -o yaml --dry-run | kubectl replace -f -
```

然后，重启相关的PODS，重启 POD 有很多方法，比如我们可以直接更新 Deployment，最简单的方法是我们可以直接手动删除 POD，然后 Deployemnt 会自动重建一个 POD起来的。
```
$ kubectl delete pod -l name=envtest
```

