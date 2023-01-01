#Steps and commands

Create seperate namespace
`kubectl create namespace monitoring`

Install Prometheus operator using helm charts
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack -n monitoring
```

Portforward the request to the service
`kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring`