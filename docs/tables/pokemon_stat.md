# Table: pokemon_stat

Stats determine certain aspects of battles. Each Pokémon has a value for each stat which grows as they gain levels and can be altered momentarily by effects in battles.

## Examples

### steampipe query

```bash

> .inspect pokemon.pokemon_stat
+-------------------+---------+-----------------------------------------------------------------------------------------------+
| column            | type    | description                                                                                   |
+-------------------+---------+-----------------------------------------------------------------------------------------------+
| affecting_moves   | jsonb   | A detail of moves which affect this stat positively or negatively.                            |
| affecting_natures | jsonb   | A detail of natures which affect this stat positively or negatively.                          |
| characteristics   | jsonb   | A list of characteristics that are set on a Pokémon when its highest base stat is this stat.  |
| game_index        | bigint  | ID the games use for this stat.                                                               |
| id                | bigint  | The identifier for this resource.                                                             |
| is_battle_only    | boolean | Whether this stat only exists within a battle.                                                |
| move_damage_class | jsonb   | The class of damage this stat is directly related to.                                         |
| name              | text    | The name for this resource.                                                                   |
| names             | jsonb   | The name of this resource listed in different languages.                                      |
+-------------------+---------+-----------------------------------------------------------------------------------------------+

```
### Basic info

```sql
select
  id,
  name,
  game_index,
  is_battle_only,
  affecting_moves,
  affecting_natures,
  characteristics,
  move_damage_class,
  names
from
  pokemon_stat
```

### List pokémon stats where game index is greater than 3 and less than 10

```sql
select 
    id, 
    name, 
    game_index, 
    move_damage_class 
from pokemon.pokemon_stat 
    where game_index > 3 
    and game_index < 10;
```
