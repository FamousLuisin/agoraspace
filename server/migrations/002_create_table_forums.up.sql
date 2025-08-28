CREATE TABLE tb_forums(
    id UUID PRIMARY KEY,
    title VARCHAR(75) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(15) DEFAULT 'active',
    is_public BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    owner UUID NOT NULL,
    CONSTRAINT fk_forums_users FOREIGN KEY (owner) REFERENCES tb_users(id)
);

CREATE TABLE tb_members(
    user_id UUID NOT NULL,
    forum_id UUID NOT NULL,
    role VARCHAR(15) DEFAULT 'participant',
    active BOOLEAN DEFAULT TRUE, 
    joined_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_members_users FOREIGN KEY (user_id) REFERENCES tb_users(id),
    CONSTRAINT fk_members_forums FOREIGN KEY (forum_id) REFERENCES tb_forums(id),
    CONSTRAINT tb_member_pkey PRIMARY KEY (user_id, forum_id)
);