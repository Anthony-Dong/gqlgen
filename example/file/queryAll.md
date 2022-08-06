## query
```graphql
query queryTodo($withFriends: Boolean!, $withoutUserInfo: Boolean!) {
  todo_list: todos {
    ...fragmentToto
  }
  user_list: users {
    ...userFragment
  }
}

fragment fragmentToto on Todo {
  id
  text
  done
  user @skip(if: $withoutUserInfo) {
    ...userFragment
  }
  extra {
    remindTime
  }
}

fragment userFragment on User {
  id
  name
  height
  hobby
  age
  gender
  friends(count: 10) @include(if: $withFriends) {
    id
    name
  }
}
```
## Case1

### variable
```json
{
  "withFriends": true,
  "withoutUserInfo": true
}
```
### response
```json
{
  "data": {
    "todo_list": [
      {
        "id": "T5577006791947779410",
        "text": "bobo 要吃饭",
        "done": false,
        "extra": {
          "remindTime": "2022-08-08T12:00:00Z"
        }
      }
    ],
    "user_list": [
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
      },
      {
        "id": 4,
        "name": "bobo",
        "height": 1.5,
        "hobby": [
          "acting cute"
        ],
        "age": 10,
        "gender": "MALE",
        "friends": []
      }
    ]
  }
}
```

## Case2
### variable
```json
{
  "withFriends": true,
  "withoutUserInfo": false
}
```
### response
```json
{
  "data": {
    "todo_list": [
      {
        "id": "T5577006791947779410",
        "text": "bobo 要吃饭",
        "done": false,
        "user": {
          "id": 4,
          "name": "bobo",
          "height": 1.5,
          "hobby": [
            "acting cute"
          ],
          "age": 10,
          "gender": "MALE",
          "friends": []
        },
        "extra": {
          "remindTime": "2022-08-08T12:00:00Z"
        }
      }
    ],
    "user_list": [
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
      },
      {
        "id": 4,
        "name": "bobo",
        "height": 1.5,
        "hobby": [
          "acting cute"
        ],
        "age": 10,
        "gender": "MALE",
        "friends": []
      }
    ]
  }
}
```
