module istio.io/client-go

go 1.15

require (
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/grpc v1.35.0 // indirect
	istio.io/api v0.0.0-20210622202155-0ee0abf2a92d
	istio.io/gogo-genproto v0.0.0-20210113155706-4daf5697332f // indirect
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.15.12
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.15.12
	k8s.io/apimachinery => k8s.io/apimachinery v0.15.12
)
