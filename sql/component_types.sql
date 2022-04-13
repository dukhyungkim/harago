create table component_types
(
    id        bigserial
        constraint component_types_pk
            primary key,
    component varchar(32) not null,
    type      varchar(32) not null
);

alter table component_types
    owner to harago;

create unique index component_types_component_uindex
    on component_types (component);

