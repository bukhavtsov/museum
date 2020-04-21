import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, Validators, FormControl } from '@angular/forms';

import { ArtifactService, Artifact } from '../../../shared/artifactService';

@Component({
  selector: 'app-add-form',
  templateUrl: './add-form.component.html',
})

export class AddFormComponent implements OnInit {

  private addForm: FormGroup;

  constructor(private artifactService: ArtifactService,
    private router: Router
  ) {}

  ngOnInit() {
    this.initForm()
  }

  private initForm() {
    this.addForm = new FormGroup({
      date: new FormControl(''),
      type: new FormControl(''),
      title: new FormControl('', Validators.compose([
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(100),
      ])),
      text: new FormControl('Text example'),
      image: new FormControl(''),
    })
  }

  addArtifact(artifact: Artifact) {
    // artifact.id = this.artifactService.getArtifact(this.artifactService.getArtifactList().length - 1).id + 1;
    // this.artifactService.add(artifact);
    //
    // this.router.navigate(['/artifact'])
  }

}
