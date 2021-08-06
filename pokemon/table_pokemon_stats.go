package pokemon

import (
	"context"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonStats(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pokemon_stat",
		Description: "Stats determine certain aspects of battles. Each Pokémon has a value for each stat which grows as they gain levels and can be altered momentarily by effects in battles.",
		List: &plugin.ListConfig{
			Hydrate: listPokemonStats,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			// TODO: Add support for 'id' key column
			//KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate: getPokemonStats,
			// Bad error message is a result of https://github.com/mtslzr/pokeapi-go/issues/29
			ShouldIgnoreError: isNotFoundError([]string{"invalid character 'N' looking for beginning of value"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The identifier for this resource.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemonStats,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "name",
				Description: "The name for this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "game_index",
				Description: "ID the games use for this stat.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "is_battle_only",
				Description: "Whether this stat only exists within a battle.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "affecting_moves",
				Description: "A detail of moves which affect this stat positively or negatively.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "affecting_natures",
				Description: "A detail of natures which affect this stat positively or negatively.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "characteristics",
				Description: "A list of characteristics that are set on a Pokémon when its highest base stat is this stat.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "move_damage_class",
				Description: "The class of damage this stat is directly related to.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemonStats,
			},
			{
				Name:        "names",
				Description: "The name of this resource listed in different languages.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPokemonStats,
			},
		},
	}
}

func listPokemonStats(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("listPokemonStats")

	offset := 0

	for true {

		// Reference - https://github.com/mtslzr/pokeapi-go#get-stats
		resources, err := pokeapi.Resource("stat", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_stat.listPokemonStats", "query_error", err)
			return nil, err
		}

		for _, pokemon_stat := range resources.Results {
			d.StreamListItem(ctx, pokemon_stat)
		}

		// No next URL returned
		if len(resources.Next) == 0 {
			break
		}

		urlOffset, err := extractUrlOffset(resources.Next)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_stat.listPokemonStats", "extract_url_offset_error", err)
			return nil, err
		}

		//Set next offset
		offset = urlOffset
	}

	return nil, nil
}

func getPokemonStats(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getPokemonStats")

	var name string
	//var id string

	if h.Item != nil {
		result := h.Item.(structs.Result)
		name = result.Name
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	logger.Debug("name", name)

	// Reference - https://github.com/mtslzr/pokeapi-go#get-stats
	// Must pass an ID (e.g. "1") or name (e.g. "hp").
	pokemon_stat, err := pokeapi.Stat(name)

	if err != nil {
		plugin.Logger(ctx).Error("pokemon_stat.pokemonGetStat", "query_error", err)
		return nil, err
	}

	return pokemon_stat, nil
}
