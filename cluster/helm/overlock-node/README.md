```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.18.2/cert-manager.crds.yaml
kubectl -n overlock create secret generic keys --from-file=priv_validator_key.json
helm -n overlock upgrade -i overlock-node --create-namespace ./
```