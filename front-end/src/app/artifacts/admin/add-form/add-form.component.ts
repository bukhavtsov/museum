import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import { FormControl, FormGroup } from '@angular/forms';
import {ArtifactsElementsCreateComponent} from './artifacts-elements-create/artifacts-elements-create.component';
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
            creator: new FormControl(),
            artifact_style: new FormControl(''),
            date_exc: new FormControl(''),
            transferred_by: new FormControl(''),
            length: new FormControl(''),
            height: new FormControl(''),
            width: new FormControl(''),
        })
    }

    addArtifact() {
        const newArtifact: Artifact = {
            creator: this.addForm.controls.creator.value,
            artifact_style: this.addForm.controls.artifact_style.value,
            date_exc: this.addForm.controls.date_exc.value,
            transferred_by: this.addForm.controls.transferred_by.value,
            artifact_measurement: {
                length: +this.addForm.controls.length.value,
                width: +this.addForm.controls.width.value,
                height: +this.addForm.controls.height.value,
            },
        };
        this.artifactService.add(newArtifact);
        this.router.navigate(['/artifacts'])
    }

}
