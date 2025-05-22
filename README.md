# 🧭 Pokedex CLI

**Pokedex-CLI** is a command-line application written in **Golang** that simulates a Pokédex. You can explore locations, find Pokémon, catch them, and inspect their stats — all from your terminal!

---

## 🎮 Features

- 🗺️ Explore different Pokémon locations
- 🎯 Catch Pokémon by name
- 🔍 Inspect details of the Pokémon you've caught
- 📚 Maintain a list of caught Pokémon
- 🧭 Navigate between multiple location areas

---

## 🚀 Usage

```bash
go run main.go
```
Once running, you can start entering commands into the terminal.

## Usage
1. Use `help` command to know about all the commands existing in the system.
2. Intially you can use `map` command to go through the locations,
3. For any location you can use the `expore <location_name>` to see all the pokemons available in that location.
4. In order to catch that pokemon use `catch <pokemon_name>`.
5. Finally you can see all the pokemons you caught using inspect command.

## 📖 CLI Commands

| Command                   | Description                                               |
| ------------------------- | --------------------------------------------------------- |
| `help`                    | Show all available commands and their usage               |
| `map`                     | Display the names of the next 20 location areas           |
| `mapb`                    | Display the names of the previous 20 location areas       |
| `explore <location-area>` | Show all Pokémon available in the specified location area |
| `catch <pokemon-name>`    | Attempt to catch a specified Pokémon                      |
| `inspect <pokemon-name>`  | Show details of a caught Pokémon                          |
| `pokedex`                 | List all Pokémon you’ve caught                            |
| `exit`                    | Exit the Pokedex CLI                                      |

## 💡 Getting Started
1. Run the app using go run main.go or build with go build and run the executable.

2. Use the help command to see the list of available actions.

3. Start exploring using the map command.

4. Visit a location using explore <location-area>.

5. Try catching Pokémon with catch <pokemon-name>.

6. Use inspect <pokemon-name> to view stats and info.

## 📦 Example Workflow
```
> help
> map
> explore kanto-route-2
> catch pikachu
> pokedex
> inspect pikachu
```
### 🧑‍💻 Author
[Harsha Vardhan Mirthinti](https://www.linkedin.com/in/harshavardhanmirthinti/)

## 💡 Contributing
Pull requests are welcome! Feel free to open issues or suggest improvements.



