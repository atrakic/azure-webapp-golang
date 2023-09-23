#!/usr/bin/env bash

set -e
set -o pipefail

# Deploy Azure infrastructure for web apps Linux containers.
#
# Usage:
# \curl -sSL https://raw.githubusercontent.com/atrakic/azure-webapp-deploy/main/infra/setup.sh | \
#  WEB_APP_NAME=app-$RANDOM IMAGE_NAME=ghcr.io/atrakic/hello-chi:latest bash -s

WEB_APP_NAME=${WEB_APP_NAME:?"You need to configure the WEB_APP_NAME environment variable; eg. app-$RANDOM"}

LOCATION_NAME=${LOCATION_NAME:-westeurope}
APP_SERVICE_PLAN_NAME=${APP_SERVICE_PLAN_NAME:-MyPlan}
RESOURCE_GROUP_NAME=${RESOURCE_GROUP_NAME:-rg-$WEB_APP_NAME}
IMAGE_NAME=${IMAGE_NAME:-docker.io/nginx:latest}
SKU_NAME=${SKU_NAME:-S1}
USERNAME=${USERNAME:-}
PASSWORD=${PASSWORD:-}

main() {
  # https://learn.microsoft.com/en-us/cli/azure/group?view=azure-cli-latest
  az group create --location "$LOCATION_NAME" --name "$RESOURCE_GROUP_NAME"

  # https://learn.microsoft.com/en-us/cli/azure/appservice/plan?view=azure-cli-latest#az-appservice-plan-create()
  az appservice plan create --name "$APP_SERVICE_PLAN_NAME" --resource-group "$RESOURCE_GROUP_NAME" --sku "$SKU_NAME" --is-linux
  prState=''
  while [[ $prState != 'Succeeded' ]];
  do
      prState=$(az appservice plan show "$APP_SERVICE_PLAN_NAME" --resource-group "$RESOURCE_GROUP_NAME" --query 'provisioningState' -o tsv)
      echo "appservice plan $APP_SERVICE_PLAN_NAME provisioningState=$prState"
      sleep 5
  done

  # https://learn.microsoft.com/en-us/cli/azure/webapp?view=azure-cli-latest
  declare -a ARGS=()
  if [[ -n "${USERNAME}" && -n "${PASSWORD}" ]]; then
    ARGS=(-s "${USERNAME}" -w "${PASSWORD}")
  fi
  az webapp create --name "$WEB_APP_NAME" --resource-group "$RESOURCE_GROUP_NAME" --plan "$APP_SERVICE_PLAN_NAME" -i "$IMAGE_NAME" "${ARGS[@]}"
}
main "$@"
