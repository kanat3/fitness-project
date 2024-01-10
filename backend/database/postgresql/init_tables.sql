INSERT INTO diet_templates (diet) VALUES
    ('Завтрак: Омлет с овощами, овсянка с фруктами; Обед: Куриный салат, рисовые хлопья; Ужин: Греческий салат, курица с овощами'),
    ('Завтрак: Греческий йогурт с орехами и медом, фруктовый салат; Обед: Тунцовый сандвич, овощной суп; Ужин: Паста с томатным соусом, свежие овощи'),
    ('Завтрак: Банановый смузи, тосты с авокадо; Обед: Овощной рисовый салат, запеченный лосось; Ужин: Цезарь салат, гриль курица с картошкой'),
    ('Завтрак: Йогурт с мюсли, фруктовый коктейль; Обед: Овощной суп, кус-кус с овощами; Ужин: Стейк из индейки, овощи на гриле');

INSERT INTO workout_templates (workout) VALUES
    ('Утренний бег 3 км, пресс; Вечерний йога и стретчинг'),
    ('Кардио тренировка: бег 5 км; Занятия на тренажерах: подтягивания, отжимания'),
    ('Силовая тренировка: жим штанги, приседания; Кардио: велотренировка 30 мин'),
    ('Плавание 1 км, занятия по методике HIIT: высокоинтенсивные интервальные тренировки');

INSERT INTO users (id_users, first_name, second_name, last_name, phone, email, profile_img, created, password)
VALUES ( '99',
  'Тестовый',
  'Тест',
  'Тестов',
  '89751614387',
  'test@mail.coom',
  'https://example.com/profile_images/some-jpg.png',
  CURRENT_TIMESTAMP,
  '$2a$14$0PLKrEp1bFk/6fy74z24SeAPcwjVZk73sBf/jq5BjkIh1vv0.NNU2'
);