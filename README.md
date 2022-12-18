# mini-wallet Backend

mini-wallet is a Mini Wallet Exercise API based on Golang, Echo and Gorm.

## Stack Requirements

- Golang
- Mysql
- dbmate `brew install dbmate`
- Makefile `brew install make`

## How to Run

### Clone

First clone this repo by run:

```sh
$ git clone git@github.com:notblessy/mini-wallet.git
```

### Environtment

- Don't forget to set `.env`, you can copy from `env.sample`

### Database Migration

- To migrate tables, ensure you have already installed `dbmate` then run

```sh
$ dbmate up
```

### Running project

- Before running, you need to create `Makefile` from the file `Makefile.sample`. Then run for debugging by `make run`.

### API Documentations

- To test the API, import `postman collection` from folder `api-docs/`. All the API is available there.

## Author

```
I Komang Frederich Blessy
```
