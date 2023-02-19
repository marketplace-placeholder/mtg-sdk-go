# Magic: The Gathering SDK

This is the Magic: The Gathering SDK Go implementation. It is a wrapper around the MTG API of [magicthegathering.io](http://magicthegathering.io/).

## Installation

Just run

`go get github.com/Enviy/mtg-sdk-go`
OR just let go mod handle dependencies for you by calling
`go mod init <yourProject>`
`go mod tidy`

Want to see what sets are in standard?
``` Go
sets, err := mtg.StandardSets()
if err != nil {
    return err
}
fmt.Println(sets)
```

## Docs

See [GoDoc](https://pkg.go.dev/github.com/Enviy/mtg-sdk-go)
