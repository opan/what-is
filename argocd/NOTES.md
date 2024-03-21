Ref: https://github.com/lucj/k8sapps?tab=readme-ov-file 

Follow installation instruction [here](https://argo-cd.readthedocs.io/en/stable/getting_started/).

Requirements:
- microk8s
- multipass

Notes when running in local via VM:

Make sure to specify `--address` when port-forwarding so that we can access it from browser 

`kubectl port-forward svc/argocd-server -n argocd 8080:443 --address 0.0.0.0`