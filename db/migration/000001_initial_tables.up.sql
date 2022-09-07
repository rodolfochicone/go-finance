create table users(
    id serial primary key not null ,
    username varchar not null,
    password varchar not null,
    email varchar unique not null,
    created_at timestamptz not null default now()
);

create table categories(
    id serial primary key not null ,
    user_id int not null references users(id),
    title varchar not null,
    type varchar not null,
    description varchar not null,
    created_at timestamptz not null default now()
);

create table accounts(
    id serial primary key not null ,
    user_id int not null references users(id),
    category_id int not null references categories(id),
    title varchar not null,
    type varchar not null,
    description varchar not null,
    value integer not null,
    date date not null,
    created_at timestamptz not null default now()
);