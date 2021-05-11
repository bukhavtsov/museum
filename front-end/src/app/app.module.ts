import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';

import {ArtifactFilterPipe} from './shared/artifact-filter.pipe';

import {AppComponent} from './app.component';

import {HeaderComponent} from './shared/header/header.component';
import {FooterComponent} from './shared/footer/footer.component';
import {MainComponent} from './shared/main/main.component';

import {ArtifactsComponent} from './artifacts/artifacts.component';

import {AdminPanelComponent} from './artifacts/admin/admin-panel/admin-panel.component';
import {AddFormComponent} from './artifacts/admin/add-form/add-form.component';
import {ViewArtifactComponent} from './artifacts/view-artifact/view-artifact.component';
import {EditFormComponent} from './artifacts/admin/edit-form/edit-form.component';
import {ArtifactListComponent} from './artifacts/artifact-list/artifact-list.component';
import {HttpClientModule} from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatTreeModule} from '@angular/material/tree';


@NgModule({
    declarations: [
        AppComponent,
        HeaderComponent,
        FooterComponent,
        MainComponent,
        ArtifactsComponent,
        ArtifactFilterPipe,
        AdminPanelComponent,
        AddFormComponent,
        ViewArtifactComponent,
        EditFormComponent,
        ArtifactListComponent,
    ],
    imports: [
        BrowserModule,
        MatTreeModule,
        AppRoutingModule,
        FormsModule,
        ReactiveFormsModule,
        HttpClientModule,
        BrowserAnimationsModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})

export class AppModule {
}
