import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ViewArtifactComponent } from './artifacts/view-artifact/view-artifact.component';
import { ArtifactListComponent } from './artifacts/artifact-list/artifact-list.component';
import { EditFormComponent } from './artifacts/admin/edit-form/edit-form.component';
import { AddFormComponent } from './artifacts/admin/add-form/add-form.component';


const routes: Routes = [
  { path: '', redirectTo: '/artifacts', pathMatch: 'full' },
  { path: 'artifacts', component: ArtifactListComponent },
  { path: 'artifacts/edit/:id', component: EditFormComponent },
  { path: 'artifacts/add', component: AddFormComponent },
  { path: 'artifacts/:id', component: ViewArtifactComponent },
  { path: '**', redirectTo: '/artifacts', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
