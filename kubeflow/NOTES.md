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

```
