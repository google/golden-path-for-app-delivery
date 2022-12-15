### TODO:
* Create project using project factory with APIs enabled and service account provisioned
* enable APIs in terraform
* automate variable replacement in main.tf
* spin up staging and prod clusters in loop
* make it driven from on main.tf and just git clone down the modules

### Configure project
1. `cd ./terraform/configure-project`
1. `terraform init`
1. `terraform plan -var="project_id=cicd-tf"`
1. `terraform apply -var="project_id=cicd-tf"`

### GKE Autopilot Steps
1. Download the GKE Autopilot Terraform and update the files to properly configure the staging and prod clusters:
    ```
    cd terraform
    git clone https://github.com/terraform-google-modules/terraform-google-kubernetes-engine
    cp -r terraform-google-kubernetes-engine/examples/simple_autopilot_public gke-staging
    cp -r terraform-google-kubernetes-engine/examples/simple_autopilot_public gke-prod
    cp gke-staging-main.tf gke-staging/main.tf
    cp gke-prod-main.tf gke-prod/main.tf
    ```
1. In both the gke-staging and gke-prod directories, edit the `outputs.tf` file and make sure the `output "ca_certificate"` section looks like the following:
    ```
    output "ca_certificate" {
      description = "The cluster ca certificate (base64 encoded)"
      sensitive   = true
      value       = module.gke.ca_certificate
    }
    ```
1. In both the gke-staging and gke-prod directories, do the following:
  * `terraform init`
  * `terraform plan -var="project_id=cicd-tf"` and press Enter
  * `terraform apply -var="project_id=cicd-tf"` and type in `yes` at the prompt

### Cloud Build Steps
1. Return to the `golden-path-for-app-delivery/terraform` directory
1. `git clone https://github.com/GoogleCloudPlatform/terraform-google-secure-cicd`
1. `cp -r terraform-google-secure-cicd/examples/app_cicd cicd`
1. `cd cicd`
1. Create a file called `terraform.tfvars` with the following contents:
    ```
    project_id = "<PROJECT_ID>"
    primary_location = "us-central1"
    deploy_branch_clusters  = {
      staging = {
        cluster               = "staging-cluster",
        project_id            = "cicd-tf",
        location              = "us-central1",
        required_attestations = ["projects/cicd-tf/attestors/build-attestor"]
        env_attestation       = "projects/cicd-tf/attestors/security-attestor"
        next_env              = ""
      },
    }
    ```
1. Edit `main.tf` and set the sources to be:
  * `../terraform-google-secure-cicd/modules/secure-ci`
  * `../terraform-google-secure-cicd/modules/secure-cd`
1. Create a service account with proper permission to run the Terraform:
    ```
    gcloud iam service-accounts create cicd-tf-svc-00
    gcloud iam service-accounts keys create cicd-tf-svc-00-credentials.json \
    --iam-account=cicd-tf-svc-00@cicd-tf.iam.gserviceaccount.com
    gcloud auth application-default login
    declare -a roles=("roles/storage.admin"
      "roles/artifactregistry.admin"
      "roles/binaryauthorization.attestorsAdmin"
      "roles/cloudbuild.builds.builder"
      "roles/cloudbuild.workerPoolOwner"
      "roles/cloudkms.admin"
      "roles/cloudkms.publicKeyViewer"
      "roles/containeranalysis.notes.editor"
      "roles/compute.networkAdmin"
      "roles/serviceusage.serviceUsageAdmin"
      "roles/source.admin"
      "roles/resourcemanager.projectIamAdmin"
      "roles/viewer"
      "roles/compute.networkAdmin"
      "roles/binaryauthorization.policyEditor"
      "roles/resourcemanager.projectIamAdmin"
      "roles/serviceusage.serviceUsageViewer"
      "roles/iam.serviceAccountUser"
      "roles/clouddeploy.developer")
    for role in "${roles[@]}"
    do
    gcloud projects add-iam-policy-binding cicd-tf \
      --member="serviceAccount:cicd-tf-svc-00@cicd-tf.iam.gserviceaccount.com" \
      --role=$role
    done
    ```
1. `export GOOGLE_APPLICATION_CREDENTIALS=$PWD/cicd-tf-svc-00-credentials.json`
1. `gcloud auth application-default login`
1. `terraform init`
1. `terraform plan` and press Enter
1. `terraform apply` and type in `yes` at the prompt