# Database side of museum app

## Schema

### artifact_element

|           Column           |          Type          | Collation | Nullable |                   Default                    |
|----------------------------|------------------------|-----------|----------|----------------------------------------------|
| id                         | integer                |           | not null | nextval('artifact_element_id_seq'::regclass) |
| artifact_id                | integer                |           | not null |               |
| artifact_element_name      | character varying(100) |           | not null |               |
| artifact_parent_element_id | integer                |           |          |               | 

### artifact_element
            
|        Column        |          Type          | Collation | Nullable |                   Default                  | 
|----------------------|------------------------|-----------|----------|--------------------------------------------|
| id                   | integer                |           | not null | nextval('artifact_master_id_seq'::regclass)|
| artifact_id          | integer                |           |          |             |
| museum_id            | integer                |           | not null |             |             |
| excavation_region_id | integer                |           |          |             |
| reg_confidence_id    | integer                |           |          |             | 
| date_exc             | date                   |           |          |             | 
| creator              | character varying(100) |           |          |             | 
| hist_culture_id      | integer                |           |          |             | 
| desc                 | text                   |           |          |             | 
| translation          | text                   |           |          |             | 
| min_age              | integer                |           |          |             | 
| max_age              | integer                |           |          |             | 
| artifact_info_photo  | text                   |           |          |             | 
| photo                | character varying(100) |           |          |             | 
| transferred_by_id    | integer                |           |          |             | 
#### Indexes:
    "artifact_master_pk" PRIMARY KEY, btree (id)
    "artifact_master_artifact_id_uindex" UNIQUE, btree (artifact_id)
#### Foreign-key constraints:
    "artifact_master_excavation_region_id_fk" FOREIGN KEY (excavation_region_id) REFERENCES excavation_region(id)
    "artifact_master_hist_culture_id_fk" FOREIGN KEY (hist_culture_id) REFERENCES hist_culture(id)
    "artifact_master_museum_id_fk" FOREIGN KEY (museum_id) REFERENCES museum(id)
    "artifact_master_reg_confidence_level_id_fk" FOREIGN KEY (reg_confidence_id) REFERENCES reg_confidence_level(id)
    "artifact_master_transferred_by_lut_id_fk" FOREIGN KEY (transferred_by_id) REFERENCES transferred_by_lut(id)
#### Referenced by:
    TABLE "artifact_element" CONSTRAINT "artifact_element_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "artifact_measurement" CONSTRAINT "artifact_measurement_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "artifact_preservation" CONSTRAINT "artifact_preservation_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "artifact_publication" CONSTRAINT "artifact_publication_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "artifact_related_people" CONSTRAINT "artifact_related_people_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "artifact_style" CONSTRAINT "artifact_style_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "collection" CONSTRAINT "collection_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "material" CONSTRAINT "material_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "object_group" CONSTRAINT "object_group_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "pb_isotope" CONSTRAINT "pb_isotope_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "provenience_intersite" CONSTRAINT "provenience_intersite_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "reference" CONSTRAINT "reference_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "restoration" CONSTRAINT "restoration_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "site_name" CONSTRAINT "site_name_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    TABLE "site_type" CONSTRAINT "site_type_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)


##### Index "public.artifact_master_artifact_id_uindex"

|   Column    |  Type   | Definition  |  
|-------------|---------|-------------|
| artifact_id | integer | artifact_id |
##### unique, btree, for table "public.artifact_master"

               Sequence "public.artifact_master_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_master.id

##### Index "public.artifact_master_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_master"

### artifact_measurement"
   Column    |  Type   | Collation | Nullable |                     Default                      
-------------|---------|-----------|----------|--------------------------------------------------
 id          | integer |           | not null | nextval('artifact_measurement_id_seq'::regclass)
 artifact_id | integer |           | not null | 
 length      | integer |           |          | 
 height      | integer |           | not null | 
 width       | integer |           |          | 
#### Indexes:
    "artifact_measurement_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "artifact_measurement_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

            Sequence "public.artifact_measurement_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_measurement.id

##### Index "public.artifact_measurement_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_measurement"

