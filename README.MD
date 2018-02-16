# dbconnect.go
----

Go言語で Gorp (O/R) マッパーを利用したDB接続サンプルです

利用方法
```
# ライブラリの取得
$ mkdir lib
$ export GOPATH=./lib
$ go get github.com/coopernurse/gorp
$ go get github.com/go-sql-driver/mysql
$ go get github.com/mattn/go-sqlite3

# DBの作成 (sqlite3の場合)
$ sqlite3 sqlite3.db
.read ddl/customer.sql
.read data/customer.sql
.quit

# ビルド
$ go build dbconn_*.go

# 実行 (sqlite3版)
$ ./dbconn_sqlite3
```

以上
