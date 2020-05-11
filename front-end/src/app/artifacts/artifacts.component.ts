import { Component } from '@angular/core';

import { ArtifactService } from '../shared/artifact.service';

@Component({
  selector: 'app-artifacts',
  templateUrl: './artifacts.component.html',
})

export class ArtifactsComponent {

  constructor(private artifactService: ArtifactService) { }
}
