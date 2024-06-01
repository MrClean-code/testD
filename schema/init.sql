-- CREATE TABLE IF NOT EXISTS deals (
--                                      id SERIAL PRIMARY KEY,
--                                      name TEXT,
--                                      owner TEXT,
--                                      price TEXT,
--                                      count_reviews TEXT,
--                                      score TEXT,
--                                      link TEXT
-- );
-- INSERT INTO deals (name, owner, price, count_reviews, score, link)
-- VALUES ('Название услуги', 'Владелец услуги', 'Цена услуги', 'Количество отзывов', 'Оценка', 'Ссылка на продавца');

-- TRUNCATE TABLE deals RESTART IDENTITY;
-- CREATE TABLE IF NOT EXISTS deals (
--                                      id SERIAL PRIMARY KEY,
--                                      name TEXT,
--                                      owner TEXT,
--                                      price INTEGER,
--                                      count_reviews INTEGER,
--                                      score DOUBLE PRECISION,
--                                      link TEXT
-- );

-- CREATE TABLE public.file_links (
--                                    id SERIAL PRIMARY KEY,
--                                    filename VARCHAR(255),
--                                    url VARCHAR(255),
--                                    size INTEGER
-- );


-- CREATE TABLE public.search_deal (
--                                    id SERIAL PRIMARY KEY,
--                                    name VARCHAR(255)
-- );

-- CREATE TABLE public.data
-- (
--     id     SERIAL PRIMARY KEY,
--     name   TEXT,
--     region TEXT,
--     seal   DOUBLE PRECISION,
--     date DATE,
--     search_deal_id INTEGER,
--     FOREIGN KEY (search_deal_id) references public.search_deal(id)
-- );

-- CREATE INDEX idx_data_name ON public.data(name);
-- CREATE INDEX idx_data_region ON public.data(region);
-- CREATE INDEX idx_data_date ON public.data(date);
-- CREATE INDEX idx_data_sale ON public.data(seal);

-- ALTER TABLE public.search_deal ADD COLUMN percent DOUBLE PRECISION;

-- SELECT data.id, data.name, data.region, data.seal, data.date,
--        search_deal.name, search_deal.percent FROM data
-- JOIN search_deal ON data.search_deal_id = search_deal.id
--                         and search_deal.name='Ремонт, окраска и пошив обуви'
--                         and data.date BETWEEN '2022-01-01 ' AND '2024-03-03'
--                         and data.region = 'Воронежская область'