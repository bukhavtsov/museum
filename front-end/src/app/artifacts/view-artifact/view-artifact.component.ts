import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';

import {Artifact, ArtifactService} from '../../shared/artifact.service'

@Component({
    selector: 'app-view-artifact',
    templateUrl: `<h1></h1>`
})
export class ViewArtifactComponent implements OnInit {

    private artifact: Artifact;

    constructor(private artifactService: ArtifactService, private route: ActivatedRoute) {
    }

    ngOnInit() {
    }
}
