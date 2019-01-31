import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ajax } from 'rxjs/ajax';

// import { gapi } from 'gapi';

// const CLIENT_ID = '404364039745-0caba0fvhaja2cogru4jvl0gqq3anf50.apps.googleusercontent.com';
const CLIENT_ID = '404364039745-0caba0fvhaja2cogru4jvl0gqq3anf50.apps.googleusercontent.com';
const API_KEY = 'AIzaSyAzcUTdn8odVZoOItYURXDwbSKB2se6HbQ';
const DISCOVERY_DOCS = ['https://www.googleapis.com/discovery/v1/apis/drive/v3/rest'];
const SCOPES = 'https://www.googleapis.com/auth/drive';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  user: any;
  googleAuth: gapi.auth2.GoogleAuth;



initClient() {
    return new Promise((resolve, reject) => {
        // gapi.load('auth2', () => {
        //     return gapi.auth2.init({

        //         client_id: CLIENT_ID,
        //         redirect_uri: 'http://localhost:3001',
        //         scope: SCOPES,
        //     }).then(() => {
        //         this.googleAuth = gapi.auth2.getAuthInstance();
        //         this.signIn();

        //         resolve();
        //     });
        // });
        gapi.load('client:auth2', () => {
                return gapi.client.init({
                    apiKey: API_KEY,
                    clientId: CLIENT_ID,
                    discoveryDocs: DISCOVERY_DOCS,
                    scope: SCOPES,
                }).then(() => {
                    this.googleAuth = gapi.auth2.getAuthInstance();
                    resolve();

                    this.signIn();
                });
            });
    });
}
signIn() {
    // return this.googleAuth.signIn({
    //     prompt: 'consent',
    //     scope: 'profile'
    // }).then((googleUser: gapi.auth2.GoogleUser) => {
    //     console.log(googleUser);
    // });
    return this.googleAuth.signIn({
            prompt: 'consent'
        }).then((googleUser: gapi.auth2.GoogleUser) => {
          console.log(googleUser);
        });
}

    constructor() {
      this.initClient();
    }
  sample(): void {
    console.log('sample');
  }
}
