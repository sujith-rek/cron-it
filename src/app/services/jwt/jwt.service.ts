import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  constructor() { }

  allowedCookies = ['user'];

  //{ id: "3d49ec24-ecf5-47ea-86be-d25d26672b7a", email: "abc@gmail.com", name: "123456", limit: 3 }
  getCookie(name: string) {
    const value = `${document.cookie}`;
    const parts = value.split(`${name}=`);
    const cookieValue = parts[1];
    if (cookieValue) {
      // Decode the URL-encoded string and parse it as JSON
      const decodedValue = decodeURIComponent(cookieValue);
      try {
        const parsedValue = JSON.parse(decodedValue);
        return parsedValue;
      } catch (e) {
        return null;
      }
    }
    return null;
  }



}
