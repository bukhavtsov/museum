-- insert country
INSERT INTO country (id, country_name)
VALUES (1, 'Беларусь');
-- insert region
INSERT INTO region (id, region_name, country_id)
VALUES (1, 'Гомельская область', 1);
INSERT INTO region (id, region_name, country_id)
VALUES (2, 'Ветковско-Стародубский', 1);

-- insert location
INSERT INTO location (id, location_name, region_id)
VALUES (1, 'Ветка', 1);
INSERT INTO location (id, location_name, region_id)
VALUES (2, 'Ветка', 2);

-- insert museum
INSERT INTO "museum" ("id", "museum_name", "location_id")
VALUES (1, 'Ветковский Музей', 1);

-- insert contacts
INSERT INTO museum_contacts ("id", "phone_number", "site_addr", "email", "museum_id")
VALUES (1, '8(02330) 4-26-05 ', 'vetka-museum.ru', 'vetkamuzejj@rambler.ru', 1);

-- insert excavation_region
INSERT INTO excavation_region ("id", location_id, "x_coordinate", "y_coordinate")
VALUES (1, 2, null, null);

-- insert reg_confidence_level
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (1, 'very low');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (2, 'low');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (3, 'normal');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (4, 'high');
INSERT INTO reg_confidence_level (id, reg_confidence_level)
VALUES (5, 'very high');

-- transferred_by_lut
INSERT INTO transferred_by_lut (id, transferred_by)
VALUES (1, 'поступила в составе коллекции Шклярова Ф.Г. 17.02.1979 г');

