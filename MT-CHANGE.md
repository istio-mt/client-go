## 使用说明
- 当前 kubernetes 集群版本是 1.15.12，依赖库版本需要修改
- 需要使用 kubernetes 1.15.12 版本的 code generator
- 改造点1：将 kubetype-gen 独立出来，使用官方的镜像编译 
```
sudo IMG=gcr.io/istio-testing/build-tools:release-1.9-2021-04-06T17-47-31 make clean generate-kubetype-gen
```
- 改造点2：将 client gen 独立出来，使用修改后的镜像编译
```
sudo IMG=swr.cn-north-4.myhuaweicloud.com/mesh/mt-build-tools-for-istio-config-controller:kubernetes-1.15.12 make generate-k8s-client-without-kubetype-gen
```
- 改造点3：go.mod 更新 k8s.io/api@v0.15.12，k8s.io/client-go@v0.15.12，k8s.io/apimachinery@v0.15.12
