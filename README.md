# Twitter like web application

Simple twitter like web app to process tweets data.

# Evalable endpoints

| endpoint | what they do |
| ------ | ------ |
| localhost:8081/tweets/put | help you to put the data inside the storage |
| localhost:8081/tweets/get | help you to retrive the data by id field from storage |
| localhost:8081/tweets/list | return all stored tweets |
| localhost:8081/tweets/time | return current time, no more |


# Tech
Go, Git, http/net, sync, context, go mod, 

# Install and run

- Download the application: 
```sh
git clone git@github.com:stevenstr/tweets_app_reforged.git
```
- Go to the application directory:
```sh
cd tweets_app_reforged\tweets\cmd\
```
- Run micriservice using the following command:
```sh
go run main.go
```



# Usage
## Put the data
 ```sh
curl --location --request PUT 'localhost:8081/tweets/put?id=1&message=I_am_giorgio'
curl --location --request PUT 'localhost:8081/tweets/put?id=2&message=I_am_PEPE'
curl --location --request PUT 'localhost:8081/tweets/put?id=3&message=SWAFFARD' 
```

### Get the data by id
```sh
curl --location 'localhost:8081/tweets/get?id=1'
```
{"message":"I am giorgio"}
```sh
curl --location 'localhost:8081/tweets/get?id=3'
```
{"message":"SWAFFARD"}

```sh
curl --location 'localhost:8081/tweets/get?id=4'
```
null

### Get the all data 
```sh
curl --location 'localhost:8081/tweets/list'
```
{"1":{"message":"I am giorgio"},"2":{"message":"I_am_PEPE"},"3":{"message":"SWAFFARD"}}

### Get the current time
```sh
curl --location 'localhost:8081/tweets/time'
```
"Thu, 30 Jan 2025 14:03:43 MSK"

# License
- MIT License