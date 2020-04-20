-- select artifact_data
SELECT artifact_master.id,
       artifact_master.creator,
       artifact_style.artifact_style_name,
       transferred_by_lut.transferred_by,
       artifact_master.date_exc,
       artifact_measurement.height,
       artifact_measurement.width,
       artifact_measurement.length,
       artifact_safety.safety
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
         INNER JOIN artifact_safety on artifact_master.id = artifact_safety.artifact_id;


-- select artifact_element table with foreign key to same table
SELECT ae1.artifact_id, ae1.artifact_element_name, ae1.artifact_parent_element_id
FROM artifact_element ae1
         LEFT JOIN artifact_element ae2 ON ae1.artifact_parent_element_id = ae2.id;

-- select artifact_
SELECT m1."%composition", m1.quantity, material_type_lut.material_type, material_type_lut.material_type
FROM material m1
         LEFT JOIN material m2 ON m1.artifact_id = m2.id
         LEFT JOIN material_type_lut ON m1.material_type_id = material_type_lut.id
