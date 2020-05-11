import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

import {Artifact, ArtifactService} from '../../../shared/artifact.service'

@Component({
    selector: 'app-edit-form',
    templateUrl: './edit-form.component.html',
})

export class EditFormComponent implements OnInit {

    private editForm: FormGroup;
    private artifact: Artifact;

    constructor(private artifactService: ArtifactService,
                private formBuilder: FormBuilder,
                private activatedRoute: ActivatedRoute,
                private router: Router
    ) {
    }

    ngOnInit() {
        let id = this.activatedRoute.snapshot.paramMap.get('id');
        this.artifact = this.artifactService.getArtifact(Number(id));
        this.initForm()
    }

    initForm() {
        this.editForm = this.formBuilder.group({
            id: this.artifact.id,
            date: "",
            type: "",
            title: ["", Validators.compose([
                Validators.required,
                Validators.minLength(10)
            ])],
            image: "",
            text: ""
        })
    }

    editArifact(artifact: Artifact) {
        this.artifactService.edit(artifact)
        this.router.navigate(['/artifact'])
    }

}

