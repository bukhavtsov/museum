-- insert country
INSERT INTO country (id, country_name)
VALUES (1, 'Беларусь');
-- insert region
INSERT INTO region (id, region_name, country_id)
VALUES (1, 'Гомельская область', 1);
INSERT INTO region (id, region_name, country_id)
VALUES (1, 'Ветковско-Стародубский', 1);

-- insert region
INSERT INTO city (id, city_name, region_id)
VALUES (1, 'Ветка', 1);
INSERT INTO city (id, city_name, region_id)
VALUES (1, 'Ветка', 2);

-- insert contacts
INSERT INTO museum_contacts ("id", "phone_number", "site_addr", "email")
VALUES (1, '8(02330) 4-26-05 ', 'vetka-museum.ru', 'vetkamuzejj@rambler.ru');

-- insert museum
INSERT INTO "museum" ("id", "museum_name", "city_id")
VALUES (1, 'Ветковский Музей', 1);

-- insert into object_group_lut
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (1, 'памятник');
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (2, 'Икона Николай чудотворец');
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (3, 'слева сверху Евдокия');
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (4, 'слева снизу Прокопий');
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (5, 'справа сверху женский образ');
INSERT INTO "object_group_lut" ("id", "object_group_name") VALUES (6, 'справа снизу ап. Павел');

-- insert object group
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 1, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 2, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 3, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 4, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 5, 1);
INSERT INTO "object_group" ("id", "object_group_id", "artifact_id") VALUES (1, 6, 1);


-- insert excavation_region
INSERT INTO "public".excavation_region ("id", city_id, "x_coordinate", "y_coordinate")
VALUES (1, 2, null, null);

-- insert artifact_master_phas
INSERT INTO "artifact_master_phas" ("id", "artifact_id", "museum_id", excavation_region_id,
                                    "reg_confidence_id",
                                    "excavator_full_name", "date_exc", "creator", "hist_culture_id", "rss_desc",
                                    "translation", "min_age",
                                    "max_age", reference_id, "artifact_info_photo", "photo")
VALUES (1, 1, 1, 1, null, 'Шкляров Ф.Г', '1979-02-17', null, null,
        'Никола́й Чудотво́рец[2]; Николай Уго́дник; Николай Мирлики́йский[1]; Святи́тель Николай (греч. Άγιος Νικόλαος — святой Николай; около 270 года, Патара[1], Ликия — около 345 года, Миры[1], Ликия) — святой в исторических церквях, архиепископ Мир Ликийских[3] (Византия). В христианстве почитается как чудотворец, на Востоке является покровителем путешествующих, заключённых и сирот[4], на Западе — покровителем практически всех слоёв общества, но в основном детей[5].', 'поясная центральноориенированная фигура святого с благословляющей десницей и открытым Евангелием. По обе стороны фигуры на уровне плеч ростовые фигуры Христа и Богородицы на облаках. Красная фелонь святителя украшена сложносоставными золотыми букетами; омофор – светло-розовый с бело-красными крестами и золотым растительным орнаментом.
               ', 219, 120, 'ru.wikipedia.org/wiki/Николай_Чудотворец',
        'утраты живописи и золотого покрытия, потёртости, осыпи по верхней кромке, вертикальная трещина.', '');

-- insert artifact_measurement
INSERT INTO "public"."artifact_measurement" ("id", "artifact_id", "length", "height", "width")
VALUES (1, 1, 527, 435, 34);

-- init material type
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (1, 'дерево');
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (2, 'паволока');
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (3, 'левкас');
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (4, 'яичная темпера');
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (5, 'творёное золото');
INSERT INTO "public"."material_type_lut" ("id", "material_type")
VALUES (6, 'сусальное золото');


-- init confidence level
INSERT INTO "public"."material_confidence_level" ("id", "material_confidence_level")
VALUES (1, 'low');
INSERT INTO "public"."material_confidence_level" ("id", "material_confidence_level")
VALUES (2, 'normal');
INSERT INTO "public"."material_confidence_level" ("id", "material_confidence_level")
VALUES (3, 'high');
INSERT INTO "public"."material_confidence_level" ("id", "material_confidence_level")
VALUES (4, 'very high');

-- init material
INSERT INTO "public"."material" ("id", "artifact_id", "material_id", "quantity", "%composition", "confidence_level_id")
VALUES (2, 1, 2, 1, null, 2);
INSERT INTO "public"."material" ("id", "artifact_id", "material_id", "quantity", "%composition", "confidence_level_id")
VALUES (3, 1, 3, 1, null, 2);
INSERT INTO "public"."material" ("id", "artifact_id", "material_id", "quantity", "%composition", "confidence_level_id")
VALUES (4, 1, 4, 1, null, 2);
INSERT INTO "public"."material" ("id", "artifact_id", "material_id", "quantity", "%composition", "confidence_level_id")
VALUES (5, 1, 5, 1, null, 2);
INSERT INTO "public"."material" ("id", "artifact_id", "material_id", "quantity", "%composition", "confidence_level_id")
VALUES (6, 1, 5, 1, null, 2);










