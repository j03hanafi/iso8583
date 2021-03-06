# iso8583
iso8583 is a golang app that handle transaction request and gonna processed to standard ISO:8583, it included database connection to save any request that it get.

## Installation
Get the program
```bash
git clone https://github.com/j03hanafi/iso8583
cd jsonGo
```
Prepare package
```bash
github.com/go-sql-driver/mysql
github.com/gorilla/mux
github.com/mofax/iso8583
github.com/rivo/uniseg
github.com/stretchr/testify
```

Running program
```bash
go run *.go
```
Build and running program
```bash
go build -o {program name}
./{program name}
```

## Available Request (For Now)
#### GET
- `/payment` : Get all transaction data
- `/payment/{processingCode}` : Get transaction data based on processingCode
- `/payment/{processingCode}/iso8583` : Get ISO8385 format of transaction data based on processingCode
- `/payment/{processingCode}/iso8583/{element}` : Get each Data Element from ISO8583 format message

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
