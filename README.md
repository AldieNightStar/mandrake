# Mandrake

<center><img src="logo.png" width="250px" /></center>

## Note
* Allows you to encrypt/decrypt any byte data
* All you need to have is password

## Install
```bash
go get github.com/AldieNightStar/mandrake
```


## Usage
* `data` - is a `[]byte` data
```go
PASSWORD := []byte("The password")

// Encrypt data
encrypted := mandrake.Encode(data, PASSWORD)

// Decrypt data
originalData := mandrake.Decode(encrypted, PASSWORD)
```


# Base 64 (Text to Bytes and Back)


## Note
* Allows you to encode/decode with base64 format
* Allows you to send over the text-only networks (For example HTTP(S))


## Usage
* `data` - is a `[]byte` data
```go
PASSWORD := []byte("The password")

// Encrypt data
base64String := mandrake.EncodeBase64(data, PASSWORD)

// Decrypt data
originalData, err := mandrake.DecodeBase64(base64String, PASSWORD)
```
