Connect to the test cluster.
Update `test-server-deployment.yaml` with your DF image. 

```
kubectl create -f test-server-deployment.yaml
kubectl create -f test-server-svc.yaml
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