# tsmap
this util package is aim to solve the problem that _original go map_ is not thread safe but the _sync.Map_ can not be used with specific type.    

TSMap is a **thread safe** map with specific key/value _**type**_.