const functions = require('firebase-functions');
const admin = require('firebase-admin');
let promise = require('promise');
const cors = require('cors')({origin: true});
const auth = require('basic-auth');
const reqest = require('request');
const algoliasearch = require('algoliasearch');
admin.initializeApp(functions.config().firebase);
const db = admin.firestore();

// Listen for a new Record Add
exports.addNewFilm = functions.firestore.document('films/{document}').onCreate(event => {
    console.log('Adding new Film..',event)

    return addToAlgolia(event.data(), 'films')
        .then(res => console.log('Added to algolia', res))
        .catch(err => console.log('Error on Algolia', err));
});

// Listen for a Edit record
exports.editFilm = functions.firestore.document('films/{document}').onUpdate(event => {
    console.log('Editing record..', event)

    return editToAlgolia(event.data(),'films')
        .then(res => console.log('Edit Success in Algolia', res))
        .catch(err => console.log('Error in update in algolia', err));
})

// Add to algolia
function addToAlgolia(object, indexName){
    const ALGOLIA_ID = functions.config().algolia.app_id;
    const ALGOLIA_ADMIN_KEY = functions.config().algolia.api_key;
    const client = algoliasearch(ALGOLIA_ID, ALGOLIA_ADMIN_KEY);
    const index = client.initIndex(indexName);

    return new Promise((resolve, reject) => {
        index.addObject(object)
            .then(res => { console.log('res GOOD', res); resolve(res); return; })
            .catch(err => { console.log('err BAD', err); reject(err) });
    })
}

// Edit Algolia record
function editToAlgolia(object, indexName) {
    const ALGOLIA_ID = functions.config().algolia.app_id;
    const ALGOLIA_ADMIN_KEY = functions.config().algolia.api_key;
    const client = algoliasearch(ALGOLIA_ID, ALGOLIA_ADMIN_KEY);
    const index = client.initIndex(indexName);

    return new Promise((resolve, reject) => {
        index.saveObject(object)
        .then(res => { console.log('res GOOD', res); resolve(res); return })
        .catch(err => { console.log('err BAD', err); reject(err) });
    });
}

