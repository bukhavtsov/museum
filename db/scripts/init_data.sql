-- insert country
INSERT INTO country (id, country_name)
VALUES (DEFAULT, 'Беларусь');
-- insert region
INSERT INTO region (id, region_name, country_id)
VALUES (DEFAULT, 'Гомельская область', 1);
INSERT INTO region (id, region_name, country_id)
VALUES (DEFAULT, 'Ветковско-Стародубский', 1);

-- insert location
INSERT INTO location (id, location_name, region_id)
VALUES (DEFAULT, 'Ветка', 1);
INSERT INTO location (id, location_name, region_id)
VALUES (DEFAULT, 'Ветка', 2);

-- insert museum
INSERT INTO "museum" ("id", "museum_name", "location_id")
VALUES (DEFAULT, 'Ветковский Музей', 1);

-- insert contacts
INSERT INTO museum_contacts ("id", "phone_number", "site_addr", "email", "museum_id")
VALUES (DEFAULT, '8(02330) 4-26-05 ', 'vetka-museum.ru', 'vetkamuzejj@rambler.ru', 1);

-- insert into object_group_lut
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'памятник');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'Икона Николай чудотворец');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'слева сверху Евдокия');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'слева снизу Прокопий');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'справа сверху женский образ');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'справа снизу ап. Павел');
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'Палеосные святые');


-- insert excavation_region
INSERT INTO excavation_region ("id", location_id, "x_coordinate", "y_coordinate")
VALUES (DEFAULT, 2, null, null);

-- insert reg_confidence_level
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (DEFAULT, 'very low');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (DEFAULT, 'low');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (DEFAULT, 'normal');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (DEFAULT, 'high');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (DEFAULT, 'very high');

-- transferred_by_lut
INSERT INTO transferred_by_lut (id, transferred_by)
VALUES (DEFAULT, 'поступила в составе коллекции Шклярова Ф.Г. 17.02.1979 г');

-- insert artifact_master
INSERT INTO artifact_master (id, artifact_id, museum_id, excavation_region_id, reg_confidence_id,
                             date_exc, creator, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (DEFAULT, 1, 1, 1, 5, '1979-02-17', null, null, null,
        'Композиция средника иконы: поясная центральноориенированная фигура святого с благословляющей десницей и открытым Евангелием. По обе стороны фигуры на уровне плеч ростовые фигуры Христа и Богородицы на облаках. Красная фелонь святителя украшена сложносоставными золотыми букетами; омофор – светло-розовый с бело-красными крестами и золотым растительным орнаментом.',
        219, 120, 'ru.wikipedia.org/wiki/Николай_Чудотворец',
        '/path/to/photo/image.jpg', 1);

-- insert object group
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 1, 1, null);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 2, 1, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 3, 1, 2);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 4, 1, 2);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 5, 1, 2);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 6, 1, 2);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 7, 1, 2);

-- insert artifact_preservation
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 1, 'утраты живописи и золотого покрытия', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 1, 'потёртости', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 1, 'осыпи по верхней кромке', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 1, 'вертикальная трещина', null);

-- insert artifact_measurement
INSERT INTO "artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (DEFAULT, 1, 527, 435, 34);

-- init material type
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'дерево');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'паволока');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'левкас');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'яичная темпера');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'творёное золото');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'сусальное золото');

-- insert material confidence level

INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (DEFAULT, 'very low');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (DEFAULT, 'low');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (DEFAULT, 'normal');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (DEFAULT, 'high');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (DEFAULT, 'very high');

-- init material
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 1, 2, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 1, 3, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 1, 4, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 1, 5, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 1, 5, 1, null, 2, null);

-- insert artifact_element
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'доска', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'ольха', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'трехсоставная', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'шпонки дубовые', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'пластевые', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'встречные', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'выступающие', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 1, 'отделанные калёвкой', 4);

