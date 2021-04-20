FROM gcr.io/cloud-builders/gcloud
RUN apt-get update && apt-get install -y wget jq
RUN wget https://github.com/open-policy-agent/conftest/releases/download/v0.23.0/conftest_0.23.0_Linux_x86_64.tar.gz && \
    tar xzf conftest_0.23.0_Linux_x86_64.tar.gz && \
    mv conftest /usr/local/bin
RUN wget https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv4.1.0/kustomize_v4.1.0_linux_amd64.tar.gz && \
    tar xzf kustomize_v4.1.0_linux_amd64.tar.gz && \
    mv kustomize /usr/local/bin
COPY test/validate.sh validate.sh
COPY test/policy/* ./policy/
COPY frontend/k8s/ /src/frontend/k8s/
COPY backend/k8s/ /src/backend/k8s/
CMD ["sh"]