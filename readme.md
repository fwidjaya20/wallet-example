# Wallet Service Example
Wallet Service with Event Sourcing.

#### Prerequisites
- Go

#### Domain
- Wallet

## Getting Started
#### How To Serve
```shell script
go run ./cmd/app
```
or
```shell script
make run_dev
```

## Tool
#### Make
1. Migration
```shell script
make migration version=YOUR_VERSION_NUMBER filename=YOUR_FILE_NAME
```
2. Domain
```shell script
make domain name=YOUR_DOMAIN_NAME
```

## Service
#### Top Up
> [POST] localhost:8001/wallet/{wallet_id}/top-up
##### Request Payload
```json
{
    "amount": 25000
}
```
##### Response
```json
{
    "data": {
        "amount": 25000,
        "job": "Top Up",
        "status": "succeed"
    },
    "metadata": {}
}
```

#### Withdraw
> [POST] localhost:8001/wallet/{wallet_id}/withdraw
##### Request Payload
```json
{
    "account_number": "5271531169",
    "account_holder": "Fredrick Widjaya",
    "bank": "bca",
    "amount": 500000
}
```
##### Response
```json
{
    "data": {
        "amount": 500000,
        "job": "Withdraw",
        "status": "succeed"
    },
    "metadata": {}
}
```

#### GetBalance
> [POST] localhost:8001/wallet/{wallet_id}/balance
##### Response
```json
{
    "data": {
        "balance": 25000
    },
    "metadata": {}
}
```

#### GetTransaction
> [POST] localhost:8001/wallet/{wallet_id}/transactions
##### Response
```json
{
    "data": [
        {
            "transaction_id": "cbf3a305-326f-40b7-a947-265c03b64ed7",
            "amount": -500000,
            "balance_type": "withdraw",
            "notes": "{\"fee\": 4000, \"amount\": 500000.000000, \"message\": \"withdraw\", \"withdraw_destination\": {\"bank\": \"bca\", \"account_holder\": \"Fredrick Widjaya\", \"account_number\": \"5271531169\"}}"
        },
        {
            "transaction_id": "9a7fa26e-e177-47b5-b6a0-a3dd2a77fb01",
            "amount": 25000,
            "balance_type": "deposit",
            "notes": "{\"message\": \"add\"}"
        },
        {
            "transaction_id": "555aaa77-a43e-453a-af25-4d376905b97c",
            "amount": 250000,
            "balance_type": "deposit",
            "notes": "{\"message\": \"add\"}"
        },
        {
            "transaction_id": "235ce377-ca65-4fd5-a060-717fede5e205",
            "amount": 150000,
            "balance_type": "deposit",
            "notes": "{\"message\": \"add\"}"
        },
        {
            "transaction_id": "00fc28d7-bff8-46fc-8ff7-1292be2977c3",
            "amount": 50000,
            "balance_type": "deposit",
            "notes": "{\"message\": \"add\"}"
        },
        {
            "transaction_id": "981284eb-88d6-445b-9f55-927dd92c59ac",
            "amount": 50000,
            "balance_type": "deposit",
            "notes": "{\"message\": \"add\"}"
        }
    ],
    "metadata": {}
}
```
