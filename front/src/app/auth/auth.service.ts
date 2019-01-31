import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { ajax } from 'rxjs/ajax';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  user: any;

  constructor(private router: Router) {
    this.getAuth();
  }
  getAuth() {
    ajax('http://localhost:3000/user').subscribe(
      res => console.log(res)
    );
  }
  sample(): void {
    console.log('sample');
  }
}