### artifact_preservation
|             Column              |  Type   | Collation | Nullable |                      Default                     |                      
---------------------------------|---------|-----------|----------|---------------------------------------------------
 id                              | integer |           | not null | nextval('artifact_preservation_id_seq'::regclass)
 artifact_id                     | integer |           | not null | 
 preservation                    | text    |           | not null | 
 artifact_preservation_parent_id | integer |           |          | 
#### Indexes:
    "artifact_preservation_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "artifact_preservation_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "artifact_preservation_artifact_preservation_id_fk" FOREIGN KEY (artifact_preservation_parent_id) REFERENCES artifact_preservation(id)
#### Referenced by:
    TABLE "artifact_preservation" CONSTRAINT "artifact_preservation_artifact_preservation_id_fk" FOREIGN KEY (artifact_preservation_parent_id) REFERENCES artifact_preservation(id)

            Sequence "public.artifact_preservation_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_preservation.id

##### Index "public.artifact_preservation_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_preservation"

### artifact_publication
   Column    |          Type          | Collation | Nullable |                     Default                      
-------------|------------------------|-----------|----------|--------------------------------------------------
 id          | integer                |           | not null | nextval('artifact_publication_id_seq'::regclass)
 artifact_id | integer                |           | not null | 
 author_name | character varying(100) |           | not null | 
 date        | date                   |           |          | 
#### Indexes:
    "artifact_publication_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "artifact_publication_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

            Sequence "public.artifact_publication_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_publication.id

##### Index "public.artifact_publication_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_publication"

### artifact_related_people
|   Column    |          Type          | Collation | Nullable |                       Default                      |                       
-------------|------------------------|-----------|----------|-----------------------------------------------------
 id          | integer                |           | not null | nextval('artifact_related_people_id_seq'::regclass)
 artifact_id | integer                |           | not null | 
 person_name | character varying(100) |           | not null | 
#### Indexes:
    "artifact_related_people_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "artifact_related_people_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

           Sequence "public.artifact_related_people_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_related_people.id

##### Index "public.artifact_related_people_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_related_people"

### artifact_style
|      Column       |  Type   | Collation | Nullable | Default | 
|-------------------|---------|-----------|----------|---------
| id                | integer |           | not null | 
| artifact_id       | integer |           | not null | 
| artifact_style_id | integer |           | not null | 
#### Indexes:
    "artifact_style_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "artifact_style_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "artifact_style_artifact_style_lut_id_fk" FOREIGN KEY (artifact_style_id) REFERENCES artifact_style_lut(id)

### artifact_style_lut
|       Column        |          Type          | Collation | Nullable |                    Default                   |                     
---------------------|------------------------|-----------|----------|------------------------------------------------
 id                  | integer                |           | not null | nextval('artifact_style_lut_id_seq'::regclass)
 artifact_style_name | character varying(100) |           | not null | 
#### Indexes:
    "artifact_style_lut_pk" PRIMARY KEY, btree (id)
    "artifact_style_lut_artifact_style_name_uindex" UNIQUE, btree (artifact_style_name)
#### Referenced by:
    TABLE "artifact_style" CONSTRAINT "artifact_style_artifact_style_lut_id_fk" FOREIGN KEY (artifact_style_id) REFERENCES artifact_style_lut(id)

##### Index "public.artifact_style_lut_artifact_style_name_uindex"

|       Column        |          Type          |     Definition    |      
---------------------|------------------------|---------------------
 artifact_style_name | character varying(100) | artifact_style_name
##### unique, btree, for table "public.artifact_style_lut"

             Sequence "public.artifact_style_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.artifact_style_lut.id

##### Index "public.artifact_style_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_style_lut"

##### Index "public.artifact_style_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.artifact_style"

### collection
|     Column      |         Type          | Collation | Nullable |                Default               |                 
-----------------|-----------------------|-----------|----------|----------------------------------------
 id              | integer               |           | not null | nextval('collection_id_seq'::regclass)
 artifact_id     | integer               |           | not null | 
 collection_name | character varying(50) |           |          | 
#### Foreign-key constraints:
    "collection_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

                 Sequence "public.collection_id_seq"
|  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache | 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.collection.id

##### Index "public.contacts_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.museum_contacts"

### country
|    Column    |          Type          | Collation | Nullable |               Default              |               
--------------|------------------------|-----------|----------|-------------------------------------
 id           | integer                |           | not null | nextval('country_id_seq'::regclass)
 country_name | character varying(100) |           | not null | 
