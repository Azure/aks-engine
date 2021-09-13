#!/bin/bash

if [ -z "$RESOURCE_GROUP" ]; then
    echo "must provide a RESOURCE_GROUP env var"
    exit 1;
fi

# Continually look for non-Succeeded VMSS instances
while true; do
  NUM_VMSS=0
  NUM_TERMINAL_VMSS=0
  for VMSS in $(az vmss list -g $RESOURCE_GROUP | jq -r '.[] | .name'); do
    ((NUM_VMSS++))
    for TERMINAL_VMSS in $(az vmss show -g $RESOURCE_GROUP -n $VMSS | jq -r '. | select(.provisioningState != "Updating") | .name'); do
      ((NUM_TERMINAL_VMSS++))
      echo VMSS $TERMINAL_VMSS is in a terminal state!
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[].resources[] | select(.name == "vmssCSE" and .provisioningState != "Succeeded" and .provisioningState != "Creating" and .provisioningState != "Deleting") | .id' | awk -F'/' '{print $9}'); do
        echo Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_}; then
          exit 1
        fi
      done
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[].resources[] | select(.publisher == "Microsoft.AKS" and .provisioningState != "Succeeded" and .provisioningState != "Creating" and .provisioningState != "Deleting") | .id' | awk -F'/' '{print $9}'); do
        echo Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_}; then
          exit 1
        fi
      done
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[] | select(.provisioningState == "Failed") | .name'); do
        echo Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_} --no-wait; then
          exit 1
        fi
      done
    done
  done
  if [ "$LOOP_FOREVER" == "true" ]; then
    sleep 300
  else
    if [[ "${NUM_VMSS}" == "${NUM_TERMINAL_VMSS}" ]]; then
      exit 0
    fi
  fi
done