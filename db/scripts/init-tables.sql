-- country table
create table country
(
    id           serial       not null
        constraint country_lut_pk
            primary key,
    country_name varchar(100) not null
);

create unique index country_country_name_uindex
    on country (country_name);


-- region table
create table region
(
    id          serial       not null
        constraint region_pk
            primary key,
    region_name varchar(100) not null,
    country_id  integer      not null
        constraint region_country_id_fk
            references country
);

-- location table
create table location
(
    id            serial      not null
        constraint location_pk
            primary key,
    location_name varchar(50) not null,
    region_id     integer     not null
        constraint location_region_id_fk
            references region
);

-- museum table
create table museum
(
    id          serial      not null
        constraint museum_pk
            primary key,
    museum_name varchar(50) not null,
    location_id integer     not null
        constraint museum_location_id_fk
            references location
);

-- museum_contacts table
create table museum_contacts
(
    id           serial not null
        constraint contacts_pk
            primary key,
    phone_number varchar(50),
    site_addr    varchar(50),
    email        varchar(50),
    museum_id    integer
        constraint museum_contacts_museum_id_fk
            references museum
);

-- excavation_region table
create table excavation_region
(
    id           serial  not null
        constraint excavation_region_pk
            primary key,
    location_id  integer not null
        constraint excavation_region_location_id_fk
            references location,
    x_coordinate integer,
    y_coordinate integer
);

-- reg_confidence_level table
create table reg_confidence_level
(
    id                   serial      not null
        constraint reg_confidence_level_pk
            primary key,
    reg_confidence_level varchar(50) not null
);

-- object_group_lut table
create table object_group_lut
(
    id                serial       not null
        constraint object_group_lut_pk
            primary key,
    object_group_name varchar(100) not null
);

create unique index object_group_lut_object_group_name_uindex
    on object_group_lut (object_group_name);

-- hist culture table
create table hist_culture
(
    id           serial      not null
        constraint hist_culture_pk
            primary key,
    hist_culture varchar(50) not null
);

create unique index hist_culture_hist_culture_uindex
    on hist_culture (hist_culture);

-- transferred_by_lut table
create table transferred_by_lut
(
    id             serial       not null
        constraint transferred_by_lut_pk
            primary key,
    transferred_by varchar(200) not null
);

create unique index transferred_by_lut_transferred_by_uindex
    on transferred_by_lut (transferred_by);

-- artifact_master table
create table artifact_master
(
    id                   serial  not null
        constraint artifact_master_pk
            primary key,
    artifact_id          integer,
    museum_id            integer
        constraint artifact_master_museum_id_fk
            references museum,
    excavation_region_id integer
        constraint artifact_master_excavation_region_id_fk
            references excavation_region,
    reg_confidence_id    integer
        constraint artifact_master_reg_confidence_level_id_fk
            references reg_confidence_level,
    date_exc             date,
    creator              varchar(100),
    hist_culture_id      integer
        constraint artifact_master_hist_culture_id_fk
            references hist_culture,
    "desc"               text,
    translation          text,
    min_age              integer,
    max_age              integer,
    artifact_info_photo  text,
    photo                varchar(100),
    transferred_by_id    integer
        constraint artifact_master_transferred_by_lut_id_fk
            references transferred_by_lut
);

create unique index artifact_master_artifact_id_uindex
    on artifact_master (artifact_id);

-- object_group table
create table object_group
(
    id                     serial  not null
        constraint object_group_pk
            primary key,
    object_group_id        integer not null
        constraint object_group_object_group_lut_id_fk
            references object_group_lut,
    artifact_id            integer not null
        constraint object_group_artifact_master_id_fk
            references artifact_master,
    object_group_parent_id integer
        constraint object_group_object_group_id_fk
            references object_group
);

-- material_type_lut table
create table material_type_lut
(
    id            serial not null
        constraint material_type_lut_pk
            primary key,
    material_type varchar(50)
);

-- material_confidence_level table
create table material_confidence_level
(
    id                        serial      not null
        constraint material_confidence_level_pk
            primary key,
    material_confidence_level varchar(50) not null
);

create unique index material_confidence_level_material_confidence_level_uindex
    on material_confidence_level (material_confidence_level);

-- material table
create table material
(
    id                      serial  not null
        constraint material_pk
            primary key,
    artifact_id             integer not null
        constraint material_artifact_master_id_fk
            references artifact_master,
    material_type_id        integer
        constraint material_material_type_lut_id_fk
            references material_type_lut,
    quantity                integer,
    "%composition"          integer,
    confidence_level_id     integer
        constraint material_material_confidence_level_id_fk
            references material_confidence_level,
    material_type_parent_id integer
        constraint material_material_id_fk
            references material
);

-- pb_isotope table
create table pb_isotope
(
    id          serial  not null
        constraint pb_isotope_pk
            primary key,
    artifact_id integer not null
        constraint pb_isotope_artifact_master_id_fk
            references artifact_master,
    isotope     varchar(50),
    value       varchar(50),
    date        date
);

-- collection table
create table collection
(
    id              serial  not null,
    artifact_id     integer not null
        constraint collection_artifact_master_id_fk
            references artifact_master,
    collection_name varchar(50)
);

