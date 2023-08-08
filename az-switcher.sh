#!/bin/bash

# Display subscription list
list_subscriptions() {
    az account list --output table
}

# Select a subscription
select_subscription() {
    subscription_id="$1"
    az account set --subscription "$subscription_id"
    echo "Subscription has been changed to $subscription_id"
}

# Switch to another tenant
switch_tenant() {
    tenant_id="$1"
    az login --tenant "$tenant_id"
    echo "Berhasil masuk ulang dengan tenant: $tenant_id"
}

