--
-- Tables
--

create table tracks
(
    id   serial not null primary key,
    name text   not null
);

create table albums
(
    id   serial not null primary key,
    name text   not null
);

create table artists
(
    id   serial not null primary key,
    name text   not null
);

create unique index artists_lower_case_name_index on artists ((lower(name)));

create table genres
(
    id   serial not null primary key,
    name text   not null
);

create unique index genres_lower_case_name_index on genres ((lower(name)));

--
-- Join Tables
--

create table track_artists
(
    track_id  int not null,
    artist_id int not null,
    primary key (track_id, artist_id),
    foreign key (track_id) references tracks (id) ON DELETE CASCADE,
    foreign key (artist_id) references artists (id) ON DELETE CASCADE
);

create index on track_artists (track_id);
create index on track_artists (artist_id);

create table track_albums
(
    track_id int not null,
    album_id int not null,
    primary key (track_id, album_id),
    foreign key (track_id) references tracks (id) ON DELETE CASCADE,
    foreign key (album_id) references albums (id) ON DELETE CASCADE
);

create index on track_albums (track_id);
create index on track_albums (album_id);

create table track_genres
(
    track_id int not null,
    genre_id int not null,
    primary key (track_id, genre_id),
    foreign key (track_id) references tracks (id) ON DELETE CASCADE,
    foreign key (genre_id) references genres (id) ON DELETE CASCADE
);

create index on track_genres (track_id);
create index on track_genres (genre_id);
