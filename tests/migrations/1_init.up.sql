INSERT INTO orders (customer_name, pizza_type, pizza_size)
VALUES ('Katia', 'Gorgonzola and Pear', 'LARGE')
ON CONFLICT DO NOTHING;
