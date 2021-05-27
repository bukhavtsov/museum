import {Component, OnDestroy, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {FormControl, FormGroup} from '@angular/forms';
import {Artifact, ArtifactElement, ArtifactService} from '../../../shared/artifactService';
import {Subscription} from 'rxjs';
import {ArtifactElementsDatasource} from '../../../shared/artifact-elements-datasource.service';

@Component({
    selector: 'app-add-form',
    templateUrl: './add-form.component.html',
    providers: [ArtifactElementsDatasource],
})

export class AddFormComponent implements OnInit, OnDestroy {

    private addForm: FormGroup;
    private artifactElements: ArtifactElement[];
    private subscription: Subscription;

    constructor(private artifactService: ArtifactService,
                private router: Router,
                private dataSource: ArtifactElementsDatasource,
    ) {
    }


    ngOnInit() {
        this.initForm()
        this.subscription = this.dataSource.dataChange.subscribe(elements => {
                this.artifactElements = elements
            }
        )
    }

    ngOnDestroy() {
        this.subscription.unsubscribe();
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
            artifact_elements: this.artifactElements
        };
        this.artifactService.add(newArtifact);
        this.router.navigate(['/artifacts'])
    }

}
