-- +goose Up
-- +goose StatementBegin

create table "coordinate"
(
    id serial primary key,
    x integer,
    y integer
);

alter table "coordinate" 
    add constraint coordinate_unique unique (x,y);

-- +goose StatementEnd

-- +goose Down
    drop table "coordinate" cascade;
-- +goose End