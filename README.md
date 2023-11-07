# Hashcash TCP-server example

This repository contains both client and server implementations of the Hashcash algorithm, a proof-of-work system commonly used for mitigating email spam and distributed denial-of-service (DDoS) attacks.

## Usage
```
make server
make client
```
Builds and runs docker containers.

## Why hashcash?
+ It is relatively simple to implement and has well-documented specifications
+ It allows you to adjust the level of computational difficulty required to generate a proof of work. This flexibility enables you to fine-tune the resource cost according to your specific needs.

## Communication example
All requests must be JSON-encoded and containing "type" field.

First request requesting a challenge
```
{"type":0}
```
Challenge response
```
{"challenge":{"Bits":3,"Date":"2023-11-07T19:34:47.079223+01:00","Rand":701215707,"Nonce":0}}
```
Submitting a solution
```
{"type":1,"solution":{"Bits":3,"Date":"2023-11-07T19:34:47.079223+01:00","Rand":701215707,"Nonce":3012}}
```
Successful response from server
```
{"quote":"If you have knowledge, let others light their candles in it."}
```

## Possible improvements
+ Move hardcoded parameters to CLI
+ Use persistent storage for used tokens. This way server can have multiple instances.
+ More unit tests coverage
+ Performance tests