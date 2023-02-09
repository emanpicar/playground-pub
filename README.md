# Playground
Public collection of random technologies [I](https://ph.linkedin.com/in/eman-picar-9447aa112) play around with. Mostly the result of my "Study and Practice" approach.

<br><br>

## TODO
List of things I plan to add on this repo.

<br>

#### Terraform

---
<br>

#### Kubernetes
1. Deploy a custom image in DockerHub
    - Go server with multiple api to handle
        - plain response
        - GET ServiceAccount user token
        - Create files (volume.txt, persistent-volume.txt)
        - GET all ConfigMaps and Secrets
        - GET/List pods
        - GET/List CRDs
1. Create a deployment using the custom image
1. Create ServiceAccount, Role, RoleBinding
    - for GET/List pods
    - for GET/List CRDs
1. Add a sidecar to the deployment
    - add random header in the rquest
1. Create Volume/PersistentVolume
    - store the files from the custom api accordingly
1. Create ConfigMaps/Secrets kinds
1. Setup Services/Ingress
1. Setup NetworkPolicy for Ingress/Egress
    - setup the ingress policy for the custom api
    - setup the egress policy for the custom api
    - create pods that curls the custom api

Advance Topics
1. Deploy a full ORY setup in the cluster
    - https://k8s.ory.sh/helm/
    - Details TBD

---
<br>

#### Azure DevOps
---
<br>

#### Ansible
---
<br>

#### Terraform/Kubernetes/Ansible/Azure DevOps
---
<br>

#### Go/Kubernetes
1. Create CRD, ConfigMap and Secret
1. Prepare Deployment and Pod yaml for controller
1. Generate the CRD controller
    - CRD create
        - print all ConfigMap/Secret
        - apply the deployment yaml
    - CRD update
        - change values in ConfigMap/Secret
        - delete old deployment
        - apply the deployment yaml
    - CRD delete
        - TBD
---
<br>

#### Go/Swagger
1. Generate Go Client from swagger.json
2. Generate swagger.json from Go (Swaggo)
    - use swaggo annotations
    - serve swagger ui
3. Generate swagger.json from Go (Fizz)
---
<br>