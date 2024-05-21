# Azctl

Azctl is a simple tool designed to streamline the management of Azure subscriptions, allowing users to switch between subscriptions easily and quickly.

## Features

- Display a list of Azure subscriptions.
- Select a specific subscription to set it as the active subscription.

## Prerequisites

- [Go](https://go.dev/doc/install) installed on your machine.
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli) installed and configured.

## Dependencies

This project uses the following Go packages:
- [cobra](https://github.com/spf13/cobra): A commander for modern Go CLI interactions.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/pandudw/azctl.git
   ```
2. Change into the project directory:
   ```
    cd azctl
    ```
3. Build binary file azctl and move to /usr/local/bin/:
    ```
    go build -o azctl main.go
    sudo mv azctl /usr/local/bin/
    export PATH=$PATH:/usr/local/bin
    source ~/.bashrc
    ```
4. Display the list of Azure subscriptions:
    ```
    azctl list
    ```
5. Select a specific subscription:
    ```
    azctl select [subs_id]
    ```

## Contributing
Feel free to contribute to azctl by opening issues or submitting pull requests. Your contributions are highly appreciated!
