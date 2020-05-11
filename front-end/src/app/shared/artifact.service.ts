// 1. model like separate file
// 2. service saves BehaviorSubject
// 3. components responsible for subscribing and getting info by service

import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';

export interface Artifact {
    id?: number;
    creator?: string;
    artifact_style?: string;
    date_exc?: string;
    transferred_by?: string;
    artifact_measurement?: ArtifactMeasurement;
    artifact_elements?: ArtifactElement[]
    artifact_materials?: ArtifactMaterial[]

}

export interface ArtifactElement {
    id: number
    artifact_id: number
    name: string
    parent_id: number
    child_elements?: ArtifactElement[]
}

export interface ArtifactMaterial {
    id: number
    artifact_id: number
    quantity?: number
    composition?: number
    material_type?: string
    parent_id: number
    child_materials?: ArtifactMaterial[]
}


export interface ArtifactMeasurement {
    width?: number
    height?: number
    length?: number
}

@Injectable({
    providedIn: 'root'
})
export class ArtifactService {
    private artifactList: Artifact[] = [];
    private readonly getCardsURL = 'http://localhost:8080/artifacts';

    constructor(private http: HttpClient) {
    }

    public getArtifactList(): Observable<Artifact[]> {
        return this.http.get<Artifact[]>(this.getCardsURL);
    }

    public getArtifact(id: number): Artifact {
        return this.artifactList.find(artifact => artifact.id === id)
    }

    public remove(id: number) {
        this.artifactList = this.artifactList.filter(event => event.id !== id);
    }

    public add(artifact: Artifact) {
        this.artifactList.push(artifact);
    }

    public edit(artifact: Artifact) {
        const id = this.artifactList.findIndex((event => event.id === artifact.id));
        this.artifactList[id] = artifact;
    }
}
