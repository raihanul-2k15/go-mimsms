# go-mimsms

Go package to interact with MIM SMS API, a Bangladeshi SMS Gateway

## Features

-   [x] Send SMS (single text to many recipients)
-   [ ] Send SMS (different text to different recipients)
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
import "github.com/raihanul-2k15/go-mimsms/mimsms"

apiKey := "YOURAPIKEY.HERE"
client := mimsms.NewClient(apiKey)

msgId, err := client.SendMessage("09601234567", []string{"01717171717"}, "Hello World", mimsms.ContentTypeText, mimsms.MessageTypeTransactional)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(msgId)
```

### Check Balance

```go
import "github.com/raihanul-2k15/go-mimsms/mimsms"

apiKey := "YOURAPIKEY.HERE"
client := mimsms.NewClient(apiKey)

balance, err := client.GetBalance()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(balance)
```

## Known Issues

Incorrect error (API Not Found, error code: 1003) is returned in case of wrong API key or expired balance due to bug in the MIM SMS API itself

**Note**: This is not a bug in this package

## Disclaimer

Author is not affiliated with MIM SMS
