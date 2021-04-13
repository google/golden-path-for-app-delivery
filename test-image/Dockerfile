FROM gcr.io/google.com/cloudsdktool/cloud-sdk:slim
RUN apt-get update && apt-get install -y jq google-cloud-sdk-local-extract
COPY . .
ENTRYPOINT [ "./hack/scan.sh" ]