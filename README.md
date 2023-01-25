<div id="top"></div>

<!--
*** Inspired by the Best-README-Template.
*** Let's create something AMAZING! :D

*** GitLab Flavored Markdown - https://gitlab.com/gitlab-org/gitlab/-/blob/master/doc/user/markdown.md
-->


<div align="center">
  <h1>Video-Manager</h1>
</div>

## 📍 About The Project

`Video-Manager` is a service that helps to manage video annotations

### Databases

- postgresql://postgres:secret@0.0.0.0:5432/video-manager?sslmode=disable

![DB MODEL](https://github.com/tobslob/Video-manager/blob/main/database-model.png?raw=true)


# Running locally

To start the application in dev mode, please run:

```sh
git clone https://github.com/tobslob/video-manager.git
```

```sh
cd video-manager
```

```sh
 go install
```

## Initial setup

Install PostgreSQL on local machine using the following command:

```sh
docker pull postgres

## 1. We will create a local folder and mount it as a data volume for our running container to store all the database files in a known location.

mkdir ${HOME}/postgres-data/
## 2. run the postgres image

docker run -d --name dev-postgres \
 --restart=always \
 -e POSTGRES_PASSWORD=secret \
 -e POSTGRES_USER=postgres \
 -e POSTGRES_DB=video-manager \
 -v ${HOME}/postgres-data/:/var/lib/postgresql/data \
 -p 5432:5432 postgres
## 3. check that the container is running
docker ps

```

```sh
 make postgres - Command to run after you have postgres container runing
 ```
 ```sh
 make createdb - Command help to creates a DB name video-manager 
 ```
 ```sh
 make migrateup - Command creates table into the database
 ```
 ```sh
 make test - Command help to run unit test
 ```
```sh
 make server - Command help to start server
```

```sh
Application is ready to receive connection @ http://localhost:8080
```
```sh
## API Documentation
https://documenter.getpostman.com/view/6225567/2s8ZDcyKSF
```
```sh
## API-ENDPOINTS

- V1

`- POST /api/v1/users Create user account`

`- POST /api/v1/users/login Login a user`

`- POST /api/v1/videos Create a video`

`- GET /api/v1/videos/<:id> Get a Annotation`

`- DELETE /api/v1/video/<:id> Delete a video`

`- POST /api/v1/annotations Create a video annotation`

`- GET /api/v1/annotations/<:id> Get a Annotation`

`- GET /api/v1/annotations/<:video_id>?page_id=1&page_size=5 Get A video Annotations`

`- DELETE /api/v1/annotations/<:id> Delete Annotation`

`- PATCH /api/v1/annotations/<:id>/<:video_id> Update an annotation`
```

## Run using Docker Image