Connect to the test cluster.

```
kubectl create -f run-test-server.yaml
kubectl create -f create-svc.yaml
```

There should already be a curl pod. If not, run

```
kubectl run curl-<YOUR NAME> --image=radial/busyboxplus:curl -i --tty --rm
```

Then

```
kubectl attach curl-<YOUR NAME> -i -t
curl -X GET df-server-test:13301/sanitycheck
```