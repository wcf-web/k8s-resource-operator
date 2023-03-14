
# operator-sdk

参考：[operator-sdk 操作](https://stackoverflow.com/questions/67585119/fata0004-failed-to-initialize-project-unable-to-scaffold-with-base-go-kubebu)

初始化项目
````
operator-sdk init --domain example.com --repo github.com/example/memcached-operator
````

创建 api
````
operator-sdk create api --group example --version v1alpha1 --kind HelloWorld --resource --controller
````

生成或更新 CRD
````
make manifests
````