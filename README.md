### Main features

## Application cycle
lets have a look at the applicaton from a high level view.

!["High level view of the application"](https://github.com/youssefhmidi/E2E_encryptedConnection/blob/main/.assets/1.png)

as you can see in the above image, it shows that if  a client sends a request it will be sent to the http server who execute some code or call a function 
in the controllers directory who calls other function which mutate or get some data from the database

> the application may seem complex for early intermediate golang develloper but it is 'for me' an easy architecture. I may got a little to far with the 
> abstraction, but it made it easy for me to add, remove or change something, and for people who may add a new feature to it.


> I made this backend application for the goal of being more comfortable with making and handling large code bases, even though it's lack a lot of features
> 'i.e adding/removing members, handling invite keys and others simple feature' my goal was to lunch a project into production as fast as possible.

!["Detailed view of the backend architecture"](https://github.com/youssefhmidi/E2E_encryptedConnection/blob/main/.assets/2.png)

all the file in the "/routes/" directory handle incomming requests and calls a function in the "/controllers/" directory who uses some features form 
the "/services/" directory or use some packages from the "/_internals/", and then those directories mutate the database and return a success or a data response, also the socket package may send someting to the end user because it streams data.

for the request/response cycle here's a example of it:

!["Detailed view of the request/response cycle"](https://github.com/youssefhmidi/E2E_encryptedConnection/blob/main/.assets/3.png)

## Endpoints

- "/login" & "/signup"  & "/refresh/" endpoints for user logic
- "/chat/" endpoint for the actual app :
    - "/chat/rooms" Shows all the public rooms
    - "/chat/new" create a new dms or a group chat
    - "/chat/@me/dms" 
    - "/chat/{room_id}" get the specified data in the request 
    - "/chat/{room_id}/join" require an invitation key if not public
- "/users/" get all users, **Note: this is an OnlyAdmin endpoint** :
    - "/users/@me": get information about the current user after providing an acces_token

## Encryption strategies
Most encryption will happen at the client side so the server will be reciving just some random text that can be decrypted, but the server side will be 
responsible for orginizing group chats and public groups  

## Usage "aka code snippets for you to copy :)

- run these commands so you can clone this repo

```bash
    git clone https://github.com/youssefhmidi/e2ee_ChatApp.git
    cd ./e2ee_ChatApp
```

- then copy and make your own .env file (no db config it uses sqlite)
- make sure to initilize a db if it didn't initilize it itself
- then run these commands 

```bash
    go -o ./build/app.exe ./cmd/main.go 
    ./build/app.exe
```

then the app will run automatically.

> Note : if you moved the executable into another directory make sure that the .env file is in the same directory as the .exe.

if your moving the executable into another directory run these commands before starting the app

```bash
    mkdire database/db/
```

**TODO : make the app initilize the directory withount needing any manual setup**