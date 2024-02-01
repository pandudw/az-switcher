package main

import (
    "fmt"
    "os"
    "os/exec"
)

var subscriptionSelected bool

func listSubscriptions() {
    cmd := exec.Command("az", "account", "list", "--output", "yaml")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func selectSubscription(subscriptionID string) {
    cmd := exec.Command("az", "account", "set", "--subscription", subscriptionID)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Failed to select subscription: %v\n", err)
        return
    }

    subscriptionSelected = true
    fmt.Printf("Subscription has been changed to %s\n", subscriptionID)
}

func main() {
    for {
        fmt.Println("WELCOME TO AZURE SWITCHER!")
        fmt.Println("Choose Actions:")
        fmt.Println("1. Display subscription list")
        fmt.Println("2. Select a subscription")
        fmt.Println("3. Exit")

        var choice string
        fmt.Scanln(&choice)

        switch choice {
        case "1":
            listSubscriptions()
        case "2":
            listSubscriptions()
            var subID string
            fmt.Print("Enter the ID of the subscription you want to select: ")
            fmt.Scanln(&subID)
            selectSubscription(subID)
            if subscriptionSelected {
                return 
            }
        case "3":
            fmt.Println("Thank You!")
            if subscriptionSelected {
                return 
            } else {
                fmt.Println("Please select a subscription before exiting.")
            }
        default:
            fmt.Println("Invalid choice.")
        }
    }
}