#### Indexes:
    "country_lut_pk" PRIMARY KEY, btree (id)
    "country_country_name_uindex" UNIQUE, btree (country_name)
#### Referenced by:
    TABLE "region" CONSTRAINT "region_country_id_fk" FOREIGN KEY (country_id) REFERENCES country(id)

##### Index "public.country_country_name_uindex"

|    Column    |          Type          |  Definition |  
--------------|------------------------|--------------
 country_name | character varying(100) | country_name
##### unique, btree, for table "public.country"

                   Sequence "public.country_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.country.id

##### Index "public.country_lut_pk"
 
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.country"

### excavation_region
|    Column    |  Type   | Collation | Nullable |                    Default                  |                    
--------------|---------|-----------|----------|-----------------------------------------------
 id           | integer |           | not null | nextval('excavation_region_id_seq'::regclass)
 location_id  | integer |           | not null | 
 x_coordinate | integer |           |          | 
 y_coordinate | integer |           |          | 
#### Indexes:
    "excavation_region_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "excavation_region_location_id_fk" FOREIGN KEY (location_id) REFERENCES location(id)
#### Referenced by:
    TABLE "artifact_master" CONSTRAINT "artifact_master_excavation_region_id_fk" FOREIGN KEY (excavation_region_id) REFERENCES excavation_region(id)

              Sequence "public.excavation_region_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache |
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.excavation_region.id

##### Index "public.excavation_region_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.excavation_region"

### hist_culture
|    Column    |         Type          | Collation | Nullable |                 Default                |                  
--------------|-----------------------|-----------|----------|------------------------------------------
 id           | integer               |           | not null | nextval('hist_culture_id_seq'::regclass)
 hist_culture | character varying(50) |           | not null | 
#### Indexes:
    "hist_culture_pk" PRIMARY KEY, btree (id)
    "hist_culture_hist_culture_uindex" UNIQUE, btree (hist_culture)
#### Referenced by:
    TABLE "artifact_master" CONSTRAINT "artifact_master_hist_culture_id_fk" FOREIGN KEY (hist_culture_id) REFERENCES hist_culture(id)

##### Index "public.hist_culture_hist_culture_uindex"
   
|    Column    |         Type          |  Definition |  
--------------|-----------------------|--------------
 hist_culture | character varying(50) | hist_culture
##### unique, btree, for table "public.hist_culture"

                Sequence "public.hist_culture_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.hist_culture.id

##### Index "public.hist_culture_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.hist_culture"

### location
|    Column     |         Type          | Collation | Nullable |               Default               |                
---------------|-----------------------|-----------|----------|--------------------------------------
 id            | integer               |           | not null | nextval('location_id_seq'::regclass)
 location_name | character varying(50) |           | not null | 
 region_id     | integer               |           | not null | 
#### Indexes:
    "location_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "location_region_id_fk" FOREIGN KEY (region_id) REFERENCES region(id)
#### Referenced by:
    TABLE "excavation_region" CONSTRAINT "excavation_region_location_id_fk" FOREIGN KEY (location_id) REFERENCES location(id)
    TABLE "museum" CONSTRAINT "museum_location_id_fk" FOREIGN KEY (location_id) REFERENCES location(id)

                  Sequence "public.location_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.location.id

##### Index "public.location_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.location"

### material
|         Column          |  Type   | Collation | Nullable |               Default               |                
-------------------------|---------|-----------|----------|--------------------------------------
 id                      | integer |           | not null | nextval('material_id_seq'::regclass)
 artifact_id             | integer |           | not null | 
 material_type_id        | integer |           |          | 
 quantity                | integer |           |          | 
 %composition            | integer |           |          | 
 confidence_level_id     | integer |           |          | 
 material_type_parent_id | integer |           |          | 
#### Indexes:
    "material_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "material_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "material_material_confidence_level_id_fk" FOREIGN KEY (confidence_level_id) REFERENCES material_confidence_level(id)
    "material_material_id_fk" FOREIGN KEY (material_type_parent_id) REFERENCES material(id)
    "material_material_type_lut_id_fk" FOREIGN KEY (material_type_id) REFERENCES material_type_lut(id)
