# SQLSender
A simple program written in [Go](https://golang.org) to parse .sql scripts files and send to a remote host. 


## How to install

Make sure to have [Go](https://golang.org) installed on your machine, then install the following libraries on the command line:
```
  go get "github.com/go-sql-driver/mysql"
  go get "github.com/moraes/config"
```

After everything is setup you can compile the application on your machine, clone the repo on your dev directory and run on the command line

```
  cd SQLSender
  go build SQLSender
 ```
 And edit the config.yml with your database info.
 
 ## Why should I use that? ?
 
 I made this tool to learn Go better and to study SQL faster than using others GUI tool to write sql.
  You can write your .sql files(using your text editor and use SQLSender send to the database with your SQL cli open. its really more faster! 
  
  
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
