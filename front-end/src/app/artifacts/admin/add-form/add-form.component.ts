import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {FormControl, FormGroup} from '@angular/forms';

import {Artifact, ArtifactService} from '../../../shared/artifactService';

@Component({
    selector: 'app-add-form',
    templateUrl: './add-form.component.html',
})

export class AddFormComponent implements OnInit {

    private addForm: FormGroup;

    constructor(private artifactService: ArtifactService,
                private router: Router
    ) {
    }

    ngOnInit() {
        this.initForm()
    }

    private initForm() {
        this.addForm = new FormGroup({
            creator: new FormControl(''),
            artifact_style: new FormControl(''),
            date_exc: new FormControl(''),
            transferred_by: new FormControl(''),
        })
    }

    addArtifact(artifact: Artifact) {
        console.log("hello from form");
        //artifact.id = this.artifactService.getArtifact(this.artifactService.getArtifactNumber() - 1).id + 1;
        this.artifactService.add(artifact);
        this.router.navigate(['/artifact'])
    }

}
