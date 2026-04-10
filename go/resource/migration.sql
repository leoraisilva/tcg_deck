CREATE TABLE IF NOT EXISTS deck (
    id SERIAL PRIMARY KEY,
    quantidade INT NOT NULL,
    tipo VARCHAR(50) NOT NULL,
    estatistica INT NOT NULL
);

CREATE TABLE IF NOT EXISTS deck_card (
    id_deck INT NOT NULL,
    card_deck INT NOT NULL,
    PRIMARY KEY (id_deck, card_deck),
    FOREIGN KEY (id_deck) REFERENCES deck(id),
    FOREIGN KEY (card_deck) REFERENCES cards (id)
);

CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    type_card VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS cards_pokemon (
    id_card INT NOT NULL,
    card_pokemon INT NOT NULL,
    PRIMARY KEY (id_card, card_pokemon),
    FOREIGN KEY (id_card) REFERENCES cards(id),
    FOREIGN KEY (card_pokemon) REFERENCES pokemon (id)
);

CREATE TABLE IF NOT EXISTS cards_apoiador (
    id_card INT NOT NULL,
    card_apoiador INT NOT NULL,
    PRIMARY KEY (id_card, card_apoiador),
    FOREIGN KEY (id_card) REFERENCES cards(id),
    FOREIGN KEY (card_apoiador) REFERENCES apoiador (id)
);

CREATE TABLE IF NOT EXISTS cards_item (
    id_card INT NOT NULL,
    card_item INT NOT NULL,
    PRIMARY KEY (id_card, card_item),
    FOREIGN KEY (id_card) REFERENCES cards(id),
    FOREIGN KEY (card_item) REFERENCES item (id)
);

CREATE TABLE IF NOT EXISTS estatistica (
    id SERIAL PRIMARY KEY,
    vitoria INT NOT NULL,
    derrota INT NOT NULL,
    total INT NOT NULL,
    pontos_ganho INT NOT NULL,
    pontos_perdido INT NOT NULL,
    media_pontos DECIMAL(10,2) NOT NULL
);