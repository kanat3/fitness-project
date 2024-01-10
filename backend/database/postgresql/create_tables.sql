create table if not exists users (
    id_users serial primary key,
    first_name varchar(25) not null,
    second_name varchar(25) not null,
    last_name varchar(25) not null,
    phone varchar(20) check (phone ~ '^(\s*)?(\+)?([- _():=+]?\d[- _():=+]?){10,14}(\s*)?$') not null UNIQUE,
    email varchar(254) not null UNIQUE,
    profile_img varchar(400),
    created timestamp not null,
    password varchar(120) not null
);

create table if not exists contacts (
    option_link varchar(100),
    tg_link varchar(100),
    inst_link varchar(100),
    vk_link varchar(100),
    refusers integer,
    foreign key (refusers) references users (id_users),
    id_contacts serial primary key
);

create table if not exists workout_plan (
    id_workout_plan serial primary key,
    day_plan text,
    dplan_created timestamp,
    week_plan text,
    wplan_created timestamp
);

create table if not exists workout_plan_list (
    refusers integer,
    foreign key (refusers) references users (id_users),
    refworkout_plan integer,
    foreign key (refworkout_plan) references workout_plan (id_workout_plan)  
);

create table if not exists diet_plan (
    id_diet_plan serial primary key,
    day_plan text,
    dplan_created timestamp,
    week_plan text,
    wplan_created timestamp
);

create table if not exists diet_plan_list (
    refusers integer,
    foreign key (refusers) references users (id_users),
    refdiet_plan integer,
    foreign key (refdiet_plan) references diet_plan (id_diet_plan)
);

create table if not exists coach (
    id_coach serial primary key,
    first_name varchar(25) not null,
    second_name varchar(25) not null,
    last_name varchar(25) not null,
    email varchar(254) not null UNIQUE,
    profile_img varchar(400),
    created timestamp not null
);

create table if not exists fit_classes (
    id_fit_classes serial primary key,
    name varchar(40) not null,
    type text
);

create table if not exists coach_classes_list (
    refcoach integer,
    foreign key (refcoach) references coach (id_coach),
    reffit_classes integer,
    foreign key (reffit_classes) references fit_classes (id_fit_classes)
);

create table if not exists bid (
    id_bid serial primary key,
    optional_goal text,
    optional_message text,
    created timestamp not null
);

create table if not exists bid_by_user (
    refid_bid integer,
    foreign key (refid_bid) references bid (id_bid),
    refusers integer,
    foreign key (refusers) references users (id_users)
);

create table if not exists fit_classes_bid (
    refid_bid integer,
    foreign key (refid_bid) references bid (id_bid),
    reffit_classes integer,
    foreign key (reffit_classes) references fit_classes (id_fit_classes)
);

create table if not exists goal (
    id_goal serial primary key,
    refbid integer,
    foreign key (refbid) references bid (id_bid),
    type varchar(100) not null
);

create table if not exists sex (
    id_sex serial primary key,
    refbid integer,
    foreign key (refbid) references bid (id_bid),
    type varchar(100) not null
);

create table if not exists centers (
    id_centers serial primary key,
    name varchar(40) not null,
    address varchar(100) not null,
    phone varchar(20) check (phone ~ '^(\s*)?(\+)?([- _():=+]?\d[- _():=+]?){10,14}(\s*)?$') not null
);

create table if not exists centers_list (
    refbid integer,
    foreign key (refbid) references bid (id_bid),
    refcenters integer,
    foreign key (refcenters) references centers (id_centers)
);

create table if not exists day_time (
    id_day_time serial primary key,
    refbid integer,
    foreign key (refbid) references bid (id_bid),
    type varchar(40) not null
);

create table if not exists status (
    id_status serial primary key,
    refbid integer,
    foreign key (refbid) references bid (id_bid),
    type varchar(40) not null
);

create table if not exists diet_templates (
    id serial primary key,
    diet text not null
);

create table if not exists workout_templates (
    id serial primary key,
    workout text not null
);