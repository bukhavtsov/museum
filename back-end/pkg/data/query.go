package data

const getBasicArtifactInfo = `
SELECT artifact_master_phas.id,
       artifact_master_phas.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by,
       artifact_master_phas.date_exc,
       artifact_measurement.height,
       artifact_measurement.width,
       artifact_measurement.length,
       artifact_safety.safety
FROM artifact_master_phas
         INNER JOIN transferred_by_lut
                    ON (artifact_master_phas.transferred_by_id = transferred_by_lut.id)
         LEFT JOIN
     (
         SELECT artifact_id, artifact_style_name
         FROM artifact_style
                  INNER JOIN artifact_style_lut on (artifact_style.id = artifact_style_lut.id)
     ) as artifact_style
     on (artifact_master_phas.id = artifact_style.artifact_id)
         INNER JOIN artifact_measurement on artifact_master_phas.id = artifact_measurement.artifact_id
         INNER JOIN artifact_safety on artifact_master_phas.id = artifact_safety.artifact_id;
`
