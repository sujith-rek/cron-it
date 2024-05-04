import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { FormsModule } from '@angular/forms';
import cookie from 'cookie';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './auth.component.html',
  styleUrl: './auth.component.scss'
})
export class AuthComponent {
  constructor(private authService: AuthService) { }

  name: string = '';
  email: string = '';
  password: string = '';
  confirmPassword: string = '';

  signUp() {
    if (this.password !== this.confirmPassword) {
      alert('Passwords do not match');
      return;
    }

    if (this.name === '' || this.email === '' || this.password === '') {
      alert('All fields are required');
      return;
    }

    this.authService.signUp({
      email: this.email,
      password: this.password,
      name: this.name
    }).subscribe({
      next: (res) => {
        alert('User registered successfully');
        console.log(res);
      },
      error: (err) => {
        alert(err.error.message);
        console.log(err);
      }
    });
  }

  signIn() {
    if (this.email === '' || this.password === '') {
      alert('All fields are required');
      return;
    }

    this.authService.signIn({
      email: this.email,
      password: this.password
    }).subscribe({
      next: (res) => {
        alert('User logged in successfully');
        console.log(res);
      },
      error: (err) => {
        alert(err.error.message);
        console.log(err);
      }
    });
  }

  logout() {
    this.authService.logout().subscribe({
      next: (res) => {
        alert('User logged out successfully');
        console.log(res);
      },
      error: (err) => {
        alert(err.error.message);
        console.log(err);
      }
    });
  }

  readMyCookie() {
    console.log(cookie.parse(document.cookie));
  }


}