#### Referenced by:
    TABLE "material" CONSTRAINT "material_material_id_fk" FOREIGN KEY (material_type_parent_id) REFERENCES material(id)

### material_confidence_level
|          Column           |         Type          | Collation | Nullable |                        Default                      |
---------------------------|-----------------------|-----------|----------|-------------------------------------------------------
 id                        | integer               |           | not null | nextval('material_confidence_level_id_seq'::regclass)
 material_confidence_level | character varying(50) |           | not null | 
#### Indexes:
    "material_confidence_level_pk" PRIMARY KEY, btree (id)
    "material_confidence_level_material_confidence_level_uindex" UNIQUE, btree (material_confidence_level)
#### Referenced by:
    TABLE "material" CONSTRAINT "material_material_confidence_level_id_fk" FOREIGN KEY (confidence_level_id) REFERENCES material_confidence_level(id)

          Sequence "public.material_confidence_level_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.material_confidence_level.id

##### Index "public.material_confidence_level_material_confidence_level_uindex"
   
|          Column           |         Type          |        Definition       |         
---------------------------|-----------------------|---------------------------
 material_confidence_level | character varying(50) | material_confidence_level
##### unique, btree, for table "public.material_confidence_level"

##### Index "public.material_confidence_level_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.material_confidence_level"

                  Sequence "public.material_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.material.id

##### Index "public.material_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.material"

### material_type_lut
|    Column     |         Type          | Collation | Nullable |                    Default                  |                    
---------------|-----------------------|-----------|----------|-----------------------------------------------
 id            | integer               |           | not null | nextval('material_type_lut_id_seq'::regclass)
 material_type | character varying(50) |           |          | 
#### Indexes:
    "material_type_lut_pk" PRIMARY KEY, btree (id)
#### Referenced by:
    TABLE "material" CONSTRAINT "material_material_type_lut_id_fk" FOREIGN KEY (material_type_id) REFERENCES material_type_lut(id)

                        Sequence "public.material_type_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.material_type_lut.id

##### Index "public.material_type_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.material_type_lut"

### museum
|   Column    |         Type          | Collation | Nullable |              Default             |               
-------------|-----------------------|-----------|----------|------------------------------------
 id          | integer               |           | not null | nextval('museum_id_seq'::regclass)
 museum_name | character varying(50) |           | not null | 
 location_id | integer               |           | not null | 
#### Indexes:
    "museum_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "museum_location_id_fk" FOREIGN KEY (location_id) REFERENCES location(id)
#### Referenced by:
    TABLE "artifact_master" CONSTRAINT "artifact_master_museum_id_fk" FOREIGN KEY (museum_id) REFERENCES museum(id)
    TABLE "museum_contacts" CONSTRAINT "museum_contacts_museum_id_fk" FOREIGN KEY (museum_id) REFERENCES museum(id)

### museum_contacts
|    Column    |         Type          | Collation | Nullable |                   Default                 |  
--------------|-----------------------|-----------|----------|---------------------------------------------
 id           | integer               |           | not null | nextval('museum_contacts_id_seq'::regclass)
 phone_number | character varying(50) |           |          | 
 site_addr    | character varying(50) |           |          | 
 email        | character varying(50) |           |          | 
 museum_id    | integer               |           | not null | 
#### Indexes:
    "contacts_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "museum_contacts_museum_id_fk" FOREIGN KEY (museum_id) REFERENCES museum(id)

               Sequence "public.museum_contacts_id_seq"
|  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache |
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.museum_contacts.id

                   Sequence "public.museum_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.museum.id

##### Index "public.museum_pk"
   
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.museum"

### object_group
|         Column         |  Type   | Collation | Nullable |                 Default                |                  
------------------------|---------|-----------|----------|------------------------------------------
 id                     | integer |           | not null | nextval('object_group_id_seq'::regclass)
 object_group_id        | integer |           | not null | 
 artifact_id            | integer |           | not null | 
 object_group_parent_id | integer |           |          | 
#### Indexes:
    "object_group_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "object_group_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "object_group_object_group_id_fk" FOREIGN KEY (object_group_parent_id) REFERENCES object_group(id)
    "object_group_object_group_lut_id_fk" FOREIGN KEY (object_group_id) REFERENCES object_group_lut(id)
