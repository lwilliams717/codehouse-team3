import { NgModule } from '@angular/core';

import { RouterModule, Routes } from '@angular/router';
import { LandingComponent } from './landing/landing.component';
import { TestComponent } from './test/test.component';
import { PeopleComponent } from './people/people.component';

const routes: Routes = [
    { path: '', component: LandingComponent },
    { path: 'testing', component: TestComponent },
    { path: 'people', component: PeopleComponent}
]; // helps with navigating

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }

