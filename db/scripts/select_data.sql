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
        INNER JOIN artifact_style_lut ON (artifact_style.id = artifact_style_lut.id)
     ) AS artifact_style
    ON (artifact_master.id = artifact_style.artifact_id)
    INNER JOIN artifact_measurement ON artifact_master.id = artifact_measurement.artifact_id;


-- select artifact_element table with foreign key to same table
SELECT child_ae.id, child_ae.artifact_id, child_ae.artifact_element_name, parent_ae.id
FROM artifact_element child_ae
    LEFT JOIN artifact_element parent_ae ON child_ae.artifact_parent_element_id = parent_ae.id;

--select object_group table with foreign key to same table
SELECT child_og.artifact_id,
    child_og_lut.object_group_name,
    parent_og_lut.object_group_name AS object_group_parent_name
FROM object_group child_og
    LEFT JOIN object_group parent_og ON child_og.object_group_parent_id = parent_og.id
    LEFT JOIN object_group_lut child_og_lut ON child_og.object_group_id = child_og_lut.id
    LEFT JOIN object_group_lut parent_og_lut ON parent_og.object_group_id = parent_og_lut.id;

--select artifact_preservation table with foreign key to same table
SELECT child_ap.artifact_id,
    child_ap.preservation,
    parent_ap.preservation
FROM artifact_preservation child_ap
    LEFT JOIN artifact_preservation parent_ap ON child_ap.artifact_preservation_parent_id = parent_ap.id;


--select material table with foreigh key to same table 
SELECT child_m.id,
    child_m.artifact_id,
    child_m.quantity,
    child_m."%composition",
    child_m_lut.material_type,
    parent_m.id as id_parent,
    parent_m.artifact_id as artifact_id_parent,
    parent_m_lut.material_type AS material_type_parent,
    parent_m.quantity AS quantity_parent,
    parent_m."%composition" AS "%composition_parent"
FROM material child_m
    LEFT JOIN material parent_m ON child_m.material_type_parent_id = parent_m.id
    LEFT JOIN material_type_lut child_m_lut ON child_m.material_type_id = child_m_lut.id
    LEFT JOIN material_type_lut parent_m_lut ON parent_m.material_type_id = parent_m_lut.id
