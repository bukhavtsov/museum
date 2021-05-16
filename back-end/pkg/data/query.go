package data

const getArtifactsWithBasicInfoQuery = `
SELECT artifact_master.id,
       artifact_master.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by,
       artifact_master.date_exc,
       ArtifactMeasurement.height,
       ArtifactMeasurement.width,
       ArtifactMeasurement.length
FROM artifact_master
         FULL JOIN transferred_by_lut
                   ON (artifact_master.transferred_by_id = transferred_by_lut.id)
         FULL JOIN
     (
         SELECT artifact_id, artifact_style_name
         FROM artifact_style
                  FULL JOIN artifact_style_lut on (artifact_style.id = artifact_style_lut.id)
     ) as artifact_style
     on (artifact_master.id = artifact_style.artifact_id)
         INNER JOIN ArtifactMeasurement on artifact_master.id = ArtifactMeasurement.artifact_id
`

const getArtifactsWithBasicInfoByIDQuery = `
SELECT artifact_master.id,
       artifact_master.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by,
       artifact_master.date_exc,
       ArtifactMeasurement.height,
       ArtifactMeasurement.width,
       ArtifactMeasurement.length
FROM artifact_master
         FULL JOIN transferred_by_lut
                    ON (artifact_master.transferred_by_id = transferred_by_lut.id)
         FULL JOIN
     (
         SELECT artifact_id, artifact_style_name
         FROM artifact_style
                  FULL JOIN artifact_style_lut on (artifact_style.id = artifact_style_lut.id)
     ) as artifact_style
     on (artifact_master.id = artifact_style.artifact_id)
         FULL JOIN ArtifactMeasurement on artifact_master.id = ArtifactMeasurement.artifact_id
WHERE artifact_master.id = $1
`

const getArtifactElementByIdQuery = `SELECT id, artifact_id, artifact_element_name, artifact_parent_element_id
FROM artifact_element
WHERE artifact_id = ?
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
SELECT id FROM transferred_by_lut WHERE transferred_by = $1
`

const insertArtifactStyleLUT = `
INSERT INTO artifact_style_lut (id, artifact_style_name) VALUES (DEFAULT, ?)
`
const selectArtifactStyleLUT = `
SELECT id FROM artifact_style_lut WHERE artifact_style_name=?
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
SELECT id FROM artifact_master WHERE creator = ? AND date_exc = ? AND transferred_by_id = ?
`

const insertMeasurement = `
INSERT INTO "ArtifactMeasurement" ("id", "artifact_id", "length", "height", "width")
VALUES (DEFAULT, ?, ?, ?, ?)
`
const selectArtifactMeasurement = `
SELECT id FROM "ArtifactMeasurement" WHERE artifact_id = ? AND length = ? AND height = ? AND width = ?
`

const getTransferredByIdFieldByName = `
SELECT id FROM transferred_by_lut WHERE transferred_by = $1
`

const updateArtifactMaster = `
UPDATE artifact_master SET creator = $1, date_exc = $2, transferred_by_id = $3 WHERE id = $4
`


// TODO: clarify, why we should have separate artifact_id? better to add to artifact_master field with artifact measurements
const updateArtifactMeasurement = `
UPDATE ArtifactMeasurement SET length = ?, height = ?, width = ? WHERE artifact_id = ?
`

const updateArtifactStyle = `
UPDATE artifact_style SET artifact_style_id = ? WHERE artifact_id = ? 
`

const deleteMeasurement = `
DELETE FROM ArtifactMeasurement WHERE artifact_id = ?
`

const deleteArtifactMaster = `
DELETE FROM artifact_master WHERE artifact_id = ?
`

const insertArtifactElement = `INSERT INTO artifact_element (artifact_id, elements) VALUES (?, ?) RETURNING id`
const selectArtifactElement = `SELECT elements FROM artifact_element WHERE artifact_id = ?`
