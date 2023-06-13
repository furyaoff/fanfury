<!-- 
order: 4
-->

# Events

The fantoken module emits the following events:
## EventIssue

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgIssue` |
| fanfury.fantoken.v1beta1.EventIssue | denom        | {denom}         |

## EventDisableMint

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgDisableMint` |
| fanfury.fantoken.v1beta1.EventDisableMint | denom        | {denom}         |

## EventMint

| Type                     | Attribute Key | Attribute Value   |
| :----------------------- | :------------ | :---------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgMint` |
| fanfury.fantoken.v1beta1.EventMint | recipient        | {recipient}         |
| fanfury.fantoken.v1beta1.EventMint | coin        | {coin}         |

## EventBurn

| Type           | Attribute Key | Attribute Value    |
| :------------- | :------------ | :----------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgBurn` |
| fanfury.fantoken.v1beta1.EventBurn | sender        | {sender}         |
| fanfury.fantoken.v1beta1.EventBurn | coin        | {coin}         |

## EventSetAuthority

| Type           | Attribute Key | Attribute Value |
| :------------- | :------------ | :-------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgSetAuthority` |
| fanfury.fantoken.v1beta1.EventTransferAuthority | denom        | {denom}         |
| fanfury.fantoken.v1beta1.EventTransferAuthority | old_authority        | {old_authority}         |
| fanfury.fantoken.v1beta1.EventTransferAuthority | new_authority        | {new_authority}         |

## EventSetMinter

| Type           | Attribute Key | Attribute Value |
| :------------- | :------------ | :-------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgSetMinter` |
| fanfury.fantoken.v1beta1.EventTransferMinter | denom        | {denom}         |
| fanfury.fantoken.v1beta1.EventTransferMinter | old_minter        | {old_minter}         |
| fanfury.fantoken.v1beta1.EventTransferMinter | new_authority        | {new_minter}         |

## EventSetUri

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/fanfury.fantoken.v1beta1.MsgSetUri` |
| fanfury.fantoken.v1beta1.EventSetUri | denom        | {denom}         |
