import {FlatTreeControl} from '@angular/cdk/tree';
import {Component, Injectable} from '@angular/core';
import {MatTreeFlatDataSource, MatTreeFlattener} from '@angular/material/tree';
import {BehaviorSubject} from 'rxjs';
import {ArtifactElement} from '../../../../shared/artifactService';
import {ArtifactElementsDatasource} from '../../../../shared/artifact-elements-datasource.service';

/** Flat Artifact Element name node with expandable and level information */
export class ArtifactElementFlatNode {
    name: string;
    level: number;
    expandable: boolean;
}

@Component({
    selector: 'app-artifacts-elements-create',
    templateUrl: `artifacts-elements-create.component.html`,
    styleUrls: ['artifacts-elements-create.component.scss'],
})
export class ArtifactsElementsCreateComponent {
    /** Map from flat node to nested node. This helps us finding the nested node to be modified */
    flatNodeMap = new Map<ArtifactElementFlatNode, ArtifactElement>();

    /** Map from nested node to flattened node. This helps us to keep the same object for selection */
    nestedNodeMap = new Map<ArtifactElement, ArtifactElementFlatNode>();

    /** The new name's name */
    newElementName = '';

    treeControl: FlatTreeControl<ArtifactElementFlatNode>;

    treeFlattener: MatTreeFlattener<ArtifactElement, ArtifactElementFlatNode>;

    dataSource: MatTreeFlatDataSource<ArtifactElement, ArtifactElementFlatNode>;

    constructor(private _database: ArtifactElementsDatasource) {
        this.treeFlattener = new MatTreeFlattener(this.transformer, this.getLevel,
            this.alwaysExpandable, this.getChildren);
        this.treeControl = new FlatTreeControl<ArtifactElementFlatNode>(this.getLevel, this.alwaysExpandable);
        this.dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

        _database.dataChange.subscribe(data => {
            this.dataSource.data = data;
        });
    }

    getLevel = (node: ArtifactElementFlatNode) => node.level;

    alwaysExpandable = (node: ArtifactElementFlatNode) => true;

    getChildren = (node: ArtifactElement): ArtifactElement[] => node.children;

    alwaysAllowCreate = (_: number, _nodeData: ArtifactElementFlatNode) => true;

    hasNoContent = (_: number, _nodeData: ArtifactElementFlatNode) => _nodeData.name === '';

    /**
     * Transformer to convert nested node to flat node. Record the nodes in maps for later use.
     */
    transformer = (node: ArtifactElement, level: number) => {
        const existingNode = this.nestedNodeMap.get(node);
        const flatNode = existingNode && existingNode.name === node.name
            ? existingNode
            : new ArtifactElementFlatNode();
        flatNode.name = node.name;
        flatNode.level = level;
        flatNode.expandable = !!node.children?.length;
        this.flatNodeMap.set(flatNode, node);
        this.nestedNodeMap.set(node, flatNode);
        return flatNode;
    }

    /* Get the parent node of a node */
    getParentNode(node: ArtifactElementFlatNode): ArtifactElementFlatNode | null {
        const currentLevel = this.getLevel(node);

        if (currentLevel < 1) {
            return null;
        }

        const startIndex = this.treeControl.dataNodes.indexOf(node) - 1;

        for (let i = startIndex; i >= 0; i--) {
            const currentNode = this.treeControl.dataNodes[i];

            if (this.getLevel(currentNode) < currentLevel) {
                return currentNode;
            }
        }
        return null;
    }

    /** Select the category so we can insert the new name. */
    addNewElement(node: ArtifactElementFlatNode) {
        const parentNode = this.flatNodeMap.get(node);
        if (parentNode.children === undefined) {
            parentNode.children = []
        }
        this._database.insertElement(parentNode!, '');
        this.treeControl.expand(node);
    }

    /** Save the node to database */
    saveNode(node: ArtifactElementFlatNode, elementValue: string) {
        const nestedNode = this.flatNodeMap.get(node);
        this._database.updateElement(nestedNode!, elementValue);
    }
}
