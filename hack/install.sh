# Create 4 clusters spinnaker,dev,staging,prod
gcloud beta container clusters create --machine-type n1-standard-4 --zone us-central1-f --enable-stackdriver-kubernetes --cluster-version 1.10 --async production

# Get tokens from each cluster
gcloud iam service-accounts create spinnaker-gcs --display-name "Spinnaker"
gcloud projects add-iam-policy-binding vic-cd-demo --member=serviceAccount:spinnaker-gcs@vic-cd-demo.iam.gserviceaccount.com --role 'roles/storage.admin'
gcloud iam service-accounts keys create --iam-account spinnaker-gcs@vic-cd-demo.iam.gserviceaccount.com ~/Downloads/vic-cd-demo-spinnaker-gcs.json

CONTEXT=spinnaker
kubectl create ns spinnaker --context $CONTEXT
kubectl apply -f ~/go/src/sample-app/hack/spinnaker-sa.yaml --context $CONTEXT
TOKEN=$(kubectl get secret --context $CONTEXT \
   $(kubectl get serviceaccount spinnaker-service-account \
       --context $CONTEXT \
       -n spinnaker \
       -o jsonpath='{.secrets[0].name}') \
   -n spinnaker \
   -o jsonpath='{.data.token}' | base64 --decode)
kubectl config set-credentials ${CONTEXT}-token-user --token $TOKEN
kubectl config set-context $CONTEXT --user ${CONTEXT}-token-user


# Get GCS keys

# install helm with RBAC
kubectl create serviceaccount tiller --namespace kube-system
kubectl create clusterrolebinding tiller --clusterrole=cluster-admin     --serviceaccount=kube-system:tiller 
helm init --service-account=tiller 

