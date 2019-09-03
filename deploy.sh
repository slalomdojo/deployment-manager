#!/bin/bash

node_count=$1
env=""

case $2 in
"null")
  env="none"
  ;;
"development")
  env="dev"
  ;;
"staging")
  env="staging"
  ;;
"production")
  env="prod"
  ;;
*)
  echo "Invalid environment."
  exit 1
  ;;
esac

cd templates/gke
/google-cloud-sdk/bin/gcloud deployment-manager deployments create gke \
    --async \
    --template cluster.py \
    --properties zone:northamerica-northeast1-a,initialNodeCount:$node_count,env:$env \
    >> deployment.txt 2>&1

if [[ $? == 0 ]]; then
  echo "Kubernetes deployed. "
else 
  cat deployment.txt
  exit 1
fi

rm -f deployment.txt

if [[ "$3" == "a database" ]]; then
  cd ../database
  /google-cloud-sdk/bin/gcloud deployment-manager deployments create db \
      --async \
      --template cloudsql.jinja \
      --properties instance_name:$env-sql-$(date +%s | base64 | tail -c6 | tr -d '\n=' | tr '[:upper:]' '[:lower:]') \
      >> deployment.txt 2>&1

  if [[ $? == 0 ]]; then
    echo "Database deployed."
  else 
    cat deployment.txt
    exit 1
  fi
fi
