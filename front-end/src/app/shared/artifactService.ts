// 1. model like separate file
// 2. service saves BehaviorSubject
// 3. components responsible for subscribing and getting info by service

import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";


export interface Artifact {
    id?: number
    creator: string
    artifact_style: string
    date_exc: string
    transferred_by: string
    artifact_measurement: artifact_measurement
    artifact_elements?: Map<string, string[]>
    artifact_object_group?: Map<string, string[]>
    artifact_materials?: Map<string, string[]>
    artifact_preservation?: Map<string, string[]>
}

export interface artifact_measurement {
    width?: number
    height?: number
    length?: number
}

@Injectable({
    providedIn: 'root'
})
export class ArtifactService {
    private artifactList: Artifact[] = [];
    private readonly artifactsURL = 'http://localhost:8080/artifacts';

    constructor(private http: HttpClient) {
    }

    public getArtifactList(): Observable<Artifact[]> {
        return this.http.get<Artifact[]>(this.artifactsURL);
    }


    public getArtifact(id: number): Artifact {
        return this.artifactList.find(artifact => artifact.id === id)
    }

    public remove(id: number) {
        this.artifactList = this.artifactList.filter(event => event.id !== id);
    }

    public add(artifact: Artifact) {
        this.http.post<Artifact>(this.artifactsURL, artifact).subscribe({
            error: error => console.error('There was an error!', error)
        })
    }

    public edit(artifactID: number, newArtifact: Artifact) {
        this.http.put<Artifact>(this.artifactsURL + `/${artifactID}`, newArtifact).subscribe({
            error: error => console.error('There was an error!', error)
        })
    }
    public delete(artifactID: number) {
        this.http.delete<Artifact>(this.artifactsURL + `/${artifactID}`).subscribe({
            error: error => console.error('There was an error!', error)
        })
    }
}