#### Referenced by:
    TABLE "object_group" CONSTRAINT "object_group_object_group_id_fk" FOREIGN KEY (object_group_parent_id) REFERENCES object_group(id)

                Sequence "public.object_group_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.object_group.id

### object_group_lut
|      Column       |          Type          | Collation | Nullable |                   Default                  |  
-------------------|------------------------|-----------|----------|----------------------------------------------
 id                | integer                |           | not null | nextval('object_group_lut_id_seq'::regclass)
 object_group_name | character varying(100) |           | not null | 
#### Indexes:
    "object_group_lut_pk" PRIMARY KEY, btree (id)
    "object_group_lut_object_group_name_uindex" UNIQUE, btree (object_group_name)
#### Referenced by:
    TABLE "object_group" CONSTRAINT "object_group_object_group_lut_id_fk" FOREIGN KEY (object_group_id) REFERENCES object_group_lut(id)

              Sequence "public.object_group_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.object_group_lut.id

##### Index "public.object_group_lut_object_group_name_uindex"
|      Column       |          Type          |    Definition   | 
-------------------|------------------------|-------------------
 object_group_name | character varying(100) | object_group_name
##### unique, btree, for table "public.object_group_lut"

##### Index "public.object_group_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.object_group_lut"

##### Index "public.object_group_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.object_group"

### pb_isotope
|   Column    |         Type          | Collation | Nullable |                Default               |                 
-------------|-----------------------|-----------|----------|----------------------------------------
 id          | integer               |           | not null | nextval('pb_isotope_id_seq'::regclass)
 artifact_id | integer               |           | not null | 
 isotope     | character varying(50) |           |          | 
 value       | character varying(50) |           |          | 
 date        | date                  |           |          | 
#### Indexes:
    "pb_isotope_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "pb_isotope_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

                 Sequence "public.pb_isotope_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.pb_isotope.id

##### Index "public.pb_isotope_pk"
 
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.pb_isotope"

### prov_category_lut
|    Column     |         Type          | Collation | Nullable | Default | 
---------------|-----------------------|-----------|----------|---------
 id            | integer               |           | not null | 
 prov_category | character varying(50) |           | not null | 
#### Indexes:
    "prov_category_lut_pk" PRIMARY KEY, btree (id)
    "prov_category_lut_prov_category_uindex" UNIQUE, btree (prov_category)
#### Referenced by:
    TABLE "provenience_intersite" CONSTRAINT "provenience_intersite_prov_category_lut_id_fk" FOREIGN KEY (p_category_id) REFERENCES prov_category_lut(id)

##### Index "public.prov_category_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.prov_category_lut"

##### Index "public.prov_category_lut_prov_category_uindex"
 
|    Column     |         Type          |  Definition  |   
---------------|-----------------------|---------------
 prov_category | character varying(50) | prov_category
##### unique, btree, for table "public.prov_category_lut"

### provenience_intersite
|    Column     |  Type   | Collation | Nullable |                      Default                     | 
---------------|---------|-----------|----------|---------------------------------------------------
 id            | integer |           | not null | nextval('provenience_intersite_id_seq'::regclass)
 artifact_id   | integer |           | not null | 
 p_category_id | integer |           |          | 
 p_info        | text    |           |          | 
#### Indexes:
    "provenience_intersite_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "provenience_intersite_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "provenience_intersite_prov_category_lut_id_fk" FOREIGN KEY (p_category_id) REFERENCES prov_category_lut(id)

            Sequence "public.provenience_intersite_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.provenience_intersite.id

##### Index "public.provenience_intersite_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.provenience_intersite"

### ref_categ_lut
| Column  |         Type          | Collation | Nullable |                  Default                 | 
---------|-----------------------|-----------|----------|-------------------------------------------
 id      | integer               |           | not null | nextval('ref_categ_lut_id_seq'::regclass)
 r_categ | character varying(50) |           | not null | 
#### Indexes:
    "ref_categ_lut_pk" PRIMARY KEY, btree (id)
    "ref_categ_lut_r_categ_uindex" UNIQUE, btree (r_categ)
