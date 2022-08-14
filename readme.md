HOW TO RUN THE SERVICE:
# In local :
- set the env, copy from .env.template
- run `docker build`
- run `docker-compose build`
- run `docker-compose up`
# In the cloud :
https://git.heroku.com/api-with-go.git

# API DOCS :
- `GET /users`
- `GET /users/:userid`

### GET /users

description: 
  get all users

Response:

- status: 200
- body:

```json
[
    {
        "Userid": "01",
        "Name": "Budi"
    },
    {
        "Userid": "02",
        "Name": "Nano"
    }
]
```

### GET /users/:userid

description: 
  get a user

params:
    userid

Response:

- status: 200
- body:

```json
{
    "Userid": "01",
    "Name": "Budi"
}
```