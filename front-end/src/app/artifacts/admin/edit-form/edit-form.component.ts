import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import {Artifact, ArtifactService} from '../../../shared/artifactService'

@Component({
    selector: 'app-edit-form',
    templateUrl: './edit-form.component.html',
})

export class EditFormComponent implements OnInit {

    private editForm: FormGroup;
    private artifact: Artifact;
    private id: string;
    constructor(private artifactService: ArtifactService,
                private formBuilder: FormBuilder,
                private activatedRoute: ActivatedRoute,
                private router: Router
    ) {
    }

    ngOnInit() {
        this.id = this.activatedRoute.snapshot.paramMap.get('id');
        // this.artifact = this.artifactService.getArtifact(Number(id));
        // console.log(this.artifact)
        this.initForm()
    }

    initForm() {
        // this.editForm = new FormGroup({
        //     creator: new FormControl(this.artifact.creator),
        //     artifact_style: new FormControl(this.artifact.artifact_style),
        //     date_exc: new FormControl(this.artifact.date_exc),
        //     transferred_by: new FormControl(this.artifact.transferred_by),
        //     length: new FormControl(this.artifact.ArtifactMeasurement.length),
        //     height: new FormControl(this.artifact.ArtifactMeasurement.height),
        //     width: new FormControl(this.artifact.ArtifactMeasurement.width),
        // })
        this.editForm = new FormGroup({
            creator: new FormControl(''),
            artifact_style: new FormControl(''),
            date_exc: new FormControl(''),
            transferred_by: new FormControl(''),
            length: new FormControl(''),
            height: new FormControl(''),
            width: new FormControl(''),
        })
    }

    editArtifact() {
        const newArtifact: Artifact = {
            creator: this.editForm.controls.creator.value,
            artifact_style: this.editForm.controls.artifact_style.value,
            date_exc: this.editForm.controls.date_exc.value,
            transferred_by: this.editForm.controls.transferred_by.value,
            artifact_measurement: {
                length: +this.editForm.controls.length.value,
                width: +this.editForm.controls.width.value,
                height: +this.editForm.controls.height.value,
            },
        };
        this.artifactService.edit(Number(this.id), newArtifact)
        this.router.navigate(['/artifacts'])
    }

}

