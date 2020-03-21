# Errata for _gorestful_

Corrections for the book [_Building RESTful Web Services with Go_](). The pages listed are for the released book itself, not for any preprints or other forms of the articles.

### First Printing, December 2017

## Page 162, Installing the PostgrSQL database.
#### sudo sh -c 'echo “deb ht<span>tp://apt.postgresql.org/pub/repos/apt/ \`lsb_release -cs\`-pgdg main” >> /etc/apt/sources.list.d/pgdg.list'

For `Ubuntu 18.04LTS` open the file `/etc/apt/sources.list.d/pgdg.list` as [root]() and if the first and
 last characters are double quotes, remove them to overcome error:

#### E: Type ‘“deb’ is not known on line 1 in source list /etc/apt/sources.list.d/pgdg.list

produced when doing:

#### sudo apt-get update
