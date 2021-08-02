# Sample App For End to End Golden Path

## Boostrapping your project

1. Enable services

```shell
gcloud services enable sourcerepo.googleapis.com \
                       cloudbuild.googleapis.com \
                       clouddeploy.googleapis.com \
                       container.googleapis.com \
                       redis.googleapis.com \
                       cloudresourcemanager.googleapis.com \
                       servicenetworking.googleapis.com
```

1. Configure Cloud Build to allow modification of Cloud Deploy delivery pipelines and deploy to GKE:

```shell
PROJECT_NUMBER=$(gcloud projects list --filter="$(gcloud config get-value project)" --format="value(PROJECT_NUMBER)")
gcloud projects add-iam-policy-binding --member="serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com" --role roles/clouddeploy.admin $(gcloud config get-value project)
gcloud projects add-iam-policy-binding --member="serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com" --role roles/container.developer $(gcloud config get-value project)
gcloud projects add-iam-policy-binding --member="serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com" --role roles/iam.serviceAccountUser $(gcloud config get-value project)
gcloud projects add-iam-policy-binding --member="serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com" --role roles/clouddeploy.jobRunner $(gcloud config get-value project)
gcloud projects add-iam-policy-binding --member="serviceAccount:${PROJECT_NUMBER}-compute@developer.gserviceaccount.com" --role roles/container.admin $(gcloud config get-value project)
```

1. Create a source repository

```shell
gcloud source repos create sample-app
```

1. Create a Cloud Build trigger for the main branch and a bucket for persisting build artifacts.

```shell
gcloud beta builds triggers create cloud-source-repositories --name="sample-app-master" \
                                                             --repo="sample-app" \
                                                             --branch-pattern="master" \
                                                             --build-config="cloudbuild.yaml"
gsutil mb gs://$(gcloud config get-value project)-gceme-artifacts/
```

1. Create a `staging` GKE Cluster:

```shell
gcloud container clusters create staging \
    --release-channel regular \
    --addons ConfigConnector \
    --workload-pool=$(gcloud config get-value project).svc.id.goog \
    --enable-stackdriver-kubernetes --region us-central1 \
    --enable-ip-alias
```

1. Create a `prod` GKE cluster:
```shell
gcloud container clusters create prod \
    --release-channel regular \
    --addons ConfigConnector \
    --workload-pool=$(gcloud config get-value project).svc.id.goog \
    --enable-stackdriver-kubernetes --region us-central1 \
    --enable-ip-alias
```

1. Configure Config Connector:

```shell
# https://cloud.google.com/config-connector/docs/how-to/install-upgrade-uninstall
gcloud iam service-accounts create sample-app-config-connector
gcloud projects add-iam-policy-binding $(gcloud config get-value project) \
    --member="serviceAccount:sample-app-config-connector@$(gcloud config get-value project).iam.gserviceaccount.com" \
    --role="roles/owner"
gcloud iam service-accounts add-iam-policy-binding \
    sample-app-config-connector@$(gcloud config get-value project).iam.gserviceaccount.com \
    --member="serviceAccount:$(gcloud config get-value project).svc.id.goog[cnrm-system/cnrm-controller-manager]" \
    --role="roles/iam.workloadIdentityUser"

cat > config-connector.yaml <<EOF
apiVersion: core.cnrm.cloud.google.com/v1beta1
kind: ConfigConnector
metadata:
  name: configconnector.core.cnrm.cloud.google.com
spec:
 mode: cluster
 googleServiceAccount: "sample-app-config-connector@$(gcloud config get-value project).iam.gserviceaccount.com"
EOF
kubectl apply -f config-connector.yaml --context gke_$(gcloud config get-value project)_us-central1_staging
kubectl annotate namespace default cnrm.cloud.google.com/project-id=$(gcloud config get-value project) \
                --context gke_$(gcloud config get-value project)_us-central1_staging
kubectl apply -f config-connector.yaml --context gke_$(gcloud config get-value project)_us-central1_prod
kubectl annotate namespace default cnrm.cloud.google.com/project-id=$(gcloud config get-value project) \
                --context gke_$(gcloud config get-value project)_us-central1_prod
# Create default network in both clusters so it can be referenced
cat > default-network.yaml <<EOF
---
apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeNetwork
metadata:
  name: default
spec:
  routingMode: REGIONAL
  autoCreateSubnetworks: true
EOF
kubectl apply -f default-network.yaml --context gke_$(gcloud config get-value project)_us-central1_staging
kubectl apply -f default-network.yaml --context gke_$(gcloud config get-value project)_us-central1_prod
```

1. Create Private Service Access for Redis:
```shell
# https://cloud.google.com/vpc/docs/configure-private-services-access
gcloud compute addresses create sample-app \
    --global \
    --purpose=VPC_PEERING \
    --prefix-length=16 \
    --description="Sample App range" \
    --network=default
gcloud services vpc-peerings connect \
    --service=servicenetworking.googleapis.com \
    --ranges=sample-app \
    --network=default \
    --project=$(gcloud config get-value project)
```

1. Push your source code to the repo:
```shell
git config --global credential.https://source.developers.google.com.helper gcloud.sh
git remote add google https://source.developers.google.com/p/$(gcloud config get-value project)/r/sample-app
git push google $(git branch --show-current):master
```
