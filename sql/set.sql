DROP TABLE IF EXISTS users          CASCADE;
DROP TABLE IF EXISTS voters         CASCADE;
DROP TABLE IF EXISTS voter_quiz     CASCADE;
DROP TABLE IF EXISTS presentation   CASCADE;
DROP TABLE IF EXISTS slideorder     CASCADE;
DROP TABLE IF EXISTS quiz           CASCADE;
DROP TABLE IF EXISTS convertedslide CASCADE;
DROP TABLE IF EXISTS vote           CASCADE;
DROP TABLE IF EXISTS question       CASCADE;

DROP FUNCTION IF EXISTS gen_random_bytes;
DROP FUNCTION IF EXISTS random_string;
DROP FUNCTION IF EXISTS unique_random;


create function gen_random_bytes(int) returns bytea as
'$libdir/pgcrypto', 'pg_random_bytes' language c strict;

create function random_string(len int) returns text as $$
declare
  chars text[] = '{0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z}';
  result text = '';
  i int = 0;
  rand bytea;
begin
  -- generate secure random bytes and convert them to a string of chars.
  rand = gen_random_bytes($1);
  for i in 0..len-1 loop
    -- rand indexing is zero-based, chars is 1-based.
    result = result || chars[1 + (get_byte(rand, i) % array_length(chars, 1))];
  end loop;
  return result;
end;
$$ language plpgsql;

-- return random string confirmed to not exist in given tablename.colname
create function unique_random(len int, _table text, _col text) returns text as $$
declare
  result text;
  numrows int;
begin
  result = random_string(len);
  loop
    execute format('select 1 from %I where %I = %L', _table, _col, result);
    get diagnostics numrows = row_count;
    if numrows = 0 then
      return result; 
    end if;
    result = random_string(len);
  end loop;
end;
$$ language plpgsql;

CREATE TABLE users (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    username    VARCHAR(50) NOT NULL,
    password    VARCHAR(200) NOT NULL,
    email       VARCHAR(50) NOT NULL,
    imgsrc      VARCHAR(50) DEFAULT '/static/profile/avatars/default.jpeg'
);

CREATE TABLE presentation (
    id                      BIGSERIAL NOT NULL PRIMARY KEY,
    creator_id              BIGINT REFERENCES users (id),
    name                    VARCHAR(128) DEFAULT 'Новая презентация' NOT NULL,
    viewmode                BOOLEAN DEFAULT FALSE NOT NULL,
    code                    VARCHAR(4) UNIQUE DEFAULT unique_random(4, 'presentation', 'code'),
    demo_idx                SMALLINT DEFAULT 0 NOT NULL,
    url                     VARCHAR(128) DEFAULT '/static/presentation/slides/' NOT NULL,
    converted_slide_num     SMALLINT DEFAULT 0 NOT NULL,
    quiz_num                SMALLINT DEFAULT 0 NOT NULL,
    emotion_like            BIGINT DEFAULT 0 NOT NULL,
    emotion_love            BIGINT DEFAULT 0 NOT NULL,
    emotion_laughter        BIGINT DEFAULT 0 NOT NULL,
    emotion_surprise        BIGINT DEFAULT 0 NOT NULL,
    emotion_sad             BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE slideorder (
    id              BIGSERIAL NOT NULL PRIMARY KEY,
    presentation_id BIGINT REFERENCES presentation (id),
    type            VARCHAR(64) NOT NULL,
    item_id         BIGINT NOT NULL,
    idx             SMALLINT NOT NULL
);

CREATE TABLE quiz (
    id              BIGSERIAL NOT NULL PRIMARY KEY,
    type            VARCHAR(64) DEFAULT 'horizontal' NOT NULL,
    question        VARCHAR(512) NOT NULL,
    runout          BOOLEAN DEFAULT TRUE NOT NULL,
    seconds_num     BIGINT DEFAULT 0 NOT NULL,
    result_after    BOOLEAN DEFAULT FALSE NOT NULL,
    price           BIGINT DEFAULT 0 NOT NULL,
    extra_points    BOOLEAN DEFAULT FALSE NOT NULL,
    start_time      TIMESTAMP DEFAULT current_timestamp NOT NULL,
    background      VARCHAR(16) NOT NULL,
    font_color      VARCHAR(16) NOT NULL,
    font_size       VARCHAR(16) NOT NULL,
    graph_color     VARCHAR(16) NOT NULL
);

CREATE TABLE convertedslide (
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    name    VARCHAR(64) NOT NULL,
    width   SMALLINT NOT NULL,
    height  SMALLINT NOT NULL
);

CREATE TABLE vote (
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    quiz_id     BIGINT REFERENCES quiz (id),
    idx         SMALLINT NOT NULL,
    option      VARCHAR(512) NOT NULL,
    correct     BOOLEAN DEFAULT FALSE NOT NULL,
    votes_num   BIGINT NOT NULL,
    color       VARCHAR(16) NOT NULL
);

CREATE TABLE question (
    id              BIGSERIAL NOT NULL PRIMARY KEY,
    presentation_id BIGINT REFERENCES presentation (id),
    idx             BIGINT DEFAULT 0 NOT NULL,
    option          VARCHAR(512) NOT NULL,
    likes           BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE voters (
    id              BIGSERIAL NOT NULL PRIMARY KEY,
    presentation_id BIGINT REFERENCES presentation (id),
    name            VARCHAR(32) DEFAULT '' NOT NULL,
    points          BIGINT DEFAULT 0 NOT NULL,
    top_place       SMALLINT DEFAULT 0 NOT NULL
);

CREATE TABLE voter_quiz (
    voter_id BIGINT             REFERENCES voters (id),
    quiz_id BIGINT              REFERENCES quiz (id),
    CONSTRAINT voter_quiz_id    PRIMARY KEY (voter_id, quiz_id)
);