#### Referenced by:
    TABLE "reference" CONSTRAINT "reference_ref_categ_lut_id_fk" FOREIGN KEY (r_category_id) REFERENCES ref_categ_lut(id)

                Sequence "public.ref_categ_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.ref_categ_lut.id

##### Index "public.ref_categ_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.ref_categ_lut"

##### Index "public.ref_categ_lut_r_categ_uindex"
 
 Column  |         Type          | Definition 
---------|-----------------------|------------
 r_categ | character varying(50) | r_categ
##### unique, btree, for table "public.ref_categ_lut"

### reference
|     Column     |  Type   | Collation | Nullable |                Default               |
----------------|---------|-----------|----------|---------------------------------------
 id             | integer |           | not null | nextval('reference_id_seq'::regclass)
 artifact_id    | integer |           | not null | 
 r_category_id  | integer |           |          | 
 reference_info | text    |           |          | 
#### Indexes:
    "reference_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "reference_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "reference_ref_categ_lut_id_fk" FOREIGN KEY (r_category_id) REFERENCES ref_categ_lut(id)

                  Sequence "public.reference_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.reference.id

##### Index "public.reference_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.reference"

### reg_confidence_level
|        Column        |         Type          | Collation | Nullable |                     Default                    | 
----------------------|-----------------------|-----------|----------|--------------------------------------------------
 id                   | integer               |           | not null | nextval('reg_confidence_level_id_seq'::regclass)
 reg_confidence_level | character varying(50) |           | not null | 
#### Indexes:
    "reg_confidence_level_pk" PRIMARY KEY, btree (id)
#### Referenced by:
    TABLE "artifact_master" CONSTRAINT "artifact_master_reg_confidence_level_id_fk" FOREIGN KEY (reg_confidence_id) REFERENCES reg_confidence_level(id)

            Sequence "public.reg_confidence_level_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.reg_confidence_level.id

##### Index "public.reg_confidence_level_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.reg_confidence_level"

### region
   Column    |          Type          | Collation | Nullable |              Default               
-------------|------------------------|-----------|----------|------------------------------------
 id          | integer                |           | not null | nextval('region_id_seq'::regclass)
 region_name | character varying(100) |           | not null | 
 country_id  | integer                |           | not null | 
#### Indexes:
    "region_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "region_country_id_fk" FOREIGN KEY (country_id) REFERENCES country(id)
#### Referenced by:
    TABLE "location" CONSTRAINT "location_region_id_fk" FOREIGN KEY (region_id) REFERENCES region(id)

                   Sequence "public.region_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.region.id

##### Index "public.region_pk"
   
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.region"

### restoration
|   Column    |          Type          | Collation | Nullable |                 Default                | 
-------------|------------------------|-----------|----------|-----------------------------------------
 id          | integer                |           | not null | nextval('restoration_id_seq'::regclass)
 artifact_id | integer                |           | not null | 
 date        | date                   |           |          | 
 updates     | character varying(100) |           |          | 
 author      | character varying(100) |           |          | 
#### Indexes:
    "restoration_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "restoration_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)

                 Sequence "public.restoration_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.restoration.id

##### Index "public.restoration_pk"
 
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.restoration"

### site_name
|    Column    |         Type          | Collation | Nullable |                Default               |
--------------|-----------------------|-----------|----------|---------------------------------------
 id           | integer               |           | not null | nextval('site_name_id_seq'::regclass)
 artifact_id  | integer               |           | not null | 
 site_name_id | integer               |           |          | 
 location     | character varying(50) |           |          | 
 country      | character varying(50) |           |          | 
 comments     | text                  |           |          | 
 start_date   | date                  |           |          | 
 finish_date  | integer               |           |          | 
#### Indexes:
    "site_name_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "site_name_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "site_name_site_name_lut_id_fk" FOREIGN KEY (site_name_id) REFERENCES site_name_lut(id)

                  Sequence "public.site_name_id_seq"
|  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.site_name.id

### site_name_lut
|  Column   |         Type          | Collation | Nullable |                  Default                 | 
-----------|-----------------------|-----------|----------|-------------------------------------------
 id        | integer               |           | not null | nextval('site_name_lut_id_seq'::regclass)
 site_name | character varying(50) |           | not null | 
#### Indexes:
    "site_name_lut_pk" PRIMARY KEY, btree (id)
    "site_name_lut_site_name_uindex" UNIQUE, btree (site_name)
