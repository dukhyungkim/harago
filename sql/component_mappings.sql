create table component_mappings
(
    id         bigint default nextval('component_mapping_id_seq'::regclass) not null
        constraint component_type_pk
            primary key,
    company    varchar(64)                                               not null,
    type       varchar(32)                                               not null,
    component  varchar(32)                                               not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

alter table component_mappings
    owner to harago;

create unique index component_type_company_component_uindex
    on component_mappings (company, type);
