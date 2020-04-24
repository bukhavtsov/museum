package data

const getArtifactsWithBasicInfoQuery = `
SELECT artifact_master.id,
       artifact_master.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by,
       artifact_master.date_exc,
       artifact_measurement.height,
       artifact_measurement.width,
       artifact_measurement.length
FROM artifact_master
         INNER JOIN transferred_by_lut
                    ON (artifact_master.transferred_by_id = transferred_by_lut.id)
         LEFT JOIN
     (
         SELECT artifact_id, artifact_style_name
         FROM artifact_style
                  INNER JOIN artifact_style_lut on (artifact_style.id = artifact_style_lut.id)
     ) as artifact_style
     on (artifact_master.id = artifact_style.artifact_id)
         INNER JOIN artifact_measurement on artifact_master.id = artifact_measurement.artifact_id
`
const getArtifactElementsQuery = `
SELECT ae1.artifact_id, ae1.artifact_element_name, ae2.artifact_element_name
FROM artifact_element ae1
         LEFT JOIN artifact_element ae2 ON ae1.artifact_parent_element_id = ae2.id
`

const getArtifactElementByIdQuery = getArtifactElementsQuery + " WHERE ae1.artifact_id = ?"