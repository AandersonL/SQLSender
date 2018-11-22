# SQLSender
A simple program written in [Go](https://golang.org), still in developement.


## How install

Make sure to have [Go](https://golang.org) installed on your machine, then install the following libraries on the command line:
```
  go get "github.com/go-sql-driver/mysql"
  go get "github.com/moraes/config"
```

After everything is setup you can compile the application, on your machine, clone the repo on your dev directory and run on the command line

```
  cd SQLSender
  go build SQLSender
 ```
 
 ## Why i'm gonna use that ?
 
 I made this tool to improve my Go capabilities and to study SQL more faster instead use a totally GUI program to write sql.
  You can write your .sql files(using your text editos and use SQLSender send to the database with your SQL cli open. its really more faster! 
  
  
  ### Here a simple vim script to make it faster
  
  ```vim
  fu! SendQuery()
	!./SQLSender %
endfunction

fu! SendAllFiles()
	!./SQLSender *.sql
endfunction
```


# Warning
This is a very simple parser, you can't forget the ";" in the end of every query! :)
