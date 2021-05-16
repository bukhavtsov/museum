-- insert country
INSERT INTO country (country_name)
VALUES ('Беларусь');
-- insert region
INSERT INTO region (region_name, country_id)
VALUES ('Гомельская область', 1);
INSERT INTO region (region_name, country_id)
VALUES ('Ветковско-Стародубский', 1);

-- insert location
INSERT INTO location (location_name, region_id)
VALUES ('Ветка', 1);
INSERT INTO location (location_name, region_id)
VALUES ('Ветка', 2);

-- insert museum
INSERT INTO "museum" ("museum_name", "location_id")
VALUES ('Ветковский Музей', 1);

-- insert contacts
INSERT INTO museum_contacts ("phone_number", "site_addr", "email", "museum_id")
VALUES ('8(02330) 4-26-05 ', 'vetka-museum.ru', 'vetkamuzejj@rambler.ru', 1);

-- insert into object_group_lut
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('памятник');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('Икона Николай чудотворец');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('слева сверху Евдокия');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('слева снизу Прокопий');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('справа сверху женский образ');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('справа снизу ап. Павел');
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('Палеосные святые');


-- insert excavation_region
INSERT INTO excavation_region (location_id, "x_coordinate", "y_coordinate")
VALUES (2, null, null);

-- insert reg_confidence_level
INSERT INTO reg_confidence_level (reg_confidence_level)
VALUES ('very low');
INSERT INTO reg_confidence_level (reg_confidence_level)
VALUES ('low');
INSERT INTO reg_confidence_level (reg_confidence_level)
VALUES ('normal');
INSERT INTO reg_confidence_level (reg_confidence_level)
VALUES ('high');
INSERT INTO reg_confidence_level (reg_confidence_level)
VALUES ('very high');

-- transferred_by_lut
INSERT INTO transferred_by_lut (transferred_by)
VALUES ('поступила в составе коллекции Шклярова Ф.Г. 17.02.1979 г');

-- insert artifact_master
INSERT INTO artifact_master (artifact_id, museum_id, excavation_region_id, reg_confidence_id,
                             date_exc, creator, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (1, 1, 1, 5, '1979-02-17', null, null, null,
        'Композиция средника иконы: поясная центральноориенированная фигура святого с благословляющей десницей и открытым Евангелием. По обе стороны фигуры на уровне плеч ростовые фигуры Христа и Богородицы на облаках. Красная фелонь святителя украшена сложносоставными золотыми букетами; омофор – светло-розовый с бело-красными крестами и золотым растительным орнаментом.',
        219, 120, 'ru.wikipedia.org/wiki/Николай_Чудотворец',
        '/path/to/photo/image.jpg', 1);

-- -- insert object group
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (1, 1, null);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (2, 1, 1);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (3, 1, 2);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (4, 1, 2);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (5, 1, 2);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (6, 1, 2);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (7, 1, 2);

-- -- insert artifact_preservation
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (1, 'утраты живописи и золотого покрытия', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (1, 'потёртости', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (1, 'осыпи по верхней кромке', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (1, 'вертикальная трещина', null);

-- insert ArtifactMeasurement
INSERT INTO "artifact_measurement" ("artifact_id", "length", "height", "width")
VALUES (1, 527, 435, 34);

-- init material type
INSERT INTO "material_type_lut" ("material_type")
VALUES ('дерево');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('паволока');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('левкас');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('яичная темпера');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('творёное золото');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('сусальное золото');

-- insert material confidence level

INSERT INTO material_confidence_level (material_confidence_level)
VALUES ('very low');
INSERT INTO material_confidence_level (material_confidence_level)
VALUES ('low');
INSERT INTO material_confidence_level (material_confidence_level)
VALUES ('normal');
INSERT INTO material_confidence_level (material_confidence_level)
VALUES ('high');
INSERT INTO material_confidence_level (material_confidence_level)
VALUES ('very high');

-- -- init material
-- INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
--                         "confidence_level_id", "material_type_parent_id")
-- VALUES (1, 2, 1, null, 2, null);
-- INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
--                         "confidence_level_id", "material_type_parent_id")
-- VALUES (1, 3, 1, null, 2, null);
-- INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
--                         "confidence_level_id", "material_type_parent_id")
-- VALUES (1, 4, 1, null, 2, null);
-- INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
--                         "confidence_level_id", "material_type_parent_id")
-- VALUES (1, 5, 1, null, 2, null);
-- INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
--                         "confidence_level_id", "material_type_parent_id")
-- VALUES (1, 5, 1, null, 2, null);

