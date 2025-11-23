# snap-sim

## Create Payment
POST /snap/v1/payments
Headers:
```
X-Client-Key: string
X-External-Id: string           # idempotency
X-Timestamp: yyyy-mm-ddTHH:MM:SSZ
X-Signature: base64(HMAC)
Content-Type: application/json
```
Body:
```
{
  "amount": 150000,
  "currency": "IDR",
  "sourceAccount": {
    "accountNo": "1234567890",
    "bankCode": "014"
  },
  "destinationAccount": {
    "accountNo": "9876543210",
    "bankCode": "014"
  },
  "description": "Pembayaran Order #123",
  "additionalInfo": {
    "orderId": "ORDER-123"
  }
}
```
Response:
```
202 Accepted
{
  "responseCode": "0001",
  "responseMessage": "Payment accepted",
  "data": {
    "snapReferenceNo": "SNAP-20251123-00001",
    "externalId": "ORDER-123",
    "status": "PENDING"
  }
}
```

## Get Payment Status
GET /snap/v1/payments/{snapReferenceNo}

Headers:

**Same as Create Payment.**

Response:
```
{
  "responseCode": "0000",
  "responseMessage": "Success",
  "data": {
    "snapReferenceNo": "SNAP-20251123-00001",
    "externalId": "ORDER-123",
    "amount": 150000,
    "currency": "IDR",
    "status": "SUCCESS",
    "settledAt": "2025-11-24T09:35:22Z"
  }
}
```
## Account Inquiry
POST /snap/v1/inquiry/account

Body:

```
{
  "accountNo": "9876543210",
  "bankCode": "014"
}
```
Response:

```
{
  "responseCode": "0000",
  "responseMessage": "Success",
  "data": {
    "accountNo": "9876543210",
    "bankCode": "014",
    "accountName": "John Doe",
    "status": "ACTIVE"
  }
}
```

## Forward to Payment Adapter
POST /internal/validated/payment

```
{
  "clientId": "MER-001",
  "externalId": "ORDER-123",
  "timestamp": "2025-11-24T09:00:15Z",
  "request": {
    "amount": 150000,
    "currency": "IDR",
    "sourceAccount": {
      "accountNo": "1234567890",
      "bankCode": "014"
    },
    "destinationAccount": {
      "accountNo": "9876543210",
      "bankCode": "014"
    },
    "description": "Pembayaran Order #123"
  }
}
```

## Forward to Inquiry Adapter
POST /internal/validated/inquiry

```
{
  "clientId": "MER-001",
  "timestamp": "2025-11-24T09:00:15Z",
  "request": {
    "accountNo": "9876543210",
    "bankCode": "014"
  }
}
```

## Payment Core (XML)
POST /internal/core/payment/process

Content-Type: application/xml

Request:

```
<PaymentRequest>
    <Header>
        <ClientId>MER-001</ClientId>
        <ExternalId>ORDER-123</ExternalId>
        <Timestamp>2025-11-24T09:00:15Z</Timestamp>
        <Channel>SNAP</Channel>
    </Header>
    <Transaction>
        <Amount>150000</Amount>
        <Currency>IDR</Currency>
        <Description>Pembayaran Order #123</Description>
    </Transaction>
    <Source>
        <AccountNo>1234567890</AccountNo>
        <BankCode>014</BankCode>
    </Source>
    <Destination>
        <AccountNo>9876543210</AccountNo>
        <BankCode>014</BankCode>
    </Destination>
</PaymentRequest>
```

Response:

```
<PaymentResponse>
    <ExternalId>ORDER-123</ExternalId>
    <SnapReferenceNo>SNAP-20251123-00001</SnapReferenceNo>
    <Status>PENDING</Status>
    <ResponseCode>0000</ResponseCode>
    <ResponseMessage>Accepted</ResponseMessage>
</PaymentResponse>
```

## Inquiry Core (XML)
POST /internal/core/inquiry/account

Request:

```
<AccountInquiryRequest>
    <ClientId>MER-001</ClientId>
    <AccountNo>9876543210</AccountNo>
    <BankCode>014</BankCode>
</AccountInquiryRequest>
```

Response:

```
<AccountInquiryResponse>
    <AccountNo>9876543210</AccountNo>
    <BankCode>014</BankCode>
    <AccountName>John Doe</AccountName>
    <Status>ACTIVE</Status>
    <ResponseCode>0000</ResponseCode>
</AccountInquiryResponse>
```

## RabbitMQ Events
(Internal messaging)

Event: Transaction Created

```
{
  "snapReferenceNo": "SNAP-20251123-00001",
  "externalId": "ORDER-123",
  "clientId": "MER-001",
  "amount": 150000,
  "currency": "IDR",
  "status": "PENDING"
}
```

Event: Callback Ready

```
{
  "snapReferenceNo": "SNAP-20251123-00001",
  "externalId": "ORDER-123",
  "clientId": "MER-001",
  "status": "SUCCESS",
  "settledAt": "2025-11-24T09:34:58Z"
}
```

## Payment Callback
POST /merchant/callback/payment

Headers:

```
X-Provider-Key: SNAP-PROVIDER-001
X-Timestamp: ...
X-Signature: base64(HMAC)
```

Body:

```
{
  "snapReferenceNo": "SNAP-20251123-00001",
  "externalId": "ORDER-123",
  "amount": 150000,
  "currency": "IDR",
  "status": "SUCCESS",
  "settledAt": "2025-11-24T09:34:58Z"
}
```
## Create Order (Merchant)
POST /merchant/orders

```
{
  "orderId": "ORDER-123",
  "amount": 150000,
  "customerId": "CUST-01"
}
```

Response:

```
{
  "orderId": "ORDER-123",
  "snapReferenceNo": "SNAP-20251123-00001",
  "status": "PENDING"
}
```