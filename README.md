### Main features

## Application cycle

## Endpoints

- "/login" & "/signup"  & "/refresh/" endpoints for user logic
- "/chat/" endpoint for the actual app :
    - "/chat/rooms" Shows all the public rooms
    - "/chat/@me/dms" 
    - "/chat/{room_id}"
    - "/chat/{room_id}/messages" if public
    - "/chat/{room_id}/members" if public
    - "/chat/{room_id}/owner" if public
    - "/chat/{room_id}/join" require an invitation key if public
- "/users/" get all users, **Note: this is an OnlyAdmin endpoint** :
    - "/users/{user_id}": get public data about a user 
    - "/users/@me": get information about the current user after providing an acces_token
    - "/users/@me/chatrooms"
    - "/users/@me/friends" 

## Encryption strategies
Most encryption will happen at the client side so the server will be reciving just some random text that can be decrypted, but the server side will be 
responsible for orginizing group chats and public groups  