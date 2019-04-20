const express = require("express");
const bodyParser =  require('body-parser');

const app = express();
const router = express.Router();
const admin = require("firebase-admin");
const serviceAccount = require("./service.json");
const request = require("request");

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true
}));

admin.initializeApp({
    credential: admin.credential.cert(serviceAccount),
    databaseURL: 'YOUR-DATABASE_URL'
});

const db = admin.firestore();

app.get('/', (req, res) => res.send('Hello World!'));

setInterval(function() {
    syncDB();
}, 300000);

// var object = {
//     "sample":"sample"
// }
// db.collection('films').doc('100').set(object);
syncDB()
// getData('9623')

async function syncDB(){

    // QPX REST API URL (I censored my api key)
    var url = "https://yts.am/api/v2/list_movies.json?page=";

    // fire request
    for(var j=1;j<562;j++){
        await request({
            url: url+j,
            json: true,
            method: 'GET'
        }, function (error, response, body) {
            if (!error && response.statusCode === 200) {
                for(var i=0;i<body['data']['movies'].length;i++){
                    console.log(body['data']['movies'][i].id);
                    firestoreSave(body['data']['movies'][i]);
                    // db.collection('films').doc(`${body['data']['movies'][i].id}`).set(body['data']['movies'][i]);
                    //console.log(body['Products'][i]);
                }
                //console.log('Syncing Films...')
            }
            else {
                console.log("error: " + error)
                console.log("response.statusCode: " + response.statusCode)
                console.log("response.statusText: " + response.statusText)
            }
        });
        //console.log('finished')
    }   
}

async function firestoreSave(object){
    await db.collection('films').doc(`${object.id}`).set(object);
}

async function getData(id){
    var docRef = db.collection('films').doc(id);
    docRef.get().then(function(doc){
        if(doc.exists){
            console.log('Data '+doc.data().id)
        }else{
            console.log('No such Doucment');
        }
    })
}

app.listen(4000, () => console.log('Express server running on port 4000'));

