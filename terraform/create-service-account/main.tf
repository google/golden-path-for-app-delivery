/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

module "service_accounts" {
  source        = "terraform-google-modules/service-accounts/google"
  project_id    = var.project_id
  prefix        = var.prefix
  names         = ["tf-svc"]
  project_roles = [
    "${var.project_id}=>roles/viewer",
    "${var.project_id}=>roles/storage.admin",
    "${var.project_id}=>roles/artifactregistry.admin",
    "${var.project_id}=>roles/binaryauthorization.attestorsAdmin",
    "${var.project_id}=>roles/cloudbuild.builds.builder",
    "${var.project_id}=>roles/cloudbuild.workerPoolOwner",
    "${var.project_id}=>roles/cloudkms.admin",
    "${var.project_id}=>roles/cloudkms.publicKeyViewer",
    "${var.project_id}=>roles/containeranalysis.notes.editor",
    "${var.project_id}=>roles/compute.networkAdmin",
    "${var.project_id}=>roles/serviceusage.serviceUsageAdmin",
    "${var.project_id}=>roles/source.admin",
    "${var.project_id}=>roles/resourcemanager.projectIamAdmin",
    "${var.project_id}=>roles/viewer",
    "${var.project_id}=>roles/compute.networkAdmin",
    "${var.project_id}=>roles/binaryauthorization.policyEditor",
    "${var.project_id}=>roles/resourcemanager.projectIamAdmin",
    "${var.project_id}=>roles/serviceusage.serviceUsageViewer",
    "${var.project_id}=>roles/iam.serviceAccountAdmin",
    "${var.project_id}=>roles/iam.serviceAccountUser",
    "${var.project_id}=>roles/clouddeploy.developer",
    "${var.project_id}=>roles/container.admin"
  ]
  display_name  = "Terraform Service Account"
  description   = "Terraform Service Account"
}