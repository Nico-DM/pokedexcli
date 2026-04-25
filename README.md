# Pokedex CLI

A CLI for [PokeAPI](https://pokeapi.co/) made for [Boot.dev](https://boot.dev).

## Usage

```
go run .
```

## Commands

- `help`: Displays a help message
- `catch <pokemon_name>`: Attempt to catch a pokemon
- `inspect <pokemon_name>`: View details about a caught pokemon
- `explore <location_name>`: Explore a location
- `map`: Get the next page of locations
- `mapb`: Get the previous page of locations
- `pokedex`: See all the pokemon you've caught
- `exit`: Exit the Pokedex

## Testing

```
go test ./...
```

## Possible Improvements

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon