import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';

import {Artifact, ArtifactService} from '../../shared/artifact.service'

@Component({
    selector: 'app-view-artifact',
    templateUrl: './view-artifact.component.html',
})
export class ViewArtifactComponent implements OnInit {

    private artifact: Artifact;

    constructor(private artifactService: ArtifactService, private route: ActivatedRoute) {
    }

    ngOnInit() {
        let id = this.route.snapshot.paramMap.get('id');
        this.artifact = this.artifactService.getArtifact(Number(id));
    }
}
