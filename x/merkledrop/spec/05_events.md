<!-- 
order: 5
-->

# Events

The merkledrop module emits the following events:
## EventCreate

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/fanfury.merkledrop.v1beta1.MsgCreate` |
| fanfury.merkledrop.v1beta1.EventCreate | owner        | {owner}         |
| fanfury.merkledrop.v1beta1.EventCreate | merkledrop_id        | {merkledrop_id}         |

## EventClaim

| Type            | Attribute Key | Attribute Value  |
| :-------------- | :------------ | :--------------- |
| message         | action        | `/fanfury.merkledrop.v1beta1.MsgClaim` |
| fanfury.merkledrop.v1beta1.EventClaim | merkledrop_id        | {merkledrop_id}         |
| fanfury.merkledrop.v1beta1.EventClaim | index        | {index}         |
| fanfury.merkledrop.v1beta1.EventClaim | coin        | {coin}         |

## EventWithdraw

| Type                     | Attribute Key | Attribute Value   |
| :----------------------- | :------------ | :---------------- |
| fanfury.merkledrop.v1beta1.EventWithdraw | merkledrop_id        | {merkledrop_id}         |
| fanfury.merkledrop.v1beta1.EventWithdraw | coin        | {coin}         |