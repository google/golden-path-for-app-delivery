FROM alpine:latest
RUN apk add curl grep
RUN curl -s https://api.github.com/repos/kubernetes-sigs/kustomize/releases/latest |\
  grep browser_download |\
  grep linux |\
  cut -d '"' -f 4 |\
  xargs curl -O -L && \
  mv kustomize_*_linux_amd64 kustomize && \
  chmod u+x kustomize
ENTRYPOINT [ "/kustomize" ]