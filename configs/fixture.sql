INSERT INTO public.customer (id, email,password_hash,billing_address,credit_card,shipping_address) VALUES
(1,'test1@mail.com','e7e0140e34e645699a30e23374ea17294c2dbf47986ed1db0bcfad28943cc49c','33 S Main Street, Warren,ar, 31631  United States','{
	"Number": "4001297386906500",
	"Holder": "Henrik Chandler",
	"Date": "10/27",
	"CVV": "691"
}','35 Kai Kane Road, Captain Cook,hi, 96304  United States'),
(2,'test2@mail.com','e7e0140e34e645699a30e23374ea17294c2dbf47986ed1db0bcfad28943cc49c','13 Duxbury Drive, Raleigh,nc, 23603  United States','{
	"Number": "4006488490726828",
	"Holder": "Annika Obialo",
	"Date": "04/28",
	"CVV": "762"
}','1 E 14th Court, Anchorage,ak, 99504  United States'),
(3,'test3@mail.com','e7e0140e34e645699a30e23374ea17294c2dbf47986ed1db0bcfad28943cc49c','25 Near Avenue, Evanston,wy, 82930  United States',NULL,'10 Prairie Street, Arena,wi, 53503  United States');

INSERT INTO public.product (id,articul,price,delivery_time_description,status,inventory,vendor,category) VALUES
(1,'mars',5,'1 day','active',3,'MARS choco','food'),
(2,'iPhone SE',700,'3-7 business days','active',30,'Apple','smartphones'),
(3,'Xiaomi A1',300,'3-7 business days','active',10,'Xiaomi','smartphones'),
(4,'Galaxy S Mini',400,'2 days','active',3,'Samsung','smartphones'),
(5,'Super Table',100,'up to 10 days','active',2,'IKEA','furniture'),
(6,'Super Chair',40,'up to 12 days','active',5,'IKEA','furniture'),
(7,'Super Wardrobe',200,'up to 18 days','inactive',2,'ALLIANCE','furniture');

INSERT INTO public.image (id,product_id) VALUES
('d8a4cebc-ce56-4b77-920d-998b62f0583c',1),
('3e586ca5-640a-4d35-a292-32eaf8da8e0d',1),
('570e96ec-c84d-4048-93d8-757d7d2e26e3',2),
('600f798f-0691-46b9-8da3-f8afc1283925',2),
('e4afd6b2-4e03-4293-babe-cfdba17b46e5',3),
('8385fd1e-8124-400d-8532-4de988def963',3),
('4aa69d89-6675-4e1c-9aeb-f05857c16681',4),
('b7f7b353-3737-4a44-99dc-a35482e7d438',4),
('d757eeb3-8797-49c7-ac0f-109097b46142',5),
('377cd8e2-bf3c-4f63-addf-900566e0bc13',5),
('b28edc77-b948-480f-a88e-f27a74b5447e',6),
('7b4f23e5-c7e6-472a-b928-b53e0a4bab3f',6),
('d336fd14-fd5c-4c5a-9a35-b2362688d774',6),
('6263590d-37c5-4c28-98af-53b779db37cb',7);

--password for all users is Q123456 (SHA256-hash)
INSERT INTO public.merchandiser (id, username,password_hash) VALUES
(1,'merchandiser01','e7e0140e34e645699a30e23374ea17294c2dbf47986ed1db0bcfad28943cc49c'),
(2,'merchandiser02','e7e0140e34e645699a30e23374ea17294c2dbf47986ed1db0bcfad28943cc49c');
