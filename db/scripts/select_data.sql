SELECT artifact_master_phas.id,
       artifact_master_phas.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by
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
