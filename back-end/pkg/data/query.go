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
const getArtifactElementByIdQuery = `
SELECT child_ae.artifact_id, child_ae.artifact_element_name, parent_ae.artifact_element_name
FROM artifact_element child_ae
         LEFT JOIN artifact_element parent_ae ON child_ae.artifact_parent_element_id = parent_ae.id
WHERE child_ae.artifact_id = ?
`

const getArtifactObjectGroupByIdQuery = `
SELECT child_og.artifact_id,
       child_og_lut.object_group_name,
       parent_og_lut.object_group_name as object_group_parent_name
FROM object_group child_og
         LEFT JOIN object_group parent_og on child_og.object_group_parent_id = parent_og.id
         LEFT JOIN object_group_lut child_og_lut on child_og.object_group_id = child_og_lut.id
         LEFT JOIN object_group_lut parent_og_lut on parent_og.object_group_id = parent_og_lut.id
WHERE child_og.artifact_id = ?
`

const getArtifactPreservationByIdQuery = `
SELECT child_ap.artifact_id, child_ap.preservation, parent_ap.preservation
FROM artifact_preservation child_ap
         LEFT JOIN artifact_preservation parent_ap ON child_ap.artifact_preservation_parent_id = parent_ap.id
WHERE child_ap.artifact_id = ?
`

const insertTransferredBy = `
INSERT INTO transferred_by_lut (transferred_by) VALUES (?)
`

const selectTransferredBy = `
SELECT id FROM transferred_by_lut WHERE transferred_by = ?
`

const insertArtifactStyleLUT = `
INSERT INTO artifact_style_lut (artifact_style_name) VALUES (?)
`
const selectArtifactStyleLUT = `
SELECT id FROM artifact_style_lut WHERE artifact_style_name = ? 
`

const insertArtifactStyle = `
INSERT INTO artifact_style (artifact_id, artifact_style_id) VALUES (?,?)
`

const selectArtifactStyle = `
SELECT id FROM artifact_style WHERE artifact_id = ? AND artifact_style_id = ?
`

const insertArtifactMaster = `
INSERT INTO artifact_master (creator, date_exc, transferred_by_id) VALUES (?, ?, ?)
`

const selectArtifactMaster = `
SELECT id FROM artifact_master where creator = ? AND date_exc = ? AND transferred_by_id = ?
`
