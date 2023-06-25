# go-mimsms

A simple, easy to use Go package to interact with [MIM SMS](https://www.mimsms.com/) API, a Bangladeshi SMS Gateway

## Features

-   [x] Send SMS (single text to many recipients)
-   [x] Check Balance
-   [ ] Check Delivery status
-   [x] Error messages returned by package does not leak the API key

## Install

```
go get github.com/raihanul-2k15/go-mimsms
```

## Usage

### Send SMS

```go
import "github.com/raihanul-2k15/go-mimsms/v2/mimsms"

apiKey := "yourapikeyhere"
apiToken := "apitoken"
client := mimsms.NewClient(apiKey, apiToken)

groupId, err := client.SendMessage("09601234567", []string{"01717171717"}, "Hello World")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(groupId)
```

### Check Balance

```go
import "github.com/raihanul-2k15/go-mimsms/v2/mimsms"

apiKey := "yourapikeyhere"
apiToken := "apitoken"
client := mimsms.NewClient(apiKey, apiToken)

balance, err := client.GetBalance()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(balance)
```

## Disclaimer

Author is not affiliated with MIM SMS
