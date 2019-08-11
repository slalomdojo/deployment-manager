#!/bin/bash

case $1 in
"gke")
  cd templates/gke
  /google-cloud-sdk/bin/gcloud deployment-manager deployments create gke \
     --async \
     --template cluster.py \
     --properties zone:northamerica-northeast1-a \
     >> deployment.txt 2>&1

  if [[ $? == 0 ]]; then
    echo "Ok"
    exit 0
  else 
    echo "Failed."
    exit 1
  fi
  ;;
"vm")
  cd templates/single-vm
  
  for i in `seq 1 $2`; do 
    /google-cloud-sdk/bin/gcloud deployment-manager deployments create vm-$i \
        --async \
        --template vm_template.py \
        --properties zone:northamerica-northeast1-a \
        >> deployment.txt 2>&1
    
    if [[ $? != 0 ]]; then
        echo "Failed."
        exit 1
    fi
  done

  if [[ $? == 0 ]]; then
    echo "Ok"
    exit 0
  else 
    echo "Failed."
    exit 1
  fi
  ;;
*)
  echo "Template unrecognized."
  exit 1
  ;;
esac
