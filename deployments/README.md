## Deploy to a k8s cluster

```
TEST_NAMESPACE="broken-image-reg"
kubectl create ns "${TEST_NAMESPACE}"

# Create a secret with your TLS certs
kubectl -n "${TEST_NAMESPACE}" create secret tls tls-cert \
--cert <path to complete chain of certs> \
--key <path to private key>

kubectl -n "${TEST_NAMESPACE}" apply -f deployment.yaml

kubectl -n "${TEST_NAMESPACE}" apply -f service.yaml

# Get the ClusterIP for service
kubectl -n "${TEST_NAMESPACE}" get svc/broken-reg -o=jsonpath='{.spec.clusterIPs[0]}'

# Give that IP a hostname corresponding to your TLS certs
... depends on DNS provider ...

# Attempts to use the registry within the cluster will be broken or otherwise :)
```
