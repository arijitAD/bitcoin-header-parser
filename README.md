# Bitcoin Header parser

## Steps
* Download the [bitcoin core](https://bitcoin.org/en/download) and sync only headers. This takes around 10 minutes. (Note: Syncing entire block data might take longer. It's around 300 GB)
* Change the path of bitcoin DB in `TestBlockCount` in `main.go`.
* Run `go test`