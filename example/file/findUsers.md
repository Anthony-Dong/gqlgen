## query
```graphql
query findUsers($friendCount: Int!){
  users{
    ... userFragment
  }
}

fragment userFragment on User {
  id
  name
  height
  hobby
  age
  gender
  friends(count: $friendCount){
    id
    name
  }
}
```

## variable
```json
{
  "friendCount": 2
}
```

## response
```json
{
  "data": {
    "users": [
      {
        "id": 1,
        "name": "tom",
        "height": 1.8,
        "hobby": [
          "coding"
        ],
        "age": 18,
        "gender": "MALE",
        "friends": [
          {
            "id": 2,
            "name": "tony"
          },
          {
            "id": 3,
            "name": "kitty"
          }
        ]
      },
      {
        "id": 2,
        "name": "tony",
        "height": 1.9,
        "hobby": [
          "basketball"
        ],
        "age": 22,
        "gender": "MALE",
        "friends": [
          {
            "id": 3,
            "name": "kitty"
          }
        ]
      },
      {
        "id": 3,
        "name": "kitty",
        "height": 1.7,
        "hobby": [
          "dance"
        ],
        "age": 12,
        "gender": "FEMALE",
        "friends": [
          {
            "id": 1,
            "name": "tom"
          }
        ]
      }
    ]
  }
}
```
