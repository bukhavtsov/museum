import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {

  private artifactsSource: string;
  constructor() { }

  private receiveArtifactsSource($event: string) {
    this.artifactsSource = $event;
  }

}
