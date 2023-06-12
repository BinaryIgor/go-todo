# Go TODO

TODO api (+ optional html version) implementation to play with go lang.

## Api requirements
Errors:
```
{
    errors: [string],
    message: string
}
```

Public endpoints:
```
POST: /users/sign-up
Body: {
    name: string,
    password: string
}
Response: { 
    id: UUID
}


POST: /users/sign-in
Body: {
    name: string,
    password: string
}
Response: {
    token: string
}
```

Private endpoints (all in the context of a current :userId):
```
POST: /todos
Body: {
    name: string,
    deadline: timestamp?,
    priority: enum, one of the: LOW, MEDIUM, HIGH,
    description: string?,
    steps: [
        {
            order: 1,
            name: string,
            description: string?
        },
        {
            order: 2,
            name: string,
            description: string?
        }
    ] 
}
Response: {
    id: UUID,
    //whole TODO
}


PUT: /todos/:todoId
Body: {
  TODO model with steps
}
Response: {
    //changed TODO
}

DELETE: /todos/todoId
Response: Empty

GET: /todos?priority=(optional, multiple)&deadlineFrom(optional, inclusive)=&deadlineTo(optional, exclusive)
Response: [
    {
        //TODO with steps
    }
]
```