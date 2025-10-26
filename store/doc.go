package main

// 参考库 - https://github.com/skizzehq/skizze
// scribble - NDJSON - github.com/nanobox-io/golang-scribble
// diskv - KV数据库 - https://github.com/peterbourgon/diskv
// s3git - Git版本库 - https://github.com/s3git/s3git-go

/**
## API - 设计

1. Create a new Domain (Collection of Sketches):

```sh
CREATE DOM demostream 10000000 100
```

2. Add values to the domain:

```sh
#ADD DOM $name $value1, $value2 ....
ADD DOM demostream zod joker grod zod zod grod
```

3. Get the cardinality of the domain:

```sh
# GET CARD $name
GET CARD demostream

# returns:
# Cardinality: 9
```

4. Get the rankings of the domain:

```sh
# GET RANK $name
GET RANK demostream

# returns:
# Rank: 1	  Value: zod	  Hits: 3
# Rank: 2	  Value: grod	  Hits: 2
# Rank: 3	  Value: joker	  Hits: 1
```
*/
