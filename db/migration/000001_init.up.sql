CREATE TABLE author (
    id INT GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
	primary key (id)
);

CREATE TABLE todo (
    id INT GENERATED ALWAYS AS IDENTITY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (id),
    author_id INT NOT NULL,
    constraint todo_author FOREIGN KEY (author_id) REFERENCES author (id)
);