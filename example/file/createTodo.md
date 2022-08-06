## query
```graphql
mutation createTodo($note: String!, $user: String!, $remind: String!) {
  createTodo(input: {text: $note, userName: $user, extra: {remindTime: $remind}}) {
    ...fragmentToto
  }
}

fragment fragmentToto on Todo {
  id
  text
  done
  user {
    id
    name
    hobby
    friends(count: 10){
      id
      name
    }
  }
  extra{
    remindTime
  }
}
```

## variables
```json
{
  "note": "bobo 要吃饭",
  "user": "bobo",
  "remind": "2022-08-08 12:00:00"
}
```
## response
```json
{
  "data": {
    "createTodo": {
      "id": "T5577006791947779410",
      "text": "bobo 要吃饭",
      "done": false,
      "user": {
        "id": 4,
        "name": "bobo",
        "hobby": [
          "acting cute"
        ],
        "friends": []
      },
      "extra": {
        "remindTime": "2022-08-08T12:00:00Z"
      }
    }
  }
}
```
