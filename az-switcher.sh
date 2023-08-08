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

# Main loop
while true; do
    echo "wELCOME TO AZURE SWITCHER!"
    echo "Choose Actions:"
    echo "1. Display subscription list "
    echo "2. Select a subscription"
    echo "3. Switch to another tenant"
    echo "4. Exit"
    read choice

    case $choice in
        1)
            list_subscriptions
            ;;
        2)
            list_subscriptions
            echo "Enter the ID of the subscription you want to select:"
            read sub_id
            select_subscription "$sub_id"
            ;;
        3)
            echo "Enter the tenant ID you want to use:"
            read tenant_id
            switch_tenant "$tenant_id"
            ;;
        4)
            echo "Thank You,!"
            break
            ;;
        *)
            echo "Invalid choice."
            ;;
    esac
done
