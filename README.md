# liquide-assignment
The repository provides an application that allows authenticated users to access a blog using rest API calls.
The assignment covers
* creating login mechanism for users
* post creation
* upvote and post functionalities

There are a few assumptions and borders drawn around the application in the interest of time and scoping. 
* Posts can only have text and no resources like images or videos in them.
* only posts can be voted. comments cannot be voted
* once a post is upvoted, it can only be downvoted by the user. the vote cannot be removed
* admin can delete any post they want but users can delete only their post 

### Open improvements to do
* Add snowflake id to posts, comments in db
* prepare redis write through for posts and comments
* prepare a complex post score for dynamic feed
* add vote management in redis and write to db in batches

## postman-collection
https://dark-station-154939.postman.co/workspace/Liquide-Assignment~bb764463-0583-4784-8c62-78b3cc1f5fd9/collection/32061457-f46f94fd-310e-4346-a873-1ceae92587bd?action=share&creator=32061457

## Getting Started

To get started to be using this project, the pre-requisites would be needed

### Prerequisites
- Go 1.20+
- Docker (optional, for containerization and/or database)
- make installation

