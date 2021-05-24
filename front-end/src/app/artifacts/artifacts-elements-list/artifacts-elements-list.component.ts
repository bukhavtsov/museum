import { NestedTreeControl } from '@angular/cdk/tree';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { SelectionModel } from '@angular/cdk/collections';
import { FlatTreeControl } from '@angular/cdk/tree';
import {AfterViewInit, Component, Injectable, Input} from '@angular/core';
import {
    MatTreeFlatDataSource,
    MatTreeFlattener
} from '@angular/material/tree';
import { BehaviorSubject } from 'rxjs';
import {ArtifactElement} from '../../shared/artifactService';

/**
 * Food data with nested structure.
 * Each node has a name and an optiona list of children.
 */
interface ArtifactElementNode {
    name: string;
    children?: ArtifactElementNode[];
    ok?: boolean;
}

/**
 * @title Tree with nested nodes
 */
@Component({
    selector: 'app-artifacts-elements-list',
    templateUrl: 'artifacts-elements-list.component.html',
    styleUrls: ['artifacts-elements-list.component.scss']
})
export class ArtifactsElementsListComponent implements AfterViewInit{
    @Input() elements: ArtifactElement[];
    treeControl = new NestedTreeControl<ArtifactElementNode>(node => node.children);
    dataSource = new MatTreeNestedDataSource<ArtifactElementNode>();

    constructor() {}

    ngAfterViewInit() {
        this.dataSource.data = this.elements;
        Object.keys(this.dataSource.data).forEach(x => {
            this.setParent(this.dataSource.data[x], null);
        });
    }

    hasChild = (_: number, node: ArtifactElementNode) =>
        !!node.children && node.children.length > 0;

    setParent(data, parent) {
        data.parent = parent;
        if (data.children) {
            data.children.forEach(x => {
                this.setParent(x, data);
            });
        }
    }

    checkAllParents(node) {
        if (node.parent) {
            const descendants = this.treeControl.getDescendants(node.parent);
            this.checkAllParents(node.parent);
        }
    }

    setChildOk(text: string, node: any) {
        node.forEach(x => {
            x.ok = x.name.indexOf(text) >= 0;
            if (x.parent) this.setParentOk(text, x.parent, x.ok);
            if (x.children) this.setChildOk(text, x.children);
        });
    }
    setParentOk(text, node, ok) {
        node.ok = ok || node.ok || node.name.indexOf(text) >= 0;
        if (node.parent) this.setParentOk(text, node.parent, node.ok);
    }
}

