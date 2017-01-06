### This is the go-fact app

#### To run

1. Bring up the cluser with docker-compose
```
docker-compose build && up -d
```

2. This runs the go-fact-api and nginx


#### Continuous integration

We use circle-ci to build the api (zahid/go-facts-api) and upload the image to docker hub. It is triggered on each push to the master branch.