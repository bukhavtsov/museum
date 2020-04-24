-- select artifact_data
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
         INNER JOIN artifact_measurement on artifact_master.id = artifact_measurement.artifact_id;


-- select artifact_element table with foreign key to same table
SELECT ae1.artifact_id, ae1.artifact_element_name, ae2.artifact_element_name
FROM artifact_element ae1
         LEFT JOIN artifact_element ae2 ON ae1.artifact_parent_element_id = ae2.id;

--select object_group
SELECT og1.artifact_id, ogl1.object_group_name, ogl2.object_group_name as object_group_parent_name
FROM object_group og1
         LEFT JOIN object_group_lut ogl1 on og1.object_group_id = ogl1.id
         LEFT JOIN object_group og2 on og1.object_group_id = og2.id
         LEFT JOIN object_group_lut ogl2 on og2.object_group_id = ogl2.id;
