import {Pipe, PipeTransform} from '@angular/core';

import {Artifact} from './artifactService'

@Pipe({
    name: 'artifactFilter'
})

export class ArtifactFilterPipe implements PipeTransform {

    transform(artifacts: Artifact[], search: string): Artifact[] {
        if (!search.trim()) {
            return artifacts;
        }
        return artifacts.filter(artifact => {
            return String(artifact.id).toLocaleLowerCase().indexOf(search.toLocaleLowerCase()) !== -1;
        });
    }
}
