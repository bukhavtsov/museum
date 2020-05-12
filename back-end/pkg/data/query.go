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
const getArtifactElementByIDQuery = `
SELECT child_ae.id, child_ae.artifact_id, child_ae.artifact_element_name, parent_ae.id
FROM artifact_element child_ae
         LEFT JOIN artifact_element parent_ae ON child_ae.artifact_parent_element_id = parent_ae.id
WHERE child_ae.artifact_id = ? ORDER BY child_ae.id asc
`

const getArtifactChildElementQuery = `
SELECT child_ae.id, child_ae.artifact_id, child_ae.artifact_element_name, parent_ae.id
FROM artifact_element child_ae
         LEFT JOIN artifact_element parent_ae ON child_ae.artifact_parent_element_id = parent_ae.id
WHERE child_ae.artifact_id = ? AND child_ae.artifact_parent_element_id = ? ORDER BY child_ae.id ASC
`

const getArtifactObjectGroupByIDQuery = `
SELECT child_og.id,
       child_og.artifact_id,
       child_og.object_group_name,
       child_og.object_group_parent_id
FROM object_group child_og
         LEFT JOIN object_group parent_og on child_og.object_group_parent_id = parent_og.id
WHERE child_og.artifact_id = ? ORDER BY child_og.id ASC
`

const getArtifactChildObjectGroupQuery = `
SELECT child_og.id,
       child_og.artifact_id,
       child_og.object_group_name,
       child_og.object_group_parent_id
FROM object_group child_og
         LEFT JOIN object_group parent_og on child_og.object_group_parent_id = parent_og.id
WHERE child_og.artifact_id = ? AND child_og.object_group_parent_id = ? ORDER BY child_og.id ASC
`

const getArtifactPreservationByIDQuery = `
SELECT child_ap.id,
	child_ap.artifact_id,
	child_ap.preservation,
	child_ap.artifact_preservation_parent_id
FROM artifact_preservation child_ap
         LEFT JOIN artifact_preservation parent_ap ON child_ap.artifact_preservation_parent_id = parent_ap.id
WHERE child_ap.artifact_id = ? ORDER BY child_ap.id ASC
`

const getArtifactChildPreservationQuery = `
SELECT child_ap.id,
	child_ap.artifact_id,
	child_ap.preservation,
	child_ap.artifact_preservation_parent_id
FROM artifact_preservation child_ap
         LEFT JOIN artifact_preservation parent_ap ON child_ap.artifact_preservation_parent_id = parent_ap.id
WHERE child_ap.artifact_id = ? AND child_ap.artifact_preservation_parent_id = ? ORDER BY child_ap.id ASC
`

const getArtifactMaterialsByIDQuery = `
SELECT child_m.id,
	child_m.artifact_id,
	child_m.quantity,
    child_m."%composition",
    child_m.material_type,
	parent_m.id as id_parent
FROM material child_m
    LEFT JOIN material parent_m ON child_m.material_type_parent_id = parent_m.id
WHERE child_m.artifact_id = ? ORDER BY child_m.id ASC
`

const getArtifactChildMaterialsQuery = `
SELECT child_m.id,
	child_m.artifact_id,
	child_m.quantity,
    child_m."%composition",
    child_m.material_type,
	parent_m.id as id_parent
FROM material child_m
    LEFT JOIN material parent_m ON child_m.material_type_parent_id = parent_m.id
WHERE child_m.artifact_id = ? AND child_m.material_type_parent_id = ? ORDER BY child_m.id ASC
`

const insertTransferredByLUTQuery = `
	INSERT INTO transferred_by_lut (id, transferred_by) VALUES (default, ?);`

const insertArtifactMasterQuery = `
INSERT INTO artifact_master (museum_id, excavation_region_id, reg_confidence_id,
                             creator, date_exc, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
`

const insertObjectGroupQuery = `
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (default, ?, ?, ?)
`

const insertMaterialTypeQuery = `
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (default, ?);
`
const insertArtifactElementsQuery = `
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (default, ?, ?, ?);
`

const artifactMeasurementQuery = `
INSERT INTO "artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (default, ?, ?, ?, ?);
`

const artifactPreservationQuery = `
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (default, ?, ?, ?);
`

const artifactElementQuery = `
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (default, ?, ?, ?);
`

const getTransferredByLUTQuery =`
SELECT id FROM transferred_by_lut WHERE transferred_by = ?; 
`