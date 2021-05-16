import {Component, Input, OnInit} from '@angular/core';

import {FlatTreeControl} from '@angular/cdk/tree';
import {MatTreeFlatDataSource, MatTreeFlattener} from '@angular/material/tree';

import {ArtifactElement} from '../../shared/artifactService'



/** Flat node with expandable and level information */
interface FlatNode {
    expandable: boolean;
    name: string;
    level: number;
}

@Component({
    selector: 'app-artifacts-elements-list',
    templateUrl: './artifacts-elements-list.component.html'
})

export class ArtifactsElementsListComponent implements OnInit {
    @Input() elements: ArtifactElement[];

    ngOnInit(): void {
        this.dataSource.data = this.elements;
    }

    private _transformer = (node: ArtifactElement, level: number) => {
        return {
            expandable: !!node.children && node.children.length > 0,
            name: node.name,
            level: level,
        };
    }

    treeControl = new FlatTreeControl<FlatNode>(
        node => node.level, node => node.expandable);

    treeFlattener = new MatTreeFlattener(
        this._transformer, node => node.level, node => node.expandable, node => node.children);

    dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);


    hasChild = (_: number, node: FlatNode) => node.expandable;
}
