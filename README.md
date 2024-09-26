This is a very simple tool that will read all the keys in your BuntDB database, marshall the data for JSON,
and write the JSON data to an export file. 

# Example
```cmd
go build main.go
.\main.exe -db-path data.db -export-path out.json
```

# References
The following references were used to help put this together, thank you [@tidwall](https://github.com/tidwall)!:
* https://github.com/tidwall/buntdb?tab=readme-ov-file#opening-a-database
* https://github.com/tidwall/buntdb/issues/59