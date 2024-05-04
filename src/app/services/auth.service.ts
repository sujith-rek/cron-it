import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { constants } from '../constants';
import { SignInInput,SignUpInput } from '../types';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  signUp(data: SignUpInput) {
    return this.http.post(`${constants.backednUrl}/register`, data);
  }

  signIn(data: SignInInput) {
    return this.http.post(`${constants.backednUrl}/login`, data, { withCredentials: true });
  }

  refreshToken() {
    return this.http.post(`${constants.backednUrl}/refresh`, {}, { withCredentials: true });
  }

  logout() {
    return this.http.post(`${constants.backednUrl}/logout`, {}, { withCredentials: true });
  }

}
