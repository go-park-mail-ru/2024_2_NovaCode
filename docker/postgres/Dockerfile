FROM postgres:16

RUN apt-get update && apt-get install -y \
    hunspell \
    hunspell-ru \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/share/postgresql/16/tsearch_data && \
    cp /usr/share/hunspell/ru_RU.aff /usr/share/postgresql/16/tsearch_data/russian.affix && \
    cp /usr/share/hunspell/ru_RU.dic /usr/share/postgresql/16/tsearch_data/russian.dict && \
    ln -s /usr/share/postgresql/tsearch_data /usr/share/postgresql/16/tsearch_data
