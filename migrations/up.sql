create extension if not exists citext with schema public;

create table if not exists salons
(
    id      bigserial primary key,
    address citext unique
);

create table if not exists employees
(
    id        bigserial primary key,
    fullname  citext,
    role      citext,
    salon_id  bigint,
    status    text,
    email     citext,
    gender    citext  default null,
    isWorking boolean default false,
    foreign key (salon_id) references salons (id)
);

create table if not exists customers
(
    id       bigserial primary key,
    fullname citext not null,
    email    citext default null,
    gender   citext default null,
    address  citext default null
);

create table if not exists haircuts
(
    id          bigserial primary key,
    name        citext not null,
    description text
);


create table if not exists prices
(
    id          bigserial primary key,
    price_value integer not null,
    haircut_id  bigint,
    foreign key (haircut_id) references haircuts (id)
);

create table if not exists deals
(
    id          bigserial primary key,
    customer_id bigint,
    haircut_id  bigint,
    employee_id bigint,
    price_id    bigint,
    date        timestamp with time zone default now(),
    foreign key (customer_id) references customers (id),
    foreign key (haircut_id) references haircuts (id),
    foreign key (employee_id) references employees (id),
    foreign key (price_id) references prices (id)
);
