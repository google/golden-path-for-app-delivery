FROM gcr.io/cloud-builders/gcloud
RUN apt-get update && apt-get install -y wget jq
RUN wget https://github.com/open-policy-agent/conftest/releases/download/v0.23.0/conftest_0.23.0_Linux_x86_64.tar.gz && \
    tar xzf conftest_0.23.0_Linux_x86_64.tar.gz && \
    mv conftest /usr/local/bin
RUN gcloud -q components install kustomize skaffold kpt
COPY test/validate.sh validate.sh
COPY test/policy/* ./policy/
COPY frontend/k8s/ frontend/k8s/
COPY backend/k8s/ backend/k8s/
CMD ["sh"]