-- insert artifact_style_lut
INSERT INTO artifact_style_lut (id, artifact_style_name)
VALUES (default, 'Ветковская школа');

-- insert artifact_style
INSERT INTO artifact_style (id, artifact_id, artifact_style_id)
VALUES (DEFAULT, 1, 1);



-- next card


-- insert into object_group_lut
INSERT INTO "object_group_lut" ("id", "object_group_name")
VALUES (DEFAULT, 'Евангелие-тетр');

-- insert country
INSERT INTO country (id, country_name)
VALUES (DEFAULT, 'Литва');

-- insert region
INSERT INTO region (id, region_name, country_id)
VALUES (DEFAULT, 'Литовская область', 2);

-- insert location
INSERT INTO location (id, location_name, region_id)
VALUES (DEFAULT, 'Вильно, дом Мамоничей', 3);

-- insert excavation_region
INSERT INTO excavation_region ("id", location_id, "x_coordinate", "y_coordinate")
VALUES (DEFAULT, 3, null, null);

-- transferred_by_lut
INSERT INTO transferred_by_lut (id, transferred_by)
VALUES (default, 'Приобретена в д. Леонтьево Добрушского р-на у Душечкина Якова Даниловича в 1981 (№ акта 209 от 28.08.1981., протокол от 24.08.1981.)
');

-- insert artifact_master
INSERT INTO artifact_master (id, artifact_id, museum_id, excavation_region_id, reg_confidence_id,
                             creator, date_exc, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (default, 2, 1, 2, 5, 'Пётр Тимофеев Мстиславец', '1575-03-30', null,
        'Переплёт: доски в коже, покрыты бархатом сиреневого цвета, на верхней крышке накладная доска с живописными наугольниками с изображениями 4-х евангелистов и металлическим литым средником (крест-распятие с предстоящими), застёжки.',
        null, 445, 320, 'ru.wikipedia.org/wiki/Николай_Чудотворец', '/path/to/photo/image.jpg', 2);

-- insert object_group
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 1, 2, null);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id", "object_group_parent_id")
VALUES (DEFAULT, 8, 2, 8);

-- insert material_type_lut
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'бумага с филигранями');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'двуцветная печать');
INSERT INTO "material_type_lut" ("id", "material_type")
VALUES (DEFAULT, 'гравюры');

-- insert material
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 2, 7, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 2, 8, 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", material_type_id, "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (DEFAULT, 2, 9, 1, null, 2, null);


-- insert object_group
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'Переплёт', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'доска', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'кожа', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'бархат', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'живопись', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'металлический средник', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'застёжки', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'фигурные жуковины', 9);

-- insert artifact_measurement
INSERT INTO "artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (DEFAULT, 2, 324, 203, 75);

-- insert artifact_preservation
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Отсутствует 1 лист', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'между 177-178 л', 5);

INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Вырван', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'между 376-381 л', 7);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Рукописные на бумаге в линейку Добрушской писчебумажной фабрики', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'лист с выходными данными отсутствует', 9);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'восстановлен рукописным', 9);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Бумага загрязнена', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'следы воска', 12);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'сырости', 12);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Бархат на переплёте', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'порван', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'выцвел', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'реставрирован современным бархатом малинового цвета', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'живопись на накладной доске почти утрачена', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2, 'Застёжки', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (DEFAULT, 2,'новодел', 20);

-- insert artifact_element
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'Орнамент', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'заставок – 10 с 10 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, '4 гравюры (после 1-го листа – евангелист Матфей', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, '105 об. – евангелист Марк', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'между 171-172 – евангелист Лука', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'между 280-281 – евангелист Иоанн)', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, '4 инициала (буквицы) с 4 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, 'маргинальных рамок – 15 с 9 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (DEFAULT, 2, ' вязь киноварная', 17);

-- insert artifact_style
INSERT INTO artifact_style (id, artifact_id, artifact_style_id)
VALUES (DEFAULT, 2, 1);
