# Pokedex-CLI
A Command Line written in Golang, is just a make-believe device that lets us look up information about Pokemon - things like their name, type, and stats.

## Usage
1. Use `help` command to know about all the commands existing in the system.
2. Intially you can use `map` command to go through the locations,
3. For any location you can use the `expore <location_name>` to see all the pokemons available in that location.
4. In order to catch that pokemon use `catch <pokemon_name>`.
5. Finally you can see all the pokemons you caught using inspect command.

## 📖 CLI Commands

| Command                          | Description                                                    |
| -------------------------------- | ---------------------------------------------------------------|
| `catch <pokemon-name>`           | Tries to Catch the given Pokemon                               | 
| `inspect <pokemon-name>`         |  Inspects the details of the Pokemons that are already caught  |
| `pokedex`                        | Displays all the pokemons that are caught                      |
| `help`                           |  Displays a all the commands available along with their usage  |
| `exit`                           | Exit the Pokedex                                               |
| `map`                            | Displays the names of 20 location areas                        |
| `mapb`                           | Displays the names of previous 20 location areas               |
| `explore <location-area>`        | Displays list of all Pokemons available in a location area     |


