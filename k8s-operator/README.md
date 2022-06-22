# Mysql Operator
- `Mysql`이름의 CRD와 Operator 생성 예제입니다.  
- 실제로 어떤 동작을 하는 것이 아닌 단순히 `Mysql` Custom Resource로 mysql 파드를 띄우는 연습을 위한 프로젝트입니다.  
- 블로그 주소 : [Operator](https://velog.io/@harryroh2003/Operator)

## Prerequisite
- operator-sdk: v1.22.0
- golang: v1.18.3
- kubernetes cluster

## Getting Started
사전에 쿠버네티스 클러스터를 준비해주세요. mac과 window는 docker-desktop을, linux는 on-premise, minikube, kind 등의 방법으로 클러스터를 생성합니다.

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/k8s-operator:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/k8s-operator:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)