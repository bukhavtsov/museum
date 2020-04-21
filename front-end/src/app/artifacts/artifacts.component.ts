import { Component } from '@angular/core';

import { ArtifactService } from '../shared/artifactService';

@Component({
  selector: 'app-artifacts',
  templateUrl: './artifacts.component.html',
})

export class ArtifactsComponent {

  constructor(private artifactService: ArtifactService) { }
}
