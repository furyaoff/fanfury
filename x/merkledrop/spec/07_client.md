<!-- 
order: 7
-->

# Client

## Transactions

The `transactions` commands allow users to `create` and `claim` for _merkledrops_.

```bash=
fanfuryd tx merkledrop --help
```
### create

```bash=
fanfuryd tx merkledrop create [account-file] [output-file] \
	--denom=ufury \
	--start-height=1 \
	--end-height=10 \
	--from=<key-name> -b block --chain-id <chain-id>
```

### claim

```bash=
fanfuryd tx merkledrop claim [merkledrop-id] \
	--proofs=[proofs-list] \
	--amount=[amount-to-claim] \
	--index=[level-index]
	--from=<key-name> -b block --chain-id <chain-id>
```

## Query

The `query` commands allow users to query the _merkledrop_ module.

```bash=
fanfuryd q merkledrop --help
```

### detail by id

```bash=
fanfuryd q merkledrop detail [id]
```

### if index and id have been claimed

```bash=
fanfuryd q merkledrop index-claimed [id] [index]
```

### params

```bash=
fanfuryd q merkledrop params
```