### TODO:
* Get CI/CD pipeline working with tutorial app instead of bank of anthos

### Configure project
1. `cd ./terraform/configure-project`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf"`
1. `terraform apply -var="project_id=cicd-tf"`

### Create terraform service account
1. `cd ../create-service-account`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf"`
1. `terraform apply -var="project_id=cicd-tf"`

### Configure gcloud to use service account to run terraform steps
1. `cd ..` (you should now be in the terraform directory)
1. `gcloud iam service-accounts keys create tf-svc-credentials.json --iam-account=tf-svc@cicd-tf.iam.gserviceaccount.com`
1. Set your ADC so Terraform runs as the service account
    ```
    export GOOGLE_APPLICATION_CREDENTIALS=$PWD/tf-svc-credentials.json
    gcloud auth application-default login
    ```

### GKE Autopilot Steps
1. `cd ../gke-staging`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf"`
1. `terraform apply -var="project_id=cicd-tf"`
1. `cd ../gke-prod`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf"`
1. `terraform apply -var="project_id=cicd-tf"`

### Create CI/CD pipeline
1. `cd ../cicd`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf" -var="primary_location=us-central1"`
1. `terraform apply -var="project_id=cicd-tf" -var="primary_location=us-central1"`