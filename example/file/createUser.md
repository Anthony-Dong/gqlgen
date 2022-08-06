## query
```graphql
## 注意: 参数是否required，取决于schema定义
mutation createUser($name:String!, $hobby: [String!], $age: Int!, $gender: Gender, $height: Float){
  createUser(input: {name:$name, hobby: $hobby, age: $age, gender: $gender, height: $height}){
    ... userFragment
  }
}

## fragment必须标注on申明来自于哪个对象
fragment userFragment on User {
  id
  name
  height
  hobby
  age
  gender
}
```

### variables
```json
{
  "name":"bobo",
  "hobby": ["acting cute"],
  "age": 10,
  "gender": "MALE",
  "height": 1.5
}
```

### response
```json
{
  "data": {
    "createUser": {
      "id": 4,
      "name": "bobo",
      "height": 1.5,
      "hobby": [
        "acting cute"
      ],
      "age": 10,
      "gender": "MALE"
    }
  }
}
```
