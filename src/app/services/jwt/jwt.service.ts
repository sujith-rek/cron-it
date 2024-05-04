import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  constructor() { }

  getCookie(name: string): any {
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
