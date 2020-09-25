import {Component, OnInit} from '@angular/core';
import {Artifact, ArtifactService} from 'src/app/shared/artifactService';
import {SearchArtifactService} from 'src/app/shared/search-artifact.service';

@Component({
    selector: 'app-artifact-list',
    templateUrl: './artifact-list.component.html',
    styleUrls: ['./artifact-list.component.scss']
})
export class ArtifactListComponent implements OnInit {

    private isAdmin: boolean;
    private search: string;

    private cards: Artifact[] = [];
    showShortDescription: Map<number, boolean>;


    constructor(private artifactService: ArtifactService, private svcSearch: SearchArtifactService) {
    }

    ngOnInit(): void {
        this.isAdmin = true;
        this.svcSearch.sharedSearch.subscribe(search => this.search = search);
        this.initMap();
    }

    private initMap() {
        this.showShortDescription = new Map<number, boolean>();
        this.artifactService.getArtifactList().subscribe(data => this.cards = data);
        this.cards.forEach(artifact => {
            this.showShortDescription.set(artifact.id, false)
        });
    }

    private alterDescriptionText(id: number) {
        this.showShortDescription.set(id, !this.showShortDescription.get(id))
    }

}
