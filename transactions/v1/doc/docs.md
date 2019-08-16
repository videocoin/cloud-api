# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [transactions/v1/action.proto](#transactions/v1/action.proto)
    - [Action](#cloud.api.transactions.v1.Action)
  
    - [ActionSource](#cloud.api.transactions.v1.ActionSource)
    - [ActionType](#cloud.api.transactions.v1.ActionType)
  
  
  

- [transactions/v1/action_service.proto](#transactions/v1/action_service.proto)
    - [GetActionRequest](#cloud.api.transactions.v1.GetActionRequest)
    - [GetActionResponse](#cloud.api.transactions.v1.GetActionResponse)
    - [GetActionsByStreamIdRequest](#cloud.api.transactions.v1.GetActionsByStreamIdRequest)
    - [GetActionsByStreamIdResponse](#cloud.api.transactions.v1.GetActionsByStreamIdResponse)
    - [GetActionsRequest](#cloud.api.transactions.v1.GetActionsRequest)
    - [GetActionsResponse](#cloud.api.transactions.v1.GetActionsResponse)
  
  
  
    - [ActionService](#cloud.api.transactions.v1.ActionService)
  

- [transactions/v1/block.proto](#transactions/v1/block.proto)
    - [Block](#cloud.api.transactions.v1.Block)
    - [BlockDetail](#cloud.api.transactions.v1.BlockDetail)
  
  
  
  

- [transactions/v1/block_service.proto](#transactions/v1/block_service.proto)
    - [GetBlockRequest](#cloud.api.transactions.v1.GetBlockRequest)
    - [GetBlockResponse](#cloud.api.transactions.v1.GetBlockResponse)
    - [GetBlocksRequest](#cloud.api.transactions.v1.GetBlocksRequest)
    - [GetBlocksResponse](#cloud.api.transactions.v1.GetBlocksResponse)
  
  
  
    - [BlockService](#cloud.api.transactions.v1.BlockService)
  

- [transactions/v1/transaction.proto](#transactions/v1/transaction.proto)
    - [Transaction](#cloud.api.transactions.v1.Transaction)
    - [TransactionDetail](#cloud.api.transactions.v1.TransactionDetail)
  
  
  
  

- [transactions/v1/transaction_service.proto](#transactions/v1/transaction_service.proto)
    - [GetAllTransactionsRequest](#cloud.api.transactions.v1.GetAllTransactionsRequest)
    - [GetAllTransactionsResponse](#cloud.api.transactions.v1.GetAllTransactionsResponse)
    - [GetTransactionRequest](#cloud.api.transactions.v1.GetTransactionRequest)
    - [GetTransactionResponse](#cloud.api.transactions.v1.GetTransactionResponse)
    - [GetTransactionsRequest](#cloud.api.transactions.v1.GetTransactionsRequest)
    - [GetTransactionsResponse](#cloud.api.transactions.v1.GetTransactionsResponse)
  
  
  
    - [TransactionService](#cloud.api.transactions.v1.TransactionService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="transactions/v1/action.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/action.proto



<a name="cloud.api.transactions.v1.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source | [ActionSource](#cloud.api.transactions.v1.ActionSource) |  |  |
| hash | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| value | [string](#string) |  |  |
| type | [ActionType](#cloud.api.transactions.v1.ActionType) |  |  |





 


<a name="cloud.api.transactions.v1.ActionSource"></a>

### ActionSource


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDETECTED | 0 |  |
| ETH | 1 |  |
| VID | 2 |  |



<a name="cloud.api.transactions.v1.ActionType"></a>

### ActionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| DEPOSIT | 1 |  |
| STREAM_CREATED | 2 |  |
| STREAM_ENDED | 3 |  |
| INPUT_CHUNK_ADDED | 4 |  |
| CHUNK_PROOF_SUBMITTED | 5 |  |
| CHUNK_PROOF_VALIDATED | 6 |  |
| CHUNK_PROOF_SCRAPPED | 14 |  |
| ACCOUNT_FUNDED | 7 |  |
| STREAM_REQUESTED | 8 |  |
| STREAM_REFUNDED | 15 |  |
| STREAM_APPROVED | 9 |  |
| VALIDATOR_ADDED | 10 |  |
| VALIDATOR_REMOVED | 11 |  |
| REFUND_ALLOWED | 12 |  |
| REFUND_REVOKED | 13 |  |


 

 

 



<a name="transactions/v1/action_service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/action_service.proto



<a name="cloud.api.transactions.v1.GetActionRequest"></a>

### GetActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |






<a name="cloud.api.transactions.v1.GetActionResponse"></a>

### GetActionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [Action](#cloud.api.transactions.v1.Action) |  |  |






<a name="cloud.api.transactions.v1.GetActionsByStreamIdRequest"></a>

### GetActionsByStreamIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream_id | [uint64](#uint64) |  |  |
| limit | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="cloud.api.transactions.v1.GetActionsByStreamIdResponse"></a>

### GetActionsByStreamIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [Action](#cloud.api.transactions.v1.Action) | repeated |  |






<a name="cloud.api.transactions.v1.GetActionsRequest"></a>

### GetActionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| limit | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="cloud.api.transactions.v1.GetActionsResponse"></a>

### GetActionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [Action](#cloud.api.transactions.v1.Action) | repeated |  |





 

 

 


<a name="cloud.api.transactions.v1.ActionService"></a>

### ActionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetActions | [GetActionsRequest](#cloud.api.transactions.v1.GetActionsRequest) | [GetActionsResponse](#cloud.api.transactions.v1.GetActionsResponse) |  |
| GetAction | [GetActionRequest](#cloud.api.transactions.v1.GetActionRequest) | [GetActionResponse](#cloud.api.transactions.v1.GetActionResponse) |  |
| GetActionsByStreamId | [GetActionsByStreamIdRequest](#cloud.api.transactions.v1.GetActionsByStreamIdRequest) | [GetActionsByStreamIdResponse](#cloud.api.transactions.v1.GetActionsByStreamIdResponse) |  |

 



<a name="transactions/v1/block.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/block.proto



<a name="cloud.api.transactions.v1.Block"></a>

### Block



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |
| number | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| num_txs | [int64](#int64) |  |  |
| gas_used | [string](#string) |  |  |
| gas_limit | [string](#string) |  |  |
| size | [double](#double) |  |  |






<a name="cloud.api.transactions.v1.BlockDetail"></a>

### BlockDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |
| parent_hash | [string](#string) |  |  |
| number | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| gas_used | [string](#string) |  |  |
| gas_limit | [string](#string) |  |  |
| size | [double](#double) |  |  |
| transactions | [string](#string) | repeated |  |
| extra_data | [string](#string) |  |  |





 

 

 

 



<a name="transactions/v1/block_service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/block_service.proto



<a name="cloud.api.transactions.v1.GetBlockRequest"></a>

### GetBlockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |






<a name="cloud.api.transactions.v1.GetBlockResponse"></a>

### GetBlockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block | [BlockDetail](#cloud.api.transactions.v1.BlockDetail) |  |  |






<a name="cloud.api.transactions.v1.GetBlocksRequest"></a>

### GetBlocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="cloud.api.transactions.v1.GetBlocksResponse"></a>

### GetBlocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blocks | [Block](#cloud.api.transactions.v1.Block) | repeated |  |





 

 

 


<a name="cloud.api.transactions.v1.BlockService"></a>

### BlockService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetBlocks | [GetBlocksRequest](#cloud.api.transactions.v1.GetBlocksRequest) | [GetBlocksResponse](#cloud.api.transactions.v1.GetBlocksResponse) |  |
| GetBlock | [GetBlockRequest](#cloud.api.transactions.v1.GetBlockRequest) | [GetBlockResponse](#cloud.api.transactions.v1.GetBlockResponse) |  |

 



<a name="transactions/v1/transaction.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/transaction.proto



<a name="cloud.api.transactions.v1.Transaction"></a>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| value | [string](#string) |  |  |






<a name="cloud.api.transactions.v1.TransactionDetail"></a>

### TransactionDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| value | [string](#string) |  |  |
| block_number | [string](#string) |  |  |
| block_hash | [string](#string) |  |  |
| nonce | [uint64](#uint64) |  |  |
| input | [string](#string) |  |  |
| gas | [uint64](#uint64) |  |  |
| gas_price | [string](#string) |  |  |
| status | [uint64](#uint64) |  |  |





 

 

 

 



<a name="transactions/v1/transaction_service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## transactions/v1/transaction_service.proto



<a name="cloud.api.transactions.v1.GetAllTransactionsRequest"></a>

### GetAllTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="cloud.api.transactions.v1.GetAllTransactionsResponse"></a>

### GetAllTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactions | [Transaction](#cloud.api.transactions.v1.Transaction) | repeated |  |






<a name="cloud.api.transactions.v1.GetTransactionRequest"></a>

### GetTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hash | [string](#string) |  |  |






<a name="cloud.api.transactions.v1.GetTransactionResponse"></a>

### GetTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [TransactionDetail](#cloud.api.transactions.v1.TransactionDetail) |  |  |






<a name="cloud.api.transactions.v1.GetTransactionsRequest"></a>

### GetTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| limit | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="cloud.api.transactions.v1.GetTransactionsResponse"></a>

### GetTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactions | [Transaction](#cloud.api.transactions.v1.Transaction) | repeated |  |





 

 

 


<a name="cloud.api.transactions.v1.TransactionService"></a>

### TransactionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTransactions | [GetTransactionsRequest](#cloud.api.transactions.v1.GetTransactionsRequest) | [GetTransactionsResponse](#cloud.api.transactions.v1.GetTransactionsResponse) |  |
| GetTransaction | [GetTransactionRequest](#cloud.api.transactions.v1.GetTransactionRequest) | [GetTransactionResponse](#cloud.api.transactions.v1.GetTransactionResponse) |  |
| GetAllTransactions | [GetAllTransactionsRequest](#cloud.api.transactions.v1.GetAllTransactionsRequest) | [GetAllTransactionsResponse](#cloud.api.transactions.v1.GetAllTransactionsResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

