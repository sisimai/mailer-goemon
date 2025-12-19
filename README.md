[![License](https://img.shields.io/badge/license-BSD%202--Clause-orange.svg)](https://github.com/sisimai/mailer-goemon/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sisimai/mailer-goemon)
[![Go Reference](https://pkg.go.dev/badge/libsisimai.org/mailer-goemon.svg)](https://pkg.go.dev/libsisimai.org/mailer-goemon)

Go Email Modules — Operations & Normalization
===================================================================================================
**mailer-goemon** is a library of Go modules extracted from [Sisimai](https://libsisimai.org/),
featuring the `address` package for email address validation and the `smtp` package for SMTP reply
code and status code analysis.

Usage
---------------------------------------------------------------------------------------------------
```go
import "libsisimai.org/address"
```

Packages and functions
===================================================================================================

address
---------------------------------------------------------------------------------------------------
Package "address" provide functions related to an email address.

### Find(text string) [3]string
`address.Find` is an email address parser with a name and comment.
```go
import "libsisimai.org/mailer-goemon/address"
func main(){
    fmt.Printf("1. %v\n", address.Find("kijitora <neko@example.jp> (Nyaan?)")
    fmt.Printf("2. %v\n", address.Find(`cat (Nyaan?) <"kijitora neko"@example.jp>`)
}
// 1. [neko@example.jp kijitora (Nyaan?)]
// 2. ["kijitora neko"@example.jp cat (Nyaan?)]
```

### ExpandVERP(text string) string
`address.ExpandVERP` gets the original recipient address from a VERP address.
```go
import "libsisimai.org/mailer-goemon/address"
func main(){
    fmt.Printf("1. %s\n", address.ExpandVERP("bounce+neko=example.jp@libsisimai.org"))
}
// 1. neko@example.jp
```

### ExpandAlias(text string) string
`address.ExpandAlias` removes string from `+` to `@` at a local part.
```go
import "libsisimai.org/mailer-goemon/address"
func main(){
    fmt.Printf("1. %s\n", address.ExpandAlias("neko+newsletter@example.jp"))
}
// 1. neko@example.jp
```

smtp/reply
---------------------------------------------------------------------------------------------------
Package `smtp/reply` provides funtions related to SMTP reply codes such as `421`, `550`.

### Find(logs, hint string) string
`reply.Find` returns an SMTP reply code found from the given string.
```go
import "libsisimai.org/mailer-goemon/address"
func main(){
    fmt.Printf("1. %s\n", reply.Find("552 5.2.3 Message size exceeds fixed maximum message size (10MB)", ""))
    fmt.Printf("2. %s\n", reply.Find("550 5.1.10 RESOLVER.ADR.RecipientNotFound; Recipient not found", "4"))
}
// 1. 552
// 2. (empty)
```

### reply.Test(code string) bool
`reply.Test` checks whether a reply code is a valid code or not.
```go
import "libsisimai.org/mailer-goemon/smtp/reply"
func main() {
    fmt.Printf("1. %t\n", reply.Test("550"))
    fmt.Printf("2. %t\n", reply.Test("490"))
}
// 1. true
// 2. false
```

smtp/status
---------------------------------------------------------------------------------------------------
Package `smtp/status` provides functions related to SMTP Status codes such as `4.2.2`, `5.1.1`.
See https://www.iana.org/assignments/smtp-enhanced-status-codes/smtp-enhanced-status-codes.xhtml

### Find(logs, hint string) string
`status.Find` returns a delivery status code found from the given string.
```go
import "libsisimai.org/mailer-goemon/smtp/status"
func main() {
    fmt.Printf("1. %s\n", status.Find("552 5.2.3 Message size exceeds fixed maximum message size (10MB)", ""))
    fmt.Printf("2. %s\n", status.Find("550 5.1.10 RESOLVER.ADR.RecipientNotFound; Recipient not found", "4"))
}
// 1. 5.2.3
// 2. (empty)
```

### Test(code string) bool
`status.Test` checks whether an SMTP status code is a valid code or not.
```go
import "libsisimai.org/mailer-goemon/smtp/status"
func main() {
	fmt.Printf("1. %t\n", status.Test("5.1.1"))
	fmt.Printf("2. %t\n", status.Test("3.1.4"))
}
// 1. true
// 2. false
```

smtp/command
---------------------------------------------------------------------------------------------------
Package `smtp/command` provides functions related to SMTP commands.

### Find(text string) string
`command.Find` returns an SMTP command found in the argument.
```go
import "libsisimai.org/mailer-goemon/smtp/command"
func main() {
	fmt.Printf("1. %s\n", command.Find("550-5.7.26 The MAIL FROM domain [v.example.jp] has an SPF record with a hard fail"))
	fmt.Printf("2. %s\n", command.Find("550 5.2.2 <sabineko@example.jp>... Mailbox Full (in reply to RCPT TO command)"))
	fmt.Printf("3. %s\n", command.Find("MAILER-DAEMON"))
}
// 1. MAIL
// 2. RCPT
// 3. (empty)
```

### Test(comm string) bool
`command.Test` checks that an SMTP command in the argument is valid or not.
```go
import "libsisimai.org/mailer-goemon/smtp/command"
func main() {
	fmt.Printf("1. %t\n", command.Test("STARTTLS")) // 1. true
	fmt.Printf("2. %t\n", command.Test("mail"))     // 2. true
	fmt.Printf("3. %t\n", command.Test("NEKO"))     // 3. false
}
```

rfc5322
---------------------------------------------------------------------------------------------------
Package `rfc5322` provides functions for email addresses, `Date:` header, `Received:` headers, and
other headers and messages related to RFC5322. https://datatracker.ietf.org/doc/html/rfc5322

### IsEmailAddress(email string) bool
`rfc5322.IsEmailAddress` checks that the argument is an email address or not.
```go
import "libsisimai.org/mailer-goemon/rfc5322"
func main() {
	fmt.Printf("1. %t\n", rfc5322.IsEmailAddress("neko@example.jp"))
	fmt.Printf("2. %t\n", rfc5322.IsEmailAddress(`"neko nyaan"@example.jp`))
	fmt.Printf("3. %t\n", rfc5322.IsEmailAddress(`neko%example.jp`))
}
// 1. true
// 2. true
// 3. false
```

### IsQuotedAddress(email string) bool
`rfc5322.IsQuotedAddress` checks that the local part of the argument is quoted address or not.
```go
import "libsisimai.org/mailer-goemon/rfc5322"
func main() {
	fmt.Printf("1. %t\n", rfc5322.IsQuotedAddress("neko@example.jp"))
	fmt.Printf("2. %t\n", rfc5322.IsQuotedAddress(`"neko nyaan"@example.jp`))
}
// 1. false
// 2. true
```

### IsComment(text string) bool
`rfc5322.IsComment` returns `true` if the string starts with `(` and ends with `)`.
```go
import "libsisimai.org/mailer-goemon/rfc5322"
func main() {
	fmt.Printf("1. %t\n", rfc5322.IsComment("(nyaan?)"))
	fmt.Printf("2. %t\n", rfc5322.IsComment("nyaaaaan?"))
}
// 1. true
// 2. false
```

### Received(rhead string) [6]string
`rfc5322.Received` convert a `Received` header to a structured data.
```go
import "libsisimai.org/mailer-goemon/rfc5322"
func main() {
	fmt.Printf("r1. %+#v\n", rfc5322.Received("from mx.example.org (c182128.example.net [192.0.2.128]) by mx.example.jp (8.14.4/8.14.4) with ESMTP id oBB3JxRJ022484 for <shironeko@example.jp>; Sat, 11 Dec 2010 12:20:00 +0900 (JST)"));
}
// [6]string{"mx.example.org", "mx.example.jp", "", "esmtp", "obb3jxrj022484", "shironeko@example.jp"}
```


See also
---------------------------------------------------------------------------------------------------
* [RFC5321 - Simple Mail Transfer Protocol](https://tools.ietf.org/html/rfc5321)
* [RFC5322 - Internet Message Format](https://tools.ietf.org/html/rfc5322)

Author
===================================================================================================
[@azumakuniyuki](https://twitter.com/azumakuniyuki) and sisimai development team

Copyright
===================================================================================================
Copyright (C) 2025 azumakuniyuki and sisimai development team, All Rights Reserved.

License
===================================================================================================
This software is distributed under The BSD 2-Clause License.

