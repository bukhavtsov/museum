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
    private maxArtifactsQty = 0;
    private currentArtifactQty: number;
    private difference: number;
    showShortDesciption: Map<number, boolean>;


    constructor(private artifactService: ArtifactService, private svcSearch: SearchArtifactService) {
    }

    ngOnInit(): void {
        this.difference = 5;
        this.currentArtifactQty = this.difference;
        this.isAdmin = true;

        this.svcSearch.sharedSearch.subscribe(search => this.search = search);

        this.initMap();
    }

    private loadMore() {
        if (this.currentArtifactQty + this.difference < this.maxArtifactsQty) {
            this.currentArtifactQty += this.difference;
        } else if (this.currentArtifactQty < this.maxArtifactsQty) {
            this.currentArtifactQty = this.maxArtifactsQty
        }
    }

    private initMap() {
        this.showShortDesciption = new Map<number, boolean>();
        this.artifactService.getArtifactList().subscribe(data => this.cards = data);
        this.cards.forEach(artifact => {
            this.showShortDesciption.set(artifact.id, false)
        });
        this.maxArtifactsQty = this.cards.length;
    }

    private alterDescriptionText(id: number) {
        this.showShortDesciption.set(id, !this.showShortDesciption.get(id))
    }

}