-- -- insert artifact_element
INSERT INTO artifact_element (artifact_id, elements)
VALUES (1, '{"name":"parent element","children":[{"name":"child 1","children":[{"name":"sub child 1"}]},{"name":"child 2"}]}');
INSERT INTO artifact_element (artifact_id, elements)
VALUES (1, '{"name":"parent element second","children":[{"name":"child 1 second","children":[{"name":"sub child 1 second"}]},{"name":"child 2 second"}]}');

-- insert artifact_style_lut
INSERT INTO artifact_style_lut (artifact_style_name)
VALUES ('Ветковская школа');

-- -- insert artifact_style
-- INSERT INTO artifact_style (artifact_id, artifact_style_id)
-- VALUES (1, 1, 1);



-- next card


-- insert into object_group_lut
INSERT INTO "object_group_lut" ("object_group_name")
VALUES ('Евангелие-тетр');

-- insert country
INSERT INTO country (country_name)
VALUES ('Литва');

-- insert region
INSERT INTO region (region_name, country_id)
VALUES ('Литовская область', 2);

-- insert location
INSERT INTO location (location_name, region_id)
VALUES ('Вильно, дом Мамоничей', 3);

-- insert excavation_region
INSERT INTO excavation_region (location_id, "x_coordinate", "y_coordinate")
VALUES (3, null, null);

-- transferred_by_lut
INSERT INTO transferred_by_lut (transferred_by)
VALUES ('Приобретена в д. Леонтьево Добрушского р-на у Душечкина Якова Даниловича в 1981 (№ акта 209 от 28.08.1981., протокол от 24.08.1981.)
');

-- insert artifact_master
INSERT INTO artifact_master (artifact_id, museum_id, excavation_region_id, reg_confidence_id,
                             creator, date_exc, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (2, 1, 2, 5, 'Пётр Тимофеев Мстиславец', '1575-03-30', null,
        'Переплёт: доски в коже, покрыты бархатом сиреневого цвета, на верхней крышке накладная доска с живописными наугольниками с изображениями 4-х евангелистов и металлическим литым средником (крест-распятие с предстоящими), застёжки.',
        null, 445, 320, 'ru.wikipedia.org/wiki/Николай_Чудотворец', '/path/to/photo/image.jpg', 2);

-- -- insert object_group
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (1, 2, null);
-- INSERT INTO "object_group" ("object_group_id", "artifact_id", "object_group_parent_id")
-- VALUES (8, 2, 8);

-- insert material_type_lut
INSERT INTO "material_type_lut" ("material_type")
VALUES ('бумага с филигранями');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('двуцветная печать');
INSERT INTO "material_type_lut" ("material_type")
VALUES ('гравюры');

-- -- insert material
INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (2, 7, 1, null, 2, null);
INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (2, 8, 1, null, 2, null);
INSERT INTO "material" ("artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (2, 9, 1, null, 2, null);

-- -- insert artifact_element
INSERT INTO artifact_element (artifact_id, elements)
VALUES (2, '{"name":"parent element","children":[{"name":"child 1","children":[{"name":"sub child 1"}]},{"name":"child 2"}]}');
INSERT INTO artifact_element (artifact_id, elements)
VALUES (2, '{"name":"parent element second","children":[{"name":"child 1 second","children":[{"name":"sub child 1 second"}]},{"name":"child 2 second"}]}');


-- insert ArtifactMeasurement
INSERT INTO "artifact_measurement" ("artifact_id", "length", "height", "width")
VALUES (2, 324, 203, 75);

-- -- insert artifact_preservation
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Отсутствует 1 лист', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'между 177-178 л', 5);
--
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Вырван', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'между 376-381 л', 7);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Рукописные на бумаге в линейку Добрушской писчебумажной фабрики', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'лист с выходными данными отсутствует', 9);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'восстановлен рукописным', 9);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Бумага загрязнена', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'следы воска', 12);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'сырости', 12);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Бархат на переплёте', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'порван', 15);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'выцвел', 15);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'реставрирован современным бархатом малинового цвета', 15);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'живопись на накладной доске почти утрачена', 15);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2, 'Застёжки', null);
-- INSERT INTO artifact_preservation ("artifact_id", "preservation", "artifact_preservation_parent_id")
-- VALUES (2,'новодел', 20);

-- -- insert artifact_element
INSERT INTO artifact_element (artifact_id, elements)
VALUES (2, '{"name":"parent element","children":[{"name":"child 1 third","children":[{"name":"sub child 1 third"}]},{"name":"child 2 third"}]}');
INSERT INTO artifact_element (artifact_id, elements)
VALUES (2, '{"name":"parent element second third","children":[{"name":"child 1 second third","children":[{"name":"sub child 1 second third"}]},{"name":"child 2 second third"}]}');

-- -- insert artifact_style
INSERT INTO artifact_style (artifact_id, artifact_style_id)
VALUES (2, 1);
