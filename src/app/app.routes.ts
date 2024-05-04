import { Routes } from '@angular/router';
import { AuthComponent } from './components/auth/auth.component';
import { AppComponent } from './app.component';

export const routes: Routes = [{
    path: '',
    component: AppComponent
}, {
    path: 'auth',
    component: AuthComponent
}];
