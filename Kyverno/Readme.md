Reference application taken from [AnaisUrlich/kyverno-example](https://github.com/AnaisUrlichs/kyverno-example)

## Steps:

1. To install kyverno helm chart: `helm repo add kyverno https://kyverno.github.io/kyverno/`

2. Update helm records: `helm repo update`

3. Install Kyverno-policy `helm install kyverno-policies kyverno/kyverno-policies -n kyverno`

4. Check the deployment `kubectl get deploy -n kyverno`

5. Alternate command to install kyverno `helm install kyverno kyverno/kyverno -n kyverno --create-namespace --set replicaSet=3`

6. Take a policy from [Kyverno Policy Tempaltes](https://kyverno.io/policies/) eg: [disallow-latest-tags](https://kyverno.io/policies/best-practices/disallow_latest_tag/disallow_latest_tag/) and paste it into your file named eg `my-policy.yml`

7. Tweak the policy a little bit like changes in following categories under `specs`
  ```yml
  validationFailureAction: enforce
  background: true
  failurePolicy: Fail
  ```
8. Apply the policy to the cluster `kubectl apply -f my-policy.yml`

9. Apply the deployment in this case `kubectl apply -f nginx.yml` now if the nginx image has `:latest` tag in it then you will get an error message on fixing the tag name your deployment will be successfully created
