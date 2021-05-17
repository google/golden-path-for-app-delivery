#
gcloud projects add-iam-policy-binding --member serviceAccount:54902187596@cloudbuild.gserviceaccount.com --role roles/owner vic-e2e-cicd

# Setup KCC
gcloud iam service-accounts create kcc-operator
gcloud projects add-iam-policy-binding vic-e2e-cicd --member="serviceAccount:kcc-operator@vic-e2e-cicd.iam.gserviceaccount.com" --role="roles/owner"
gcloud iam service-accounts add-iam-policy-binding kcc-operator@vic-e2e-cicd.iam.gserviceaccount.com --member="serviceAccount:vic-e2e-cicd.svc.id.goog[cnrm-system/cnrm-controller-manager]" --role="roles/iam.workloadIdentityUser"

kctx gke_vic-e2e-cicd_us-west2_prod
kubectl annotate namespace default cnrm.cloud.google.com/project-id=vic-e2e-cicd
kubectl apply -f configconnector.yaml

kctx gke_vic-e2e-cicd_us-west2_staging
kubectl annotate namespace default cnrm.cloud.google.com/project-id=vic-e2e-cicd
kubectl apply -f configconnector.yaml