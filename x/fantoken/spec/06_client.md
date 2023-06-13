<!-- 
order: 6
-->

# Client

## Transactions

The `transactions` commands allow users to `issue`, `mint`, `burn`, `disable minting`, `transfer minting and editing capabilities` for _fan tokens_.

```bash=
fanfuryd tx fantoken --help
```
### issue

```bash=
fanfuryd tx fantoken issue \
    --name "fantoken name" \
    --symbol "bitangel" \
    --max-supply 100000000000 \
    --uri "ipfs://...." \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### mint

```bash=
fanfuryd tx fantoken mint [amount][denom] \
    --recipient <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### burn

```bash=
fanfuryd tx fantoken burn [amount][denom] \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-authority

```bash=
fanfuryd tx fantoken set-authority [denom] \
    --new-authority <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-minter

```bash=
fanfuryd tx fantoken set-minter [denom] \
    --new-minter <address> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### set-uri

```bash=
fanfuryd tx fantoken set-uri [denom] \
    --uri <uri> \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

### disable-mint

```bash=
fanfuryd tx fantoken disable-mint [denom] \
    --from <key-name> -b block --chain-id <chain-id> --fees <fee>
```

## Query

The `query` commands allow users to query the `fantoken` module.

```bash=
fanfuryd q fantoken --help
```

### denom

```bash=
fanfuryd q fantoken denom <denom>
```

### authority

```bash=
fanfuryd q fantoken authority <address>
```

### params

```bash=
fanfuryd q fantoken params
```