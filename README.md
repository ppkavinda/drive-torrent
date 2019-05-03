# Drive torrent

Simple torrent client to download torrents into Google Drive. Also we maintain a Drive link library for easy to download torrent.

Demo: [Drive-torrent.tk:3000](http://drive-torrent.tk:3000)

Inspired by [cloud-torrent](https://github.com/jpillora/cloud-torrent)

## Getting Started

Drive torrent is written in [Go1.10.4](https://golang.org) and [Vue](https://vuejs.org) On top of [anacrolix/torrent](https://github.com/anacrolix/torrent).These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

* [Node.js](https://nodejs.org/)
* [Go 1.10.4](https://github.com/golang/go/releases/tag/go1.10.4) (This version is required)
* [Firebase](firebase.google.com) Project with [Flame Plan](https://firebase.google.com/pricing/) or [Blaze Plan](https://firebase.google.com/pricing/) ( Not Spark Plan )
* [Algolia](https://algolia.com) Account (Free Plan is ok)

### Installation

#### Go server

Download Dependancy

```
go get github.com/ppkavinda/drive-torrent

go run main.go
```

Test it , Login to http://localhost:3000 from a web browser


#### Vue Pages

open **static/config/firebase.sample.config** and add your firebase project credentials

Change the directory and install dependancies

```
cd front
npm install
npm run dev
```

#### Sync YTS Films torrent Database to firebase

save your firebase service.json in **syncer** directory

open **syncer/server.js** and paste your Database url and add how many pages you need to feed to your database

next install dependancies and run the script

```
cd syncer
npm install
node server.js
```

#### Firebase Functions

With firebase functions we sync our film Database to Algolia to proper indexing.

For a getting start with Cloud Functions follow [this guide](https://firebase.google.com/docs/functions/get-started)

```
cd functions
npm install -g firebase-tools
firebase login
npm install
firebase deploy
```

## Deployment
You can easiely deploy with docker image

```
docker build
docker run drive-torrent
```

## Authors

* Prasad Kavinda [https://github.com/ppkavinda](https://github.com/ppkavinda)
* Danushka Herath [https://github.com/danushka96](https://github.com/danushka96)

PRs are welcome