-- insert artifact_master
INSERT INTO artifact_master (id, museum_id, excavation_region_id, reg_confidence_id,
                             date_exc, creator, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (1, 1, 1, 5, '1979-02-17', null, null, null,
        'Композиция средника иконы: поясная центральноориенированная фигура святого с благословляющей десницей и открытым Евангелием. По обе стороны фигуры на уровне плеч ростовые фигуры Христа и Богородицы на облаках. Красная фелонь святителя украшена сложносоставными золотыми букетами; омофор – светло-розовый с бело-красными крестами и золотым растительным орнаментом.',
        219, 120, 'ru.wikipedia.org/wiki/Николай_Чудотворец',
        '/path/to/photo/image.jpg', 1);

-- insert object group
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (1, 'памятник', 1, null);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (2, 'Икона Николай чудотворец', 1, 1);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (3, 'слева сверху Евдокия', 1, 2);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (4, 'слева снизу Прокопий', 1, 2);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (5, 'справа сверху женский образ', 1, 2);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (6, 'справа снизу ап. Павел', 1, 2);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (7, 'Палеосные святые', 1, 2);

-- insert artifact_preservation
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (1, 1, 'утраты живописи и золотого покрытия', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (2, 1, 'потёртости', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (3, 1, 'осыпи по верхней кромке', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (4, 1, 'вертикальная трещина', null);

-- insert artifact_measurement
INSERT INTO "artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (1, 1, 527, 435, 34);

-- insert material confidence level

INSERT INTO "material_confidence_level" (id, material_confidence_level)
VALUES (1, 'very low');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (2, 'low');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (3, 'normal');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (4, 'high');
INSERT INTO material_confidence_level (id, material_confidence_level)
VALUES (5, 'very high');

-- insert material
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (1, 1, 'паволока', 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (2, 1, 'левкас', 1, null, 2, 1);
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (3, 1, 'яичная темпера', 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (4, 1, 'творёное золото', 1, null, 2, 3);

-- insert artifact_element
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (1, 1, 'доска', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (2, 1, 'ольха', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (3, 1, 'трехсоставная', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (4, 1, 'шпонки дубовые', 1);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (5, 1, 'пластевые', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (6, 1, 'встречные', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (7, 1, 'выступающие', 4);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (8, 1, 'отделанные калёвкой', 4);

-- insert artifact_style_lut
INSERT INTO artifact_style_lut (id, artifact_style_name)
VALUES (1, 'Ветковская школа');

-- insert artifact_style
INSERT INTO artifact_style (id, artifact_id, artifact_style_id)
VALUES (1, 1, 1);



-- next card

-- insert country
INSERT INTO country (id, country_name)
VALUES (2, 'Литва');

-- insert region
INSERT INTO region (id, region_name, country_id)
VALUES (3, 'Литовская область', 2);

-- insert location
INSERT INTO location (id, location_name, region_id)
VALUES (3, 'Вильно, дом Мамоничей', 3);

-- insert excavation_region
INSERT INTO excavation_region ("id", location_id, "x_coordinate", "y_coordinate")
VALUES (2, 3, null, null);

-- transferred_by_lut
INSERT INTO transferred_by_lut (id, transferred_by)
VALUES (2, 'Приобретена в д. Леонтьево Добрушского р-на у Душечкина Якова Даниловича в 1981 (№ акта 209 от 28.08.1981., протокол от 24.08.1981.)
');

-- insert artifact_master
INSERT INTO artifact_master (id, museum_id, excavation_region_id, reg_confidence_id,
                             creator, date_exc, hist_culture_id, "desc", translation,
                             min_age, max_age, artifact_info_photo, photo, transferred_by_id)
VALUES (2, 1, 2, 5, 'Пётр Тимофеев Мстиславец', '1575-03-30', null,
        'Переплёт: доски в коже, покрыты бархатом сиреневого цвета, на верхней крышке накладная доска с живописными наугольниками с изображениями 4-х евангелистов и металлическим литым средником (крест-распятие с предстоящими), застёжки.',
        null, 445, 320, 'ru.wikipedia.org/wiki/Николай_Чудотворец', '/path/to/photo/image.jpg', 2);

-- insert object_group
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (8, 'памятник', 2, null);
INSERT INTO "object_group" ("id", "object_group_name", "artifact_id", "object_group_parent_id")
VALUES (9, 'Евангелие-тетр', 2, 8);

-- insert material
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (5, 2, 'бумага с филигранями', 1, null, 2, null);
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (6, 2, 'двуцветная печать', 1, null, 2, 5);
INSERT INTO "material" ("id", "artifact_id", "material_type", "quantity", "%composition",
                        "confidence_level_id", "material_type_parent_id")
VALUES (7, 2, 'гравюры', 1, null, 2, 5);


-- insert object_group
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (9, 2, 'Переплёт', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (10, 2, 'доска', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (11, 2, 'кожа', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (12, 2, 'бархат', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (13, 2, 'живопись', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (14, 2, 'металлический средник', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (15, 2, 'застёжки', 9);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (16, 2, 'фигурные жуковины', 9);

-- insert artifact_measurement
INSERT INTO "artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (2, 2, 324, 203, 75);

-- insert artifact_preservation
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (5, 2, 'Отсутствует 1 лист', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (6, 2, 'между 177-178 л', 5);

INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (7, 2, 'Вырван', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (8, 2, 'между 376-381 л', 7);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (9, 2, 'Рукописные на бумаге в линейку Добрушской писчебумажной фабрики', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (10, 2, 'лист с выходными данными отсутствует', 9);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (11, 2, 'восстановлен рукописным', 9);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (12, 2, 'Бумага загрязнена', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (13, 2, 'следы воска', 12);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (14, 2, 'сырости', 12);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (15, 2, 'Бархат на переплёте', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (16, 2, 'порван', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (17, 2, 'выцвел', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (18, 2, 'реставрирован современным бархатом малинового цвета', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (19, 2, 'живопись на накладной доске почти утрачена', 15);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (20, 2, 'Застёжки', null);
INSERT INTO artifact_preservation ("id", "artifact_id", "preservation", "artifact_preservation_parent_id")
VALUES (21, 2, 'новодел', 20);

-- insert artifact_element
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (17, 2, 'Орнамент', null);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (18, 2, 'заставок – 10 с 10 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (19, 2, '4 гравюры (после 1-го листа – евангелист Матфей', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (20, 2, '105 об. – евангелист Марк', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (21, 2, 'между 171-172 – евангелист Лука', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (22, 2, 'между 280-281 – евангелист Иоанн)', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (23, 2, '4 инициала (буквицы) с 4 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (24, 2, 'маргинальных рамок – 15 с 9 досок', 17);
INSERT INTO artifact_element (id, artifact_id, artifact_element_name, artifact_parent_element_id)
VALUES (25, 2, ' вязь киноварная', 17);