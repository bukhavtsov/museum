import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';

import { SearchArtifactService } from 'src/app/shared/search-artifact.service';

@Component({
  selector: 'app-admin-panel',
  templateUrl: './admin-panel.component.html',
})
export class AdminPanelComponent implements OnInit {

  private search: string;
  private searchForm: FormGroup;
  @Output() ArtifactSource = new EventEmitter<string>();

  constructor(private svcSearch: SearchArtifactService,
    private formBuilder: FormBuilder,
  ) { }

  ngOnInit(): void {
    this.initForm();
    this.ArtifactSource.emit('Local')
    this.svcSearch.sharedSearch.subscribe(search => this.search = search);
  }

  submitSearch(value: { title: string; }) {
    this.svcSearch.nextSearch(value.title)
  }

  initForm() {
    this.searchForm = this.formBuilder.group({
      title: ''
    })
  }

  changeArtifactSource(artifactSource: string): void {
    this.ArtifactSource.emit(artifactSource)
  }
}
