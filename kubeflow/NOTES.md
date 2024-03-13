# Kubeflow

## Install and setup Kubeflow in local environment

### 1. Install Multipass

Follow the [guide](https://github.com/canonical/multipass)

### 2. Create Ubuntu VM

```
# using jammy/22.04
multipass launch jammy -n kubeflow -m 8G -d 40G -c 4
```

### 3. Enter the VM

```
multipass shell kubeflow
```

### 4. Install Microk8s

https://microk8s.io/

```
# will install the latest stable version
sudo snap install microk8s --classic

# switch to root
sudo su

# make alias for kubectl
snap alias microk8s.kubectl kubectl

mkdir -p $HOME/.kube
kubectl config view --raw > $HOME/.kube/config

# wait for k8s cluster become ready
microk8s status --wait-ready

# enabling microk8s services/addon
microk8s enable dns hostpath-storage ingress metallb:10.64.140.43-10.64.140.49 rbac

```

### 5. Install Kubeflow

You can follow the instruction [here](https://charmed-kubeflow.io/docs/get-started-with-charmed-kubeflow).

Or follow more compact version below:

```
# install juju
sudo snap install juju --classic --channel=3.1/stable

# ensure the directory exist
mkdir -p ~/.local/share

# configure microk8s to work properly with juju
microk8s config | juju add-k8s my-k8s --client

# deploy juju controller to the kubernetes we set up with microk8s
juju bootstrap my-k8s uk8sx

# add a model for kubeflow to the controller
juju add-model kubeflow

# ensure inotify limits high enough, kubeflow use this to interact with the filesystem
sudo sysctl fs.inotify.max_user_instances=1280
sudo sysctl fs.inotify.max_user_watches=655360

# to persist the inotify changes after every machine restart, add line below to /etc/sysctl.conf
fs.inotify.max_user_instances=1280
fs.inotify.max_user_watches=655360

# deploy charmed kubeflow
juju deploy kubeflow --trust  --channel=1.8/stable

```

When we get this message `Deploy of bundle completed.`, it means the deploy commands complete


https://kind.sigs.k8s.io/docs/user/quick-start/#installation
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-22-04
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

https://github.com/kubeflow/manifests?tab=readme-ov-file#install-with-a-single-command


kubectl create secret generic regcred \
    --from-file=.dockerconfigjson=/root/.docker/config.json \
    --type=kubernetes.io/dockerconfigjson


alias kustomize="kubectl kustomize"
alias k="kubectl"
source /root/.bashrc

https://github.com/kubeflow/manifests?tab=readme-ov-file#install-individual-components

kubectl wait --for=condition=ready pod -l 'app in (cert-manager,webhook)' --timeout=180s -n cert-manager
kustomize common/cert-manager/kubeflow-issuer/base | kubectl apply -f -

kustomize common/istio-1-17/istio-crds/base | kubectl apply -f -
kustomize common/istio-1-17/istio-namespace/base | kubectl apply -f -
kustomize common/istio-1-17/istio-install/base | kubectl apply -f -


kustomize common/oidc-client/oidc-authservice/base | kubectl apply -f -

kustomize common/dex/overlays/istio | kubectl apply -f -

kustomize common/knative/knative-serving/overlays/gateways | kubectl apply -f -
kustomize common/istio-1-17/cluster-local-gateway/base | kubectl apply -f -

kustomize common/kubeflow-namespace/base | kubectl apply -f -

kustomize common/kubeflow-roles/base | kubectl apply -f -

kustomize common/istio-1-17/kubeflow-istio-resources/base | kubectl apply -f -

kustomize apps/pipeline/upstream/env/cert-manager/platform-agnostic-multi-user | kubectl apply -f -

kustomize contrib/kserve/kserve | kubectl apply -f -

kustomize contrib/kserve/models-web-app/overlays/kubeflow | kubectl apply -f -

kustomize apps/katib/upstream/installs/katib-with-kubeflow | kubectl apply -f -

kustomize apps/centraldashboard/upstream/overlays/kserve | kubectl apply -f -

kustomize apps/admission-webhook/upstream/overlays/cert-manager | kubectl apply -f -
oci-image: docker.io/kubeflownotebookswg/poddefaults-webhook:v1.8.0: https://github.com/canonical/admission-webhook-operator/releases


kustomize apps/jupyter/notebook-controller/upstream/overlays/kubeflow | kubectl apply -f -

kustomize apps/jupyter/jupyter-web-app/upstream/overlays/istio | kubectl apply -f -

kustomize apps/pvcviewer-controller/upstream/default | kubectl apply -f -

kustomize apps/profiles/upstream/overlays/kubeflow | kubectl apply -f -

kustomize apps/volumes-web-app/upstream/overlays/istio | kubectl apply -f -

kustomize apps/tensorboard/tensorboards-web-app/upstream/overlays/istio | kubectl apply -f -

kustomize apps/tensorboard/tensorboard-controller/upstream/overlays/kubeflow | kubectl apply -f -

kustomize apps/training-operator/upstream/overlays/kubeflow | kubectl apply -f -

kustomize common/user-namespace/base | kubectl apply -f -

kubectl get pods -n cert-manager
kubectl get pods -n istio-system
kubectl get pods -n auth
kubectl get pods -n knative-eventing
kubectl get pods -n knative-serving
kubectl get pods -n kubeflow
kubectl get pods -n kubeflow-user-example-com

s-go-sy-scp-test-kubeflow-a-custom-01