DROP TABLE SENDER_INFORMATION cascade;
DROP TABLE Receiver_Information cascade;
DROP TABLE BOX cascade;
DROP TABLE List_of_items cascade;
DROP TABLE TOTAL_VALUE_ITEM cascade;

CREATE TABLE SENDER_INFORMATION
(
    sender_id  serial primary key not null,
    first_name text               not null,
    last_name  text               not null,
    main_phone text               not null,
    email      text               not null,
    address    text               not null,
    city       text               not null,
    country    text default 'United Kingdom',
    post_code  int                not null
);

CREATE TABLE Receiver_Information
(
    receiver_id serial primary key not null,
    sender_id   int references sender_information (sender_id)
);

CREATE TABLE BOX
(
    box_id                 serial primary key not null,
    first_name             text               not null,
    last_name              text               not null,
    passport_serial_number text               not null,
    main_phone             text               not null,
    address                text               not null,
    city                   text               not null,
    country                text default 'Uzbekistan',
    receiver_id            int references Receiver_Information (receiver_id)
);

CREATE TABLE LIST_OF_ITEMS
(
    list_id       serial primary key not null,
    item_name     text               not null,
    item_quantity text               not null,
    box_id        int references box (box_id)
);

CREATE TABLE TOTAL_VALUE_ITEM
(
    total_id    serial primary key not null,
    total_value int                not null,
    box_id      int references box (box_id)
);