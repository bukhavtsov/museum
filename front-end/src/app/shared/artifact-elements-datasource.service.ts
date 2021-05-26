import {Injectable} from '@angular/core';
import {BehaviorSubject} from 'rxjs';
import {ArtifactElement} from './artifactService';

@Injectable()
export class ArtifactElementsDatasource {
    dataChange = new BehaviorSubject<ArtifactElement[]>([]);

    get data(): ArtifactElement[] {
        return this.dataChange.value;
    }

    constructor() {
        this.initialize();
    }

    initialize() {
        const data = [{name: ''}]
        this.dataChange.next(data);
    }

    insertElement(parent: ArtifactElement, name: string) {
        if (parent.children) {
            parent.children.push({name: name} as ArtifactElement);
            this.dataChange.next(this.data);
        }
    }

    updateElement(node: ArtifactElement, name: string) {
        node.name = name;
        this.dataChange.next(this.data);
    }
}
