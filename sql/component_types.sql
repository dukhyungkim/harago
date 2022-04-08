create table component_types
(
    id         bigserial
        constraint component_type_pk
            primary key,
    company    varchar(64) not null,
    component  varchar(32) not null,
    selection  varchar(32) not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

alter table component_types
    owner to postgres;

create unique index component_type_company_component_uindex
    on component_types (company, component);
