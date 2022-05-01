create table rpc_service_tab (
    id int auto_increment primary key comment 'auto inc id',
    name varchar(256) comment 'service name, can duplicate under different hierarchy',
    service_key varchar(767) comment 'service auth key',
    is_service bool comment 'false indicates a directory entry',
    parent_id int comment 'refers to auto inc id',
    complete_path varchar(767) unique key comment 'the fully qualified service name'
) comment 'rpc service meta tab'