-- prov_category_lut table
create table prov_category_lut
(
    id            integer     not null
        constraint prov_category_lut_pk
            primary key,
    prov_category varchar(50) not null
);

create unique index prov_category_lut_prov_category_uindex
    on prov_category_lut (prov_category);


-- provenience_intersite table
create table provenience_intersite
(
    id            serial  not null
        constraint provenience_intersite_pk
            primary key,
    artifact_id   integer not null
        constraint provenience_intersite_artifact_master_id_fk
            references artifact_master,
    p_category_id integer
        constraint provenience_intersite_prov_category_lut_id_fk
            references prov_category_lut,
    p_info        text
);

-- artifact_measurement table
create table artifact_measurement
(
    id          serial  not null
        constraint artifact_measurement_pk
            primary key,
    artifact_id integer not null
        constraint artifact_measurement_artifact_master_id_fk
            references artifact_master,
    length      integer,
    height      integer not null,
    width       integer
);

-- site_name_lut table
create table site_name_lut
(
    id        serial      not null
        constraint site_name_lut_pk
            primary key,
    site_name varchar(50) not null
);


create unique index site_name_lut_site_name_uindex
    on site_name_lut (site_name);

-- site_name table
create table site_name
(
    id           serial  not null
        constraint site_name_pk
            primary key,
    artifact_id  integer not null
        constraint site_name_artifact_master_id_fk
            references artifact_master,
    site_name_id integer
        constraint site_name_site_name_lut_id_fk
            references site_name_lut,
    location     varchar(50),
    country      varchar(50),
    comments     text,
    start_date   date,
    finish_date  integer
);

-- ref_categ_lut table
create table ref_categ_lut
(
    id      serial      not null
        constraint ref_categ_lut_pk
            primary key,
    r_categ varchar(50) not null
);

create unique index ref_categ_lut_r_categ_uindex
    on ref_categ_lut (r_categ);

-- reference table
create table reference
(
    id             serial  not null
        constraint reference_pk
            primary key,
    artifact_id    integer not null
        constraint reference_artifact_master_id_fk
            references artifact_master,
    r_category_id  integer
        constraint reference_ref_categ_lut_id_fk
            references ref_categ_lut,
    reference_info text
);

-- site_type_lut table
create table site_type_lut
(
    id        serial      not null
        constraint site_type_lut_pk
            primary key,
    site_type varchar(50) not null
);

create unique index site_type_lut_site_type_uindex
    on site_type_lut (site_type);


-- site_type table
create table site_type
(
    id           serial  not null
        constraint site_type_pk
            primary key,
    artifact_id  integer not null
        constraint site_type_artifact_master_id_fk
            references artifact_master,
    site_type_id integer
        constraint site_type_site_type_lut_id_fk
            references site_type_lut
);

-- artifact_preservation table
create table artifact_preservation
(
    id                              serial  not null
        constraint artifact_preservation_pk
            primary key,
    artifact_id                     integer not null
        constraint artifact_preservation_artifact_master_id_fk
            references artifact_master,
    preservation                    text    not null,
    artifact_preservation_parent_id integer
        constraint artifact_preservation_artifact_preservation_id_fk
            references artifact_preservation
);

-- artifact_element table
create table artifact_element
(
    id                         serial       not null
        constraint artifact_element_pk
            primary key,
    artifact_id                integer      not null
        constraint artifact_element_artifact_master_id_fk
            references artifact_master,
    artifact_element_name      varchar(100) not null,
    artifact_parent_element_id integer
        constraint artifact_element_artifact_element_id_fk
            references artifact_element
);

-- restoration table
create table restoration
(
    id          serial  not null
        constraint restoration_pk
            primary key,
    artifact_id integer not null
        constraint restoration_artifact_master_id_fk
            references artifact_master,
    date        date,
    updates     varchar(100),
    author      varchar(100)
);

-- artifact_related_people table
create table artifact_related_people
(
    id          serial       not null
        constraint artifact_related_people_pk
            primary key,
    artifact_id integer      not null
        constraint artifact_related_people_artifact_master_id_fk
            references artifact_master,
    person_name varchar(100) not null
);

-- artifact_style_lut table
create table artifact_style_lut
(
    id                  serial       not null
        constraint artifact_style_lut_pk
            primary key,
    artifact_style_name varchar(100) not null
);
create unique index artifact_style_lut_artifact_style_name_uindex
    on artifact_style_lut (artifact_style_name);


-- artifact_style table
create table artifact_style
(
    id                serial not null
        constraint artifact_style_pk
            primary key,
    artifact_id       integer not null
        constraint artifact_style_artifact_master_id_fk
            references artifact_master,
    artifact_style_id integer not null
        constraint artifact_style_artifact_style_lut_id_fk
            references artifact_style_lut
);

-- artifact_publication table
create table artifact_publication
(
    id          serial       not null
        constraint artifact_publication_pk
            primary key,
    artifact_id integer      not null
        constraint artifact_publication_artifact_master_id_fk
            references artifact_master,
    author_name varchar(100) not null,
    date        date
);










