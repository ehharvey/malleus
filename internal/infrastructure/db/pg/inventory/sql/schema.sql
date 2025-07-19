-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE domains (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE
);
 
-- CREATE TABLE nodes (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

--     domain_id UUID NOT NULL,
--     FOREIGN KEY (domain_id) REFERENCES domains
-- );

-- CREATE TABLE network_interfaces (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

--     node_id UUID NOT NULL,
--     FOREIGN KEY (node_id) REFERENCES nodes,

--     network_interface_type_id UUID NOT NULL,
--     FOREIGN KEY (network_interface_type_id) REFERENCES network_interface_types,

--     network_bridge_id UUID,
--     FOREIGN KEY (network_bridge_id) REFERENCES network_bridges,

--     mac MACADDR UNIQUE
-- );

-- CREATE TABLE network_interface_types (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

--     name TEXT NOT NULL UNIQUE
-- );

-- CREATE TABLE networks (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

--     addr CIDR
-- );

-- CREATE TABLE network_ips (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

--     addr INET,

--     network_id UUID NOT NULL,
--     FOREIGN KEY (network_id) REFERENCES networks,

--     network_interface_id UUID NOT NULL,
--     FOREIGN KEY (network_interface_id) REFERENCES network_interfaces
-- );

-- CREATE TABLE network_bridges (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
-- );