#### Referenced by:
    TABLE "site_name" CONSTRAINT "site_name_site_name_lut_id_fk" FOREIGN KEY (site_name_id) REFERENCES site_name_lut(id)

                Sequence "public.site_name_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.site_name_lut.id

##### Index "public.site_name_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.site_name_lut"

##### Index "public.site_name_lut_site_name_uindex"
 
  Column   |         Type          | Definition 
-----------|-----------------------|------------
 site_name | character varying(50) | site_name
##### unique, btree, for table "public.site_name_lut"

##### Index "public.site_name_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.site_name"

### site_type"
|    Column    |  Type   | Collation | Nullable |                Default               |
--------------|---------|-----------|----------|---------------------------------------
 id           | integer |           | not null | nextval('site_type_id_seq'::regclass)
 artifact_id  | integer |           | not null | 
 site_type_id | integer |           |          | 
#### Indexes:
    "site_type_pk" PRIMARY KEY, btree (id)
#### Foreign-key constraints:
    "site_type_artifact_master_id_fk" FOREIGN KEY (artifact_id) REFERENCES artifact_master(id)
    "site_type_site_type_lut_id_fk" FOREIGN KEY (site_type_id) REFERENCES site_type_lut(id)

                  Sequence "public.site_type_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.site_type.id

### site_type_lut
|  Column   |         Type          | Collation | Nullable |                  Default                |  
-----------|-----------------------|-----------|----------|-------------------------------------------
 id        | integer               |           | not null | nextval('site_type_lut_id_seq'::regclass)
 site_type | character varying(50) |           | not null | 
#### Indexes:
    "site_type_lut_pk" PRIMARY KEY, btree (id)
    "site_type_lut_site_type_uindex" UNIQUE, btree (site_type)
#### Referenced by:
    TABLE "site_type" CONSTRAINT "site_type_site_type_lut_id_fk" FOREIGN KEY (site_type_id) REFERENCES site_type_lut(id)

                Sequence "public.site_type_lut_id_seq"
  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache 
---------|-------|---------|------------|-----------|---------|-------
 integer |     1 |       1 | 2147483647 |         1 | no      |     1
##### Owned by: public.site_type_lut.id

##### Index "public.site_type_lut_pk"

 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.site_type_lut"

##### Index "public.site_type_lut_site_type_uindex"
 
  Column   |         Type          | Definition 
-----------|-----------------------|------------
 site_type | character varying(50) | site_type
##### unique, btree, for table "public.site_type_lut"

##### Index "public.site_type_pk"
  
 Column |  Type   | Definition 
--------|---------|------------
 id     | integer | id
##### primary key, btree, for table "public.site_type"

### transferred_by_lut
|     Column     |          Type          | Collation | Nullable |                    Default                     
|----------------|------------------------|-----------|----------|------------------------------------------------
| id             | integer                |           | not null | nextval('transferred_by_lut_id_seq'::regclass)
| transferred_by | character varying(200) |           | not null | 
#### Indexes:
    "transferred_by_lut_pk" PRIMARY KEY, btree (id)
    "transferred_by_lut_transferred_by_uindex" UNIQUE, btree (transferred_by)
#### Referenced by:
    TABLE "artifact_master" CONSTRAINT "artifact_master_transferred_by_lut_id_fk" FOREIGN KEY (transferred_by_id) REFERENCES transferred_by_lut(id)

             Sequence "public.transferred_by_lut_id_seq"
|  Type   | Start | Minimum |  Maximum   | Increment | Cycles? | Cache | 
|---------|-------|---------|------------|-----------|---------|-------|
| integer |     1 |       1 | 2147483647 |         1 | no      |     1 |
##### Owned by: public.transferred_by_lut.id

##### Index "public.transferred_by_lut_pk"

| Column |  Type   | Definition |
|--------|---------|------------|
| id     | integer | id         |
##### primary key, btree, for table "public.transferred_by_lut"

##### Index "public.transferred_by_lut_transferred_by_uindex"
 
|     Column     |          Type          |   Definition |   
----------------|------------------------|----------------
 transferred_by | character varying(200) | transferred_by
##### unique, btree, for table "public.transferred_by_lut"
