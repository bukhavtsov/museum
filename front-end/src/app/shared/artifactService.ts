// 1. model like separate file
// 2. service saves BehaviorSubject
// 3. components responsible for subscribing and getting info by service

import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";

export interface Artifact {
    id?: number
    creator?: string
    artifact_style?: string
    date_exc?: string
    transferred_by?: string
    artifact_measurement?: artifact_measurement
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
    private artifactNumber: number = 0; // temp var, is not work
    private artifactList: Artifact[] = [];
    private readonly getCardsURL = "http://localhost:8080/artifacts";

    constructor(private http: HttpClient) {
    }

    public getArtifactList(): Observable<Artifact[]> {
        return this.http.get<Artifact[]>(this.getCardsURL);
    }

    public getArtifactNumber(): number {
        return this.artifactNumber;
    }

    public getArtifact(id: number): Artifact {
        return this.artifactList.find(artifact => artifact.id === id)
    }

    public remove(id: number) {
        this.artifactNumber--;
        this.artifactList = this.artifactList.filter(event => event.id !== id);
    }

    public add(artifact: Artifact) {
        console.log("Hello from service");
        console.log(artifact);
        this.http.post<Artifact>(this.getCardsURL, artifact).subscribe({
            error: error => console.error('There was an error!', error)
        })
    }

    public edit(artifact: Artifact) {
        const id = this.artifactList.findIndex((event => event.id === artifact.id));
        this.artifactList[id] = artifact;
    }
}
